package main

/*
The Product Owner delivered the following story:

As a user, I can see an overview of services in my organization. Acceptance criteria include:

User can see the name, a brief description, and versions available for a given service
User can navigate to a given service from its card
User can search for a specific service
User can paginate and sort through services
The Assignment

You're responsible for the data model and API portions of this story.
Implement a Services API that can be used to implement this dashboard widget.
It should support
Returning a list of services
support filtering, sorting, pagination
Fetching a particular service
including a method for retrieving its versions

The API can be read-only. Choose a persistence mechanism that is appropriate for this feature.

This project must be written in Go.

We'll evaluate the design, implementation choices, and functionality of your project. The code should be as production ready as you can make it, even if that means reducing the total features you are able to add].

*/

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
