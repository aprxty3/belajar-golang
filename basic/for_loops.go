package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		if counter == 10 {
			fmt.Println("Stop")
		} else {
			fmt.Println("Perulangan ke", counter)
		}
		counter++
	}

	//simplify
	for i := 0; i < 10; i++ {
		fmt.Println("Angka", i)
	}

	names := []string{"Aji", "Prasetyo", "Ganteng"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}
