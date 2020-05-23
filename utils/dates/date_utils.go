package dates

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

// GetNow is to get the current UTC datetime
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString is to get the current UTC datetime in the format defined in apiDateLayout
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
