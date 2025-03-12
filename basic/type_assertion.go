package main

import "fmt"

func Random() any {
	return "Ok"
}

func main() {
	result := Random()
	resultString := result.(string)
	fmt.Println(resultString)

	// var toInt int = result.(int)
	// fmt.Println(toInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
