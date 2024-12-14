package main

import "fmt"

func main() {
	var bilangan, j int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	for j = 1; j <= bilangan; j += 1 {
		if j%2 != 0 {
			fmt.Println(j, " ")
		}
	}
}
