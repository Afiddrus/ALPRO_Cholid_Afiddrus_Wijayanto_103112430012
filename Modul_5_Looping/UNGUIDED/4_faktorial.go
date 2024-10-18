package main

import "fmt" // Mengimpor package fmt untuk input/output format.

func main() {
	var input int // Mendeklarasikan variabel 'input' bertipe integer.

	fmt.Print("Masukkan bilangan bulat non-negatif: ") // Mencetak pesan untuk meminta input dari pengguna.
	fmt.Scan(&input)                                   // Membaca input dari pengguna dan menyimpannya dalam variabel 'input'.

	// Inisialisasi hasil faktorial sebagai 1
	hasil := 1

	// Hitung faktorial menggunakan loop
	for i := 1; i <= input; i++ {
		hasil *= i
	}

	// Menampilkan hasil faktorial
	fmt.Print("Hasil dari faktorial ", input, " adalah: ", hasil)
}
