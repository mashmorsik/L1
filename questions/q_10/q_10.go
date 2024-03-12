package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	// 1, тк является указателем на а
	fmt.Println(*p)
	update(p)
	// 1, тк внутри функции update создалась копия p
	fmt.Println(*p)
}
