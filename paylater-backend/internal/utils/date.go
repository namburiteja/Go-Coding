package utils

import "time"

func CalculatePaymentDueDate(now time.Time) time.Time {

	year, month, day := now.Date()
	location := now.Location()

	if day <= 5 {
		return time.Date(year, month, 5, 0, 0, 0, 0, location)
	}

	return time.Date(year, month+1, 5, 0, 0, 0, 0, location)
}