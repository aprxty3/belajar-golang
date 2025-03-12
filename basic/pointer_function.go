package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	adress := Address{}
	ChangeCountryToIndonesia(&adress)

	fmt.Println(adress)

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	ChangeCountryToIndonesia(&address1)

	fmt.Println(address1)
	fmt.Println(address2)

}
