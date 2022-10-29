package utils

import (
	"fmt"
	"time"
)

func GoTimeToYYYYMMDD(t *time.Time) (string, string, string, string, string) {

	yyyy := fmt.Sprintf("%04d", t.Year())
	mm := fmt.Sprintf("%02d", int(t.Month()))
	dd := fmt.Sprintf("%02d", t.Day())
	hh := fmt.Sprintf("%02d", t.Hour())

	dateStr := fmt.Sprintf("%s%s%s%s", yyyy, mm, dd, hh)

	return yyyy, mm, dd, hh, dateStr
}
