package main

import (
	"context"
	//"errors"
	"fmt"
	"sync"
	"time"
)
import "golang.org/x/sync/errgroup"

type User struct {
	Name string
}

func fetch(ctx context.Context, user User) (string, error) {
	//if user.Name == "Ann" {
	//	return "", errors.New("invalid name")
	//}

	ch := make(chan any)

	go func() {
		time.Sleep(time.Millisecond * 10)
		close(ch)
	}()

	return user.Name, nil
}

func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, len(users))
	mu := &sync.Mutex{}

	egroup, ectx := errgroup.WithContext(ctx)
	//egroup.SetLimit(1)

	for _, u := range users {
		egroup.Go(func() error {
			name, err := fetch(ectx, u)
			if err != nil {
				return err
			}
			mu.Lock()
			names[name] = names[name] + 1
			mu.Unlock()

			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		return nil, err
	}

	return names, nil
}

func main() {

	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background()

	start := time.Now()
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an error occurred:", err.Error())
	}
	fmt.Println("time:", time.Since(start))
	fmt.Println(res)
}
