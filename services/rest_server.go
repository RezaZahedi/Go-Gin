package services

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"log"
)

// RestServer configures the REST API router and returns a service
// "misc ...interface{}" are used in an attempt to unify the servers' configuration functions signatures
func RestServer (afterStart func() error, misc ... interface{}) web.Service {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.Address("127.0.0.1:8080"),
	)

	// Init will parse the command line flags.
	service.Init(
		web.AfterStart(afterStart),
		)

	if len(misc) < 1 {
		log.Fatal("RestServer: not enough arguments")
	}
	router, ok := misc[0].(*gin.Engine)
	if !ok {
		log.Fatalf("RestServer: invalid argument, got %T, want *gin.Engine", misc[0])
	}

	// Register Handler
	service.Handle("/", router)

	return service
}
