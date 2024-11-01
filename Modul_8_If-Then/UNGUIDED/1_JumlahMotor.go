package main

// package fmt untuk fungsi input dan output
import (
	"fmt"
)

func main() {
	var jumlahOrang, jumlahMotor int // deklarasi variabel jumlahOrang dan jumlahMotor bertipe integer (bilangan bulat)

	// Memasukkan jumlah orang
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlahOrang)

	// Menghitung jumlah motor yang diperlukan
	if jumlahOrang%2 == 0 {
		jumlahMotor = jumlahOrang / 2
	}
	if jumlahOrang%2 != 0 {
		jumlahMotor = jumlahOrang/2 + 1
	}

	// Menampilkan hasil jumlah motor yang diperlukan
	fmt.Println("Jumlah motor yang diperlukan:", jumlahMotor)
}
