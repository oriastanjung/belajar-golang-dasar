package main

import "fmt"

type Man struct {
	Name string
}

// belum pakai pointer

func (man Man) married() {
	man.Name = "Mr. " + man.Name
}

//  pakai pointer

func (man *Man) marriedWithPointer() { // otomatis di atur oleh go lang
	man.Name = "Mr. " + man.Name
}

func main() {
	orias := Man{"Orias"}
	orias.married() // tidak berubah nama kita
	fmt.Println(orias)
	orias.marriedWithPointer() //  berubah nama kita
	fmt.Println(orias)
}
