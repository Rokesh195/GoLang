package main

import "fmt"

type booktrainseats struct {
	//from-to-date-Fare-sort by-list of trains-day-train-stations-125km
	layout map[string]map[string]map[string]map[string]map[string]map[string]string
}

func Train() {
	busSeat := make(map[string]map[string]booktrainseats)
	busSeat["from"] = make(map[string]booktrainseats)
	busSeat["from"]["to"] = booktrainseats{layout: make(map[string]map[string]map[string]map[string]map[string]map[string]string)}
	busSeat["from"]["to"].layout["Date"] = make(map[string]map[string]map[string]map[string]map[string]string)
	busSeat["from"]["to"].layout["Date"]["Fare"] = make(map[string]map[string]map[string]map[string]string)
	busSeat["from"]["to"].layout["Date"]["Fare"]["sort by"] = make(map[string]map[string]map[string]string)
	busSeat["from"]["to"].layout["Date"]["Fare"]["sort by"]["List of Trains"] = make(map[string]map[string]string)
	busSeat["from"]["to"].layout["Date"]["Fare"]["sort by"]["List of Trains"]["Train"] = make(map[string]string)
	busSeat["from"]["to"].layout["Date"]["Fare"]["sort by"]["List of Trains"]["Train"]["stations"] = "125km"

	fmt.Println((busSeat))
}

func main() {
	Train()
}
