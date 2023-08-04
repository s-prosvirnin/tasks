package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	n := runtime.NumCPU()

	if n < 2 {
		log.Fatal("Atleast 2 CPUs needed")
	}

	// Expose CPUs to scheduler
	runtime.GOMAXPROCS(n)

	wg := &sync.WaitGroup{}

	// Use CPUs/2 to give the Go scheduler something to work with
	l := n / 2

	wg.Add(l)

	for i := 0; i < l; i++ {
		go func(wg *sync.WaitGroup, v int) {
			defer wg.Done()

			runtime.LockOSThread()

			defer runtime.UnlockOSThread()

			log.Printf("Hello from %d", v)

			time.Sleep(time.Second)
		}(wg, i)
	}

	wg.Wait()
}
