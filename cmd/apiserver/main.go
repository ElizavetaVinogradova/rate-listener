package main

import (
	"fmt"
	"rates-listener/cmd"
	"rates-listener/internal/apiserver"
	"rates-listener/internal/repo/mysql"
	"rates-listener/internal/service"

	"errors"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()
	
	repository, err := mysql.NewTickRepository(cmd.BuildMySqlConfig())
	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}
	defer repository.Close()

	tickService := service.NewTickService(repository)
	
	server := apiserver.NewApiServer(cmd.BuildApiServerConfig(), *tickService)
	err = server.Start()

	if errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed")
	} else if err != nil {
		log.Errorf("error starting server: %s", err)
		os.Exit(1)
	}
}
