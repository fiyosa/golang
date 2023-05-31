package helper

import (
	"time"
)

const TimeFormatSQL = "2006-01-02 15:04:05"
const TimeFormatInput = "2006-01-02T15:04:05Z"

func TimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return now.Format(TimeFormatSQL)
}

func TimeFormat(data string, formatInput string, format string) string {
	date, err := time.Parse(formatInput, data)
	if err != nil {
		return ""
	}
	return date.Format(format)
}
