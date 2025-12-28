package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://httpbin.org/status/404",
	}

	statusCodes := process(urls)

	fmt.Println(statusCodes)
}

var maxCon = 3

func worker(wg *sync.WaitGroup, client http.Client, jobs <-chan string, results chan<- int) {
	defer wg.Done()
	for u := range jobs {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			fmt.Printf("Ошибка создания запроса для %s: %s\n", u, err)
			cancel()
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Ошибка при выполнении запроса: %s\n", err)
			cancel()
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
		cancel()

	}
}

// реализовать параллельные запросы по адресам из списка
// подсчитать количество для каждого StatusCode ответа
// предусмотреть возможность отмены запроса по таймауту
func process(urls []string) map[int]int {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	jobs := make(chan string, len(urls))
	results := make(chan int, len(urls))

	wg := sync.WaitGroup{}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	for i := 0; i < maxCon; i++ {
		wg.Add(1)
		go worker(&wg, client, jobs, results)

	}

	wg.Wait()
	close(results)

	statusCounts := make(map[int]int)
	for statusCode := range results {
		statusCounts[statusCode]++
	}

	return statusCounts

}
