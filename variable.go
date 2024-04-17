package main

import "fmt"

func main() {
	var name string

	name = "Orias"
	fmt.Println(name)

	name = "oriastanjung"
	fmt.Println(name)

	var fullname = "O. Riastanjung"
	fmt.Println(fullname)

	nickname := "orias"
	fmt.Println(nickname)

	nickname = "ori"
	fmt.Println(nickname)

	var (
		firstname = "oriiii"
		lastname  = "tanjung"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	const phi = 3.14
	fmt.Println(phi)
	// phi=2.12 error

}
