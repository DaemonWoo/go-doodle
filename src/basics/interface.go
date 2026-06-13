package basics

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func makeSound(s Speaker) {
	s.Speak()
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

type emp any

func echo(e emp) string {
	return e.(CoffeePot).String()
}

func freakOut() {
	defer calmDown()
	panic("oh no")
}
func calmDown() {
	recover()
}
func c() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	dog := Dog{}
	makeSound(dog)

	err := fmt.Errorf("ex")
	println(err)

	freakOut()
}

type Player interface {
	Play(string)
	Stop()
}
