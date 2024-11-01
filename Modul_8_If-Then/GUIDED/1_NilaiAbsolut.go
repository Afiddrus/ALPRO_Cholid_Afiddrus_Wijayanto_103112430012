package main

// Import package fmt untuk fungsi input output
import "fmt"

// Fungsi utama
func main() {
	var nilaiAbsolut int                   // deklarasi variabel nilaiAbsolut bertipe integer
	fmt.Print("Masukkan Nilai Absolut = ") // menampilkan teks masukkan nilai absolut
	fmt.Scan(&nilaiAbsolut)                // membaca dan menyimpan nilai absolut
	// percabangan if-then
	if nilaiAbsolut < 0 {
		nilaiAbsolut = -nilaiAbsolut
	}
	// memasuki kondisi then percabangan
	fmt.Println("Nilai absolut adalah ", nilaiAbsolut)
}
