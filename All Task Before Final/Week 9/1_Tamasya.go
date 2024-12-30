package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&N)

	const kapasitasMobil = 7

	var jumlahMobil int
	if N%kapasitasMobil == 0 {
		jumlahMobil = N / kapasitasMobil
	} else {
		jumlahMobil = (N / kapasitasMobil) + 1
	}

	fmt.Printf("Jumlah mobil yang diperlukan: %d\n", jumlahMobil)
}
