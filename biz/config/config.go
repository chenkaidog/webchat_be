package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Init() {
	content, err := os.ReadFile("./conf/deploy.local.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(content, &globalConfig); err != nil {
		panic(err)
	}
}

func GetMySQLConf() MySQLConf {
	return globalConfig.MySQL
}

func GetRedisConf() RedisConf {
	return globalConfig.Redis
}

func GetBaiduConf() BaiduAppConf {
	return globalConfig.Baidu
}

func GetOpenAIConf() OpenaiConf {
	return globalConfig.Openai
}

func GetHttpProxyConf() HttpProxyConf {
	return globalConfig.HttpProxy
}

var globalConfig ServiceConf

type ServiceConf struct {
	MySQL     MySQLConf     `yaml:"mysql"`
	Redis     RedisConf     `yaml:"redis"`
	Baidu     BaiduAppConf  `yaml:"baidu"`
	Openai    OpenaiConf    `yaml:"openai"`
	HttpProxy HttpProxyConf `yaml:"http_proxy"`
}

type MySQLConf struct {
	DBName   string `yaml:"db_name"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConf struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type BaiduAppConf struct {
	AppKey    string `yaml:"app_key"`
	AppSecret string `yaml:"app_secret"`
}

type OpenaiConf struct {
	ApiKey string `yaml:"api_key"`
}

type HttpProxyConf struct {
	Url string `yaml:"url"`
}
