package main

import (
	"fmt"
	"sync"
)

/* Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout. */

func main() {
	cNum := make(chan int)
	cSq := make(chan int)

	nums := []int{5, 7, 8, 3, 9, 4}

	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	go func() {
		for num := range cNum {
			cSq <- num * num
		}
		close(cSq)
	}()

	go func() {
		for sq := range cSq {
			fmt.Println(sq)
		}
		wg.Done()
	}()

	for i := 0; i < len(nums); i++ {
		cNum <- nums[i]
	}
	close(cNum)
}
