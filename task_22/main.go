package main

import (
	"fmt"
	"math/big"
)

/* Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20. */

func main() {
	a := big.NewInt(3897456)
	b := big.NewInt(4100876)

	fmt.Println(MathOps(a, b, "subtract").String())
}

func MathOps(a *big.Int, b *big.Int, operation string) *big.Int {
	var c big.Int

	switch operation {
	case "add":
		c.Add(a, b)
	case "subtract":
		c.Sub(a, b)
	case "multiply":
		c.Mul(a, b)
	case "divide":
		c.Div(a, b)
	}
	return &c
}
