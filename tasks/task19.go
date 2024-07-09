package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s) // конвертируем в руны чтобы корректно итерироваться по юникод символам

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // меняем элементы местами, i идет с начала, j идет с конца
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(reverse("привет"))
	fmt.Println(reverse("hello"))
	fmt.Println(reverse("as⃝df̅"))
	fmt.Println(reverse("␠␡␢␣␤␥␦t"))
}
