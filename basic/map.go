package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Aji",
		"age":     "23",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Aji Prasetyo"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(book["title"])

	delete(book, "ups")

	fmt.Println(book)
}
