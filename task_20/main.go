package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

func main() {
	fmt.Println(FlipperStr("snow dog sun"))
	fmt.Println(FlipperStrBuilder("snow dog sun"))
}

// Делаем массив из слов, переворачиваем - возвращаем строку
func FlipperStr(str string) string {
	words := strings.Split(str, " ")
	var result string

	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 {
			result += words[i]
		} else {
			result += words[i] + " "
		}
	}
	return result
}

// Оптимизируем
func FlipperStrBuilder(str string) string {
	words := strings.Fields(str)
	var result strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		if i != 0 {
			result.WriteByte(' ')
		}
	}
	return result.String()
}
