package main //Mendnadakan bahwa file ini berada di package main

import (
	"fmt" //Import modul fmt yang digunakan untuk Input/Output standar
)

func main() { //Fungsi utama yang akan dijalankan

	var a, b, c, d, e int

	fmt.Print("Masukkan Nilai a: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&a)                    //Membaca inputan user dan menyimpan ke variabel a

	fmt.Print("Masukkan Nilai b: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&b)                    //Membaca inputan user dan menyimpan ke variabel b

	fmt.Print("Masukkan Nilai c: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&c)                    //Membaca inputan user dan menyimpan ke variabel c

	fmt.Print("Masukkan Nilai d: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&d)                    //Membaca inputan user dan menyimpan ke variabel d

	fmt.Print("Masukkan Nilai e: ") //Menampilkan isi variabel nama ke layar
	fmt.Scan(&e)                    //Membaca inputan user dan menyimpan ke variabel e

	jumlah := a + b + c + d + e                          //Menghitung jumlah semua nilai
	fmt.Printf("Hasil Penjumlahan adalah: %d\n", jumlah) //Menampilkan hasil jumlah ke layar
}
