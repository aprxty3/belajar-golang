package main

import (
	"belajar-golang/database"
	_ "belajar-golang/internal"
	"fmt"
)

func main() {

	fmt.Println(database.GetDatabase())
}
