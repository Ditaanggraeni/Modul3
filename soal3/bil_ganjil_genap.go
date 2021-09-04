package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("Masukkan Nilai Awal: \n")
	fmt.Scan(&a)
	fmt.Printf("Masukkan Nilai Akhir: \n")
	fmt.Scan(&b)

	fmt.Print("Read Number!!\n", "Mulai Dari Angka", a, "Sampai", b, "This Number Your Put")

	if ((a > 0) && (b > 0)) || (a < b) {
		fmt.Print("\nResult -> Deret Bilangan: \n\t\t")
		for j := a; j <= b; j++ {
			fmt.Print(j, " ")
			if j%10 == 0 {
				fmt.Printf("\n\t\t")
			}
		}
	}

	var y, i, g int
	fmt.Print("\n\t\tBilangan Ganjilnya -> ")
	for y = a; y <= b; y++ {
		i = 0
		for g = 1; g <= y; g++ {
			if y%2 != 0 {
				i++
			}
		}

		if y%2 != 0 {
			fmt.Print(y, ", ")
		}
		if y%10 == 0 {
			fmt.Printf("\n\t\t\t\t")
		}
	}
	fmt.Print("\n\t\tBilangan Genapnya -> ")
	for y = a; y <= b; y++ {
		i = 0
		for g = 1; g <= y; g++ {
			if y%g == 0 {
				i++
			}
		}

		if y%2 == 0 {
			fmt.Print(y, ", ")
		}
		if y%10 == 0 {
			fmt.Printf("\n\t\t\t\t")
		}
	}
}
