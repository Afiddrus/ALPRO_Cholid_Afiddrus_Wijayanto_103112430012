package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah penumpang: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Jumlah penumpang harus lebih besar dari 0.")
		return
	}

	var totalWeight float64 = 0
	for i := 1; i <= N; i++ {
		var weight float64
		fmt.Printf("Masukkan berat koper penumpang %d: ", i)
		fmt.Scan(&weight)
		totalWeight += weight
	}

	averageWeight := totalWeight / float64(N)
	fmt.Printf("Rata-rata berat koper: %.2f\n", averageWeight)
}
