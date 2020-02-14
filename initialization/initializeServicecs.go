package initialization

import (
	"fmt"
	"github.com/RezaZahedi/Go-Gin/services"
	"sync"
)

// runner is used for looping through servers
type runner interface {
	Run() error
}

func InitializeServices(mWG *sync.WaitGroup) {
	// each service sends a signal on this channel when it is initialized
	// so we can start services sequentially one after another, and also
	// to start them in separate goroutines.
	// (when fiboServer is completely initialized and is running, it sends
	//  the signal so that restServer starts and registers himself as a client
	//  of the fiboServer)
	initChan := make(chan struct{})
	callback := func() error {
		initChan <- struct{}{}
		return nil
	}

	// get and run the servers
	fiboServer := services.FiboServer(callback)
	restServer := services.RestServer(callback, RouterREST())

	srvs := []runner{fiboServer, restServer}

	for _, srv := range srvs {
		mWG.Add(1)
		go func() {
			if err := srv.Run(); err != nil {
				fmt.Println(err)
			}
			mWG.Done()
		}()
		<-initChan
	}

	// get the RPC function and swap it out with the dummy function
	actualFunction := services.NewFiboClient("greeter.client")
	InitialDummyFunc = actualFunction
}
