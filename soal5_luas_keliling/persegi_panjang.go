package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukkan Panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan Lebar: ")
	fmt.Scan(&l)

	luas := p * l
	keliling := 2 * (p + l)

	fmt.Printf("Luas Persegi Panjang = %d \n", luas)
	fmt.Printf("Keliling Persegi Panjang = %d \n", keliling)
}
