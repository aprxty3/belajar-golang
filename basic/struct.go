package main

import "fmt"

type Customer struct {
	FirstName, LastName, Address string
	Age                           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.FirstName)
}

func main() {
	customer := Customer{
		FirstName: "Aji",
		LastName:  "Sukma",
		Address:   "Jl. Imam Bonjol",
		Age:       20,
	}
	fmt.Println(customer)

	var customer2 Customer
	customer2.FirstName = "Budi"
	customer2.LastName = "Santoso"
	customer2.Address = "Jl. Imam Bonjol"
	customer2.Age = 25

	fmt.Println(customer2)

	budi := Customer{
		FirstName: "Budi",
		LastName:  "Santoso",
		Address:   "Jl. Imam Bonjol",
		Age:       25,
	}

	budi.sayHello("Aji")
	customer.sayHello("Budi")
}
