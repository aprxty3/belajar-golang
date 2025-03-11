package main

import "fmt"

func sayHelloGuys(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloGuys("Aji", spamFilter)
	sayHelloGuys("Anjing", spamFilter)
}
