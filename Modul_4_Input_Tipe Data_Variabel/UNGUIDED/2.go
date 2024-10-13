package main

import (
	"fmt"
)

func main() {
	var bmi, tinggiBadan float64

	// Input nilai BMI
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)

	// Input tinggi badan dalam meter
	fmt.Print("Masukkan tinggi badan (dalam meter): ")
	fmt.Scan(&tinggiBadan)

	// Menghitung berat badan berdasarkan rumus: Berat Badan = BMI * (Tinggi Badan ^ 2)
	beratBadan := bmi * (tinggiBadan * tinggiBadan)

	// Output berat badan
	fmt.Printf("Berat badan Anda adalah: %.2f kg\n", beratBadan)
}
