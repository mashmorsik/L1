package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func main() {
	arr1 := []string{"cat", "dog", "hamster", "fish"}
	arr2 := []string{"fish", "spider", "dog"}

	fmt.Println(FindMatch(arr1, arr2))
}

func FindMatch(arr1 []string, arr2 []string) []string {
	count := make(map[string]int)
	var result []string

	for _, w := range arr1 {
		count[w]++
	}
	for _, w := range arr2 {
		count[w]++
		if count[w] == 2 {
			result = append(result, w)
		}
	}

	return result
}
