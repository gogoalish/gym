package config

import (
	"encoding/json"
	"os"
)

const CONFIG_FILE_NAME = "./config/config.json"

// write config
type Config struct {
	Server struct {
		Port           string `json:"port"`
		WriteTimeout   string `json:"writeTimeout"`
		ReadTimeout    string `json:"readTimeout"`
		MaxHeaderBytes string `json:"maxHeaderBytes"`
	} `json:"server"`
	Path struct {
		DB       string `json:"db"`
		Template string `json:"template"`
		Static   string `jsin:"static"`
	} `json:"path"`
}

// Global config.
var C Config

// Write config.
func ReadConfig() error {
	data, err := os.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &C)
}
