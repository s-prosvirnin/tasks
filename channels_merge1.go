package main

import (
	"fmt"
	"sync"
)

/*
Даны два канала, в которые поступают числа в возрастающем порядке. Нужно реализовать функцию, которая принимает эти
два канала и возвращает массив, в котором числа из обоих каналов записаны в возрастающем порядке (отсортированы).
*/
func sortChan(chans ...chan int) []int {

	// запустим прослушивание в канал в отдельных горутинах
	// и обработаем закрытие прослушивания, чтобы не заблокироваться
	// ивенты из каналов будем писать в канал general
	// когда все события из исходных каналов будут вычитаны, закроем general и сделаем wg.Done в дефере каждого писателя

	// general будет прослушиваться отдельной горутиной до момента закрытия канала
	// результат из general будем обрабатывать алгоритмом мержа двух массивов - только с 3ей переменной
	// можно сделать буфер в general на капасити исходнй массивов, чтобы не блокировать писателей в этот канал
	// в горутину где будем читать general тоже сделаем wg.Done() когда канал general закроется

	var result []int
	// общий канал
	general := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		ch := ch
		go func() {
			// читаем из ch и пишем в канал general
			defer wg.Done()
			for event := range ch {
				general <- event
			}
		}()
	}

	go func() {
		// когда из исходных массивов все вычитано в канал general закрываем его
		// general должен закрыться, когда завершаться писатели
		wg.Wait()
		close(general)
	}()

	var temp int
	// читаем из канала general и сортируем
	// выходим, когда закрывается general
	for num := range general {
		if num < temp {
			result = append(result, num)
		} else {
			result = append(result, temp)
			temp = num
		}
	}

	return result
}

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			a <- i
		}
		close(a)
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			b <- i
		}
		close(b)
	}()
	fmt.Println(sortChan(a, b))
}
