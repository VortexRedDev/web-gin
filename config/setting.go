package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Profile struct {
	Profile string `yaml:"profile"`
}

type Setting struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
	Env      string   `yaml:"env"`
}

type database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
}

type server struct {
	Port string `yaml:"port"`
}

var conf = &Setting{}
var envConfig = &Profile{}

func init() {
	loadConfig()
}

func loadConfig() {
	config, err := ioutil.ReadFile("./config/profile.yml")
	if err != nil {
		fmt.Errorf("配置文件profile.yml读取错误", err)
		panic("配置文件profile.yml读取错误")
	}
	err = yaml.Unmarshal(config, envConfig)

	file, err := ioutil.ReadFile("./config/" + envConfig.Profile + ".yml")
	if err != nil {
		fmt.Errorf("配置文件"+envConfig.Profile+"读取错误", err)
		panic("配置文件" + envConfig.Profile + "读取错误")
	}
	err = yaml.Unmarshal(file, conf)

	if err != nil {
		fmt.Errorf("配置文件"+envConfig.Profile+"读取错误", err)
		panic("配置文件" + envConfig.Profile + "读取错误")
	}
}

func Conf() *Setting {
	return conf
}
