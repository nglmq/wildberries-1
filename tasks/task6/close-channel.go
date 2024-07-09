package main

import (
	"fmt"
	"time"
)

func closeChannel(ch chan struct{}) {
	for {
		select {
		case _, open := <-ch:
			if !open {
				fmt.Println("channel closed, stop goroutine")
				return
			}
		default:
			fmt.Println("running goroutine")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{})
	go closeChannel(ch)

	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("main goroutine exit")
}
