package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		<-sigCh
		fmt.Println("context done")
		cancel()
	}()

	wg.Add(4)

	go Worker1(ctx, &wg)
	go Worker2(ctx, &wg)
	c := make(chan int)
	go sendToChan(c, &wg)
	go Worker3(c, &wg)

	wg.Wait()
}

// Завершаем горутину после отмены контекста
func Worker1(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("ctx done, finishing")
			return
		default:
			fmt.Printf("Keep workin worker1")
			time.Sleep(time.Second)
		}
	}
}

// Завершаем горутину через контекст таймаут
func Worker2(ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done, finishing")
			wg.Done()
			return
		default:
			fmt.Printf("Keep workin worker2")
			time.Sleep(time.Second)
		}
	}
}

// Завершаем горутину после чтения данных из канала и его закрытия
func Worker3(c chan int, wg *sync.WaitGroup) {
	for {
		val, ok := <-c
		if !ok {
			wg.Done()
			break
		}
		fmt.Printf("Keep workin worker3, val: %v", val)
	}
}

func sendToChan(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	defer wg.Done()
}
