package typings

import "fmt"

type TTime struct {
	Yyyy int
	Dd   int
	Mm   int
	Hh   int
}

func NewTTime(y, m, d, h int) *TTime {
	return &TTime{
		Yyyy: y,
		Dd:   d,
		Mm:   m,
		Hh:   h,
	}
}

var mmInYr int = 12

func (t *TTime) jumpMonth() {

	nextMth := (t.Mm + 1)
	if nextMth > mmInYr {
		nextMth = 1
		t.jumpYear()
	}
	t.Mm = nextMth
}

func (t *TTime) jumpYear() {
	t.Yyyy += 1
}

func (t *TTime) AsString() string {
	return t.yString() + t.mString() + t.dString() + t.hString()
}

func (t *TTime) yString() string {
	return fmt.Sprintf("%04d", t.Yyyy)
}

func (t *TTime) mString() string {
	return fmt.Sprintf("%02d", t.Mm)
}

func (t *TTime) dString() string {
	return fmt.Sprintf("%02d", t.Dd)
}

func (t *TTime) hString() string {
	return fmt.Sprintf("%02d", t.Hh)
}

var hrsInDay int = 24

func (t *TTime) JumpHour() {

	nextH := (t.Hh + 1)
	if nextH == hrsInDay {
		nextH = 0
		t.jumpDay()
	}
	t.Hh = nextH
}

var daysInYr int = 31

func (t *TTime) jumpDay() {

	nextDay := (t.Dd + 1)
	if nextDay > daysInYr {
		nextDay = 1
		t.jumpMonth()
	}
	t.Dd = nextDay
}
