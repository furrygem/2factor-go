package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/furrygem/authentication_server/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	config.StoreConfig.DbAddr = "127.0.0.1"
	config.StoreConfig.DbPasswordFile = "internal/app/store/testing_password.txt"
	config.StoreConfig.DbPort = 15432
	config.StoreConfig.SSLMode = "disable"
	toml.DecodeFile(configPath, config)
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
