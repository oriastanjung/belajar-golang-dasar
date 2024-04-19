package main

import "fmt"

type Blacklist func(string) bool

func registerName(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// anonymous func yaitu function tanpa perlu di deklarasi
	registerName("Orias", func(name string) bool {
		if name == "Anjing" {
			return true
		} else {
			return false
		}
	})
	registerName("Ekoo", func(name string) bool {
		if name == "Ekoo" {
			return true
		} else {
			return false
		}
	})
}
