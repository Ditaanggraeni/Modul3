package main

import (
	"fmt"
	"strings"
)

func main() {
	var kalimat string
	var jumlah, jmlh int

	fmt.Print("Input text: ")
	fmt.Scanf("%s", &kalimat)

	jumlah = strings.Count(kalimat, "") - 1
	jmlh = strings.Count(kalimat, "a")

	fmt.Printf("Jumlah Karakter: %d \n", jumlah)
	fmt.Printf("Jumlah Karakter Yang Sama: %d", jmlh)
}
