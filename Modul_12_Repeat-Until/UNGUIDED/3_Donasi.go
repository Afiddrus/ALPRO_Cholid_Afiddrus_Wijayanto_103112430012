package main

import "fmt"

func main() {
	var target, totalDonasi, jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	selesai := false

	for selesai == false {
		var donasi int
		fmt.Printf("Donatur %d: Masukkan jumlah donasi: ", jumlahDonatur+1)
		fmt.Scan(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)

		selesai = totalDonasi >= target
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
