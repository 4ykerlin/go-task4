package main

func search(source string, result chan<- string, errs chan<- error) {
	time.Sleep(time.Duration(len(source)*100) * time.Millisecond)

	if source == "source2" {
		errs <- errors.New(source + " не найден")
		return
	}
	result <- "Результат из " + source
}

func main() {
	sources := []string{"source1", "source2", "source3"}

	result := make(chan string, len(sources))
	errs := make(chan error, len(sources))

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
