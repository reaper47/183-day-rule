package main

import (
	"github.com/reaper47/183-day-rule/internal/input"
	"github.com/reaper47/183-day-rule/internal/ioutils"
	"github.com/reaper47/183-day-rule/internal/models"
)

func main() {
	history := ioutils.ParseHistory()
	future := input.AskDates(history)

	dates := models.NewDates(history, future)
	dates.Print()

	rule := dates.Calc183Rule()
	rule.Print()
}
