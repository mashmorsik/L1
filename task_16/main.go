package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	fmt.Println(QuickSort([]int{5, 1, 9, 0, 3}))
}

// QuickSort использует алгоритм быстрой сортировки
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	// опорный элемент
	pivot := nums[0]
	// два среза, которые будут содержать элементы, меньшие и большие (или равные) опорному элементу соответственно.
	var less, greater []int
	// цикл проходит по всем элементам среза, начиная со второго элемента, и проверяет их отношение к опорному элементу.
	for _, n := range nums[1:] {
		if n <= pivot {
			less = append(less, n)
		} else {
			greater = append(greater, n)
		}
	}
	/* Рекурсивно вызывается функция QuickSort для среза less, а затем к полученному результату
	добавляется опорный элемент pivot. */
	res := append(QuickSort(less), pivot)
	/* Рекурсивно вызывается функция QuickSort для среза greater, и результат добавляется к срезу res. */
	res = append(res, QuickSort(greater)...)
	return res
}
