package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// имеется функция которая работает неопределенно долго(до 100 секунд)
func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// написать обертку для этой функции которая будет прерывать выполнение, если
// функция работает дольше 3 секунд, и возвращать ошибку
func predictableTimeWork() error {
	s := make(chan struct{})

	go func() {
		randomTimeWork()
		close(s)
	}()

	select {
	case <-time.After(3 * time.Second):
		return errors.New("time out")
	case <-s:
		return nil
	}
}

func main() {
	e := predictableTimeWork()

	fmt.Println(e)
}
