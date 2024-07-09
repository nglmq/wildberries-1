package main

import "fmt"

func changeBit(num, pos, bit int64) int64 {
	if bit == 0 { // Установить бит в 0
		return num &^ (1 << pos)
	}

	return num | (1 << pos) // Установить бит в 1
}

func main() {
	var num int64 = 4 //0100

	fmt.Println(changeBit(num, 0, 1)) // 0101 = 5
	fmt.Println(changeBit(num, 0, 0)) // 0100 = 4
	fmt.Println(changeBit(num, 1, 1)) // 0110 = 6
	fmt.Println(changeBit(num, 1, 0)) // 0100 = 4
	fmt.Println(changeBit(num, 2, 1)) // 0100 = 4
	fmt.Println(changeBit(num, 2, 0)) // 0000 = 0
}
