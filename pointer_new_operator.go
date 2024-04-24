package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

func main() {

	var alamat1 *Address = new(Address) // hanya pointer ke data kosong di memory
	alamat2 := alamat1                  // langsung menjadi pointer

	alamat2.city = "Batam"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
