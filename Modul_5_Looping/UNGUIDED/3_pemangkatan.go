package main

import "fmt" // Mengimpor package fmt untuk input/output format.

// Fungsi utama, titik awal eksekusi program.
func main() {
	var a, b, hasil int // Deklarasi variabel a, b untuk bilangan input dan hasil untuk menyimpan hasil perhitungan dengan tipe data integer.

	fmt.Print("Masukkan bilangan pertama: ") // Mencetak pesan untuk meminta input dari pengguna.
	fmt.Scan(&a)                             // Membaca input dari pengguna dan menyimpannya dalam variabel 'a'.

	fmt.Print("Masukkan bilangan kedua: ") // Mencetak pesan untuk meminta input dari pengguna.
	fmt.Scan(&b)                           // Membaca input dari pengguna dan menyimpannya dalam variabel 'b'.

	hasil = 1 // Menginisialisasi variabel hasil dengan nilai 1, karena pangkat dihitung dengan perkalian.

	// Loop untuk menghitung a dipangkatkan dengan b, yaitu mengalikan a sebanyak b kali.
	for i := 0; i < b; i++ {
		hasil *= a
	}

	// Menampilkan hasil dari perhitungan pangkat.
	fmt.Print("Hasil dari bilangan ", a, " dipangkatkan dengan bilangan ", b, " adalah ", hasil)

}
