package main

/* Поменять местами два числа без создания временной переменной. */

import "fmt"

func main() {
	a, b := -2, -14

	// двойное присваивание
	a, b = b, a

	// сложение-вычитание
	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
