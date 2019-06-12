package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Dropbox OAuthConfig `json:"dropbox"`
}

type OAuthConfig struct {
	AuthToken string `json:"auth_token"`
}

func LoadConfig() Config {
	f, err := os.Open(".config")
	if err != nil {
		panic(err)
	}

	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}
