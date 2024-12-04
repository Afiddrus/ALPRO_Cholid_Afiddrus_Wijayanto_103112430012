package main

import "fmt"

func main() {
	var number int
	var countinueLoop bool
	for countinueLoop = true; countinueLoop; {
		fmt.Scan(&number)
		countinueLoop = number <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", number)
}
