package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// имитация создания событий
	produce := func(worksChan chan int) {
		// если не закрыть канал, будет дедлок, т.к. консьюмеры будут бесконечно ожидать сообщения в канале
		defer close(worksChan)
		for i := 0; i < 20; i++ {
			worksChan <- i
		}
	}
	// количество воркеров
	workersCount := 5
	wg.Add(workersCount)
	consume := func(worksChan chan int, workersCount int) {
		// запускаем пять горутин (воркеров)
		// каждая горутина слушает общий канал и берет задачи из него
		for workerNum := 1; workerNum <= workersCount; workerNum++ {
			go func(workerNum int) {
				defer wg.Done()
				for val := range worksChan {
					fmt.Println(workerNum, val)
				}
			}(workerNum)
		}
	}
	worksChan := make(chan int)
	go produce(worksChan)
	go consume(worksChan, workersCount)
	wg.Wait()
}
