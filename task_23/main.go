package main

import "fmt"

/* Удалить i-ый элемент из слайса. */

func main() {
	fmt.Println(DeleteElement(2, []int{0, 1, 2, 3}))
}

func DeleteElement(i int, nums []int) []int {
	return append(nums[:i], nums[i+1:]...)
}
