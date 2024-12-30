package main

import (
	"fmt"
)

func main() {
	var n int
	var massaJenis, volume, massaTotal float64

	fmt.Print("Masukkan jumlah bola (n): ")
	fmt.Scan(&n)

	massaTotal = 0.0

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan massa jenis dan volume bola ke-%d (pisahkan dengan spasi): ", i)
		fmt.Scan(&massaJenis, &volume)
		massaTotal += massaJenis * volume
	}

	rataRata := massaTotal / float64(n)

	fmt.Printf("Rata-rata berat bola: %.2f kg\n", rataRata)
}
