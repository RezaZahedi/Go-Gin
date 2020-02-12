package services

import (
	"github.com/RezaZahedi/Go-Gin/fibonacci"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"github.com/micro/go-micro/v2"
)

func FiboServer(afterStart func() error, misc ... interface{}) micro.Service {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init(
		micro.AfterStart(afterStart),
	)

	// Register handler
	fibo_model.RegisterGetFibonacciNumberHandler(service.Server(), new(fibonacci.FiboGenerator))

	return service
}