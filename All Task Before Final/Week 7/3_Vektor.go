package main

import "fmt"

func main() {
	var x1, y1, z1 float64
	var x2, y2, z2 float64
	var x3, y3, z3 float64

	fmt.Println("Masukkan koordinat X, Y, Z untuk vektor 1:")
	fmt.Scanln(&x1, &y1, &z1)
	fmt.Println("Masukkan koordinat X, Y, Z untuk vektor 2:")
	fmt.Scanln(&x2, &y2, &z2)
	fmt.Println("Masukkan koordinat X, Y, Z untuk vektor 3:")
	fmt.Scanln(&x3, &y3, &z3)

	avgX := (x1 + x2 + x3) / 3
	avgY := (y1 + y2 + y3) / 3
	avgZ := (z1 + z2 + z3) / 3

	fmt.Printf("Rata-rata vektor (X, Y, Z): %.2f, %.2f, %.2f\n", avgX, avgY, avgZ)
}
