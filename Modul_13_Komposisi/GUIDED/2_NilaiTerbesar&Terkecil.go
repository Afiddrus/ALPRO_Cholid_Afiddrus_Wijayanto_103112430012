package main

import "fmt"

// Fungsi untuk mencari nilai terbesar
func findMax(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	}
	return c
}

// Fungsi untuk mencari nilai terkecil
func findMin(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	}
	return c
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Memanggil fungsi findMax dan findMin
	max := findMax(a, b, c)
	min := findMin(a, b, c)

	fmt.Println("Nilai terbesar adalah", max)
	fmt.Println("Nilai terkecil adalah", min)
}
