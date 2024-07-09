package main

import (
	"fmt"
)

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)

	//copy(slice[index:], slice[index+1:])
	//slice[len(slice)-1] = 0 // or the zero value
	//slice = slice[:len(slice)-1]

	//return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5} // [1 2 3 4 5]

	slice = deleteElement(slice, 2)

	fmt.Println(slice) // [1 2 4 5]
}
