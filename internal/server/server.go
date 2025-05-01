package server

import (
	"log"

	"github.com/NonthapatKim/many-tooth-api/infrastructure"
	"github.com/NonthapatKim/many-tooth-api/internal/adapter/handler"
	"github.com/NonthapatKim/many-tooth-api/internal/adapter/repository"
	"github.com/NonthapatKim/many-tooth-api/internal/core/service"
	"github.com/NonthapatKim/many-tooth-api/internal/router"
)

func RunServer() error {
	mysql, err := infrastructure.NewMySQL()
	if err != nil {
		log.Fatal("Error initializing MySQL database", err)
	}
	// init repository
	repo := repository.New(mysql)

	// init service
	svc := service.New(repo)

	// init handler
	hdl := handler.New(svc)

	router, err := router.NewRouter(hdl)
	if err != nil {
		log.Fatal("Error initializing router:", err)
	}

	return router.Start()
}
