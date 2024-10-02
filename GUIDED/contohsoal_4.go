package main

import "fmt"

func main() {

	// Deklarasi variabel
	var nilai float64

	//Input parameters
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	hasil := (2/(nilai+5) + 5) //Rumus

	fmt.Printf("Nilainya Adalah %.f\n", hasil) // %.2f digunakan untuk 2 angka desimal
}
