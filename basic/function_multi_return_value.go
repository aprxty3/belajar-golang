package main

import "fmt"

func getBiodata()(string, int){
	return "Aji", 20
}

func main()  {
	name, age := getBiodata()
	fmt.Println("Nama saya", name, "dan umur saya", age, "tahun")

	namaSaya, _ := getBiodata()
	fmt.Println("Nama saya", namaSaya)
}