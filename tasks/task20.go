package main

import (
	"fmt"
	"strings"
)

func swapWords(s string) string {
	str := strings.Split(s, " ") // делаем слайс

	fmt.Println(len(str) - 1)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 { // меняем начало и конец местами
		str[i], str[j] = str[j], str[i]
	}

	return strings.Join(str, " ") // соединяем слайс в строку
}

func main() {
	fmt.Println(swapWords("hello world i am a string"))
	fmt.Println(swapWords("snow dog sun"))
}
