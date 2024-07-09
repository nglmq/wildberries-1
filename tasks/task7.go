package main

import (
	"fmt"
	"sync"
)

type Map struct {
	rw sync.RWMutex
	m  map[string]int
}

func (m *Map) writeToMap(key string, value int) {
	m.rw.Lock()
	defer m.rw.Unlock()

	fmt.Println(key, value)

	m.m[key] = value
}

func main() {
	var wg sync.WaitGroup
	m := Map{m: make(map[string]int)}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			m.writeToMap(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	wg.Wait()
	fmt.Println(m.m)

	var syncMap sync.Map

	syncMap.Store("key1", 1)
	syncMap.Store("key2", 2)
	syncMap.Store("key3", 3)
	fmt.Println(syncMap.Load("key1"))
}
