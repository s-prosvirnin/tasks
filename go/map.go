package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int, 2)
	l := &sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		// Если уменьшить количество итераций до 10, то вероятность ошибки будет намного меньше.
		for i := 0; i < 10000; i++ {
			l.Lock()
			m[2]++
			l.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			l.Lock()
			m[2]++
			l.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			_ = m[2]
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(m[2])
}
