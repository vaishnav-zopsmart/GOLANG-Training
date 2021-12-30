package main

import "fmt"

type Animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age int
}

func (l dog) breathe() {
	fmt.Println("Dog breathes")
}

func (l dog) walk() {
	fmt.Println("Dog walk")
}

func NewDog(age int) Animal {
	return dog{
		age: age,
	}
}

func NewLion(age int) Animal {
	return lion{
		age: age,
	}
}

func main() {
	a := NewLion(10)
	a.breathe()
	a.walk()

	a = NewDog(5)
	a.breathe()
	a.walk()
}
