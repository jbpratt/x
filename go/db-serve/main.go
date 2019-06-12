package main

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
)

func main() {
	cfg := LoadConfig()
	config := dropbox.Config{
		Token: cfg.Dropbox.AuthToken,
	}

	dbx := users.New(config)
	_ = dbx
}
