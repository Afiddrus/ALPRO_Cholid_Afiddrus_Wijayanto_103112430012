package main

import (
	"fmt"     // Import package untuk fungsi input dan output
	"strings" // Import package untuk membuat huruf menjadi tidak peka besar/kecil
)

// Fungsi utama
func main() {

	var a_103112430012, b string // deklarasi variabel a dan b yang menjadi input dengan tipe data string

	// Meminta input a dan b dari user
	fmt.Print("Masukkan nilai a dan b: ")
	fmt.Scanln(&a_103112430012, &b)

	// Mengubah kedua input menjadi huruf kecil dengan menggunankan package strings
	a_103112430012 = strings.ToLower(a_103112430012)
	b = strings.ToLower(b)

	//Perulangan If-Else untuk mengecek jika nilai a dan b tidak sama (!=) maka outputnya akan false, dan jika sama maka akan true.
	if a_103112430012 != b {
		fmt.Println("Tidak")
	} else {
		fmt.Println("Ya")
	}

}
