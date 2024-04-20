package main

import "fmt"

func logging() {
	fmt.Println("selesai aplikasi")
}

func runApplicationWithoutDefer() {
	fmt.Println("Aplikasi Berjalan")
	logging() // posisi di akhir
}

func runApplication() {
	defer logging() // akan di run setelah function ini selesai
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApplicationWithoutDefer()
	runApplication()
}
