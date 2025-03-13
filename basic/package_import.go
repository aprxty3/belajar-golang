package main

import (
	"belajar-golang/helper"
	"fmt"
)

func main() {
	fmt.Println(helper.HeyHello("Aji"))

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)             // tidak bisa diakses karena dia private
	// fmt.Println(helper.heyGoodbBye("Aji"))  // tidak bisa diakses karena dia private
}
