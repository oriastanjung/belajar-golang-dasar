package main

import "fmt"

func main() {
	name := "orias"

	switch name {
	case "orias":
		fmt.Println("nama orias")

	case "ori":
		fmt.Println("nama ori")

	default:
		fmt.Println("nama ga ada")
	}

	//bisa short statement juga

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama kepanjangan")
	default:
		fmt.Println("nama cukup")
	}

	length := 100
	fmt.Println(length)
}
