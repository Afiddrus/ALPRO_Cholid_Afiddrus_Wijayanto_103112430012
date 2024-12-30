package main

import "fmt"

func sumOddNumbers(a, b int) int {
	sum := 0

	if a%2 == 0 {
		a++
	}

	for i := a; i <= b; i += 2 {
		sum += i
	}

	return sum
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	result := sumOddNumbers(a, b)
	fmt.Printf("Hasil penjumlahan bilangan ganjil dari %d sampai %d adalah: %d\n", a, b, result)
}
