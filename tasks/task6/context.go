package main

import (
	"context"
	"fmt"
	"time"
)

func cancelContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is cancelled")
			return
		default:
			fmt.Println("running goroutine")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //context for goroutine

	go cancelContext(ctx) // start goroutine

	time.Sleep(2 * time.Second)
	cancel() //send cancel context
	time.Sleep(time.Second)
	fmt.Println("main goroutine exit")
}
