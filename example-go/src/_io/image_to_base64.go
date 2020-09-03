package _io

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
)

type (
	Image struct {
		url    string
		base64 string
	}
)

func NewImage(url string) *Image {
	f, err := os.Open(url)
	if err != nil {
		panic("图片打开失败")
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic("图片读取失败")
	}

	encode := base64.StdEncoding.EncodeToString(content)
	return &Image{
		url:    url,
		base64: encode,
	}
}

/**
	将图片转换为动画化，返回base64字符串
 */
func (image *Image) ToCartoonBase64() (string, error) {

	return "", errors.New("")
}

/**
	将图片转换为动画化，返回url
 */
func (image *Image) ToCartoonUrl() (string, error) {

	return "", errors.New("")
}
