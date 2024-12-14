package main

import "fmt"

type HitungGanjil struct {
	limit int
}

func NewHitungGanjil(limit int) HitungGanjil {
	return HitungGanjil{limit: limit}
}

func (oc HitungGanjil) TotalGanjil() int {
	jumlah := 0
	for i := 1; i <= oc.limit; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Bilangan bulat harus lebih besar dari 0.")
		return
	}

	hitungGanjil := NewHitungGanjil(n)
	hasil := hitungGanjil.TotalGanjil()

	fmt.Printf("Terdapat %d bilangan ganjil\n", hasil)
}
