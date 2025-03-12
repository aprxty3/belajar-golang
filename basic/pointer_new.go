package main

import "fmt"

type PointerNew struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	person1 := new(PointerNew)
	person2 := person1
	person2.FirstName = "Aji"
	person2.LastName = "Kusuma"
	person2.Age = 20

	fmt.Println(person1)
	fmt.Println(person2)
}
