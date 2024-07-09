package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr)) // буферизованный канал с емкостью равной длине массива

	wg.Add(len(arr)) // количество горутин, которые будут выполняться

	for _, num := range arr {
		go func(num int) {
			defer wg.Done() // метка, что горутина выполнена

			square := num * num
			ch <- square
		}(num)
	}

	wg.Wait() // ожидание завершения всех горутин
	close(ch) // закрытие канала для записи

	sum := 0
	for sq := range ch { //чтение заначений из канала
		sum += sq
	}

	fmt.Println(sum)
}
