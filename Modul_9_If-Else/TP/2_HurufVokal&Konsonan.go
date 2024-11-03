package main

// Import package fmt untuk fungsi input output dan package strings untuk memanipulasi string
import (
	"fmt"
	"strings"
)

// Fungsi utama
func main() {
	var input string // Deklarasi variabel input dengan tipe data string
	//Input huruf
	fmt.Print("Masukkan satu huruf: ")
	fmt.Scanln(&input)

	// Ubah input menjadi huruf kapital
	input = strings.ToUpper(input)

	// Cek apakah input merupakan huruf vokal
	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}
