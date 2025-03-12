package main

import "fmt"

type AsteriskOperator struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	person1 := AsteriskOperator{"Aji", "Kusuma", 20}
	fmt.Println(person1)

	person2 := &person1
	person2.FirstName = "Budi"

	person3 := &person1

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println("================")
	person2 = &AsteriskOperator{"Sadikin", "Kusuma", 20}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println("================")

	*person3 = AsteriskOperator{"Suaib", "Sadikin", 22}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}
