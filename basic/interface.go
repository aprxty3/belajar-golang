package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func Any() any {
	return "Any"
}

func main() {
	animal := Animal{
		Name: "Dog",
	}

	person := Person{
		Name: "Aji",
	}

	SayHello(animal)
	SayHello(person)

	fmt.Println(Any())
}
