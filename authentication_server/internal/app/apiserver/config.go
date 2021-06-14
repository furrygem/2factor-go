package apiserver

import (
	"github.com/furrygem/authentication_server/internal/app/store"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	StoreConfig *store.Config
}

/*
	Creating new apiserver config and populating it with defualt values
*/
func NewConfig() *Config {
	store_config := store.NewConfig() // generating new store config populated with default values and getting its address

	// returning default config with some varibles set to default values
	return &Config{
		BindAddr:    "127.0.0.1:8090",
		LogLevel:    "info",
		StoreConfig: store_config, // passing pointer to store config to server config
	}
}
