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


	fiboClient := services.NewFiboClient("greeter.client")
	answer, err := fiboClient("Reza")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
	answer, err = fiboClient("Reza2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}
