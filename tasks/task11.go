package main

import "fmt"

func filterArrays(array1 []int, array2 []int) []int {
	m := make(map[int]int)
	var intersection []int // пересечение

	for _, v := range array1 {
		m[v] = v // добавление значений в мапу
	}

	for _, v := range array2 {
		if _, ok := m[v]; ok { // если значение есть в мапе, добавляем в пересечение
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func main() {
	arr1 := []int{2, 4, 3, 5, 6, 40, 2, 7, 20}
	arr2 := []int{3, 4, 6, 7, 8, 10, 40, 20}

	in := filterArrays(arr1, arr2)
	fmt.Println(in) // [3 4 6 7 40 20]
}
