package main

import "fmt"

/*
var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}

func main() {
    someFunc()
}

justString занимает один и тот же блок памяти вместе с v.
Сборщик мусора не сможет удалить v, память будет использоваться неэффективно.

Использование глобальной переменной не всегда оправдано, тк к ним можно обращаться из любой функции.
*/

func someFunc() string {
	justString := createHugeString(1 << 10)[:100]

	return justString
}

func createHugeString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

func main() {
	result := someFunc()
	fmt.Println(result)
}

//Также можно создать копию строки
/*
var justString string

func createHugeString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

func someFunc() {
	v := createHugeString(1 << 10)
	println(v)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
*/
