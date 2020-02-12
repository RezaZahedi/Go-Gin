package initialization

import (
	"github.com/gin-gonic/gin"
)
var InitialDummyFunc = func(a int) int { return a }

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


	// Initialize the routes
	//REST_api.InitializeRoutes(router, productAPI, userAPI)
	initUserBookREST(router, &InitialDummyFunc)
	return router
}