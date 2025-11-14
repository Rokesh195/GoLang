package main

import "fmt"

type product struct {
	//Category - varieties - devices - variant - price
	layout map[string]map[string]map[string]int
}

func amazon() {
	productData := make(map[string]product)

	productData["Electronics"] = product{layout: make(map[string]map[string]map[string]int)}

	productData["Electronics"].layout["phone"] = make(map[string]map[string]int)

	productData["Electronics"].layout["phone"]["iphone"] = make(map[string]int)

	productData["Electronics"].layout["phone"]["iphone"]["128gb"] = 75000

	fmt.Println(productData)
}

func main() {
	amazon()
}
