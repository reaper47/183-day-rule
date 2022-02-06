package models

import "fmt"

type rule183Results struct {
	DaysInCountry int
	DaysAbroad    int
	IsPlanOK      bool
}

// Print prints 183-day rule results nicely.
func (m rule183Results) Print() {
	fmt.Println()
	fmt.Println("Rule 183 results:")
	fmt.Println("Days in the country:", m.DaysInCountry)
	fmt.Println("Days abroad:", m.DaysAbroad)

	isOK := "yes"
	if !m.IsPlanOK {
		isOK = "no"
	}
	fmt.Println("Is the rule respected?", isOK)
}
