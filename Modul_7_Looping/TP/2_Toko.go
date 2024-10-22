package main

import (
	"fmt" // Mengimpor package fmt untuk input dan output
)

func main() {
	// Deklarasi variabel
	var jumlahBarang, totalPoin int

	// Input jumlah barang
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Menghitung total poin
	if jumlahBarang <= 5 {
		// Jika barang <= 5, setiap barang mendapat 10 poin
		totalPoin = jumlahBarang * 10
	} else {
		// Jika barang > 5, barang pertama sampai ke-5 mendapat 10 poin
		// Barang ke-6 dan seterusnya mendapat 15 poin
		totalPoin = 5*10 + (jumlahBarang-5)*15
	}

	// Output hasil
	fmt.Printf("Total poin yang didapat adalah %d poin\n", totalPoin)
}
