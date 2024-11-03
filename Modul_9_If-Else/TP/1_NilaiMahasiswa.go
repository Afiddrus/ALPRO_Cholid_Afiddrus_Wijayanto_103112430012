package main

// Import package fmt untuk fungsi input dan output
import "fmt"

// Fungsi utama
func main() {
	var nilai int // Deklarasi variabel nilai dengan tipe data integer
	// Input nilai
	fmt.Print("Masukkan Nilai : ")
	fmt.Scan(&nilai)
	if nilai > 90 { // Memeriksa apakah nilai lebih dari 90
		fmt.Println("Nilai mendapatkan indeks A") // menampilkan teks "Nilai mendapatkan indeks A"
	} else if nilai > 80 { // Memeriksa apakah nilai lebih dari 80
		fmt.Println("Nilai mendapatkan indeks AB") // menampilkan teks "Nilai mendapatkan indeks AB"
	} else if nilai > 70 { // Memeriksa apakah nilai lebih dari 70
		fmt.Println("Nilai mendapatkan indeks B") // menampilkan teks "Nilai mendapatkan indeks B"
	} else if nilai > 60 { // Memeriksa apakah nilai lebih dari 60
		fmt.Println("Nilai mendapatkan indeks C") // menampilkan teks "Nilai mendapatkan indeks C"
	}
}
