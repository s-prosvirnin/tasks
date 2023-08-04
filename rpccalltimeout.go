package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	rpcTimeout  = 700 * time.Millisecond
	rpcExecTime = 800 * time.Millisecond
)

func main() {
	fmt.Print(rpcCallWrapper(rpcTimeout))
}

// Безопасная обертка над RPC вызовом с таймаутом.
func rpcCallWrapper(timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	// Буфер обязательно нужен, чтобы горутина не зависла, если выпадем по таймауту (читать канал никто не будет).
	waiter := make(chan int, 1)
	go func() {
		defer close(waiter)
		waiter <- rpcCall()
	}()

	// Ждем выполнение RPC запроса или выходим по таймауту.
	select {
	case res := <-waiter:
		return res, nil
	case <-ctx.Done():
		return 0, errors.New("timeout error")
	}
}

// Долгий RPC вызов.
func rpcCall() int {
	rnd := rand.Intn(5000)
	time.Sleep(rpcExecTime)

	return rnd
}
