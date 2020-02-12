package services

import (
	"context"
	fibo_model2 "github.com/RezaZahedi/Go-Gin/model/proto/fibo_model"
	"github.com/micro/go-micro/v2"
	"log"
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

func NewFiboClient(serviceName string) func(number int) (string, error) {
	// Create a new service
	service := micro.NewService(micro.Name(serviceName))

	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := fibo_model2.NewGetFibonacciNumberService("greeter", service.Client())

	return func(number int) (s string, err error) {
		// Call the greeter
		rsp, err := greeter.GenerateNumber(context.TODO(), &fibo_model2.Request{Input: int32(number)})
		if err != nil {
			return "-1", err
		}
		log.Println("hi", rsp.Output)
		return rsp.Output, nil
	}
}