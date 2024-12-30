package main

import "fmt"

func main() {
	var totalBelanja float64
	fmt.Print("Masukkan total belanja awal (dalam rupiah): ")
	fmt.Scanln(&totalBelanja)

	var ikutCLO bool
	fmt.Print("Apakah mahasiswa sedang mengikuti asesmen CLO? (true/false): ")
	fmt.Scanln(&ikutCLO)

	var totalAkhir float64
	if ikutCLO {
		totalAkhir = totalBelanja * 0.65
	} else {
		totalAkhir = totalBelanja
	}

	fmt.Printf("Total belanja akhir yang harus dibayar: IDR %.2f\n", totalAkhir)
}
