package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := [5]int{2, 4, 6, 8, 10}

	wg.Add(len(arr)) // количество горутин, которые будут выполняться

	for _, num := range arr {
		go func(num int) {
			defer wg.Done() // метка, что горутина выполнена

			fmt.Println(num * num)
		}(num)
	}

	wg.Wait() // ожидание завершения всех горутин
}
