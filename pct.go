package pct

import (
	"errors"
	"strconv"
	"strings"
)

func ParsePercent(s string) (float64, error) {
	var err error
	var f float64

	addPoint := func(s string) string {
		r := ""
		switch l := len(s); l {
		case 0:
			r = "0.00"
		case 1:
			r = "0.0" + s
		case 2:
			r = "0." + s
		default:
			r = s[:l-2] + "." + s[l-2:]
		}
		return r
	}

	ss := strings.Split(s, "%")
	if len(ss) != 2 {
		err = errors.New("too many percent symbols, or too few")
	}
	s = ss[0]
	ss = strings.Split(s, ".")

	switch l := len(ss); l {
	case 1:
		s = addPoint(ss[0])
	case 2:
		s = addPoint(ss[0]) + ss[1]
	default:
		err = errors.New("too many decimal points")
	}

	if err == nil {
		f, err = strconv.ParseFloat(s, 64)
	}
	return f, err
}
