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
		parsedDate, err := time.Parse(timeLayout, date)
		if err != nil {
			return "", errors.New("invalid date." + err.Error())
		}

		newDate := parsedDate.AddDate(0, 0, days)

		for newDate.Before(now) {
			newDate = newDate.AddDate(0, 0, days)
		}

		return newDate.Format(timeLayout), nil

	} else if repeat == "y" {
		parsedDate, err := time.Parse(timeLayout, date)
		if err != nil {
			return "", errors.New("invalid date." + err.Error())
		}

		newDate := parsedDate.AddDate(1, 0, 0)

		for newDate.Before(now) {
			newDate = newDate.AddDate(1, 0, 0)
		}
		return newDate.Format(timeLayout), nil

	} else if strings.Contains(repeat, "w ") {
		parsedDate, err := time.Parse(timeLayout, date)
		if err != nil {
			return "", errors.New("invalid weekday." + err.Error())
		}
		weekday := int(parsedDate.Weekday())

		var newDate time.Time
		var weekdays []int
		for _, weekdayString := range strings.Split(strings.TrimPrefix(repeat, "w "), ",") {
			weekdayInt, _ := strconv.Atoi(weekdayString)
			if weekdayInt < 1 || weekdayInt > 7 {
				return "", errors.New("weekdays must be between 1 and 7")
			}
			weekdays = append(weekdays, weekdayInt)
		}

		updated := false
		for _, v := range weekdays {
			if weekday < v {
				newDate = parsedDate.AddDate(0, 0, v-weekday)
				updated = true
				break
			}
		}

		if !updated {
			newDate = parsedDate.AddDate(0, 0, 7-weekday+weekdays[0])
		}

		for newDate.Before(now) || newDate == now {
			weekday = int(newDate.Weekday())

			if weekday == weekdays[0] {
				for _, v := range weekdays {
					if weekday < v {
						newDate = newDate.AddDate(0, 0, v-weekday)
						weekday = int(newDate.Weekday())
					}
				}
			} else {
				newDate = newDate.AddDate(0, 0, 7-weekday+weekdays[0])
			}
		}

		return newDate.Format(timeLayout), nil

	} else {
		return "", errors.New("invalid repeat value")
	}
}
