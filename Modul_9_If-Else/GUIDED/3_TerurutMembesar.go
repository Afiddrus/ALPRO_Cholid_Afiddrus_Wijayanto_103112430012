package main

import "fmt"

func main() {
	//Program untuk menentukan apakah digit dalam suatu bilangan terurut membesar, mengecil, atau tidak terurut. Bilangan hanya terdiri dari empat digit saja atau lebih besar atau sama dengan 1000 dan lebih kecil  atau sama dengan 9999.
	var bilangan, digit1, digit2, digit3, digit4 int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	if bilangan >= 1000 && bilangan <= 9999 {
		digit1 = bilangan / 1000
		digit2 = bilangan % 1000 / 100
		digit3 = bilangan % 100 / 10
		digit4 = bilangan % 10
		if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
			fmt.Println("Digit pada bilangan", bilangan, "terurut membesar")
		} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
			fmt.Println("Digit pada bilangan", bilangan, "terurut mengecil")
		} else {
			fmt.Println("Digit pada bilangan", bilangan, "tidak terurut")
		}
	} else {
		fmt.Println("Bilangan harus memiliki empat digit atau lebih besar atau sama dengan 1000 dan kecil sama atau sama dengan 9999.")
	}
}
