package main

import "fmt"

func main() {

	// Deklarasi variabel
	var sisi int
	var volume int

	//Input parameters
	fmt.Print("Masukkan nilai sisi: ")
	fmt.Scan(&sisi)

	// Rumus konversi Fahrenheit ke Kelvin
	volume = sisi * sisi * sisi

	// Output parameters
	fmt.Printf("Volume Kubus Adalah %.2f\n ", volume)
}
