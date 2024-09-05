#!/bin/bash

plugin_type=driver

fileName="${PWD##*/}"
plugin_suffix=$plugin_type



########### 下面的不建议修改 ###########
# 检查环境变量 MY_ENV_VAR
if [ "$IS_DEBUG" = "true" ]; then
    go build -tags=prod -gcflags="all=-N -l" -buildmode=plugin -o ${pluginPath}/${plugin_type}/${fileName}.${plugin_suffix}
else
    go build -tags=prod -ldflags "-s -w" -buildmode=plugin -o ${pluginPath}/${plugin_type}/${fileName}.${plugin_suffix}
#    echo "   go build -tags=prod -ldflags \"-s -w\" -buildmode=plugin -o ${pluginPath}/${plugin_type}/${fileName}.${plugin_suffix}"
fi

