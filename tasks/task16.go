package main

import (
	"fmt"
	"slices"
	"sort"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		// recursively sort elements before partition and after partition
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1, 492, 9857, 554, 33, 491, 6, 5, 2, 2, 3, 492, 100, 1000, 100}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	sort.Slice(arr, func(i, j int) bool { return arr[j] > arr[i] }) // quick sort in sort package
	slices.Sort(arr)
	fmt.Println(arr)
}
