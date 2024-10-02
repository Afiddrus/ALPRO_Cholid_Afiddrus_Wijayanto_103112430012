package main

import "fmt"

func main() {
	// Mendefinisikan konstanta pi
	const pi = 3.1415926535

	// Input jari-jari
	var r float64
	fmt.Print("Jejari = ")
	fmt.Scanln(&r)

	// Rumus menghitung volume bola
	volume := (4.0 / 3.0) * pi * r * r * r

	// Rumus menghitung luas permukaan bola
	luas := 4 * pi * r * r

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
