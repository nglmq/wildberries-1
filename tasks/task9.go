package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {
	c := gen(4, 5, 6, 7)
	out := square(c)

	for num := range out {
		fmt.Println(num)
	}
}
