package main

import (
	"fmt"
)

func main() {
	var totalBelanja, diskon int

	// Input total belanja
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)

	// Input besaran diskon dalam persen
	fmt.Print("Masukkan diskon (dalam persen): ")
	fmt.Scan(&diskon)

	// Menghitung total harga setelah diskon
	totalAkhir := totalBelanja - (totalBelanja * diskon / 100)

	// Output total harga akhir
	fmt.Printf("Total belanja setelah diskon: %d\n", totalAkhir)
}
