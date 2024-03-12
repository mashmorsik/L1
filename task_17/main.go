package main

import (
	"fmt"
)

/* Реализовать бинарный поиск встроенными методами языка. */

func main() {
	fmt.Println(BinarySearch([]int{4, 3, 6, 8, 12, 1, -1}, 8))
}

func BinarySearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middleIn := (left + right) / 2
		middleNum := arr[middleIn]
		if middleNum == num {
			return middleIn
		} else if num > middleNum {
			left = middleIn + 1
		} else {
			right = middleIn - 1
		}
	}
	return 0
}
