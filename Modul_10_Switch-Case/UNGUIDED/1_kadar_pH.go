package main

import (
	"fmt"
)

func main() {
	var ph float64

	fmt.Print("Masukkan kadar pH pada air: ")
	fmt.Scan(&ph)

	switch {
	case ph < 0 || ph > 14:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak minum")
	default:
		fmt.Println("Air tidak layak minum")
	}
}
