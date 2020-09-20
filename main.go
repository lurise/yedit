package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var (
	service   string
	imagename string
	config    string
)

func init() {
	flag.StringVar(&service, "service", "", "the service to edit.")
	flag.StringVar(&imagename, "imagename", "", "the name of  image  to change")
	flag.StringVar(&config, "config", "./docker-compose.yml", "the config path to read")
	flag.Parse()
}

func main() {
	viper.SetConfigFile(config)
	err:=viper.ReadInConfig()
	if err != nil {
		log.Fatal("err to read the config.")
	}

	viper.Set("services."+service+".image",imagename)
	err=viper.WriteConfig()
	if err != nil {
		log.Fatal("error to write the config.")
	}
}
