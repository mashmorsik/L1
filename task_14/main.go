package main

import "fmt"

func main() {
	var a interface{} = 21
	var b interface{} = true
	var c interface{} = "Sad"
	var d interface{} = make(chan string)

	fmt.Println(SwitchCase(a))
	fmt.Println(SwitchCase(b))
	fmt.Println(SwitchCase(c))
	fmt.Println(SwitchCase(d))
}

func SwitchCase(i interface{}) string {
	switch i.(type) {
	case int:
		return "This is int."
	case string:
		return "This is string."
	case bool:
		return "This is bool."
	case chan string:
		return "This is chan."
	default:
		return "I don't know what this is."
	}
}
