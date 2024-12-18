package main

import "fmt" // Package fmt unuk fungsi input dan output

// Fungsi utama
func main() {
	var n_103112430012 int // Deklarasi variabel n dengan tipe data int (bilangan bukat) yang nantinya akan digunakan sebagai iterasi untuk pola X

	// Meminta nilai n dari user
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n_103112430012)

	// Perulangan untuk membuat pola x dari 1 sampai dengan n
	for i := 1; i <= n_103112430012; i++ { // Sebuah perulangan dengan inisialisasi i dan akan terus berulang hingga i lebih besar dari n yang berfungsi untuk membuat pola x dari 1 sampai dengan n
		for j := 1; j <= n_103112430012; j++ { //  Sebuah perulangan dengan inisialisasi j dan akan terus berulang hingga j lehih besar dari n yang berfungsi untuk membuat pola x dari 1 sampai dengan n
			if i == j || i+j == n_103112430012+1 { // Sebuah percabangan dengan kondisi jika i sama dengan j atau i ditambah j sama dengan n + 1
				fmt.Print(i) // Maka akan menampilkan i yaitu bilangan dari 1 hingga ke n
			} else { // Memasuki kondisi else karena kondisi if tidak terpenuhi
				fmt.Print(" ") // Memberi spasi agar membentuk pola X
			}
		}
		fmt.Println() // Mencetak hasil
	}
}
