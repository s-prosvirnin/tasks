package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(predictableFunc())
}

func predictableFunc() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)

	defer cancel()

	waiter := make(chan int64)
	go func() {
		waiter <- unpredictableFunc()
	}()

	select {
	case res := <-waiter:
		return res
	case <-ctx.Done():
	}

	return 0
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}
