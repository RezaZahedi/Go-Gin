package main

import (
	"github.com/RezaZahedi/Go-Gin/database"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	bookDB := database.NewBookDB()
	userDB := database.NewUserDB()
	if err := database.Init(userDB, bookDB); err != nil {
		log.Fatal("failed to initialize databases")
	}

	// TODO:
	productAPI := InitProductAPI(bookDB)

	r := gin.Default()
 	r.GET(/books, productAPI.FindAll)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
