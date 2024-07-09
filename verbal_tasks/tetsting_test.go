//package main
//
//import (
//	"fmt"
//	"log"
//)
//
//func main() {
//
//	/*
//	   Для чтения входных данных необходимо получить их
//	   из стандартного потока ввода.
//	   Данные во входном потоке соответствуют описанному
//	   в условии формату. Обычно входные данные состоят
//	   из нескольких строк.
//	   Можно использовать несколько методов:
//	   * fmt.Scanln(), fmt.Scan(), fmt.Scanf() -- функции, позволяющие считывать отдельные слова в отдельные переменные;
//	   * bufio.Reader -- позволяет читать строку входного текста или один символ;
//	   * bufio.Scanner -- предоставляет удобный интерфейс для чтения строк текста.
//	   Чтобы прочитать из строки стандартного потока:
//	   * число --
//	   var number int64
//	   _, err := fmt.Scan(&number)
//	   if err != nil {
//	       log.Fatal(err)
//	   }
//	   * строку --
//	   scanner := bufio.NewScanner(os.Stdin)
//	   scanner.Scan()
//	   err := scanner.Err()
//	   if err != nil {
//	       log.Fatal(err)
//	   }
//	   * массив чисел известной длины --
//	   var arr = make([]int, 10)
//	   for i := 0; i < len(arr); i++ {
//	       _, err := fmt.Scan(&arr[i])
//	       if err != nil {
//	           log.Fatal(err)
//	       }
//	   }
//	   * последовательность слов до конца файла --
//	   scanner := bufio.NewScanner(os.Stdin)
//	   scanner.Split(bufio.ScanWords)
//	   var words []string
//	   for scanner.Scan() {
//	       words = append(words, scanner.Text())
//	   }
//	   Чтобы вывести результат в стандартный поток вывода,
//	   можно использовать fmt.Println() или fmt.Print().
//	   Возможное решение задачи "Вычислите сумму A+B":
//	   var a, b int64
//	   _, err := fmt.Scan(&a)
//	   if err != nil {
//	       log.Fatal(err)
//	   }
//	   _, err = fmt.Scan(&b)
//	   if err != nil {
//	       log.Fatal(err)
//	   }
//	   fmt.Println(a + b)
//	*/
//
//	var arr = make([]int, 3)
//	for i := 0; i < len(arr); i++ {
//		_, err := fmt.Scan(&arr[i])
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	sortInts(arr)
//
//	fmt.Println(arr[1])
//}
//
//func sortInts(arr []int) []int {
//	for i := 0; i < len(arr)-1; i++ {
//		for j := 0; j < len(arr)-i-1; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//
//	return arr
//}

//package main
//
//import "fmt"
//
//func fifoFor(arr []int, newElement int) ([]int, int) {
//	arr = append(arr, newElement) // Добавление в конец
//	firstElement := arr[0]
//	arr = arr[1:] // Удаление из начала
//
//	return arr, firstElement
//}
//
//func main() {
//	arr := []int{1, 2, 3}
//
//	arr, first := fifoFor(arr, 4)
//
//	fmt.Println(arr, first)
//
//	{
//		arr1 := []int{1, 2, 3}
//		fmt.Println(arr1)
//	}
//
//	//fmt.Println(arr1)
//}

//package main
//
//import "fmt"
//
//func F(n int) func() int {
//	return func() int {
//		n++
//		return n
//	}
//}
//
//func main() {
//	f := F(5)
//
//	defer func() {
//		fmt.Println("defer1", f())
//	}()
//	defer fmt.Println("defer2", f())
//	f()
//}

package main

import (
	"runtime"
	"sync"
	"testing"
)

type Counter struct {
	count int
}

func (c Counter) Increment() {
	for i := 0; i < 100_000; i++ {
		c.count++
	}
}

func counter() {
	c := 0

	for i := 0; i < 100_000; i++ {
		c++
	}
}

func BenchmarkAnonymousFunction(b *testing.B) { // 23962 ns/op
	for n := 0; n < b.N; n++ {
		c := 0
		func(c int) {
			for i := 0; i < 100_000; i++ {
				c++
			}
		}(c)
	}
}

func BenchmarkMethodCall(b *testing.B) { // 23480 ns/op
	counter := Counter{}
	for i := 0; i < b.N; i++ {
		counter.Increment()
	}
}

func BenchmarkFunctionCall(b *testing.B) { // 23326 ns/op
	for n := 0; n < b.N; n++ {
		counter()
	}
}

func BenchmarkGoRoutineAnonymousFunction(b *testing.B) { //63396 ns/op
	for n := 0; n < b.N; n++ {
		wg := sync.WaitGroup{}
		wg.Add(runtime.NumCPU())

		for i := 0; i < runtime.NumCPU(); i++ {
			go func() {
				defer wg.Done()
				counter()
			}()
		}

		wg.Wait()
	}
}
