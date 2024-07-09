package main

import (
	"fmt"
	"sync"
	"time"
)

func variableChannel(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stop {
			fmt.Println("receive signal, stopping goroutine")
			return
		}
		fmt.Println("running goroutine")
		time.Sleep(time.Second)
	}
}

func main() {
	var stop bool
	var wg sync.WaitGroup
	wg.Add(1)

	go variableChannel(&stop, &wg)

	time.Sleep(2 * time.Second)
	stop = true
	wg.Wait()
	fmt.Println("main goroutine finished")
}
