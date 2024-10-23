package main

import "fmt" // Import package fmt untuk input dan output

func main() {
	// Deklarasi variabel N sebagai integer
	var N int

	// Input nilai N
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	// Looping dari bilangan 1 sampai ke N
	for i := 1; i <= N; i++ {
		if i == N { // Jika i sama dengan N
			fmt.Print(i * i) // Cetak kuadrat i tanpa koma
		} else {
			fmt.Print(i*i, ",") // Cetak kuadrat i dengan koma
		}
	}
}
