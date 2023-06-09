package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// что выведет данный код
func main() {
	x := X{123}
	defer x.S()
	x.V = 456
}
