package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"India":     188181,
		"Pakisthan": 837478,
		"SriLanka":  8775744,
	}
	statePopulation["USA"] = 113312
	fmt.Println("Recent", statePopulation["USA"])
	fmt.Println(statePopulation)
}
