package repository

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const timeLayout = "20060102"

func titleValidate(title string) error {
	if title == "" {
		return errors.New("title is required")
	}
	return nil
}

func repeatValidate(repeat string) error {
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
	if !strings.Contains(repeat, "d ") && repeat != "y" {
		return errors.New("repeat rule is invalid")
	}
	return nil
}

func dateValidate(date string) error {
	_, err := time.Parse(timeLayout, date)
	if err != nil {
		return errors.New("invalid date." + err.Error())
	}
	return nil
}
