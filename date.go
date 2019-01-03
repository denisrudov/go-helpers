package go_helpers

import "time"

// Addition n days to a specific Time
func AddDays(startDate time.Time, d int) time.Time {
	return startDate.AddDate(0, 0, d)
}

// Subtraction n days from a specific Time
func SubDays(startDate time.Time, d int) time.Time {
	return startDate.AddDate(0, 0, -d)
}
