package main //Mendnadakan bahwa file ini berada di package main

import (
	"fmt" //Import modul fmt yang digunakan untuk Input/Output standar
)

func main() { //Fungsi utama yang akan dijalankan

	var x float64
	var fx float64

	fmt.Print("Masukkan Nilai X: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&x)                    //Membaca inputan user dan menyimpan ke variabel a

	fx = (2 / (x + 5)) + 5                    //Menghitung nilai f(x)
	fmt.Printf("Hasilnya adalah: %.2f\n", fx) //Menampilkan hasil jumlah ke layar
}
