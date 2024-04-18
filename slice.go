package main

import "fmt"

func main() {
	names := [...]string{"ori", "omi", "ate", "joko", "budi", "eko"}
	slice := names[4:6]
	fmt.Println(slice)

	slice1 := names[:6]
	fmt.Println(slice1)
	slice2 := names[3:]
	fmt.Println(slice2)
	var slice3 []string = names[:] //slice sperti ni aslinya, tak perlu kasih capacity
	fmt.Println(slice3)

	panjangSlice := len(slice)
	capacitySlice := cap(slice)
	fmt.Println(panjangSlice)
	fmt.Println(capacitySlice)
	slice[0] = "ori-baru" // maka data names array akan berubah
	fmt.Println(names)

	newSlice := append(slice, "ate-baru") // akan mnambah data ke slice, tidak ke array names
	fmt.Println(newSlice)
	fmt.Println(names)

	var newSliceMake []string = make([]string, 2, 5)
	newSliceMake[0] = "12"
	newSliceMake[1] = "14"
	fmt.Println(newSliceMake)
	fmt.Println(len(newSliceMake))
	fmt.Println(cap(newSliceMake))
	// newSLiceMake[2] = 15 error karena kita tentuin lengthnya 2
	// harus menggunakan append
	appendedData := append(newSliceMake, "27", "28", "29", "30")
	// jika data melebhihi capacity maka capacity otomatis x2
	fmt.Println(appendedData)
	fmt.Println(len(appendedData))
	fmt.Println(cap(appendedData))

	fromSlice := names[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{4, 5, 6}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
