package main

import (
	"fmt"
)

func main() {
	var n, alas, tinggi int

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan Alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan Tinggi: ")
		fmt.Scan(&tinggi)
		luas := alas * tinggi / 2
		print("Luas: ", luas)
	}
}
