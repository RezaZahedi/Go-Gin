package main

import (
	"github.com/RezaZahedi/Go-Gin/REST_api"
	"github.com/RezaZahedi/Go-Gin/database"
	"log"
	"github.com/gin-gonic/gin"
)

// TODO: writing tests for product package

func main() {
	bookDB := database.NewBookDB()
	userDB := database.NewUserDB()
	if err := database.Init(userDB, bookDB); err != nil {
		log.Fatal("failed to initialize databases")
	}

	// TODO:
	productAPI := InitProductAPI(bookDB)
	userAPI := InitUserAPI(userDB)
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	REST_api.InitializeRoutes(router, &productAPI, &userAPI)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
