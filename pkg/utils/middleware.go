package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const timeLayout = "20060102"

func TimeValidate(date string) error {
	if _, err := time.Parse(timeLayout, date); err != nil {
		return err
	}
	return nil
}

func GetNextDate(now time.Time, date, repeat, timeLayout string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat cannot be empty")
	}

	parsedDate, err := time.Parse(timeLayout, date)
	if err != nil {
		return "", errors.New("invalid date: " + err.Error())
	}

	switch {
	case strings.Contains(repeat, "d "):
		return GetNextDateByDays(now, parsedDate, repeat, timeLayout)
	case repeat == "y":
		return GetNextDateByYears(now, parsedDate, timeLayout)
	case strings.Contains(repeat, "w "):
		return GetNextDateByWeekdays(now, parsedDate, repeat, timeLayout)
	default:
		return "", errors.New("invalid repeat value")
	}
}

func GetNextDateByDays(now, parsedDate time.Time, repeat, timeLayout string) (string, error) {
	days, err := strconv.Atoi(strings.Trim(repeat, "d "))
	if err != nil {
		return "", err
	}
	if days > 400 {
		return "", errors.New("days cannot be greater than 400")
	}

	newDate := parsedDate.AddDate(0, 0, days)
	for newDate.Before(now) {
		newDate = newDate.AddDate(0, 0, days)
	}
	return newDate.Format(timeLayout), nil
}

func GetNextDateByYears(now, parsedDate time.Time, timeLayout string) (string, error) {
	newDate := parsedDate.AddDate(1, 0, 0)
	for newDate.Before(now) {
		newDate = newDate.AddDate(1, 0, 0)
	}
	return newDate.Format(timeLayout), nil
}

func GetNextDateByWeekdays(now, parsedDate time.Time, repeat, timeLayout string) (string, error) {
	weekdays, err := ParseWeekdays(repeat)
	if err != nil {
		return "", err
	}

	newDate := findNextWeekday(parsedDate, weekdays)
	for newDate.Before(now) || newDate.Equal(now) {
		newDate = findNextWeekday(newDate, weekdays)
	}
	return newDate.Format(timeLayout), nil
}

func ParseWeekdays(repeat string) ([]int, error) {
	var weekdays []int
	for _, weekdayString := range strings.Split(strings.TrimPrefix(repeat, "w "), ",") {
		weekdayInt, err := strconv.Atoi(weekdayString)
		if err != nil || weekdayInt < 1 || weekdayInt > 7 {
			return nil, errors.New("weekdays must be between 1 and 7")
		}
		weekdays = append(weekdays, weekdayInt)
	}
	return weekdays, nil
}

func findNextWeekday(parsedDate time.Time, weekdays []int) time.Time {
	weekday := int(parsedDate.Weekday())
	for _, v := range weekdays {
		if weekday < v {
			return parsedDate.AddDate(0, 0, v-weekday)
		}
	}
	return parsedDate.AddDate(0, 0, 7-weekday+weekdays[0])
}
