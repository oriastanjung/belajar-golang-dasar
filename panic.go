package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Berhenti")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR") // akan menghentikan aplikasi tapi defer func tetap dieksekusi
	}
}

func main() {
	runApp(false)
	runApp(true)

}
