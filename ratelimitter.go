package main

import (
	"fmt"
	"time"
)

const rateLimit = time.Second

// Интерфейс для внешнего вызова.
type callable interface {
	call(payload)
}

// Сообщение, передаваемое во внешний вызов
type payload struct{ val int }

// Rate limitter. Ограничиваем внешний вызов запуском в n-ое время (rateLimit)
func rateLimitCall(client callable, payloads []payload, rateLimit time.Duration) {
	ticker := time.Tick(rateLimit)
	for _, p := range payloads {
		<-ticker          // ждем временной период
		go client.call(p) // внешний вызов в горутине
	}
}

func main() {
	client := clientMock{executionTime: time.Second * 2}
	payloads := make([]payload, 10)
	for i := 0; i < 10; i++ {
		payloads[i] = payload{i}
	}

	rateLimitCall(client, payloads, rateLimit)
}

// Мок для внешнего вызова.
type clientMock struct {
	// задержка выполнения
	executionTime time.Duration
}

func (c clientMock) call(p payload) {
	time.Sleep(c.executionTime)
	fmt.Println(p.val)
}
