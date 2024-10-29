package main

// Import package fmt untuk fungsi input output
import "fmt"

// Fungsi utama
func main() {
	var nilaiUjian int                   // deklarasi variabel nilaiUjian bertipe integer
	fmt.Print("Masukkan Nilai Ujian = ") // menampilkan teks masukkan nilai ujian
	fmt.Scan(&nilaiUjian)                // membaca dan menyimpan nilai ujian

	// percabangan
	if nilaiUjian >= 70 {
		fmt.Println("Anda Lulus")
	} else {
		fmt.Println("Anda Tidak Lulus")
	}
}
