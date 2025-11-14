package main

import "fmt"

type bookseats struct {
	//from-to-date-list of buses-Category-seat-price
	layout map[string]map[string]map[string]int
}

func redBus() {
	busSeat := make(map[string]map[string]map[string]bookseats)
	busSeat["from"] = make(map[string]map[string]bookseats)
	busSeat["from"]["to"] = make(map[string]bookseats)
	busSeat["from"]["to"]["date"] = bookseats{layout: make(map[string]map[string]map[string]int)}
	busSeat["from"]["to"]["date"].layout["list of buses"] = make(map[string]map[string]int)
	busSeat["from"]["to"]["date"].layout["list of buses"]["Category"] = make(map[string]int)
	busSeat["from"]["to"]["date"].layout["list of buses"]["Category"]["Seats"] = 2000

	fmt.Println((busSeat))
}

func main() {
	redBus()
}
