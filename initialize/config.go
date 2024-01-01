package initialize

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"wechat/config"
	"wechat/global"
)

var ViperConfig config.Server

func InitViperConfig() {
	configPathDir := fmt.Sprintf("%s%s", "./", global.CONFIG_DEFAULT)
	data, err := ioutil.ReadFile(configPathDir)
	if err != nil {
		log.Fatal(configPathDir + " Error:" + err.Error())
	}
	//yaml.Unmarshal会根据yaml标签的字段进行赋值
	err = yaml.Unmarshal(data, &ViperConfig)
	if err != nil {
		log.Fatal(configPathDir + " Unmarshal Error:" + err.Error())
	}
	global.GVA_CONFIG = &ViperConfig
}
