package initialization

import (
	"fmt"
	"github.com/RezaZahedi/Go-Gin/services"
	"sync"
)

type runner interface {
	Run() error
}

func InitializeServices(mWG *sync.WaitGroup) {
	// each service sends a signal on this channel when it is initialized
	initChan := make(chan struct{})
	callback := func() error {
		initChan <- struct{}{}
		return nil
	}

	fiboServer := services.FiboServer(callback)
	router := RouterREST()
	restServer := services.RestServer(callback, router)

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

	actualFunction := services.NewFiboClient("greeter.client")
	InitialDummyFunc = actualFunction

}
