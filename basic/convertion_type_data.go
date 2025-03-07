package main

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	println(nilai32)
	println(nilai64)
	println(nilai8)

	var name = "Aji Prasetyo"
	var e = name[0]
	var a = name[1]
	var eString = string(e)
	var aString = string(a)
	println(eString, aString)
}
