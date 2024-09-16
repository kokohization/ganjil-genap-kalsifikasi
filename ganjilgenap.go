package main

import "fmt"

func main() {

	var angka int16
	var InfoGanjilGenap string
	var HabisTerbagiDua bool

	fmt.Print("Input Angka: ")
	fmt.Scanf("%d", &angka)

	HabisTerbagiDua = (angka%2 == 0)

	if HabisTerbagiDua == true {
		InfoGanjilGenap = "Genap"

	} else {
		InfoGanjilGenap = "Ganjil"
	}

	fmt.Println(InfoGanjilGenap)

}
