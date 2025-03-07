package main

func main() {
	var names [2]string
	names[0] = "Aji"
	names[1] = "Prasetyo"

	println(names[0], names[1])

	var names2 = [2]string{
		"Aji",
		"Prasetyo",
	}

	println(names2[0], names2[1])

	var values = [3]int{
		90,
		95,
		80,
	}

	println(values[0], values[1], values[2])
}
