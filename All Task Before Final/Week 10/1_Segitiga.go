package main

import "fmt"

func main() {
	var input1, input2, input3 int

	fmt.Print("Masukkan nilai sisi 1: ")
	fmt.Scan(&input1)
	fmt.Print("Masukkan nilai sisi 2: ")
	fmt.Scan(&input2)
	fmt.Print("Masukkan nilai sisi 3: ")
	fmt.Scan(&input3)

	if input1 == input2 && input2 == input3 {
		fmt.Println("Segitiga Sama Sisi")
	} else if input1 == input3 {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
