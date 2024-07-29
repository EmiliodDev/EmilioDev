package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    Server struct{
        Port        int
    }
    DB struct {
        Driver      string
        Host        string
        Port        int
        User        string
        Password    string
        Name        string
    }
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("error reading config file, %v", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("unable to decode into a struct, %v", err)
    }
}
