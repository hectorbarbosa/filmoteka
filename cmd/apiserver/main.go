package main

import (
	"log"

	"filmoteka/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

const (
	configPath = "config/apiserver.toml"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
