package main

import (
	"fmt"
	"time"
)

func helloWorld() {
	fmt.Println("Hello, World!")
}

func throttle(f func(), t time.Duration) func() {
	// ВАЖНО: в идеале еще использовать context, так же канал не закрывается, будет утечка.
	// Канал для блокировки выполнения. Блокируемся на записи в канал.
	blockChan := make(chan struct{}, 1)
	return func() {
		// Асинхронно пытаемся выполнить f().
		go func() {
			select {
			// Пытаемся записать в канал, если буфер полон (кто-то уже записал и выполняет f()), то блокируемся тут и ждем выполнения f() и таймаут.
			case blockChan <- struct{}{}:
				f()
				time.Sleep(t)
				// Читаем то, что записали, чтобы разблокировать остальных ждунов.
				<-blockChan
			}
			// Нельзя тут делать default, т.к. тогда мы не будем блокироваться и ждать, а сразу выйдем из функции (получим только одно выполнение f()).
		}()
	}
}

func main() {
	throttledFunc := throttle(helloWorld, time.Second)
	// Функции throttledFunc вызываются асинхронно.
	throttledFunc()
	throttledFunc()
	throttledFunc()
	time.Sleep(time.Second * 6)
}
