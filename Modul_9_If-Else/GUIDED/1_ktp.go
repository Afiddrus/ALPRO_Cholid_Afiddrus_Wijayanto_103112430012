package main

import "fmt"

func main() {
	//Program untuk mengecek apakah user bisa membuat ktp atau tidak, input terdiri dari 2 yaitu umur yang bertipe int dan juga kartu keluarga bertipe boolean, kemudian jika usia dibawah 17 tahun maka memunculkan output belum bisa membuat ktp
	var ktp int
	var kk bool

	fmt.Print("Masukkan Umur Anda:")
	fmt.Scan(&ktp)
	fmt.Print("Apakah Mempunyai KK? (true/false):")
	fmt.Scan(&kk)

	if ktp >= 17 && kk {
		fmt.Println("Bisa membuat ktp")
	} else {
		fmt.Println("Tidak bisa membuat ktp")

	}

}
