package main

import "fmt"

func main() {
	ch := make(chan int)
	data := []interface{}{"wb", 1, true, ch}

	for _, v := range data {
		fmt.Printf("%T\n", v)
	}
}
