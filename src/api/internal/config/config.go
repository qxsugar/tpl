package config

// 统一配置管理
// 统一地方读取变量

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		Default string `yaml:"default"`
	} `yaml:"database"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
}

var conf Config

func loadConfigFromFile(fileName string) error {
	log.Println("load config from file", "filename", fileName)
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("load config failed", "error", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("parse config failed", "error", err)
		return err
	}

	return nil
}

func loadConfig() {
	viper.AutomaticEnv()
	runEnv := viper.GetString("RUN_ENV")
	log.Println("load config from env", "env", runEnv)
	switch runEnv {
	case "dev":
		err := loadConfigFromFile("./config/dev.yaml")
		if err != nil {
			panic(err)
		}
	case "product":
		err := loadConfigFromFile("./config/product.yaml")
		if err != nil {
			panic(err)
		}
	default:
		err := loadConfigFromFile("./config/dev.yaml")
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	loadConfig()
}

func GetConfig() Config {
	return conf
}
