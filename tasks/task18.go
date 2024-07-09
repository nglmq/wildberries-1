package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type MutexCounter struct {
	Value int
	mx    sync.Mutex // mutex for sync
}

func (c *MutexCounter) Increment() {
	c.mx.Lock()         // block access to val for other goroutines
	defer c.mx.Unlock() // unblock access
	c.Value++
}

func (c *MutexCounter) Get() int {
	c.mx.Lock()         // block access to val for other goroutines
	defer c.mx.Unlock() // unblock access
	return c.Value
}

type AtomicCounter struct {
	Value atomic.Int32 // atomic value
}

func (c *AtomicCounter) Increment() {
	c.Value.Add(1)
}

func (c *AtomicCounter) Get() int32 {
	return c.Value.Load()
}

func main() {
	var wg sync.WaitGroup
	counterMx := MutexCounter{}
	counterAtomic := AtomicCounter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				counterMx.Increment()
				counterAtomic.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Value of mutex counter: %d\n", counterMx.Get())
	fmt.Printf("Value of atomic counter: %d\n", counterAtomic.Get())
}
