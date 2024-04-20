package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Berhenti")
	message := recover() // harus di defer function klo mau recover, karena func ini akan selalu dijalankan
	fmt.Println("Terjadi erorr : ", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")

	}
}

func main() {
	runApp(true)
	fmt.Println("Bisa berjalan lagi") // akan tampil meski sudah error
}
