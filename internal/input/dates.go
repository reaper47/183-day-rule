package input

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/reaper47/183-day-rule/internal/constants"
	"github.com/reaper47/183-day-rule/internal/models"
)

// AskDates asks the user to input all travel dates.
func AskDates(history []models.Date) []models.Date {
	var dates []models.Date
	if !isQuestion("Add dates?") {
		return dates
	}
	fmt.Println()

	lastDate := history[len(history)-1]
	i := 1

	for {
		fmt.Println("Date #" + strconv.Itoa(i) + " (" + lastDate.End.Format(constants.Layout) + " to dd/mm/yyyy):")

		dates = append(dates, askNewDate(lastDate))

		if !isQuestion("Do you want to add another date?") {
			break
		}

		i++
		lastDate = dates[len(dates)-1]
		fmt.Println()
	}

	fmt.Println()
	return dates
}

func askNewDate(lastDate models.Date) models.Date {
	for {
		d := models.Date{
			Start:  lastDate.End,
			End:    askDate("Enter the end date:"),
			Abroad: askIsAbroad(),
		}

		if d.Verify() {
			return d
		}

		if !isQuestion("Start date must be before end date. Retry?") {
			os.Exit(1)
		}
		fmt.Println()
	}
}

func askDate(msg string) time.Time {
	for {
		fmt.Printf(msg + " ")

		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			continue
		}

		parts := strings.SplitN(str, "/", 3)
		if len(parts) != 3 {
			continue
		}

		day := parts[0]
		if len(day) != 2 {
			day = "0" + string(day[0])
		}

		month := parts[1]
		if len(month) != 2 {
			month = "0" + string(month[0])
		}

		str = strings.Join([]string{day, month, parts[2]}, "/")

		t, err := time.Parse(constants.Layout, str)
		if err == nil {
			return t
		}
		fmt.Println(err)
	}
}

func askIsAbroad() bool {
	for {
		fmt.Printf("Were you abroad during this time? [y/n]: ")

		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			continue
		}

		if len(str) > 0 {
			return strings.ToLower(str)[0] == 'y'
		}
	}

}
