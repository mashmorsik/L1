package main

import (
	"fmt"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10), и выведет их квадраты в stdout. */

func main() {
	// добавляем WaitGroup для ожидания выполнения нескольких горутин
	var wg sync.WaitGroup

	nums := []int{2, 4, 6, 8, 10}

	for _, num := range nums {
		// увеличиваем счетчик wg для каждой горутины
		wg.Add(1)
		go CountSquare(num, &wg)
	}

	// ждем, пока запущенные горутины не выполнятся
	wg.Wait()
}

func CountSquare(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)

	// оповещаем WaitGroup о завершении горутины
	wg.Done()
}
