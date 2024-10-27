package main

import (
	"fmt"
)

func main() {
	var N int

	// Input bilangan bulat positif dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Menampilkan faktor
	for d := 1; d <= N; d++ {
		var s bool

		// Cek apakah d adalah faktor dari N
		if N%d == 0 {
			s = true
		} else {
			s = false
		}

		// Output d dan status s
		fmt.Printf("%d %t\n", d, s)
	}
}
