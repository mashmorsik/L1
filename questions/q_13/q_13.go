package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

// Выведет [100, 2, 3, 4, 5], тк слайс ссылается на массив, мы можем менять его элементы, а append создает копию массива,
// чтобы вернулась эта копия нужен return
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
