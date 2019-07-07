package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct{
		Addr string `yaml:"addr"`
		Grpc string `yaml:"grpc"`
		Debug bool `yaml:"debug"`
	}
	Postgresql struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Dbname string `yaml:"dbname"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Sslmode string `yaml:"sslmode"`
		MaxIdleConns int `yaml:"maxIdleConns"`
		MaxOpenConns int `yaml:"maxOpenConns"`
		Debug bool `yaml:"debug"`
	}
	Redis struct{
		Addr string `yaml:"addr"`
		Password string `yaml:"password"`
		Db int `yaml:"db"`
	}
	Oss struct{
		Endpoint string `yaml:"endpoint"`
		AccessKey string `yaml:"accessKey"`
		AccessSecretKey string `yaml:"accessSecretKey"`
		BucketName string `yaml:"bucketName"`
	}
	Email struct{
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Host string `yaml:"host"`
		Subject string `yaml:"subject"`
	}
	Jwt struct{
		Secret string `yaml:"secret"`
	}
}

var YamlConfig Config

func init() {
	// fileName := os.Getenv("CONFIGPATH")
	config , err := ioutil.ReadFile("./config.yaml")
	if err != nil{
		panic("can not find CONFIGPATH in envs")
	}
	err = yaml.Unmarshal(config, &YamlConfig)
	if err != nil{
		log.Println(err)
	}
}
