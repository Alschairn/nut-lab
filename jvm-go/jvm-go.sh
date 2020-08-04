#!/bin/bash
base_addr=$(cd `dirname $0`; pwd)
echo '执行文件目录: ' ${base_addr}

go_path=${base_addr}
export GOPATH=${go_path}
echo 'go path: ' ${go_path}

go_bin=${base_addr}/bin
export GOBIN=${go_bin}
echo 'go bin: ' ${go_bin}

main_path=${base_addr}/src/main/main.go
go build ${main_path}
echo '主函数: ' ${main_path}

