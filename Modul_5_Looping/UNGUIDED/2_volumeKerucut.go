package main

import (
	"fmt"  // Mengimpor package fmt untuk input/output format.
	"math" // Mengimpor package math untuk menggunakan konstanta dan fungsi matematika, disini digunakan untuk nilai .
)

// Fungsi utama, titik awal eksekusi program.
func main() {
	var n int // Mendeklarasikan variabel n bertipe int untuk jumlah kerucut yang akan dihitung.

	fmt.Print("Masukkan jumlah N (Kerucut): ") // Mencetak pesan untuk meminta input dari pengguna.
	fmt.Scan(&n)                               // Membaca input dari pengguna dan menyimpannya dalam variabel 'n'.

	// Mengulangi sebanyak n kali untuk setiap kerucut.
	for i := 0; i < n; i++ {
		var panjang, tinggi float32 // Mendeklarasikan variabel panjang dan tinggi bertipe float32.

		fmt.Print("Masukkan panjang jari-jari alas kerucut: ") // Mencetak permintaan input untuk panjang jari-jari.
		fmt.Scan(&panjang)                                     // Membaca input panjang dari pengguna dan menyimpannya dalam variabel 'panjang'.

		fmt.Print("Masukkan tinggi kerucut: ") // Mencetak permintaan input untuk tinggi kerucut.
		fmt.Scan(&tinggi)                      // Membaca input panjang dari pengguna dan menyimpannya dalam variabel 'tinggi'.

		// Hitung volume kerucut menggunakan rumus: V = (1/3) * π * r² * t.
		volume := (1.0 / 3.0) * math.Pi * float64(panjang) * float64(panjang) * float64(tinggi)

		// Format output untuk menampilkan volume kerucut dengan 10 angka di belakang koma.
		fmt.Printf("Volume kerucut: %.10f\n", volume) // Menggunakan format %.10f untuk 10 angka desimal.
	}
}
