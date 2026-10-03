package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
	}

	jobs := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup

	// Пул из 3 воркеров
	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for file := range jobs {
				f, err := os.Open(file)
				if err != nil {
					results <- file + ": ошибка " + err.Error()
					continue
				}

				h := md5.New()
				_, err = io.Copy(h, f)
				f.Close()

				if err != nil {
					results <- file + ": ошибка " + err.Error()
					continue
				}

				results <- fmt.Sprintf("%s: %x", file, h.Sum(nil))
			}
		}()
	}

	go func() {
		for _, file := range files {
			jobs <- file
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}