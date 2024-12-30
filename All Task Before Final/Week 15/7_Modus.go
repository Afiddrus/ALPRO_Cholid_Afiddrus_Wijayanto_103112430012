package main

import "fmt"

func main() {
	var x int
	var input int
	var modus int

	var countX int = 0 // frekuensi kemunculan x
	var count0 int = 0 // frekuensi kemunculan 0

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	countX++

	fmt.Println("Masukkan 9 bilangan berikutnya:")
	for i := 0; i < 9; i++ {
		fmt.Scan(&input)
		// Hitung frekuensi
		if input == x {
			countX++
		} else if input == 0 {
			count0++
		}
	}

	// Jika frekuensi x lebih besar dari 0, maka x adalah modus
	if countX > count0 {
		modus = x
	} else {
		// Jika frekuensi 0 lebih besar atau sama dengan x, maka 0 adalah modus
		modus = 0
	}

	fmt.Printf("Modus = %d\n", modus)
}
