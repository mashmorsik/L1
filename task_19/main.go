package main

import (
	"bytes"
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode. */

func main() {
	fmt.Println(Flipper("главрыба"))
	fmt.Println(FlipperRune("главрыба"))
	fmt.Println(FlipperByte("главрыба"))
}

// Делаем массив из букв, переворачиваем - возвращаем строку
func Flipper(str string) string {
	letters := strings.Split(str, "")
	var result string

	for i := len(letters) - 1; i >= 0; i-- {
		result += letters[i]
	}
	return result
}

// Делаем массив рун, переворачиваем - возвращаем строку
func FlipperRune(str string) string {
	runeStr := []rune(str)
	for i := 0; i < len(runeStr)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-1-i] = runeStr[len(runeStr)-1-i], runeStr[i]
	}
	return string(runeStr)
}

// Делаем массив байт, переворачиваем - возвращаем строку
func FlipperByte(str string) string {
	letters := []byte(str)
	var result bytes.Buffer

	for i := len(letters) - 1; i >= 0; i-- {
		result.WriteByte(letters[i])
	}
	return result.String()
}
