package main

import "fmt"

func main() {
	var siswa [3]string
	siswa[0] = "orias"
	siswa[1] = "tanjung"
	siswa[2] = "riastan"
	fmt.Println(siswa)

	var values = [3]int{
		10, 20, 30,
	}
	values[2] = 24
	fmt.Println(values)
	panjangArray := len(values)

	fmt.Println(panjangArray)

	// bisa kyk gini juga

	var books = [...]string{
		"halo",
		"world",
	}
	books[1] = "news"
	fmt.
}
