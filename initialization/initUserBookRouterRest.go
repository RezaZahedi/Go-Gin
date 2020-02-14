package initialization

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// InitialDummyFunc is used to break the deliberated loop in calling the
// Fibonacci service and REST API service. Will be swapped out with a
// Fibonacci Client after both services are started and running
var InitialDummyFunc = func(a int) (string, error) {
	return "", errors.New("not implemented yet, wait and try again")
}

// RouterREST initialized the router, loads the HTML files, and registers the
// routes on the router
func RouterREST() *gin.Engine {
	//bookDB := database.NewBookDB()
	//userDB := database.NewUserDB()
	//if err := database.Init(userDB, bookDB); err != nil {
	//	log.Fatal("failed to initialize databases")
	//}
	//
	//productAPI := initProductAPI(bookDB)
	//userAPI := initUserAPI(userDB)

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize User database, Book database, Fibonacci service, and then
	// register routes on the router.
	//REST_api.InitializeRoutes(router, productAPI, userAPI)
	initUserBookREST(router, &InitialDummyFunc)

	return router
}
