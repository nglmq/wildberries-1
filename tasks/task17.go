package main

import (
	"fmt"
	"slices"
	"sort"
)

func binarySearch(nums []int, target int) (int, bool) { // work only in sorted array
	index := -1
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2 // начинаем с середины слайса

		if nums[mid] <= target { // если элемент меньше или равен target,
			low = mid + 1 // первый элемент переставляется на вторую половину слайса
			index = mid
			if nums[index] == target { // если нашли элемент, возвращает его и true
				return mid, true
			}
		} else {
			high = mid - 1
		}
	}

	return index, false // если элемент не найден, возвращаем последний элемент слайса и false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(arr, 10))

	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= 30 }) //binary search in sort package
	fmt.Println(index)

	ind, found := slices.BinarySearch(arr, 30)
	fmt.Println(ind, found)
}
