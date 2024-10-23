
package main

import "fmt"

func main() {
	var qirat int
	fmt.Scan(&qirat)

	dinar := qirat / 600     // 1 dinar = 600 qirat
	sisaQirat := qirat % 600 // sisa qirat setelah dikonversi ke dinar

	dirham := sisaQirat / 60   // 1 dirham = 60 qirat
	sisaQirat = sisaQirat % 60 // sisa qirat setelah dikonversi ke dirham

	fals := sisaQirat / 6 // 1 fals = 6 qirat
	qirat = sisaQirat % 6 // sisa qirat setelah dikonversi ke fals

	fmt.Printf("%d %d %d %d\n", dinar, dirham, fals, qirat)
}
