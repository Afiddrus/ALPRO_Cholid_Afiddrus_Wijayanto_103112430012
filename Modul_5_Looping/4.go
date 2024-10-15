package main

import (
	"fmt"       // Package untuk input dan output
	"math/rand" // Package untuk menghasilkan angka acak
)

func main() {
	// Memilih angka acak antara 1 hingga 100
	targetNumber := rand.Intn(100) + 1 // Fungsi `rand.Intn(100)` menghasilkan angka acak dari 0 hingga 99, jadi +1 untuk membuat rentangnya 1 hingga 100
	var guess int                      // Variabel untuk menyimpan tebakan dari pengguna
	attempts := 5                      // Menentukan jumlah percobaan yang diberikan kepada pengguna

	// Menampilkan pesan pembuka
	fmt.Println("Selamat datang di permainan Tebak Angka!")
	fmt.Println("Kami telah memilih angka antara 1 hingga 100.")
	fmt.Printf("Kamu punya %d kesempatan untuk menebak angka tersebut.\n", attempts)

	// Loop sebanyak jumlah kesempatan (5 kali)
	for i := 1; i <= attempts; i++ {
		// Menampilkan pesan untuk meminta tebakan pengguna
		fmt.Printf("Percobaan %d: Masukkan tebakan Anda: ", i)
		// Membaca input dari pengguna dan menyimpannya ke variabel `guess`
		fmt.Scan(&guess)

		// Logika untuk membandingkan tebakan dengan angka target
		if guess < targetNumber {
			fmt.Println("Tebakan kamu terlalu kecil.") // Memberi tahu jika tebakan terlalu kecil
		} else if guess > targetNumber {
			fmt.Println("Tebakan kamu terlalu besar.") // Memberi tahu jika tebakan terlalu besar
		} else {
			// Jika tebakan benar, tampilkan pesan sukses dan keluar dari program
			fmt.Printf("Selamat! Kamu berhasil menebak angka %d dengan benar!\n", targetNumber)
			return
		}
	}

	// Jika pengguna gagal menebak dalam 5 percobaan, tampilkan pesan ini
	fmt.Printf("Kesempatan kamu habis! Angka yang benar adalah %d.\n", targetNumber)
}
