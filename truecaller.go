package main

import "fmt"

type Data struct {
	logs      []string
	spamScore int
}

type calls struct {
	layout map[string]map[string]Data
	//tag - location - spam score - logs
}

func truecaller() {
	callData := make(map[string]map[string]calls)

	callData["9042175705"] = make(map[string]calls)

	callData["9042175705"]["Rokesh"] = calls{layout: make(map[string]map[string]Data)}

	callData["9042175705"]["Rokesh"].layout["tag"] = make(map[string]Data)

	callData["9042175705"]["Rokesh"].layout["tag"]["location"] = Data{
		logs:      []string{"09:45", "10:05"},
		spamScore: 0,
	}

	fmt.Println(callData)
}

func main() {
	truecaller()
}
