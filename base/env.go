package base

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	ElasticHost string `mapstructure:"ELASTIC_HOST"`
}

var globalEnv = Env{}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Cannot Read Configuration")
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Fatal("environment can't be loaded", err)
	}

	log.Printf("%#v \n", &globalEnv)

	return globalEnv
}
