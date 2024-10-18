package main

// untuk Input/Output
import "fmt"

func main() {
	// Deklarasi variabel `input` untuk menampung angka dari user, dan `sum` untuk menampung hasil penjumlahan
	var input, sum int

	// Mencetak pesan untuk meminta input dari user
	fmt.Print("Masukkan angka: ")
	// Membaca input dari user dan menyimpannya ke variabel `input`
	fmt.Scan(&input)

	// Loop dari 1 hingga nilai `input`
	for a := 1; a <= input; a++ {
		// Menambahkan nilai `a` ke variabel `sum` dalam setiap iterasi
		sum += a
	}

	// Mencetak hasil penjumlahan dari angka 1 hingga `input`
	fmt.Println("Hasil penjumlahan dari angka 1 hingga angka", input, "adalah:", sum)
}
