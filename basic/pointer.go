package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
	address2 := address1 // copy value

	address2.City = "Semarang"

	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah

	// pass by reference
	address3 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
	address4 := &address3 // pointer

	address4.City = "Semarang"

	fmt.Println(address3) // ikut berubah
	fmt.Println(address4) // berubah
}
