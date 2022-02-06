package main

import (
	"github.com/reaper47/travel/internal/input"
	"github.com/reaper47/travel/internal/models"
)

func main() {
	history := []models.Date{
		models.NewDate("10/01/2021", "01/05/2021", false),
		models.NewDate("01/05/2021", "27/07/2021", true),
		models.NewDate("27/07/2021", "11/12/2021", false),
		models.NewDate("11/12/2021", "08/02/2022", true),
	}
	future := input.AskDates(history)

	dates := models.NewDates(history, future)
	rule := dates.Calc183Rule()
	rule.Print()
}
