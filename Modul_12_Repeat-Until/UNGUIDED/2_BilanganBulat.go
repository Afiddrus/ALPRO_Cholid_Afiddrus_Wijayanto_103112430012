package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	target := int(input)
	if float64(target) < input {
		target++
	}

	sum := 0.0
	selesai := false

	for selesai == false {
		sum += 0.1 // Penambahan 0.1 setiap iterasi
		fmt.Printf("%.1f ", sum)
		selesai = int(sum) >= target
	}
	fmt.Println(target)
}
