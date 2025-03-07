package main

func main() {
	const nilaiKkn = 90
	const nilaiAbsen = 80

	lulusNilaiKkn := nilaiKkn >= 80
	lulusNilaiAbsen := nilaiAbsen >= 80

	lulus := lulusNilaiKkn && lulusNilaiAbsen
	println(lulus)
}
