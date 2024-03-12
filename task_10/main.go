package main

import (
	"fmt"
)

/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	fmt.Println(SortTemp([]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}))
}

func SortTemp(temp []float32) map[int][]float32 {
	set := make(map[int][]float32)

	for _, t := range temp {
		x := (int(t/10) % 10) * 10
		set[x] = append(set[x], t)
	}

	return set
}
