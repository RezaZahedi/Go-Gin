package main

import (
	"github.com/RezaZahedi/Go-Gin/initialization"
	"sync"
)

// TODO: writing tests for product package

func main() {
	var mWG = new(sync.WaitGroup)

	initialization.InitializeServices(mWG)

	mWG.Wait()
}
