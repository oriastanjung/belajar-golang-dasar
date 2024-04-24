package main

import "fmt"

// interface : kumpulan method yg harus dimiliki berdasarkan suatu kontrak

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong ini bisa disebut tipe data any

func printSomething(value interface{}) interface{} {
	return value
}

func printAnyThing(value any) any {
	return value
}
func main() {
	person := Person{Name: "Orias"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)

	// penggunaan inteface kosong atau any
	fmt.Println(printSomething("Helooo"))
	fmt.Println(printSomething(123))
	fmt.Println(printSomething(person))
	fmt.Println(printSomething(animal))

	fmt.Println(printAnyThing("Helooo"))
	fmt.Println(printAnyThing(123))
	fmt.Println(printAnyThing(person))
	fmt.Println(printAnyThing(animal))

}
