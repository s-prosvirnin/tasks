package main

import "fmt"

func main() {
	writeToChanInSelect()
}

func writeToChanInSelect() {
	c := make(chan int, 5)
	for done := false; !done; {
		select {
		case c <- 3:
			fmt.Print(3)
		case <-c:
			fmt.Print(2)
			close(c)
		default:
			fmt.Print(1)
			done = true
		}
	}
}
