package main

import (
	"fmt"
	"rates-listener/cmd"
	"rates-listener/internal/apiserver"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	server := apiserver.NewApiServer(cmd.BuildApiServerConfig())
	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
