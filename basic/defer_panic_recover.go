package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Error with message:", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Error")
	}

	fmt.Println("App Running")
}

func main() {
	runApp(true)
	// runApp(false)
}
