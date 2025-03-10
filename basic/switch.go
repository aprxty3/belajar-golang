package main

import "fmt"

func main() {
	name := "Aji"

	switch name {
	case "Aji":
		fmt.Println("Hello " + name)
	case "Prasetyo":
		fmt.Println("Hello " + name)
	default:
		fmt.Println("Hello")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name is good")
	}

	name = "Prasetyo"
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name is too long")
	case length > 10:
		fmt.Println("Name is very long")
	default:
		fmt.Println("Name is good")
	}
}
