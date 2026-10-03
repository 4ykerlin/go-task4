package main

import (
	"fmt"
	"time"
)

func main() {
	backends := []string{"backend1", "backend2", "backend3"}

	backendChans := make([]chan string, len(backends))

	// Запускаем каждый backend
	for i := range backends {
		backendChans[i] = make(chan string)

		go func(name string, ch chan string) {
			for req := range ch {
				fmt.Println(name, "обработал", req)
			}
		}(backends[i], backendChans[i])
	}

	requests := make(chan string)

	// Round-robin распределение
	go func() {
		i := 0

		for req := range requests {
			backendChans[i] <- req
			i = (i + 1) % len(backendChans)
		}
	}()

	// Отправляем 10 запросов
	for j := 1; j <= 10; j++ {
		requests <- fmt.Sprintf("запрос-%d", j)
	}
	close(requests)

	time.Sleep(500 * time.Millisecond)
}