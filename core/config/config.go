package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Config *Yaml

// 注意命名要和配置文件的相同
type Yaml struct {
	MysqlIp       string `yaml:"mysqlIp"`
	MysqlPort     int    `yaml:"mysqlPort"`
	MysqlUser     string `yaml:"mysqlUser"`
	MysqlPassword string `yaml:"mysqlPassword"`
	MysqlDb       string `yaml:"mysqlDb"`
	RedisIp       string `yaml:"redisIp"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDb       int    `yaml:"redisDb"`
}

func GetConfig() Yaml {
	return *Config
}

func Init() {
	env := os.Getenv("ENV")
	configPath := "conf/config.yaml"
	if env == "prod" {
		log.Warningln("*************prod环境：**************")
		configPath = "conf/config.prod.yaml"
	} else {
		log.Warningln("*************开发环境**************")
	}
	var yaml = new(Yaml)
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败")
	}
	err1 := viper.Unmarshal(yaml)
	if err1 != nil {
		log.Fatal("配置失败")
	}
	Config = yaml
	log.Infoln("初始化配置文件成功")
	log.Infoln("******")
	log.Infoln(Config)
	log.Infoln("******")
}
