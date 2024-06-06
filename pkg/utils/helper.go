package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func TitleValidate(title string) error {
	if title == "" {
		return errors.New("title is required")
	}
	return nil
}

func RepeatValidate(repeat string) error {
	if repeat == "" {
		return errors.New("repeat is empty")
	}
	if strings.Contains(repeat, "d ") {
		days, err := strconv.Atoi(strings.Trim(repeat, "d "))
		if err != nil {
			return err
		}
		if days > 400 {
			return errors.New("days cannot be greater than 400")
		}
	}
	if strings.Contains(repeat, "w ") {
		for _, weekdayString := range strings.Split(strings.TrimPrefix(repeat, "w "), ",") {
			weekdayInt, _ := strconv.Atoi(weekdayString)
			if weekdayInt < 1 || weekdayInt > 7 {
				return errors.New("weekdays must be between 1 and 7")
			}
		}
		return nil
	}
	if !strings.Contains(repeat, "d ") && repeat != "y" {
		return errors.New("repeat rule is invalid")
	}
	return nil
}

func DateValidate(date string) error {
	_, err := time.Parse(TimeLayout, date)
	if err != nil {
		return errors.New("invalid date." + err.Error())
	}
	return nil
}
