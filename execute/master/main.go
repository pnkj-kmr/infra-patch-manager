package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/restapi"
	"github.com/pnkj-kmr/infra-patch-manager/utility"

	_ "github.com/pnkj-kmr/infra-patch-manager/docs"
)

// @title Patch API
// @version 1.0
// @description Patch APIs helps to describe the available routes for patch master server.
// @contact.name PANKAJ KUMAR
// @license.name MIT Licence
// @license.url https://www.github.com/pnkj-kmr/infra-patch-manager/README.md
func main() {
	address := flag.String("address", "", "the server port")
	flag.Parse()

	config, err := utility.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load configuration file", err)
	}

	fiberConfig := restapi.FiberConfig(config)

	// Define a new Fiber app with config.
	app := fiber.New(fiberConfig)

	// Middlewares - TODOs
	// restapi.FiberMiddleware(app) // Register Fiber's default middleware

	// Routes.
	restapi.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	restapi.PublicRoutes(app)  // Register a public routes for app.
	restapi.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	var serverAddress string
	if len(*address) > 0 {
		serverAddress = *address
	} else {
		serverAddress = fmt.Sprintf("0.0.0.0:%s", config.Port)
	}
	// server.StartServer(app, serverAddress)
	restapi.StartServerWithGracefulShutdown(app, serverAddress)
}
