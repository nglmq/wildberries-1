package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func readFromChannel(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // check if context is done
			fmt.Println(ctx.Err())
			return
		case data := <-ch:
			fmt.Printf("got %d\n", data)
		}
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second)) // set deadline
	defer cancel()                                                                           // cancel context

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go readFromChannel(ctx, &wg, ch)

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

	wg.Wait()
	fmt.Println("Done")
}
