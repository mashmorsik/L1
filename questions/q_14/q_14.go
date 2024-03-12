package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		// здесь меняется указатель на массив
		slice = append(slice, "a")
		// изменения происходят по новому указателю, но мы этот слайс не возвращаем, поэтому изменения не сохранятся
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
