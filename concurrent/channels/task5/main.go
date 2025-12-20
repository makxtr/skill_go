package main

import (
	"math/rand"
	"time"
)

func unpredictableFunc() int {
	n := rand.Intn(40)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func predictableFunc() (int, error) {

}

func main() {
	_, _ = predictableFunc()
}
