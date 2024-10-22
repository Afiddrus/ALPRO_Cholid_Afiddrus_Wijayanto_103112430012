package main

import (
	"fmt" // Mengimpor package fmt untuk input dan output
)

func main() {
	// Deklarasi variabel
	var N int
	// Input bilangan bulat positif
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&N) // Menyimpan input bilangan bulat positif

	// Perulangan dari 1 hingga N
	for i := 1; i <= N; i++ {
		if i == N { // Jika i sama dengan N
			fmt.Print(i * i) // Cetak kuadrat i tanpa koma
		} else {
			fmt.Print(i*i, ",") // Cetak kuadrat i dengan koma
		}
	}
}
