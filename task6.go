package main

import (
	"errors"
	"fmt"
	"time"
)

func search(source string, result chan<- string, errs chan<- error) {
	// Имитация разного времени поиска
	time.Sleep(time.Duration(len(source)*100) * time.Millisecond)

	// Например, source2 всегда не находит
	if source == "source2" {
		errs <- errors.New(source + " не найден")
		return
	}

	result <- "Результат из " + source
}

func main() {
	sources := []string{"source1", "source2", "source3"}

	result := make(chan string)
	errs := make(chan error)

	for _, s := range sources {
		go search(s, result, errs)
	}

	for i := 0; i < len(sources); i++ {
		select {
		case r := <-result:
			fmt.Println("Первый успех:", r)
			return

		case e := <-errs:
			fmt.Println("Ошибка:", e)
		}
	}

	fmt.Println("Нет успешных результатов")
}