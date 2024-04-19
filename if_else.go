package main

import "fmt"

func main() {
	name := "oriast"

	if name == "ori" {
		fmt.Println("nama nya ori")
	} else if name == "orias" {
		fmt.Println("nama adalah orias")
	} else {
		fmt.Println("bukan ori")
	}

	//short statement
	if length := len(name); length >= 5 {
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("nama sudah cukup")
	}

	length := 100

	fmt.Println(length) // bisa buat variable length lagi
}
