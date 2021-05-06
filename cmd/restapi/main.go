package main

import (
	"flag"
	"log"

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
	if err != nil {
		log.Fatal(err)
	}
	if err := restapi.Start(config); err != nil {
		log.Fatal(err)
	}
}
