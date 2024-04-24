// package itu nama folder nya
package helper

import "fmt"

// ga wajib pakai func main() karena bukan di package main

func ConvertToCurrency(number int) string { // jika diawal pakai kapital huruf, maka lgsg di ekspor
	currrency := fmt.Sprintf("Rp%d", number)
	return currrency
}
