package main

import (
	"fmt"
	"sync"
)

// Программа выведет deadlock, потому что мы передаём в функцию копию WaitGroup, а не указатель на неё
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
