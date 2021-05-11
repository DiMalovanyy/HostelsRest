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
	if port != "" { // for heroku
		config.BindAddress = ":" + port
		config.DatabaseURL = "postgres://pzsumgcucebrbb:675b162b381c16af13becd943e26f926a74fd5e095204d2031031c1f717d40d3@ec2-54-90-211-192.compute-1.amazonaws.com:5432/dbfsp44r075phg"
	}

	if err != nil {
		log.Fatal(err)
	}
	if err := restapi.Start(config); err != nil {
		log.Fatal(err)
	}
}
