#!/bin/bash

# 一般只需要改这里的
type=bms
plugin_type=driver

fileName="${PWD##*/}"
plugin_suffix=$plugin_type



########### 下面的不建议修改 ###########

# 检查环境变量 MY_ENV_VAR
if [ "$IS_DEBUG" = "true" ]; then
    go build -gcflags="all=-N -l" -buildmode=plugin -o ${pluginPath}/${plugin_type}/${type}_${fileName}.${plugin_suffix}
else
    go build -ldflags "-s -w" -buildmode=plugin -o ${pluginPath}/${plugin_type}/${type}_${fileName}.${plugin_suffix}
fi

