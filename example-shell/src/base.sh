#!/usr/bin/env bash

# 获取当前执行文件绝对路径
BASE_PATH=$(cd "$(dirname "$0")"; pwd)
echo ${BASE_PATH}

# 输出函数(两种方式)，多字符拼接
echo `date '+%Y-%m-%d %H:%M:%S'` ' append1 ' ' append2  '
echo $(date '+%Y-%m-%d %H:%M:%S') ' append1 ' ' append2  '

# 覆盖写
for VAR in 1 2 3 ; do
    echo ${VAR} > ./temp.txt
done
cat ${BASE_PATH}/temp.txt

# 清空文件
cat /dev/null > ${BASE_PATH}/temp.txt

# 追加写
for VAR in 1 2 3 ; do
    echo ${VAR} >> ${BASE_PATH}/temp.txt
done

cat ${BASE_PATH}/temp.txt

rm -f ${BASE_PATH}/temp.txt