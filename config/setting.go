package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Profile string `yaml:"profile"`
}

type Setting struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Env      string   `yaml:"env"`
}

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
}

type Server struct {
	Port string `yaml:"port"`
}

var setting = &Setting{}

func init() {
	loadConfig()
}

func loadConfig() {
	config, err := os.ReadFile("./config/profile.yml")
	if err != nil {
		panic("配置文件profile.yml读取错误")
	}
	var envConfig = &Profile{}
	err = yaml.Unmarshal([]byte(config), &envConfig)
	if err != nil {
		panic("配置文件" + envConfig.Profile + "读取错误")
	}
	file, err := os.ReadFile("./config/" + envConfig.Profile + ".yml")
	if err != nil {
		panic("配置文件" + envConfig.Profile + "读取错误")
	}
	err = yaml.Unmarshal(file, &setting)
	if err != nil {
		panic("配置文件" + envConfig.Profile + "读取错误")
	}
}

func Getsetting() *Setting {
	return setting
}
