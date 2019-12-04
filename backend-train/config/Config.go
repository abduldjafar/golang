package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	Postgres postgres
	Mail     mail
	Oauth    oauth
	Server	 server
}

type postgres struct {
	Url      string
	Ip       string
	Port     string
	Db       string
	User     string
	Password string
}

type mail struct {
	Smtp     string
	From     string
	Password string
	Port     string
}

type oauth struct {
	Rurl         string
	ClientID     string
	ClientSecret string
}

type server struct {
	Port string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		fmt.Println(err)
	}
}
