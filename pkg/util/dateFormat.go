package util

import "time"

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}
