package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct juga bisa memiliki method function
// func (rujukan struct) namaFunc() {}
func (customer Customer) sayHello() {
	fmt.Println(customer.Name + " Mengucapkan Halo!")
}
func main() {

	var orias Customer
	orias.Name = "orias"
	orias.Address = "indonesia"
	orias.Age = 29

	// atau

	customer1 := Customer{
		Name:    "customer1",
		Address: "Indonesia",
		Age:     10,
	}
	fmt.Println(customer1)

	// atau

	customer2 := Customer{"Customer2", "Indonesia", 27} // urutan harus sesuai di struct nya
	fmt.Println(customer2)

	orias.sayHello()
	customer1.sayHello()
	customer2.sayHello()
}
