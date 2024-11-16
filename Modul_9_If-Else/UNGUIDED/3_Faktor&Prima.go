package main

import "fmt"

func main() {
	var b int
	var hitungFaktor int

	// Menerima input bilangan b
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	// Menampilkan faktor-faktor dan menghitung jumlah faktor
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			hitungFaktor++
		}
	}
	fmt.Println()

	// Menentukan apakah b adalah bilangan prima berdasarkan jumlah faktor
	if hitungFaktor == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
