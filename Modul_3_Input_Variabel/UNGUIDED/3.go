package main

import (
	"fmt"
)

func main() {
	// Input temperatur Celsius
	var celsius float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	// Konversi suhu Celcius ke Reamur
	reamur := celsius * 4 / 5

	// Konversi suhu Celcius ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Konversi suhu Celcius ke Kelvin
	kelvin := celsius + 273.15

	// Menampilkan hasil
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
