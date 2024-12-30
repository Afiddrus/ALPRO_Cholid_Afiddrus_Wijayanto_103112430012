package main

import (
	"fmt"
)

func main() {
	var h1, m1, h2, m2 int
	fmt.Println("Masukkan jam dan menit mulai parkir (h1 m1):")
	fmt.Scan(&h1, &m1)
	fmt.Println("Masukkan jam dan menit selesai parkir (h2 m2):")
	fmt.Scan(&h2, &m2)

	startMinutes := (h1%12)*60 + m1
	endMinutes := (h2%12)*60 + m2

	if endMinutes < startMinutes {
		endMinutes += 12 * 60
	}

	duration := endMinutes - startMinutes
	hours := duration / 60
	minutes := duration % 60

	fmt.Printf("%d jam %d menit\n", hours, minutes)
}
