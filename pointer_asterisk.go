package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	address2.city = "Depok"
	fmt.Println(address1)
	fmt.Println(address2)
	// coba buat data baru
	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // data pertama tidak ikut berubah ketika adress2 kita arahkan ke pointer reference baru
	fmt.Println(address2)

	// kita gunakan asterisk * pointer biar merubah datanya juga
	address2 = &address1
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// maka address1 valuenya berubah juga ke nilai baru
	fmt.Println(address1) // berubah juga
	fmt.Println(address2)

}
