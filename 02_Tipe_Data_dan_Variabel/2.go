package main

import "fmt"

func main() {
	var satu, dua, tiga string
	var temp string

	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal =", satu, dua, tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir =", satu, dua, tiga)
}

// Program ini melakukan pertukaran nilai dari tiga variabel string.
// Fungsi utama program adalah untuk mendemonstrasikan konsep pertukaran nilai dalam pemrograman, yang sering digunakan dalam algoritma pengurutan dan manipulasi data.
// Dengan alur:
// 1. Program meminta pengguna untuk memasukkan tiga buah string.
// 2. Kemudian, program menukar posisi dari ketiga string tersebut dan menampilkan hasil akhir setelah pertukaran.
