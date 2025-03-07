package main

func main() {
	type NamaKamu string

	var nama = "Aji"
	var namaKamu NamaKamu = "Prasetyo"
	var namaKamuLagi NamaKamu = NamaKamu(nama)

	println(nama)
	println(namaKamu)
	println(namaKamuLagi)
}
