package base

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	ElasticHost  string `mapstructure:"ELASTIC_HOST"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
	PostgresDB   string `mapstructure:"POSTGRES_DB"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
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
