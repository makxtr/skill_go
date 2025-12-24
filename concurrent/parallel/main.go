package main

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

// // Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// // в stdout в отдельной строке url: , code:
// // В случае ошибки выводит в отдельной строке url: , code:
// // Функция должна завершаться при отмене контекста.
func fetch(ctx context.Context, url string) (string, error) {
	select {
	case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
		n := rand.Intn(10)
		if n < 5 {
			return "", errors.New("500 Internal Server Error")
		}
		return "200 OK", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchParallel(ctx context.Context, urls []string) {
	
}

func main() {
	urls := []string{"url1", "url2", "url3", "url4", "url5", "url6", "url7", "url8", "url9"}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	fetchParallel(ctx, urls)
}
