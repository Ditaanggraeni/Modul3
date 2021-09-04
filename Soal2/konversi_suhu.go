package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Masukkan Nilai Celcius: ")
	fmt.Scanf("%f", &c)
	fahrenheit := (c * 1.8) + 32
	kelvin := c + 273.15

	fmt.Print("Nilai Konversi Dari Celcius Ke Fahreheit: ", fahrenheit)
	fmt.Print("\nNilai Konversi Dari Celcius Ke Kelvin: ", kelvin)

}
