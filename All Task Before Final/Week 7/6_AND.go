package main

import (
	"fmt"
)

func main() {
	var n int
	var input bool
	result := true // Inisialisasi hasil operasi AND

	fmt.Print("Masukkan jumlah nilai boolean (n): ")
	fmt.Scan(&n)

	fmt.Println("Masukkan nilai boolean (true/false) sebanyak", n, ":")
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		result = result && input
	}

	fmt.Printf("Hasil operasi AND: %t\n", result)
}
