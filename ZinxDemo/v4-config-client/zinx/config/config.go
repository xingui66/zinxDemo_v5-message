package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	IP string
	Port uint32
	Name string
	TcpVersion string
}

var GlobalConfig Config

func init(){
	err :=LoadConfig()
	if err != nil {
		fmt.Println("加载json失败")
		return
	}
}

func LoadConfig() error {
	fmt.Println("开始读取配置文件...")

	//1. 读取配置文件
	//go run server_main.go启动时，目录是基于程序运的位置
	info, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		fmt.Printf("读取配置文件失败:", err)
		return err
	}

	fmt.Printf("配置文件信息: %v\n", info)

	//2. json解析配置文件，将数据放置到一个全局的配置结构中
	err = json.Unmarshal(info, &GlobalConfig)
	if err != nil {
		fmt.Println("配置文件解析失败:", err)
		return err
	}

	fmt.Println("配置文件解析解析成功:", GlobalConfig)
	return nil
}
