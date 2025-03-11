package main

import "fmt"

type Customer struct {
	FirstName, LastName, Address string
	Age                           int
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
	
}
