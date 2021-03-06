package conf

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// ConfigFile ..
var ConfigFile string

// Config ...
type Config struct {
	DBMysql       string `yaml:"db_mysql"`
	RedisAddr     string `yaml:"redis_addr"`
	RedisPassword string `yaml:"redis_password"`
	RunPort       int    `yaml:"run_port"`
}

// YAMLLoad 加载文件并解析，包含加密项的自动解密
func YAMLLoad(fn string, v *Config) error {
	dat, err := ioutil.ReadFile(fn)
	if err != nil {
		return fmt.Errorf("read config file %v error. err = %v", fn, err)
	}

	err = yaml.Unmarshal(dat, v)
	if err != nil {
		return fmt.Errorf("parse config file %v error. err = %v", fn, err)
	}
	log.Printf("config initialize success. config = %v\n", v)
	return nil
}

// Init 传入带有默认值的 config, 加载配置到 config 中
func InitConfig(f string, v *Config) {
	err := YAMLLoad(f, v)
	if err != nil {
		glog.Fatalf("Call YAMLLoad failed,err:%+v", err)
	}
}

//DefaultConfig .
var DefaultConfig = Config{
	DBMysql:       "thewall:0Thewall!2022@tcp(114.55.139.105:3306)/thewall?charset=utf8mb4",
	RedisAddr:     "localhost:6379",
	RedisPassword: "Aa123.",
	RunPort:       15004,
}

//127.0.0.1:6379>
//127.0.0.1:6379> auth biterIam007
//OK
//127.0.0.1:6379> config set requirepass a123.
//OK
