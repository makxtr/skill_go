package main

import (
	"context"

	"fmt"

	"sync"

	"time"
)

/**

* Improve the Do function, which should work within 10 ms.
* On the first error, Do should return an error and exit.
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
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	names := make(map[string]int64)
	mu := sync.Mutex{}

	errChan := make(chan error, 1)

	wg := sync.WaitGroup{}

	for _, u := range users {
		wg.Add(1)
		go func(user User) {
			defer wg.Done()

			name, err := fetch(ctx, user)
			if err != nil {
				select {
				case errChan <- err:
					cancel()
				default:
				}
				return
			}

			mu.Lock()
			names[name]++
			mu.Unlock()
		}(u)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case err := <-errChan:
		return nil, err
	case <-done:
		return names, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
