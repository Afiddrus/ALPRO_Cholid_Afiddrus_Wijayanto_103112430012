package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Print("Masukkan jumlah gula (n), jumlah kopi (m), gula per cangkir (x), kopi per cangkir (y): ")
	fmt.Scanln(&n, &m, &x, &y)

	jumlahCangkir := 0

	for n >= x && m >= y {
		n -= x
		m -= y
		jumlahCangkir++
	}

	fmt.Printf("Banyaknya cangkir kopi yang berhasil dibuat: %d\n", jumlahCangkir)
}
