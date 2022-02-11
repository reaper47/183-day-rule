// Package ioutils provides functions related to I/O.
package ioutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/reaper47/183-day-rule/internal/constants"
	"github.com/reaper47/183-day-rule/internal/models"
)

func ParseHistory() (dates []models.Date) {
	f, err := os.Open("./history.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	i := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 {
			fmt.Printf("Incorrect format for line %d: %s. Must be: [start] [end] [is abroad]. Skipping...\n", i, line)
			continue
		}

		str := "Date %s must be in the the format dd/mm/yyyy, i.e. 01/02/1992. Skipping...\n"
		if len(parts[0]) != 10 {
			fmt.Printf(str, parts[0])
			continue
		}

		if len(parts[1]) != 10 {
			fmt.Printf(str, parts[1])
			continue
		}

		date := models.NewDate(parts[0], parts[1], strings.HasPrefix(parts[2], "y"))
		if len(dates) > 0 && date.Start != dates[len(dates)-1].End {
			msg := "Start date of line %d (%s) must be the end date of line %d (%s). Exiting...\n"
			fmt.Printf(msg, i, date.Start.Format(constants.Layout), i-1, dates[len(dates)-1].End.Format(constants.Layout))
			os.Exit(1)
		}

		dates = append(dates, date)
		i++
	}

	err = s.Err()
	if err != nil {
		log.Fatalln(err)
	}
	return dates
}
