// Package models provides structs related to the 183-day rule.
package models

import (
	"fmt"
	"time"

	"github.com/reaper47/183-day-rule/internal/constants"
)

// Date holds information on a date.
type Date struct {
	Start  time.Time
	End    time.Time
	Abroad bool
}

// Verify checks whether the start date is before the end date.
func (m Date) Verify() bool {
	return m.Start.Before(m.End)
}

// IsBetween checks whether the date is between the Date's range.
func (m Date) IsBetween(date time.Time) bool {
	return date.After(m.Start) && date.Before(m.End)
}

// Print prints a date nicely to the standard output.
func (m Date) Print() {
	var isAbroad string
	if !m.Abroad {
		isAbroad = "not "
	}
	fmt.Printf("%s -> %s (%sabroad)\n", m.Start.Format(constants.Layout), m.End.Format(constants.Layout), isAbroad)
}

// NewDate creates a Date model from the given paramters.
func NewDate(start, end string, isAbroad bool) Date {
	s, _ := time.Parse(constants.Layout, start)
	e, _ := time.Parse(constants.Layout, end)
	return Date{Start: s, End: e, Abroad: isAbroad}
}
