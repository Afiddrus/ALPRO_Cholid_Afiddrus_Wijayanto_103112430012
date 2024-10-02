package main

import "fmt"

func main() {
	// Deklarasi variabel angka1, angka2, operator, dan hasil
	var angka1, angka2 float64
	var operator string
	var hasil float64

	// Input parameters
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)

	// Input operator
	fmt.Print("Masukkan operator (+, -, *, /, %): ")
	fmt.Scanln(&operator)

	// Input parameters
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	// Melakukan operasi berdasarkan operator yang dipilih
	switch operator {
	case "+":
		hasil = angka1 + angka2
		fmt.Printf("Hasil dari penjumlahan %.f dan %.f adalah %.f\n", angka1, angka2, hasil)

	case "-":
		hasil = angka1 - angka2
		fmt.Printf("Hasil dari pengurangan %.f dan %.f adalah %.f\n", angka1, angka2, hasil)

	case "*":
		hasil = angka1 * angka2
		fmt.Printf("Hasil dari perkalian %.f dan %.f adalah %.f\n", angka1, angka2, hasil)

	// Validasi
	case "/":
		if angka2 != 0 {
			hasil = angka1 / angka2
			fmt.Printf("Hasil dari pembagian %.f dengan %.f adalah %.f\n", angka1, angka2, hasil)
		} else {
			fmt.Println("Error: Tidak dapat dibagi dengan 0")
			return
		}
	// Validasi
	case "%":
		if angka2 != 0 {
			hasil = float64(int(angka1) % int(angka2)) // Konversi ke int untuk modulus
			fmt.Printf("Hasil dari modulus %.f dan %.f adalah %.f\n", angka1, angka2, hasil)
		} else {
			fmt.Println("Error: Pembagi tidak boleh 0")
			return
		}

	default:
		fmt.Println("Operator tidak valid")
		return
	}
}
