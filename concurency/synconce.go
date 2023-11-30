package main

import (
	"sync/atomic"
	"time"
)

type Once struct {
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		f()
	}
}

func main() {
	o := Once{}
	f := func() {
		println("lala")
	}
	for i := 0; i < 1000; i++ {
		go o.Do(f)
	}
	time.Sleep(time.Second)
}
