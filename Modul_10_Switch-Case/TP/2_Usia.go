package main

// Import package fmt untuk fungsi input dan output
import "fmt"

// Fungsi utama
func main() {
	var usia int // Deklarasi variabel usia dengan tipe data integer

	// Input usia
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	// Fungsi switch-case
	switch {
	case usia >= 0 && usia <= 12: // Jika usia lebih besar atau sama dengan 0 dan usia kurang dari atau sama dengan 12 maka menampilkan kategori anak-anak
		fmt.Println("Kategori Usia: Anak-anak")

	case usia >= 13 && usia <= 17: // Jika usia lebih besar atau sama dengan 13 dan usia kurang dari atau sama dengan 17 maka menampilkan kategori remaja
		fmt.Println("Kategori Usia: Remaja")

	case usia >= 18 && usia <= 64: // Jika usia lebih besar atau sama dengan 18 dan usia kurang dari atau sama dengan 64 maka menampilkan kategori dewasa
		fmt.Println("Kategori Usia: Dewasa")

	case usia >= 65: // Jika usia lebih besar atau sama dengan 65 maka menampilkan kategori lansia
		fmt.Println("Kategori Usia: Lansia")

	default: // Selain itu maka menampilkan usia tidak valid
		fmt.Println("Usia tidak valid.")
	}
}
