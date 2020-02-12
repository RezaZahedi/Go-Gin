package services

import (
	"context"
	fibo_model2 "github.com/RezaZahedi/Go-Gin/proto/fibo_model"
	"github.com/micro/go-micro/v2"
)
//
//func FiboClient(name string) (string, error) {
//	// Create a new service
//	service := micro.NewService(micro.Name("greeter.client"))
//	// Initialise the client and parse command line flags
//	service.Init()
//
//	// Create new greeter client
//	greeter := fibo_model2.NewGreeterService("greeter", service.Client())
//
//	// Call the greeter
//	rsp, err := greeter.Hello(context.TODO(), &fibo_model2.Request{Name: name})
//	if err != nil {
//		return "", err
//	}
//
//	return rsp.Greeting, nil
//}

func NewFiboClient(serviceName string) func(name string) (string, error) {
	// Create a new service
	service := micro.NewService(micro.Name(serviceName))

	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := fibo_model2.NewGreeterService("greeter", service.Client())

	return func(name string) (s string, err error) {
		// Call the greeter
		rsp, err := greeter.Hello(context.TODO(), &fibo_model2.Request{Name: name})
		if err != nil {
			return "", err
		}
		return rsp.Greeting, nil
	}
}