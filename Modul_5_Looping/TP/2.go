package main

// Untuk Input/Output
import "fmt"

func main() {
	// Deklarasi variabel `input` untuk menampung angka dari user, dan `counter` untuk menghitung jumlah bintang di setiap baris
	var input, counter int

	// Mencetak pesan untuk meminta input dari user
	fmt.Print("Masukkan Angka: ")
	// Membaca input dari user dan menyimpannya ke variabel `input`
	fmt.Scan(&input)

	// Loop untuk membuat baris sebanyak nilai `input`
	for i := 1; i <= input; i++ {
		// Reset nilai counter menjadi 0 untuk memulai cetakan bintang baru di baris berikutnya
		counter = 0
		// Loop untuk mencetak bintang sebanyak nilai `i` pada baris tersebut
		for counter < i {
			// Mencetak bintang pada baris yang sama
			fmt.Print("*")
			// Increment `counter` untuk melanjutkan hingga mencapai nilai `i`
			counter++
		}
		// Mencetak baris baru setelah semua bintang pada baris tersebut tercetak
		fmt.Println()
	}
}
