package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Println("Masukkan nilai a, b, c, d:")
	fmt.Scanln(&a, &b, &c, &d)

	determinant := (a * d) - (b * c)

	fmt.Printf("Determinannya adalah: %.2f\n", determinant)
}
