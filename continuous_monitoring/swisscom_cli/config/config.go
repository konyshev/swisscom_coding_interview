package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents url to REST API end point
type Config struct {
	Url string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
