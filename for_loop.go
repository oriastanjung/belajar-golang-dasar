package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// for range, untuk array slice map

	names := []string{"Eko", "Ori", "Omi", "Akin"}
	// for loop manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// dengan for range
	for index, item := range names {
		fmt.Println("index ==", index, " item >>", item)
	}

	//kita ga mau pake index nya ganti dengan _
	for _, item := range names {
		fmt.Println(" item >>", item)
	}

	type Book struct {
		Name   string
		author string
		year   int16
	}

	books := []Book{
		{Name: "Book1", author: "Oriastanjung", year: 2020},
		{Name: "Book2", author: "Karim mazakir", year: 2021},
		{Name: "Book3", author: "Usman Manul", year: 2022},
	}
	fmt.Println("No\t\tTitle\t\tAuthor\t\tYear")
	for idx, item := range books {
		fmt.Println(idx, "\t\t", item.Name, "\t\t", item.author, "\t\t", item.year)
	}
}
