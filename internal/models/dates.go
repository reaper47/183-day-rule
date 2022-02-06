package models

type dates struct {
	Values []Date
}

// Calc183Rule calculates the 183-day rule.
func (m dates) Calc183Rule() rule183Results {
	m.adjustDates()
	return m.calc183Rule()
}

func (m *dates) adjustDates() {
	first := m.Values[len(m.Values)-1].End.AddDate(-1, 0, 0)

	var dates []Date
	var isBetweenFound bool

	for _, d := range m.Values {
		if !isBetweenFound && (d.IsBetween(first) || first == d.Start) {
			d.Start = first
			isBetweenFound = true
		}

		if isBetweenFound {
			dates = append(dates, d)
		}
	}

	m.Values = make([]Date, len(dates))
	copy(m.Values, dates)
}

func (m dates) calc183Rule() rule183Results {
	var daysIn, daysOut int
	for _, d := range m.Values {
		days := int(d.End.Sub(d.Start).Hours() / 24)
		if d.Abroad {
			daysOut += days
		} else {
			daysIn += days
		}
	}

	return rule183Results{
		DaysInCountry: daysIn,
		DaysAbroad:    daysOut,
		IsPlanOK:      daysIn >= 183,
	}
}

// NewDates creates a new dates struct containing all dates provided.
func NewDates(allDates ...[]Date) dates {
	var d dates
	for _, date := range allDates {
		d.Values = append(d.Values, date...)
	}
	return d
}
