package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	fibo_model "github.com/RezaZahedi/Go-Gin/fibonacci_service/proto"
	"github.com/RezaZahedi/Go-Gin/initialization"
	"github.com/micro/go-micro/v2"
)

// TODO: writing tests for product package

func main() {
	var mWG = new(sync.WaitGroup)

	mWG.Add(1)
	go initializeServices(mWG)
	//mWG.Add(1)

	// each service sends a signal on this channel when it is initialized
	initChan := make(chan struct{})
	services := make([]func(mWG *sync.WaitGroup, initChan chan<- struct{}), 0)
	services = append(services, test1, test2)
	for _, service := range services {
		mWG.Add(1)
		go service(mWG, initChan)
		<-initChan
	}
	//go test1(mWG)

	//mWG.Add(1)
	//go test2(mWG)

	mWG.Wait()
}

func initializeServices(mWG *sync.WaitGroup) {
	router := initialization.RouterREST()

	initializeWG := new(sync.WaitGroup)

	initializeWG.Add(1)
	go func() {
		if err := router.Run(); err != nil {
			log.Fatal(err)
		}
		initializeWG.Done()
	}()

	initializeWG.Wait()

	mWG.Done()
}

// TODO: for testing proto and gRPC

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *fibo_model.Request, rsp *fibo_model.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func test1(mWG *sync.WaitGroup, initChan chan<- struct{}) {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init(
		micro.AfterStart(func() error {
			initChan <- struct{}{}
			return nil
		}),
	)

	// Register handler
	fibo_model.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	mWG.Done()

}

func test2(mWG *sync.WaitGroup, initChan chan<- struct{}) {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init(micro.AfterStart(func() error {
		initChan <- struct{}{}
		return nil
	}),
	)

	// Create new greeter client
	greeter := fibo_model.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &fibo_model.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
	mWG.Done()
}
