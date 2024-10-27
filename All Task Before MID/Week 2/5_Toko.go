package main

import (
	"fmt"
)

func main() {
	var hargaBeli1, hargaBeli2, hargaBeli3 int
	fmt.Print("Masukkan harga beli barang 1, barang 2, dan barang 3: ")
	fmt.Scan(&hargaBeli1, &hargaBeli2, &hargaBeli3)

	// Menghitung harga jual dengan keuntungan 5%
	hargaJual1 := float64(hargaBeli1) + (float64(hargaBeli1) * 0.05)
	hargaJual2 := float64(hargaBeli2) + (float64(hargaBeli2) * 0.05)
	hargaJual3 := float64(hargaBeli3) + (float64(hargaBeli3) * 0.05)

	fmt.Printf("Harga jual barang 1: %.2f\n", hargaJual1)
	fmt.Printf("Harga jual barang 2: %.2f\n", hargaJual2)
	fmt.Printf("Harga jual barang 3: %.2f\n", hargaJual3)
}
