package main

import (
	"fmt"
)

func jumlahFaktor(bilangan int) int {
	jumlah := 0
	for i := 1; i <= bilangan/2; i++ {
		if bilangan%i == 0 {
			jumlah += i
		}
	}
	return jumlah
}

func cekBilanganSempurna(bilangan int) bool {
	if bilangan <= 0 {
		return false
	}
	return jumlahFaktor(bilangan) == bilangan
}

func printFaktor(bilangan int) {
	fmt.Printf("Faktor-faktornya adalah: ")
	isFirst := true
	for i := 1; i <= bilangan/2; i++ {
		if bilangan%i == 0 {
			if isFirst {
				fmt.Printf("%d", i)
				isFirst = false
			} else {
				fmt.Printf(", %d", i)
			}
		}
	}
	fmt.Println()
}

func main() {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	sempurna := cekBilanganSempurna(bilangan)

	if sempurna {
		fmt.Printf("%d adalah bilangan sempurna\n", bilangan)
		printFaktor(bilangan)
		fmt.Printf("Jumlah faktor-faktornya: %d\n", jumlahFaktor(bilangan))
	} else {
		fmt.Printf("%d bukan bilangan sempurna\n", bilangan)
		printFaktor(bilangan)
		fmt.Printf("Jumlah faktor-faktornya: %d\n", jumlahFaktor(bilangan))
	}
}
