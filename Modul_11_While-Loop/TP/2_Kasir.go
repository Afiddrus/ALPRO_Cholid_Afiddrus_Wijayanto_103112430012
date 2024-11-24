package main

import "fmt"

func main() {
	var totalBelanja float32 = 0
	var lanjut string

	fmt.Println("Selamat datang di Program Kasir")

	for {
		var namaBarang string
		var hargaBarang float32

		fmt.Print("Masukkan nama barang: ")
		fmt.Scanln(&namaBarang)

		fmt.Print("Masukkan harga barang: Rp ")
		fmt.Scanln(&hargaBarang)

		totalBelanja += hargaBarang

		fmt.Printf("Total belanja saat ini: Rp %.2f\n", totalBelanja)

		fmt.Print("Tambah barang lagi? (y/n): ")
		fmt.Scanln(&lanjut)

		if lanjut != "y" && lanjut != "Y" {
			break
		}
	}

	fmt.Println("\n=== Struk Belanja ===")
	fmt.Printf("Total belanja: Rp %.2f\n", totalBelanja)
	fmt.Println("Terima kasih telah berbelanja!")
}
