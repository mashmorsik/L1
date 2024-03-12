package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования). */

type Human struct {
	Name      string
	Sex       string
	Age       int
	Languages []string
}

func (h *Human) Work(job string) string {
	work := fmt.Sprintf("%s works as a %s.", h.Name, job)
	return work
}

// Action "наследует" структуру Human
type Action struct {
	Human
}

func main() {
	Mark := &Human{
		Name:      "Mark",
		Sex:       "Male",
		Age:       32,
		Languages: []string{"English", "French"},
	}

	a := &Action{*Mark}
	b := &Action{Human{
		Name:      "Lucy",
		Sex:       "female",
		Age:       20,
		Languages: []string{"Russian"},
	}}

	// вызываем метод Work на экземпляре структуры Human (переменная Mark)
	fmt.Println(Mark.Work("firefighter"))

	// вызываем метод Work на экземпляре структуры Action (переменная a), используя данные переменной Mark
	fmt.Println(a.Work("cook"))

	/* вызываем метод Work на экземпляре структуры Action (переменная b), используя данные нового экземпляра Human,
	проинициализированного внутри переменной b */
	fmt.Println(b.Work("designer"))
}
