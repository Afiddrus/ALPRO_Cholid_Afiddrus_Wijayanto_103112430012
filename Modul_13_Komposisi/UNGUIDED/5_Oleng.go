package main

import (
	"fmt"
	"math"
)

func VersiPertama() {
	for {
		var kantongKiri, kantongKanan float64
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantongKiri, &kantongKanan)

		if kantongKiri >= 9 || kantongKiri >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}

func VersiKedua() {
	for {
		var kantongKiri, kantongKanan float64
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantongKiri, &kantongKanan)

		if kantongKiri < 0 || kantongKanan < 0 {
			fmt.Println("Program selesai.")
			break
		}

		totalBerat := kantongKiri + kantongKanan
		if totalBerat >= 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisihBerat := math.Abs(kantongKiri - kantongKanan)

		akanOleng := selisihBerat >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)
	}
}

func main() {
	/// VersiPertama() Untuk versi pertama
	VersiKedua() // Untuk versi kedua
}
