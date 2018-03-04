package main

import (
	"log"
	"os"
	"strconv"
)

type Settings struct {
	Port    int
	Debug   bool
	OnlyAPI bool
}

func GetSettings() Settings {
	var port int64
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		var err error
		port, err = strconv.ParseInt(portEnv, 10, 64)
		if err != nil {
			log.Fatal("Wrong env variable PORT!")
		}
	} else {
		port = 8080 // default
	}

	var debug bool
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" {
		if debugEnv != "false" {
			debug = true
		} else {
			debug = false
		}
	} else {
		debug = false // default
	}

	var api bool
	apiEnv := os.Getenv("ONLY_API")
	if apiEnv != "" {
		if apiEnv != "false" {
			api = true
		} else {
			api = false
		}
	} else {
		api = false // default
	}

	settings := Settings{
		Port:    int(port),
		Debug:   debug,
		OnlyAPI: api,
	}

	return settings
}
