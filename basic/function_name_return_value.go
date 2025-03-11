package main

import "fmt"

func namedReturnvValue() (firstName, lastName string, age int) {
	firstName = "Aji"
	lastName = "Kurniawan"
	age = 20

	return
}

func main() {
	firstName, lastName, age := namedReturnvValue()
	fmt.Println(firstName, lastName, age)
}
