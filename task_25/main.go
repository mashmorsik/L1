package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	ctx := context.Background()
	cctx, _ := context.WithTimeout(ctx, time.Second*5)

	Sleep(cctx, time.Second, func() {
		Add(2, 5)
	})
}

// Sleep запускает таймер, который по истечении времени запускает выполнение функции, можно добавить контекст,
// на случай, если хотим прервать выполнение функции сна
func Sleep(ctx context.Context, duration time.Duration, function func()) {
	fmt.Printf("started: %v\n", time.Now())
	var wg sync.WaitGroup

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("context done called: %v\n", time.Now())
				wg.Done()
				return
			case <-ticker.C:
				fmt.Printf("func called: %v\n", time.Now())
				function()
				wg.Done()
			}
		}
	}()
	wg.Wait()

	fmt.Printf("finished: %v\n", time.Now())
}

func Add(a, b int) int {
	return a + b
}
