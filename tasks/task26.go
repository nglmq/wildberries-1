package main

import (
	"strings"
)

func checkString(s string) bool {
	s = strings.ToLower(s) // заглавные буквы превращаем в маленькие
	m := make(map[rune]bool)

	for _, v := range s { // проход по символам строки, если такой символ уже есть в мапе, то возвращаем false
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = true
	}

	return true
}

func main() {
	println(checkString("abcd"))
	println(checkString("abCdefAaf"))
	println(checkString("aabcd"))
}
