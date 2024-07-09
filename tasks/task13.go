package main

import "fmt"

func main() {
	a, b := 3, 4
	fmt.Println("before: ", a, " ", b)

	a, b = b, a //swap
	fmt.Println("after: ", a, " ", b)
}
