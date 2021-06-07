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
	toml.DecodeFile(configPath, config)
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
