//newHash.go - print all newHash numbers between 10 and 60
//Usage: go build newHash.go; ./newHash
//Copyright (c) 2020 rndMemex

package main

import "log"

func main() {
	revenue := map[string]int{
		"Jan": 11000,
		"Feb": 16000,
		"Mar": 10000,
		"Apr": 23000,
	}
	costs := map[string]int{
		"Feb": 15000,
		"Mar": 7000,
		"Apr": 5000,
		"May": 1520,
	}
	//create a new hash called profits, and use a for loop to store the profit for the overlab months Febraury to April.
	var profits = make(map[string]int)
	for k, v := range revenue {
		if _, ok := costs[k]; ok {
			profits[k] = v - costs[k]
			log.Printf("The profit for month of %s is £%d\n", k, profits[k])
		}

	}

}
