package typings

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
