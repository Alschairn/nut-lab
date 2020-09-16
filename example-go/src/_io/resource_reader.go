package _io

import (
	"errors"
	"github.com/spf13/viper"
)

type Properties struct {
	properties map[string]string
}

func (p *Properties) Put(k string, v string) {
	if p.properties == nil {
		p.properties = make(map[string]string)
	}
	p.properties[k] = v
}

func (p *Properties) Get(k string) (string, error) {
	if p.properties == nil {
		return "", errors.New("properties 未初始化")
	}
	return p.properties[k], nil
}

/**
读取properties文件
*/
func ReaderProperties(path string, filename string) (*Properties, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("properties")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("load error")
	}

	properties := Properties{
		properties: make(map[string]string),
	}
	for k, v := range viper.AllSettings() {
		properties.Put(k, v.(string))
	}
	return &properties, nil
}
