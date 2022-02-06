// Package models provides structs related to the 183-day rule.
package models

import (
	"time"
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

// NewDate creates a Date model from the given paramters.
func NewDate(start, end string, isAbroad bool) Date {
	s, _ := time.Parse("02/01/2006", start)
	e, _ := time.Parse("02/01/2006", end)
	return Date{Start: s, End: e, Abroad: isAbroad}
}
