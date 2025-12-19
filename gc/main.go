package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	a []byte
}

//go:noinline
func getData() *data {
	arr := make([]byte, 1024*10) // 10 КБ

	return &data{
		a: arr,
	}
}

func main() {
	t1 := time.NewTimer(100 * time.Microsecond)
	go func() {
		for range t1.C {
			getData()
		}
	}()

	t := time.NewTicker(1 * time.Second)

	var m runtime.MemStats
	now := time.Now()
	for curr := range t.C {
		runtime.ReadMemStats(&m)

		fmt.Printf(
			"GC_enabled %v GC_runs %v Live_now %v\n"+
				"Pause_total_ms %.2f Time %5.0f sec\n",
			m.EnableGC,
			m.NumGC,
			m.Mallocs-m.Frees,
			float64(m.PauseTotalNs)/1000/1000,
			curr.Sub(now).Seconds(),
		)
	}
}
