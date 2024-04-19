package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}

// menggunakan type declaratiion agar parameter func mudah dibaca

type Filter func(string) string

func sayHelloWithFilter1(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func main() {
	// membuat function sebagai value

	goodbye1 := getGoodBye
	goodbye2 := getGoodBye
	fmt.Println(goodbye1("Orias"))
	fmt.Println(goodbye2("Oriastanjung"))

	// function sebagai parameter

	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Orias", spamFilter)

	sayHelloWithFilter1("Anjing", spamFilter)
}
