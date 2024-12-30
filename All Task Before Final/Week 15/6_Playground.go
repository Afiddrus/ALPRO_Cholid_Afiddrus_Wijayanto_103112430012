package main

import "fmt"

func main() {
	var membership string
	var durasi int
	var totalTarif int

	fmt.Print("Masukkan jenis membership (Gold/Silver/None): ")
	fmt.Scan(&membership)
	fmt.Print("Masukkan durasi bermain (jam): ")
	fmt.Scan(&durasi)

	if durasi <= 2 {
		totalTarif = durasi * 65000
	} else {
		totalTarif = (2 * 65000) + ((durasi - 2) * 20000)
	}

	if membership == "Gold" {
		totalTarif = totalTarif / 2 // Diskon 50%
	} else if membership == "Silver" {
		diskon := totalTarif * 25 / 100
		totalTarif = totalTarif - diskon // Diskon 25%
	}

	// Tampilkan hasil
	fmt.Printf("Total tarif: IDR %d\n", totalTarif)
}
