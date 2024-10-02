package main

import "fmt"

func main() {

	// Deklarasi variabel
	var fahrenheit float64

	//Input parameters
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Rumus konversi Fahrenheit ke Kelvin
	kelvin := (fahrenheit-32)*5/9 + 273.15

	fmt.Printf("%.2f derajat Fahrenheit sama dengan %.2f derajat Kelvin\n", fahrenheit, kelvin)
}
