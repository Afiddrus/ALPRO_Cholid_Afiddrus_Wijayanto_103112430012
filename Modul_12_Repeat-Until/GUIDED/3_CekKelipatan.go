package main

import "fmt"

func main() {
	var x, y int
	var selesai bool
	fmt.Scan(&x, &y)
	for selesai = false; !selesai; {
		x = x - y
		fmt.Println(x)
		selesai = x <= 0
	}

	fmt.Println(x == 0)
}
