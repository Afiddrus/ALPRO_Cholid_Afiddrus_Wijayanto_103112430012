package main

import (
	"fmt"
)

func main() {
	var s, minggu int

	fmt.Print("Masukkan jumlah hari (s): ")
	fmt.Scan(&s)

	minggu = s / 7
	if s%7 == 0 {
		fmt.Printf("Minggu ke-: %d\n", minggu)
	} else {
		fmt.Printf("Minggu ke-: %d\n", minggu+1)
	}
}
