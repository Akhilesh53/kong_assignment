package main

import (
	"fmt"
	"ka/pkg/logging"
	"ka/src/controllers"
	"ka/src/middlewares"

	repository "ka/src/daos"
	"ka/src/services"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	serviceRepo := repository.NewServiceRepository()
	serviceService := services.NewServiceService(serviceRepo)
	serviceHandler := controllers.NewServiceHandler(serviceService)

	logging.InitialiseLogger()
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true

	router.Use(middlewares.GenerateUUID())

	router.GET("/services", middlewares.AuthMiddleware(), serviceHandler.GetServices)
	router.GET("/services/:id", middlewares.AuthMiddleware(), serviceHandler.GetService)
	router.GET("/services/:id/versions", middlewares.AuthMiddleware(), serviceHandler.GetServiceVersions)

	fmt.Println("Server started successfully at port :: 8080")
	if err := router.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
