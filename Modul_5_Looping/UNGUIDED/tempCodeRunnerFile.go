package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah N: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		var panjang, tinggi float32

		fmt.Print("Masukkan panjang persegi panjang: ")
		fmt.Scan(&panjang)

		fmt.Print("Masukkan tinggi persegi panjang: ")
		fmt.Scan(&tinggi)

		volume := math.Pi * panjang * panjang * tinggi * 1 / 3
		fmt.Printf("Volume persegi panjang: %.2f\n", volume)
	}
}
