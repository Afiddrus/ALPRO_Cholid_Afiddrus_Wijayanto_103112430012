package main

import "fmt"

func main() {
	var n, j int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	j = n

	for j > 1 { //j =5
		fmt.Print(j, " x ")
		j = j - 1 //j--
	}
	fmt.Println(1)
}
