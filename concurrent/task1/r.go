package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 10
	// Создаем массив каналов для синхронизации "цепочкой"
	chs := make([]chan struct{}, n)
	for i := range chs {
		chs[i] = make(chan struct{})
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Ждем сигнала от предыдущей горутины (кроме первой)
			if i > 0 {
				<-chs[i-1]
			}

			fmt.Println(i)

			// Даем сигнал следующей горутине
			if i < n-1 {
				chs[i] <- struct{}{}
			}
		}(i)
	}

	wg.Wait()
}

// ключевой момент в том что main может завершится быстрее
