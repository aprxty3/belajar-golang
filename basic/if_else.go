package main

func main() {
	name := "Aji"
	// name := "Pras"

	if name == "Aji" {
		println("Hello " + name)
	} else if name == "Pras" {
		println("Hello Stranger")
	} else {
		println("Hello World")
	}

	//if short statement condition
	if length := len(name); length > 5 {
		println("Name is too long")
	} else {
		println("Name is good")
	}
}
