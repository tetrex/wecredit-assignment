package main

import (
	"log"

	"github.com/tetrex/wecredit-assignment/utils/config"
	"github.com/tetrex/wecredit-assignment/utils/logger"
)

// @title			server api
// @version			1.0
// @description		This is a backend api server
// @contact.name	github.com/tetrex
// @license.name	MIT License
//
// @host			localhost:8000
// @basePath		/
func main() {
	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config ")
		log.Fatal(err.Error())
	}

	// logger
	l := logger.New(config.AppEnv)

}
