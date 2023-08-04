package main

import (
	"fmt"
	"sync"
)

func main() {
	const addCount = 1000
	const deleteCount = 300
	const getCount = 1500

	wg := &sync.WaitGroup{}
	m := newCmap()
	for i := 0; i < addCount; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			m.add(i, -i)
		}()
	}
	for i := 0; i < deleteCount; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			m.delete(i)
		}()
	}
	for i := 0; i < getCount; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			val, isExist := m.get(i)
			fmt.Printf("key: %d; val: %d; isExist: %v; len: %d\n", i, val, isExist, m.len())
		}()
	}
	wg.Wait()
}

type cmap struct {
	m    map[int]int
	lock sync.RWMutex
}

func newCmap() *cmap {
	return &cmap{
		m:    make(map[int]int),
		lock: sync.RWMutex{},
	}
}

func (cm *cmap) add(key int, val int) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.m[key] = val
}

func (cm *cmap) get(key int) (val int, isExist bool) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	val, isExist = cm.m[key]

	return val, isExist
}

func (cm *cmap) delete(key int) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.m, key)
}

func (cm *cmap) len() int {
	return len(cm.m)
}
