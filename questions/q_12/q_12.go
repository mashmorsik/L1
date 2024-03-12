package main

import "fmt"

// Выведет 0, тк переменные, инициализированные внутри if не видны за его пределами
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
