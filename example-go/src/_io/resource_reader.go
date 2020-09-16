package _io

import (
	"errors"
	"github.com/spf13/viper"
)

type Properties struct {
	properties map[string]string
}

func (p *Properties) SetProperties(k string, v string) {
	if p.properties == nil {
		p.properties = make(map[string]string)
	}
	p.properties[k] = v
}

func (p *Properties) GetProperties(k string) (string, error) {
	if p.properties == nil {
		return "", errors.New("properties not init")
	}
	v := p.properties[k]
	if v == "" {
		return "", errors.New("properties key not found")
	}
	return v, nil
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
		return nil, errors.New("read properties file error")
	}

	properties := Properties{
		properties: make(map[string]string),
	}

	for _, k := range viper.AllKeys() {
		properties.SetProperties(k, viper.GetString(k))
	}
	return &properties, nil
}
