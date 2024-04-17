package main

import "fmt"

func main() {
	type NoKTP string

	var ktpOrias NoKTP = "21720420129124"
	fmt.Println(ktpOrias)

	var ktpOri2 = NoKTP(107) // test konversi int to NoKTP
	fmt.Println(ktpOri2)
}
