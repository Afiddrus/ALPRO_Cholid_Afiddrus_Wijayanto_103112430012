package main

import "fmt"

func main() {
	// Membaca input untuk panjang
	var panjang float64
	fmt.Print("Masukkan panjang prisma: ")
	fmt.Scanln(&panjang)

	// Membaca input untuk lebar
	var lebar float64
	fmt.Print("Masukkan lebar prisma: ")
	fmt.Scanln(&lebar)

	// Membaca input untuk tinggi
	var tinggi float64
	fmt.Print("Masukkan tinggi prisma: ")
	fmt.Scanln(&tinggi)

	// Menghitung volume prisma
	volume := panjang * lebar * tinggi

	// Menghitung luas permukaan prisma
	// Luas permukaan = 2 * (panjang * lebar + panjang * tinggi + lebar * tinggi)
	luasPermukaan := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)

	// Menampilkan hasil volume dan luas permukaan
	fmt.Printf("Volume prisma: %.2f\n", volume)
	fmt.Printf("Luas permukaan prisma: %.2f\n", luasPermukaan)
}
