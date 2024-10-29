package main

// Import package fmt untuk fungsi input output
import "fmt"

// Fungsi utama
func main() {
	var angka int                  // deklarasi variabel angka bertipe integer
	fmt.Print("Masukkan angka = ") // menampilkan teks masukkan angka
	fmt.Scan(&angka)               // membaca dan menyimpan angka

	// percabangan
	if angka%2 == 0 {
		fmt.Println("Angka tersebut adalah bilangan genap")
	} else {
		fmt.Println("Angka tersebut adalah bilangan ganjil")
	}
}
