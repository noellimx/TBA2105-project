package utils

import (
	"fmt"
	"time"
)

func GoTimeToYYYYMMDDHH(t *time.Time) (yyyy string, mm string, dd string, hh string, dateStr string) {

	yyyy_ := fmt.Sprintf("%04d", t.Year())
	mm_ := fmt.Sprintf("%02d", int(t.Month()))
	dd_ := fmt.Sprintf("%02d", t.Day())
	hh_ := fmt.Sprintf("%02d", t.Hour())

	dateStr_ := fmt.Sprintf("%s%s%s%s", yyyy_, mm_, dd_, hh_)
	return yyyy_, mm_, dd_, hh_, dateStr_
}
