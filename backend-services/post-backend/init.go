package main

import (
	"post-backend/app/log"
	"post-backend/app/server"
	"post-backend/src/handler"
	"post-backend/src/handler/router"
	"post-backend/src/mapper"
	"post-backend/src/repository"
	"post-backend/src/usecase"

	"fmt"
	"strings"
)

func RunApplication() {
	fmt.Println("Init Configuration")
	svcProperties := getServiceProperties()
	fmt.Printf("%s SERVICE\n", strings.ToUpper(svcProperties.ServiceName))

	log.SetupLogging()

	db := databaseConnect(svcProperties.ServiceName, svcProperties.Database)

	iSetupServer := server.NewSetupServer(svcProperties)
	fiberRouter := iSetupServer.InitServerConfiguration()

	// POST

	iPostMapper := mapper.NewPostMapper()
	iPostPostRepo := repository.NewPostRepo(db)
	iPostUsecase := usecase.NewPostUsecase(svcProperties, iPostMapper, iPostPostRepo)
	iPostRouter := router.NewPostRouter(iPostUsecase)
	iPostHttpHandler := handler.NewPostHttpHandler(fiberRouter, iPostRouter)
	iPostHttpHandler.Routers()

	// setup server
	iSetupServer.InitServer()
}
