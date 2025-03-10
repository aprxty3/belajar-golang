package main

import "fmt"

func main() {
	names := [...]string{"Aji", "Prasetyo", "Budi", "Cahyo", "Dwi", "Eko", "Fajar", "Gatot", "Hari", "Indra"}
	slice1 := names[2:5]
	slice2 := names[:3]
	slice3 := names[4:]
	slice4 := names[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Hari Baru")
	fmt.Println(daySlice2)
	daySlice2[0] = "Hoooreay"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Aji"
	newSlice[1] = "Prasetyo"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Budi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[2:5]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
