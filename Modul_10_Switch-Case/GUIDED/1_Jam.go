package main

import "fmt"

func main() {
	var jam12, jam24 int
	var label string
	fmt.Scan(&jam24)
	switch {
	case jam24 == 0:
		jam12 = 12
		label = "AM"
	case jam24 < 12 && jam24 > 0:
		jam12 = jam24
		label = "AM"
	case jam24 == 12:
		jam12 = 12
		label = "PM"
	case jam24 > 12 && jam24 < 23:
		jam12 = jam24 - 12
		label = "PM"
	default:
		fmt.Print("Tidak termasuk jam")
	}
	fmt.Print(jam12, " ", label)
}
