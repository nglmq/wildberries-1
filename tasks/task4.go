package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		case val := <-ch:
			fmt.Printf("worker %d got %d\n", id, val)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var workers int

	// read number of workers
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scanf("%d\n", &workers)
	if err != nil {
		fmt.Printf("error while reading number of workers: %v", err)
		return
	}

	ch := make(chan int)
	interrupt := make(chan os.Signal) // channel to receive interrupt signal
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background()) //context for workers

	for i := 1; i <= workers; i++ { // create workers
		wg.Add(1)
		go worker(ctx, i, &wg, ch)
	}

	go func() {
		gen := rand.New(rand.NewSource(time.Now().UnixNano())) //random data
		for {
			select {
			case <-ctx.Done(): // check if context cancelled, if so, close channel and return
				close(ch)
				return
			default: // write to channel
				ch <- gen.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-interrupt // wait for interrupt signal
	fmt.Println("Interrupt signal received")
	cancel()  //cancel context for workers
	wg.Wait() // wait all workers to stop

	fmt.Println("All workers stopped")
}
