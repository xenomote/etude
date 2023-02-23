package num

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Num interface {
	Min() int
	Max() int
	Width() int
	Precision() int
}

const keyword = "num"

func Parse(s string) (num Num, err error) {
	if !strings.HasPrefix(s, keyword) {
		return nil, fmt.Errorf(`must start with "%s"`, keyword)
	}
	s = strings.TrimPrefix(s, keyword)

	if len(s) == 0 {
		return nil, fmt.Errorf(`must specify either width or full range`)
	}

	prec := 0
	var width, min, max *int

	i := strings.IndexRune(s, '[')
	j := strings.IndexRune(s, ']')

	// there is one bracket without the other or in the wrong order
	if ((i > -1) != (j > -1)) || j < i {
		return nil, fmt.Errorf(`brackets must match`)
	}

	// there is a range/precision specifier
	if i > -1 {
		if len(s) > j+1 {
			return nil, fmt.Errorf(`extra values after brackets not permitted, got "%s"`, s[j+1:])
		}

		min, max, prec, err = minmaxprecs(s[i+1 : j])
		if err != nil {
			return nil, err
		}
		s = s[:i]
	}

	// there is a width specifier
	if len(s) > 0 {
		width, err = widths(s)
		if err != nil {
			return nil, err
		}
	}

	if width != nil {
		if min != nil && max != nil {
			return nil, fmt.Errorf(`cannot specify both width and full range`)
		}

		if min != nil {
			return NumWidthMin{*width, *min, prec}, nil
		}

		if max != nil {
			return NumWidthMin{*width, *max, prec}, nil
		}

		return NumWidth{*width, prec}, nil
	}

	if min == nil || max == nil {
		return nil, fmt.Errorf(`must specify either a width or full range`)
	}

	return NumMinMax{*min, *max, prec}, nil
}

var pattern = regexp.MustCompile(`^[1-9][0-9]*$`)

func widths(s string) (*int, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf(`width cannot be empty if provided`)
	}

	if !pattern.MatchString(s) {
		return nil, fmt.Errorf(`width must match the format "^[1-9][0-9]*$"`)
	}

	width, err := strconv.Atoi(s)
	if err != nil {
		return nil, fmt.Errorf(`width must be an integer, got "%s"`, s)
	}

	if width <= 0 {
		return nil, fmt.Errorf(`width must be greater than zero, got "%d"`, width)
	}

	if width%8 != 0 {
		return nil, fmt.Errorf(`width must be a multiple of 8, got "%d`, width)
	}
	width /= 8

	return &width, nil
}

func minmaxprecs(s string) (min, max *int, prec int, err error) {
	if len(s) == 0 {
		err = fmt.Errorf(`must have range and/or precision, or omit brackets entirely`)
		return
	}

	var a, b string
	c := strings.IndexRune(s, ',')
	if c > -1 {
		a = s[:c]
		b = s[c+1:]

		if len(a) == 0 {
			err = fmt.Errorf(`cannot provide empty range with comma`)
			return
		}

		if len(b) == 0 {
			err = fmt.Errorf(`cannot provide empty precision with comma`)
			return
		}
	} else if strings.ContainsRune(s, ':') {
		a = s
	} else {
		b = s
	}

	// there is a range specifier
	if len(a) > 0 {
		min, max, err = minmaxs(a)
		if err != nil {
			return
		}
	}

	// there is a precision specifier
	if len(b) > 0 {
		prec, err = precs(b)
		if err != nil {
			return
		}
	}

	return min, max, prec, nil
}

func minmaxs(s string) (min *int, max *int, err error) {
	l := strings.IndexRune(s, ':')
	if l < 0 {
		err = fmt.Errorf(`range must have a ":" symbol`)
		return
	}

	a := s[:l]
	b := s[l+1:]

	// there is a min value
	if len(a) > 0 {
		n, errs := strconv.Atoi(a)
		if errs != nil {
			err = fmt.Errorf(`range min must be an integer, got "%s"`, a)
			return
		}

		min = &n
	}

	// there is a max value
	if len(b) > 0 {
		n, errs := strconv.Atoi(b)
		if errs != nil {
			err = fmt.Errorf(`range max must be an integer, got "%s`, b)
			return
		}

		max = &n
	}

	// there is a full range
	if min != nil && max != nil && !(*min < *max) {
		err = fmt.Errorf(`range min must be less than max, got %d:%d`, *min, *max)
		return
	}

	return min, max, nil
}

func precs(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf(`precision string cannot be empty`)
	}

	switch s[0] {
	case '+', '-':
	default:
		return 0, fmt.Errorf(`precision must start with + or -`)
	}

	prec, errs := strconv.Atoi(s)
	if errs != nil {
		return 0, fmt.Errorf(`precision must be an integer, got "%s"`, s)
	}

	if prec == 0 {
		return 0, fmt.Errorf(`precision must be non-zero`)
	}

	return prec, nil
}
