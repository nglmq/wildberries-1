package main

import "fmt"

func filter(temperatures []float64) map[int][]float64 {
	m := make(map[int][]float64)

	for _, temp := range temperatures {
		key := int(temp/10) * 10 // int(24.5/10) = 2, 2 * 10 = 20, получаем группу в которую надо поместить данные
		m[key] = append(m[key], temp)
	}
	return m
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := filter(temperatures)
	fmt.Println(m)
}
