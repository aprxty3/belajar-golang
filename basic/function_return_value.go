package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main(){
	name1 := getHello("Aji")
	name2 := getHello("")

	fmt.Println(name1)
	fmt.Println(name2)
}
