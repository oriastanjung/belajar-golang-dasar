package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

// blum pakai passByReference atau pointer

func changeCountryToIndonesia(addr Address) {
	addr.country = "Indonesia"
}

// pakai passByReference atau pointer

func changeCountryToIndonesiaByPointer(addr *Address) {
	addr.country = "Indonesia"
}

func main() {
	address := Address{"Hong Kong", "China", "China"}

	changeCountryToIndonesia(address) // address tidak berubah
	fmt.Println(address)

	changeCountryToIndonesiaByPointer(&address) // address berubah
	fmt.Println(address)

	//atau
	pointerAddress := &address
	changeCountryToIndonesiaByPointer(pointerAddress)
	fmt.Println(address)
}
