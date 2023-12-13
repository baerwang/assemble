package main

import (
	"assemble/config"
	"assemble/pkg/router"
)

func main() {
	config.ConfigInit()
	// config.LoadDb(config.GetConfig().Database)
	if err := router.InitRouter().Run(":8000"); err != nil {
		panic(err)
	}
}
