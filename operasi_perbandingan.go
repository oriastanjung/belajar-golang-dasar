package main

import "fmt"

func main() {
	a := 10 < 12
	b := 10 <= 10
	c := 2 < 0
	a1 := "orias"
	a2 := "ori"
	boleantest := a1 == a2

	bool1 := true && true
	bool2 := false || false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(boleantest)
	fmt.Println(bool1)
	fmt.Println(bool2)

}
