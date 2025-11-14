//memory management
//memory operations in go lang
//OOPS concepts in go lang

//1. bookmyshow
//2. where is my train
//3. amazon //array of objects
//4. whatsapp
//5. instagram
//6. swiggy || redbus
//7. truecaller

package main

import "fmt"

func principal() {
	var pointer *string = new(string)
	*pointer = "sairam"

	fmt.Println("value at pointer", *pointer)
	fmt.Println("value at pointer", pointer)
}

type seats struct {
	layout map[string]map[string]map[string]map[string]int
	// theater time seat category value
	// pvr nexus 9:30 A1 360
}

func bookmyshow() {
	movieData := make(map[string]map[string]seats)
	// 4dx      mon
	movieData["4dx"] = make(map[string]seats)

	movieData["4dx"]["mon"] = seats{layout: make(map[string]map[string]map[string]map[string]int)}

	movieData["4dx"]["mon"].layout["pvr nexus"] = make(map[string]map[string]map[string]int)

	movieData["4dx"]["mon"].layout["pvr nexus"]["9:30"] = make(map[string]map[string]int)

	movieData["4dx"]["mon"].layout["pvr nexus"]["9:30"]["A1"] = make(map[string]int)

	movieData["4dx"]["mon"].layout["pvr nexus"]["9:30"]["A1"]["prime"] = 1

	fmt.Println(movieData)
}

func main() {
	bookmyshow()
	principal()
}
