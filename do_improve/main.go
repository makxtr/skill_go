package main

import (
	"context"

	"fmt"

	"sync"

	"time"
)

/**

* Доработать функцию Do, что бы она отрабатывала приблизительно за 10мс

* при первой ошибке Do должна возвращать ошибку и завершаться

 */

type User struct {
	Name string
}

func main() {

	curTime := time.Now()

	m, err := Do(context.Background(), []User{{"Paul"}, {"Jack"}, {"Mike"}, {"Jack"}})

	fmt.Println("time:", time.Since(curTime))

	fmt.Println(m, err)

}

func fetch(ctx context.Context, u User) (string, error) {
	timer := time.NewTimer(time.Millisecond * 10)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-timer.C:
		return u.Name, nil
	}
}

func Do(ctx context.Context, users []User) (map[string]int64, error) {
	// Создаем дочерний контекст, который мы сможем отменить вручную
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // Гарантируем очистку ресурсов

	names := make(map[string]int64)
	mu := sync.Mutex{}

	// Канал для сбора первой ошибки
	errChan := make(chan error, 1)

	wg := sync.WaitGroup{}

	for _, u := range users {
		wg.Add(1)
		// Передаем u в горутину явно, чтобы избежать проблем с захватом переменной
		go func(user User) {
			defer wg.Done()

			// Передаем ctx в fetch. Если кто-то вызовет cancel(), fetch должен это обработать
			name, err := fetch(ctx, user)
			if err != nil {
				// Пытаемся отправить ошибку. Если канал полон — значит ошибка уже есть.
				select {
				case errChan <- err:
					cancel() // Отменяем остальные операции!
				default:
				}
				return
			}

			mu.Lock()
			names[name]++
			mu.Unlock()
		}(u)
	}

	// Создаем канал для сигнала о завершении всех горутин
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Ждем либо ошибки, либо успешного завершения всех, либо отмены внешнего контекста
	select {
	case err := <-errChan:
		return nil, err
	case <-done:
		return names, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
