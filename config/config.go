package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// Define your configuration fields here
}

var AppConfig Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic("Unable to read config file")
	}

	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		panic("Unable to parse config file")
	}
}
