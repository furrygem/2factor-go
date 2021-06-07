package main

import (
	"flag"

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

}
