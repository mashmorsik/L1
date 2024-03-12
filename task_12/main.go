package main

import "fmt"

/* Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество. */

func main() {
	fmt.Println(MakeSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func MakeSet(words []string) map[string]bool {
	set := make(map[string]bool)

	for _, w := range words {
		set[w] = true
	}
	return set
}
