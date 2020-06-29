#!/bin/bash

# 设置gopath到当前目录
export GOPATH=$PWD && export GOBIN=$PWD/bin

# 构建main
cd src/main && go build main.go
