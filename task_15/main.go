package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func main() {
	someFunc()
	fmt.Println(justString)
}

// Проблема в том, что переменная justString ссылается на переменную v и не смотря на то, что нам нужна только часть
// изначальной строки, мы храним ее в памяти целиком

// Лучше использовать метод Clone, чтобы создавать новую аллокацию и хранить только необходимую нам часть от изначальной
// строчки, а память, которую занимала переменная v, освободит GC
func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func createHugeString(i int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arr := make([]byte, i)
	for a := range arr {
		arr[a] = letters[rand.Intn(len(letters))]
	}
	return string(arr)
}
