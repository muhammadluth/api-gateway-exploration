package main

import (
	"comment-backend/app/log"
	"comment-backend/app/server"
	"comment-backend/src/handler"
	"comment-backend/src/handler/router"
	"comment-backend/src/mapper"
	"comment-backend/src/repository"
	"comment-backend/src/usecase"

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

	// COMMENT
	iCommentMapper := mapper.NewCommentMapper()
	iCommentCommentRepo := repository.NewCommentRepo(db)
	iCommentUsecase := usecase.NewCommentUsecase(svcProperties, iCommentMapper, iCommentCommentRepo)
	iCommentRouter := router.NewCommentRouter(iCommentUsecase)
	iCommentHttpHandler := handler.NewCommentHttpHandler(fiberRouter, iCommentRouter)
	iCommentHttpHandler.Routers()

	// setup server
	iSetupServer.InitServer()
}
