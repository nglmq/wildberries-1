package main

import (
	"fmt"
	"time"
)

func signalChannel(ch chan bool) {
	for {
		select {
		case <-ch:
			if true {
				fmt.Println("receive signal, stopping goroutine")
				return
			}
		default:
			fmt.Println("running goroutine")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan bool)

	go signalChannel(ch)

	time.Sleep(2 * time.Second)
	ch <- true
	time.Sleep(time.Second)
	fmt.Println("main goroutine exit")
}
