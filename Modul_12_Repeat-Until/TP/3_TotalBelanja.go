package main

import "fmt"

func main() {
	var price, total int

	for {

		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scanln(&price)

		// Cek apakah input = 0
		if price == 0 {
			break // keluar dari loop karena input = 0
		}

		// Tambahkan harga ke total
		total += price
	}

	fmt.Printf("Total belanja Anda: %d\n", total)
}
