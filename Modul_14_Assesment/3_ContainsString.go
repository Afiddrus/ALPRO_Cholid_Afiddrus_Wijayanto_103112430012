package main

import (
	"fmt"     // Import package fmt untuk fungsi input dan output
	"strings" // Import package strings untuk membuat huruf menjadi tidak peka besar/kecil
)

// Fungsi utama
func main() {

	var input_103112430012 string // Deklarasi variabel input bertipe string yang digunakan untuk meminta string atau karakter dari user

	// Meminta nilai input dari user
	fmt.Print("Masukkan string: ")
	fmt.Scanln(&input_103112430012)

	// Mengubah input menjadi huruf kecil agar tidak peka huruf besar/kecil
	input_103112430012 = strings.ToLower(input_103112430012)

	// Mencetak jumlah string "go" pada input
	fmt.Println(strings.Count(input_103112430012, "go"))

}
