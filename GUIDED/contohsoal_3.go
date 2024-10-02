package main

import "fmt"

func main() {

	// Deklarasi variabel
	var rupiah, konversi int

	//Input parameters
	fmt.Print("Masukkan nilai Rupiah: ")
	fmt.Scan(&rupiah)

	konversi = rupiah / 15000 //Konversi dollar US

	//Output parameters
	fmt.Println("Konversi ke Dollar Adalah ", konversi) // %.2f digunakan untuk 2 angka desimal
}
