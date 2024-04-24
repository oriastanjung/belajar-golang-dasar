package main

import "fmt"

// func Contoh(name string) string { error karena nil bsa di gunakan di interface, map, slice, channel, pointer dll, bukan untuk satu return value data
// 	if name == ""{
// 		return nil
// 	}else{
// 		return name
// 	}
// }

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": "orias",
		}
	}
}

func main() {
	data := newMap("orias")
	if data == nil {
		fmt.Println("data nil")
	} else {
		fmt.Println(data["name"])
	}
}
