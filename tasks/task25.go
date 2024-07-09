package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration) // After waits for the duration to elapse
	// and then sends the current time on the returned channel.
}

func main() {
	fmt.Println("hello ", time.Now())
	sleep(5 * time.Second)
	fmt.Println("world ", time.Now())
}
