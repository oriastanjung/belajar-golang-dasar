package main

import "fmt"

func sayHello() {
	fmt.Println("Hallllooooo")
}

func sayName(firstname string, lastName string) {
	fmt.Println(firstname, lastName)
}

func sum2number(a int, b int) int {
	return a + b
}

func getFullName() (string, string) {
	return "orias", "tanjung"
}

func getNumbers() (number1 int, number2 float32, number3 string) {
	number1 = 12
	number2 = 3.14
	number3 = "27"
	return number1, number2, number3

}

func sumAll(argsNumbers ...int) int {
	total := 0

	for _, number := range argsNumbers {
		total += number
	}
	return total
}

type Book struct {
	name   string
	year   int
	author string
}

func filterBooks(keyword string, books ...Book) []Book {
	result := []Book{}

	for _, item := range books {
		if keyword == item.author {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	sayHello()
	sayName("Orias", "Tanjung")
	sayName("EKo", "Kennedy")

	a := 5
	b := 10
	result := sum2number(a, b)
	fmt.Println(result)

	firstname, lastname := getFullName()
	fmt.Println(firstname)
	fmt.Println(lastname)

	// kita ga mau ambil sesuatu dari multiple return value

	_, lastname1 := getFullName()
	fmt.Println(lastname1)

	x, y, z := getNumbers()

	fmt.Println(x, y, z)

	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	//slice sebagai parameter

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	total = sumAll(numbers...) // kita passing suatu slice sebagai params

	books := []Book{
		{name: "Book1", author: "Oriastanjung", year: 2020},
		{name: "Book121", author: "Oriastanjung", year: 2020},
		{name: "Book2", author: "Karim mazakir", year: 2021},
		{name: "Book233", author: "Karim mazakir", year: 2021},
		{name: "Book3", author: "Usman Manul", year: 2022},
		{name: "Book213", author: "Usman Manul", year: 2022},
	}
	filteredBooks := filterBooks("Oriastanjung", books...)
	fmt.Println("No\t\tTitle\t\tAuthor\t\tYear")
	for idx, item := range filteredBooks {
		fmt.Println(idx, "\t\t", item.name, "\t\t", item.author, "\t\t", item.year)
	}
}
