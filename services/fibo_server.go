package services

import (
	"context"
	"github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *fibo_model.Request, rsp *fibo_model.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

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
	fibo_model.RegisterGreeterHandler(service.Server(), new(Greeter))

	return service
}