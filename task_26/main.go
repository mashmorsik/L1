package main

import (
	"fmt"
	"slices"
	"strings"
)

/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой. */

func main() {
	fmt.Println(IsUnique("akdA"))
}

func IsUnique(str string) bool {
	lower := strings.Split(strings.ToLower(str), "")
	slices.Sort(lower)

	count := 1
	first := lower[0]

	/* проходимся по отсортировнному массиву и сравниваем первую букву со второй, увеличиваем счетчик, если они равны,
	проверяем, равен ли счетчик двум, если да, то возвращаем false. Если первая и вторая буква не равны, то
	сравниваем вторую и третью и т.д. */
	for i := 1; i < len(lower); i++ {
		if lower[i] == first {
			count++
			if count == 2 {
				return false
			}
		} else {
			count = 1
		}
		first = lower[i]
	}

	return true
}
