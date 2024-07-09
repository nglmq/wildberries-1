package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) Speak() { // метод структуры Human
	fmt.Printf("Hello my name is %s\n", h.Name)
}

type Action struct {
	Human  // встраиваем структуру
	Action string
}

func (a Action) Do() { // метод структуры Action
	fmt.Printf("%s doing action: %s\n", a.Name, a.Action)
}

func main() {
	action := Action{
		Human: Human{
			Name: "John",
		},
		Action: "eating",
	}
	action.Speak() // вызов метода Speak структуры Human через структуру Action
	action.Do()
}
