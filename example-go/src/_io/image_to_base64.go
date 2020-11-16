package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
)

const (
	resourcePath = "./resource"        // 资源地址
	filename     = "application"       // 资源文件名
	accessKey    = "tencent.secretKey" // 访问key
	accessId     = "tencent.secretId"  // 访问id
)

var properties Properties


func init() {
	p, err := ReaderProperties(resourcePath, filename)
	if p == nil || err != nil {
		panic("load properties error")
	}
	properties = *p
}

func ToBase64(url string) string {
	f, err := os.Open(url)
	if err != nil {
		panic("Image file open error")
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic("Image file read error")
	}
	return base64.StdEncoding.EncodeToString(content)
}

/**
将图片转换为动画化，返回base64字符串
*/
func ToCartoonBase64() (string, error) {

	return "", errors.New("")
}

/**
将图片转换为动画化，返回url
*/
func ToCartoonUrl() (string, error) {

	return "", errors.New("")
}
