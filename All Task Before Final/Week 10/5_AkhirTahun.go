package main

import "fmt"

func main() {
	var totalBelanja int
	var bersediaKartu bool
	fmt.Println("Masukkan total belanja:")
	fmt.Scan(&totalBelanja)
	fmt.Println("Apakah bersedia dibuatkan kartu? (true/false):")
	fmt.Scan(&bersediaKartu)

	kartu := bersediaKartu && totalBelanja >= 100000
	diskon := totalBelanja >= 100000
	cashback := kartu && totalBelanja >= 200000

	finalBelanja := totalBelanja
	if diskon {
		finalBelanja -= totalBelanja / 10 // Diskon 10%
	}
	if cashback {
		finalBelanja -= 75000 // Cashback Rp. 75.000
	}

	fmt.Printf("Kartu? %t\n", kartu)
	fmt.Printf("Diskon? %t\n", diskon)
	fmt.Printf("Cashback? %t\n", cashback)
	fmt.Printf("Rp. %d\n", finalBelanja)
}
