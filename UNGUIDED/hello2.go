package main

import "fmt"

func main() {
	var mk string = "Algoritma dan Pemrograman"
	var kode, sks int
	fmt.Print("Tuliskan kode MK dan SKS: ")
	fmt.Scan(&kode, &sks)
	fmt.Println("Kredit MK", kode, "-", mk, "1 adalah", sks, "SKS")
}
