package main

import (
	"github.com/connormakh/pwHashApi/app/launcher"
	"os"
	"sync"
)

func main() {
	c := make(chan os.Signal, 1)
	wg := &sync.WaitGroup{}
	launcher.Initialize(c, wg)
}
