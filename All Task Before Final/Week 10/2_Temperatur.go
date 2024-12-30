package main

import "fmt"

func main() {
	var temp1, temp2, temp3, temp4, temp5 float64
	fmt.Println("Masukkan 5 temperatur:")

	fmt.Scan(&temp1)
	fmt.Scan(&temp2)
	fmt.Scan(&temp3)
	fmt.Scan(&temp4)
	fmt.Scan(&temp5)

	isAscending, isDescending := true, true

	if temp1 < temp2 {
		isDescending = false
	} else if temp1 > temp2 {
		isAscending = false
	}

	if temp2 < temp3 {
		isDescending = false
	} else if temp2 > temp3 {
		isAscending = false
	}

	if temp3 < temp4 {
		isDescending = false
	} else if temp3 > temp4 {
		isAscending = false
	}

	if temp4 < temp5 {
		isDescending = false
	} else if temp4 > temp5 {
		isAscending = false
	}

	switch {
	case isAscending:
		fmt.Println("Stabil naik")
	case isDescending:
		fmt.Println("Stabil turun")
	default:
		fmt.Println("Tidak stabil")
	}
}
