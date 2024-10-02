package main

import "fmt"

func main() {

	// Deklarasi variabel
	var alas float32
	var tinggi float32

	//Input parameters
	fmt.Print("Masukkan nilai Alas: ")
	fmt.Scan(&alas)

	//Input parameters
	fmt.Print("Masukkan nilai Tinggi: ")
	fmt.Scan(&tinggi)

	// Rumus luas segitiga
	luas := 0.5 * alas * tinggi

	//Ouptut parameter
	fmt.Printf("Luas Segitiga Adalah %.2f\n", luas) // %.2f digunakan untuk 2 angka desimal
}
