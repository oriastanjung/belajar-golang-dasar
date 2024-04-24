package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

func updateNumber(val int) {
	val = 10
	fmt.Println("updated to 10")
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // pass by value klo seperti ini

	address2.city = "Bogor"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2)

	numb1 := 15
	fmt.Println(numb1)
	updateNumber(numb1) // tetap tidak ngupdate
	fmt.Println(numb1)

	// dengan pointer kita bisa pass by reference atau pass dari adress memory
	address3 := &address1 // pointer ke address1
	address3.city = "Bandung"
	fmt.Println(address1) // data ikut berubah
	fmt.Println(address3) // data ikut berubah

	// atau bisa denggan

	var address4 *Address = &address1
	address4.city = "Depok"
	fmt.Println(address1) // data ikut berubah
	fmt.Println(address4) // data ikut berubah
}
