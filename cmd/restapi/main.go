package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/UniverOOP/internal/app/restapi"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/restapi/config.toml", "Path to config file")
}

func main() {
	flag.Parse()
	config := restapi.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	port := os.Getenv("PORT")
	config.BindAddress = ":" + port

	log.Print(config.BindAddress)

	if err != nil {
		log.Fatal(err)
	}
	if err := restapi.Start(config); err != nil {
		log.Fatal(err)
	}
}
