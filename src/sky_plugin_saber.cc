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


#include "sky_plugin_saber.h"
#include "php_skywalking.h"
#include "segment.h"
#include "sky_utils.h"
extern void (*ori_execute_ex)(zend_execute_data *execute_data);

Span *sky_plugin_saber(zend_execute_data *execute_data, const std::string &class_name, const std::string &function_name) {

    auto *segment = sky_get_segment(execute_data, -1);

    auto *span = segment->createSpan(SkySpanType::Exit, SkySpanLayer::Http, 8002);
    uint32_t arg_count = ZEND_CALL_NUM_ARGS(execute_data);

    if (arg_count == 1) {
        zval *options = ZEND_CALL_ARG(execute_data, 1);
        std::string uri = sky_hashtable_default(options,"uri","");
        std::string method = sky_hashtable_default(options,"method","");

        zval *args;
        //post的传参
        sky_hashtable_default(options,"data",&args);
        std::string request_args = sky_json_encode(args);

//        php_printf(request_args.c_str());
//        php_printf(method.c_str());
//        php_printf("\n");
//        php_printf("===================================");
//        php_printf("\n");

        span->setOperationName("saber::"+method);
        span->addTag("url", uri);
        span->addTag("http.params", request_args);

        ori_execute_ex(execute_data);
        span->setEndTIme();

        return span;
    }

    return nullptr;
}

