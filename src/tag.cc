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


#include "tag.h"

#include <utility>

Tag::Tag(std::string key, std::string value) : _key(std::move(key)), _value(std::move(value)) {
    _value.erase(_value.find_last_not_of(' ') + 1);
}

std::string Tag::getKey() {
    return _key;
}

std::string Tag::getValue() {
    // 标签内容最多支持2k
    int len = _value.length();
    if (len > 2048){
        _value.resize(2048);
        _value.append("... total len: ");
        _value.append(std::to_string(len));
    }
    return _value;
}