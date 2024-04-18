package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}
	person["name"] = "oriastanjung"
	fmt.Println(person)

	books := map[string]string{
		"author": "Oriastanjung",
	}
	books["year"] = "2020"
	fmt.Println(books)
	fmt.Println(books["author"])

	totalKeysData := len(books)
	fmt.Println(totalKeysData)
	newBooks := make(map[string]string)
	newBooks["author"] = "newOrias"
	newBooks["year"] = "2024"
	fmt.Println(newBooks)
	totalKeysData1 := len(newBooks)
	fmt.Println(totalKeysData1)
	delete(newBooks, "year")
	fmt.Println(newBooks)

}
