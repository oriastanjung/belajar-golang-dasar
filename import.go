package main

// import namaModule/namaPackage
import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	number := 10000
	currency := helper.ConvertToCurrency(number)
	fmt.Println(currency)
}
