package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	// var resultString string = result.(string) // konversi dari any ke string
	// fmt.Println(resultString)

	// var resultInt = result.(int) // akan error dan panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		{
			fmt.Println("String >> ", value)
		}
	case int:
		{
			fmt.Println("Int >> ", value)
		}
	default:
		{
			fmt.Println("Unknown data type")
		}
	}

}
