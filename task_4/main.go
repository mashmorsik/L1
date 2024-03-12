package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/* Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества
воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров. */

func main() {

	var (
		c     = make(chan int)
		sigCh = make(chan os.Signal, 1)

		numWorkers = 2
		wg         sync.WaitGroup

		ctx, cancel = context.WithCancel(context.Background())
	)

	signal.Notify(sigCh, os.Interrupt, os.Kill)

	// Отменяем контекст, если получаем сигнал завершения
	go func() {
		s := <-sigCh
		fmt.Printf("got signal %v\n", s)
		cancel()
	}()

	// Создаем воркеров и добавляем их в WaitGroup
	for j := 0; j < numWorkers; j++ {
		wg.Add(1)
		go Worker(ctx, j, c, &wg)
	}

	// Записываем данные в канал с промежутком в секунду
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}
		close(c)
	}()
	wg.Wait()
}

func Worker(ctx context.Context, num int, c chan int, wg *sync.WaitGroup) {
	for {
		select {
		// Завершаем горутину, если нам приходит сообщение об отмене контекста
		case <-ctx.Done():
			fmt.Printf("worker %d, context canceled\n", num)
			wg.Done()
			return
		case data, ok := <-c:
			// Завершаем горутину, если канал закрывается
			if !ok {
				fmt.Printf("worker %d, ch closed\n", num)
				wg.Done()
				return
			}
			fmt.Println(data)
		}
	}
}
