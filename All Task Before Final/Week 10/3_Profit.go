package main

import "fmt"

func main() {
	var prevProfit, currProfit float64
	fmt.Println("Masukkan keuntungan bulan sebelumnya:")
	fmt.Scan(&prevProfit)
	fmt.Println("Masukkan keuntungan bulan ini:")
	fmt.Scan(&currProfit)

	change := currProfit - prevProfit

	if change > 0 {
		fmt.Printf("Peningkatan sebesar %.2f\n", change)
	} else if change < 0 {
		fmt.Printf("Penurunan sebesar %.2f\n", -change)
	} else {
		fmt.Println("Tetap")
	}
}
