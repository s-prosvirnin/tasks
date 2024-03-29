package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// эту функцию можно пропустить, она просто пишет в каналы упорядоченные числа
func prepareChannels() (ch1, ch2 chan int) {
	const l = 10
	ch1 = make(chan int, l)
	ch2 = make(chan int, l)

	go func() {
		defer close(ch1)

		nums := make([]int, l)
		for i := 0; i < l; i++ {
			nums[i] = rand.Intn(l * 5)
		}
		sort.Ints(nums)
		for _, n := range nums {
			ch1 <- n
		}
	}()

	go func() {
		defer close(ch2)

		nums := make([]int, l)
		for i := 0; i < l; i++ {
			nums[i] = rand.Intn(l * 5)
		}
		sort.Ints(nums)
		for _, n := range nums {
			ch2 <- n
		}
	}()

	return ch1, ch2
}

func main() {
	ch1, ch2 := prepareChannels()

	ch3 := make(chan int)

	go func() {
		for n := range ch3 {
			fmt.Println(n)
		}
	}()

	res := mergeByLoop(ch1, ch2)
	fmt.Println(res)
}

func mergeByLoop(ch1, ch2 chan int) []int {
	n1, ok1 := <-ch1
	n2, ok2 := <-ch2
	var res []int
	for {
		if (ok2 && n2 <= n1) || (ok2 && !ok1) {
			res = append(res, n2)
			n2, ok2 = <-ch2
			continue
		}
		if (ok1 && n1 < n2) || (ok1 && !ok2) {
			res = append(res, n1)
			n1, ok1 = <-ch1
			continue
		}

		break
	}

	return res
}
