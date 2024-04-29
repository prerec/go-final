package handler

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func timeValidate(date string) error {
	if _, err := time.Parse(timeLayout, date); err != nil {
		return err
	}
	return nil
}

func getNextDate(now time.Time, date, repeat, timeLayout string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat cannot be empty")
	}

	if strings.Contains(repeat, "d ") {
		days, err := strconv.Atoi(strings.Trim(repeat, "d "))
		if err != nil {
			return "", err
		}
		if days > 400 {
			return "", errors.New("days cannot be greater than 400")
		}
		parseDate, err := time.Parse(timeLayout, date)
		if err != nil {
			return "", errors.New("invalid date." + err.Error())
		}

		newDate := parseDate.AddDate(0, 0, days)

		for newDate.Before(now) {
			newDate = newDate.AddDate(0, 0, days)
		}

		return newDate.Format(timeLayout), nil

	} else if repeat == "y" {
		parseDate, err := time.Parse(timeLayout, date)
		if err != nil {
			return "", errors.New("invalid date." + err.Error())
		}

		newDate := parseDate.AddDate(1, 0, 0)

		for newDate.Before(now) {
			newDate = newDate.AddDate(1, 0, 0)
		}
		return newDate.Format(timeLayout), nil
	} else {
		return "", errors.New("invalid repeat value")
	}
}
