package main

import (
	"fmt"
)

func main() {
	// Input tahun
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	// Mengecek apakah tahun kabisat
	kabisat := false
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		kabisat = true
	}

	// Menampilkan hasil
	fmt.Printf("Kabisat: %t\n", kabisat)
}
