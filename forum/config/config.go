package config

import (
	"encoding/json"
	"os"
)

type path struct {
	Template string
}

type config struct {
	P path
}

var C config

func ReadConfig() error {
	file, err := os.ReadFile("/config/config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &C)
}
