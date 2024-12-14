package main

import "fmt"

type CekPrima struct {
	bilangan int
}

func NewCekPrima(bilangan int) CekPrima {
	return CekPrima{bilangan: bilangan}
}

func (pp CekPrima) ApakahPrima() bool {
	if pp.bilangan <= 1 {
		return false
	}
	for i := 2; i*i <= pp.bilangan; i++ {
		if pp.bilangan%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Bilangan bulat harus lebih besar dari 0.")
		return
	}

	pengecek := NewCekPrima(n)
	if pengecek.ApakahPrima() {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
