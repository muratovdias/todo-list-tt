package service

import (
	"time"
	"unicode/utf8"
)

func validateTitle(title string) error {
	if utf8.RuneCount([]byte(title)) > 200 {
		return ErrInvalidTitle
	}
	return nil
}

func validateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ErrInvalidDate
	}

	return nil
}

func isWeekend(date string) (bool, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, ErrInvalidDate
	}

	day := d.Weekday().String()
	if day == "Sunday" || day == "Saturday" {
		return true, nil
	}

	return false, nil
}
