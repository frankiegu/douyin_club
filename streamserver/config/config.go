package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LBAddr string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	//构造一个基于配置文件的解码器
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	//将解码器的内容输出到结构体中
	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}








