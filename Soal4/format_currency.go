package main

import (
	"fmt"

	"github.com/paimanbandi/rupiah"
)

func main() {
	var uang int
	fmt.Print("Masukkan Jumlah Uang : ")
	fmt.Scan(&uang)
	usd := 0.00007 * float32(uang)
	fmt.Println("Rupiah :", rupiah.FormatFloat64ToRp(float64(uang)))
	fmt.Println("USD :", "$", usd)
}
