package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// 3 воркера
	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for n := range jobs {
				results <- n * n
			}
		}()
	}

	// 10 задач
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Закрыть results после завершения всех воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	// Собрать результаты
	for res := range results {
		fmt.Println(res)
	}
}