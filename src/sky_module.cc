/*
 * Copyright 2021 SkyAPM
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */


#include <map>
#include <iostream>
#include "sky_module.h"
#include <boost/interprocess/ipc/message_queue.hpp>
#include <fstream>
#include <string>
#include <sstream>
#include <ostream>

#include "segment.h"
#include "sky_utils.h"
#include "sky_plugin_curl.h"
#include "sky_execute.h"
#include "manager.h"
#include "sky_plugin_error.h"
#include "sky_log.h"
#include "base64.h"


extern struct service_info *s_info;

extern void (*ori_execute_ex)(zend_execute_data *execute_data);

extern void (*ori_execute_internal)(zend_execute_data *execute_data, zval *return_value);

extern void (*orig_curl_exec)(INTERNAL_FUNCTION_PARAMETERS);

extern void (*orig_curl_setopt)(INTERNAL_FUNCTION_PARAMETERS);

extern void (*orig_curl_setopt_array)(INTERNAL_FUNCTION_PARAMETERS);

extern void (*orig_curl_close)(INTERNAL_FUNCTION_PARAMETERS);

void sky_module_init() {

    std::ostringstream strPid;
    strPid << getpid();

    std::ostringstream strPPid;
    strPPid << getppid();
    sky_log("父进程 : pid = " + strPPid.str());


    ori_execute_ex = zend_execute_ex;
    zend_execute_ex = sky_execute_ex;

    ori_execute_internal = zend_execute_internal;
    zend_execute_internal = sky_execute_internal;

    if (SKYWALKING_G(error_handler_enable)) {
        sky_plugin_error_init();
    }


    // bind curl
    zend_function *old_function;
    if ((old_function = SKY_OLD_FN("curl_exec")) != nullptr) {
        orig_curl_exec = old_function->internal_function.handler;
        old_function->internal_function.handler = sky_curl_exec_handler;
    }
    if ((old_function = SKY_OLD_FN("curl_setopt")) != nullptr) {
        orig_curl_setopt = old_function->internal_function.handler;
        old_function->internal_function.handler = sky_curl_setopt_handler;
    }
    if ((old_function = SKY_OLD_FN("curl_setopt_array")) != nullptr) {
        orig_curl_setopt_array = old_function->internal_function.handler;
        old_function->internal_function.handler = sky_curl_setopt_array_handler;
    }
    if ((old_function = SKY_OLD_FN("curl_close")) != nullptr) {
        orig_curl_close = old_function->internal_function.handler;
        old_function->internal_function.handler = sky_curl_close_handler;
    }

    ManagerOptions opt;
    opt.version = SKYWALKING_G(version);
    opt.code = SKYWALKING_G(app_code);
    opt.grpc = SKYWALKING_G(grpc);
    opt.grpc_tls = SKYWALKING_G(grpc_tls_enable);
    opt.root_certs = SKYWALKING_G(grpc_tls_pem_root_certs);
    opt.private_key = SKYWALKING_G(grpc_tls_pem_private_key);
    opt.cert_chain = SKYWALKING_G(grpc_tls_pem_cert_chain);
    opt.authentication = SKYWALKING_G(authentication);

    sprintf(s_info->mq_name, "skywalking_queue_tanikawa");
//    if (std::string(sapi_module.name) == "cli") {
//        sprintf(s_info->mq_name, "skywalking_queue_tanikawa");
//    }else{
//        sprintf(s_info->mq_name, "skywalking_queue_%d", getpid());
//    }


//    sky_log("log 95: pid = " + strPid.str());

    try {
//        boost::interprocess::message_queue::remove(s_info->mq_name);
        boost::interprocess::message_queue(
                boost::interprocess::open_or_create,
//                boost::interprocess::create_only,
                s_info->mq_name,
                5120,
                SKYWALKING_G(mq_max_message_length),
                boost::interprocess::permissions(0666)
        );
    } catch (boost::interprocess::interprocess_exception &ex) {
        sky_log("message create error init : pid = " + std::string(ex.what()));
        php_error(E_WARNING, "%s %s", "[skywalking] create queue fail ", ex.what());
    }

    std::string queue_name = s_info->mq_name;
    sky_log("module init : pid = " + strPid.str() + ", queue_name: " + queue_name);
    char *module_name;
//    sprintf(module_name, "%s", sapi_module.name);
    sky_log("module init module type: " + std::string(sapi_module.name));

    // swoole 容器场景多次模块初始化会有问题,限定在1号进程才启用队列消费者
    // php-fpm方式会有fpm进程数量的消费者处理消息
    if (!(std::string(sapi_module.name) == "cli" && strPid.str() != "1")){
        sky_log("consumer init : pid = " + strPid.str() + "\n");
        new Manager(opt, s_info);
    }

//    sky_log("process name" + get_current_process_name());

}

void sky_module_cleanup() {
    char mq_name[32];

    sprintf(mq_name, "skywalking_queue_%d", getpid());

    sky_log("module clean: queue_name: " + std::string(mq_name));
    // 频繁sigbus，统一使用相同共享内存名称，并且不再不再销毁，随docker停止消除
//    if (strcmp(s_info->mq_name, mq_name) == 0 && getpid() == 1) {
//        sky_log("removed");
//        boost::interprocess::message_queue::remove(s_info->mq_name);
//    }

//    if (std::string(sapi_module.name) != "cli") {
//        sky_log("queue removed");
//        boost::interprocess::message_queue::remove(s_info->mq_name);
//    }
}

bool sky_request_init(zval *request, uint64_t request_id) {
    sky_log("http请求->  sky_request_init: ...");
    sky_log(sapi_module.name);
    array_init(&SKYWALKING_G(curl_header));

    zval *carrier = nullptr;
    zval *sw, *peer_val;
    std::string header;
    std::string uri;
    std::string peer;

    if (request != nullptr) {
        sky_log("http请求->  获取header,server: ...");

        zval *swoole_header = sky_read_property(request, "header", 0);
        zval *swoole_server = sky_read_property(request, "server", 0);

        if (SKYWALKING_G(version) == 8) {
            sky_log("http请求->  获取sw8头");
            sw = zend_hash_str_find(Z_ARRVAL_P(swoole_header), "sw8", sizeof("sw8") - 1);
        } else {
            sw = nullptr;
        }

        header = (sw != nullptr ? Z_STRVAL_P(sw) : "");
        sky_log("header: " + header);
        uri = Z_STRVAL_P(zend_hash_str_find(Z_ARRVAL_P(swoole_server), "request_uri", sizeof("request_uri") - 1));
        sky_log("uri: " + uri);
        std::string gray_list = SKYWALKING_G(trace_gray_list);

        sky_log("灰度列表:" + gray_list);
        if(!gray_list.empty() && !string_delim_contains(gray_list, uri, ",")){
            sky_log("不在灰度列表,返回false");
            return false;
        }
        peer_val = zend_hash_str_find(Z_ARRVAL_P(swoole_header), "host", sizeof("host") - 1);
        if (peer_val != nullptr) {
            peer = Z_STRVAL_P(peer_val);
            sky_log("header 中获取到host: " + peer);
        } else {
            char hostname[HOST_NAME_MAX + 1];
            if (gethostname(hostname, sizeof(hostname))) {
                hostname[0] = '\0';
            }
            peer_val = zend_hash_str_find(Z_ARRVAL_P(swoole_server), "server_port", sizeof("server_port") - 1);
            peer += hostname;
            peer += ":";
            peer += std::to_string(Z_LVAL_P(peer_val));
            sky_log("本机地址和port拼接 获取到host: " + peer);
        }
    } else {
        sky_log("http请求->  request是null，这个是异常的？？？");
        zend_bool jit_initialization = PG(auto_globals_jit);

        if (jit_initialization) {
            zend_string *server_str = zend_string_init("_SERVER", sizeof("_SERVER") - 1, 0);
            zend_is_auto_global(server_str);
            zend_string_release(server_str);
        }
        carrier = zend_hash_str_find(&EG(symbol_table), ZEND_STRL("_SERVER"));

        if (SKYWALKING_G(version) == 5) {
            sw = zend_hash_str_find(Z_ARRVAL_P(carrier), "HTTP_SW3", sizeof("HTTP_SW3") - 1);
        } else if (SKYWALKING_G(version) == 6 || SKYWALKING_G(version) == 7) {
            sw = zend_hash_str_find(Z_ARRVAL_P(carrier), "HTTP_SW6", sizeof("HTTP_SW6") - 1);
        } else if (SKYWALKING_G(version) == 8) {
            sw = zend_hash_str_find(Z_ARRVAL_P(carrier), "HTTP_SW8", sizeof("HTTP_SW8") - 1);
        } else {
            sw = nullptr;
        }

        header = (sw != nullptr ? Z_STRVAL_P(sw) : "");
        uri = get_page_request_uri();
        peer = get_page_request_peer();
        sky_log("get_page_request_uri， peer" + uri + "," + peer);
    }
    sky_log("初始化链路信息");
    std::map<uint64_t, Segment *> *segments;
    if (SKYWALKING_G(segment) == nullptr) {
        segments = new std::map<uint64_t, Segment *>;
        SKYWALKING_G(segment) = segments;
    } else {
        segments = static_cast<std::map<uint64_t, Segment *> *>SKYWALKING_G(segment);
    }
    sky_log("初始化segment");
    auto *segment = new Segment(SKYWALKING_G(app_code), s_info->service_instance, SKYWALKING_G(version), header);
    sky_log("segment写入到请求id segment数组");
    auto const result = segments->insert(std::pair<uint64_t, Segment *>(request_id, segment));
    if (not result.second) {
        result.first->second = segment;
    }

    sky_log("segment下新建span");
    auto *span = segments->at(request_id)->createSpan(SkySpanType::Entry, SkySpanLayer::Http, 8001);
    span->setOperationName(uri);
    span->setPeer(peer);
    span->addTag("url", uri);
    segments->at(request_id)->createRefs();
    sky_log("获取request_method,添加请求方法");
    zval *request_method = zend_hash_str_find(Z_ARRVAL(PG(http_globals)[TRACK_VARS_SERVER]), ZEND_STRL("REQUEST_METHOD"));
    if (request_method != NULL) {
        span->addTag("http.method", Z_STRVAL_P(request_method));
    }

    sky_log("初始化链路信息初始化完成");
    return true;
}

void sky_rpc_init(uint64_t request_id, swoft_json_rpc rpcData) {
    std::string header;
    // 初始化全局字典
    std::map<uint64_t, Segment *> *segments;
    if (SKYWALKING_G(segment) == nullptr) {
        segments = new std::map<uint64_t, Segment *>;
        SKYWALKING_G(segment) = segments;
    } else {
        segments = static_cast<std::map<uint64_t, Segment *> *>SKYWALKING_G(segment);
    }

    if (!rpcData.ext.traceid.empty()){
        // 按照skywalking header拼接串
        header.append("0").append("-") \
        .append(Base64::encode(rpcData.ext.traceid)).append("-") \
        .append(Base64::encode(rpcData.ext.traceid)).append("-") \
        .append(rpcData.ext.spanid).append("-") \
        .append(Base64::encode(rpcData.ext.serviceName)).append("-") \
        .append(Base64::encode(rpcData.ext.ServiceInstance)).append("-") \
        .append(Base64::encode(rpcData.ext.endpoint)).append("-") \
        .append(Base64::encode(rpcData.ext.address));
    }else{
        return;
//        header = "";
    }

    php_printf(header.c_str());
    auto *segment = new Segment(s_info->service, s_info->service_instance, SKYWALKING_G(version), header);
    auto const result = segments->insert(std::pair<uint64_t, Segment *>(request_id, segment));
    if (not result.second) {
        result.first->second = segment;
    }
    std::string logStr = std::to_string(request_id);

    auto *span = segments->at(request_id)->createSpan(SkySpanType::Entry, SkySpanLayer::Http, 8001);
    span->setOperationName(rpcData.method);
    span->setPeer(rpcData.ext.endpoint);
    span->addTag("url", rpcData.method);
    segments->at(request_id)->createRefs();
    span->addTag("http.method", "rpc");
    php_printf((logStr + "--------------------------------------------------------------------\n").c_str());
}

// 清除本次请求的消息
void sky_request_destory(uint64_t request_id){
    auto *segment = sky_get_segment(nullptr, request_id);
    std::string logStr = "\nclean\n";

    if (segment != nullptr){
        logStr.append("no null\n");
        delete segment;
    }

    php_printf(logStr.c_str());
}

void sky_request_flush(zval *response, uint64_t request_id) {
    auto *segment = sky_get_segment(nullptr, request_id);

    if (segment == nullptr){
        return;
    }
    if (response == nullptr) {
        segment->setStatusCode(SG(sapi_headers).http_response_code);
    }
    std::string msg = segment->marshal(SKYWALKING_G(app_code));
    delete segment;

    int msg_length = static_cast<int>(msg.size());

    int max_length = SKYWALKING_G(mq_max_message_length);
    if (msg_length > max_length) {
        sky_log("message is too big: " + std::to_string(msg_length) + ", mq_max_message_length=" + std::to_string(max_length));
        return;
    }

    sky_log("正在写入队列" + std::string(s_info->mq_name));
    try {
        boost::interprocess::message_queue mq(
                boost::interprocess::open_only,
                s_info->mq_name
        );
        if (!mq.try_send(msg.data(), msg.size(), 0)) {
            sky_log("sky_request_flush message_queue is fulled");
        }else{
            sky_log("成功写入队列"+ std::string(s_info->mq_name));
        }

    } catch (boost::interprocess::interprocess_exception &ex) {
        sky_log("sky_request_flush message_queue ex" + std::string(ex.what()));
//        php_error(E_WARNING, "%s %s", "[skywalking] open queue fail ", ex.what());
    }
}
