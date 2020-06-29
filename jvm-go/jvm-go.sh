#!/bin/bash

# 设置go path到当前目录
export GOPATH=$(cd `dirname $0`; pwd)
export GOBIN=$(cd `dirname $0`; pwd)/bin

# 构建main函数
cd $(cd `dirname $0`; pwd)/src/main && go build main.go
