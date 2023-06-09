package main

import (
	"fmt"
	"time"
)

const (
	timeout      = 700
	funcExecTime = 500
)

func main() {
	doneChan := make(chan bool)
	go func() {
		f()
		doneChan <- true
		close(doneChan)
	}()

	select {
	case <-doneChan:
		fmt.Println("func done")
	case <-time.After(time.Millisecond * timeout):
		fmt.Println("timeout")
	}
}

func f() {
	time.Sleep(funcExecTime * time.Millisecond)
}
