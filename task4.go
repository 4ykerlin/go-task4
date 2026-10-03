package main

func main() {
	urls := []string{
		"https://example.com",
		"https://go.dev",
		"https://golang.org",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
	}

	jobs := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				resp, err := http.Get(url)
				if err != nil {
					results <- url + " ошибка: " + err.Error()
					continue
				}
				results <- url + " -> " + resp.Status
				resp.Body.Close()
			}
		}()
	}

	go func() {
		for _, url := range urls {
			jobs <- url
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
