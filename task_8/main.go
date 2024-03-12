package main

import (
	"fmt"
)

/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

func main() {
	var num int64 = 673
	fmt.Printf("%b\n", num)
	changed := ZerosOnes(num, 5, 0)
	fmt.Printf("%b\n", changed)

}

func ZerosOnes(num int64, ind int, bit int) int64 {
	if bit == 1 {
		num |= 1 << ind
	} else {
		num &^= 1 << ind
	}
	return num
}
