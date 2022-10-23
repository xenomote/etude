package num_test

import (
	"testing"

	"github.com/xenomote/etude/internal/num"
)

func TestParseSuccess(t *testing.T) {
	tests := []string{
		`num8`, // without range/precision

		`num16[1:]`,  // with min offset
		`num16[-1:]`, // with negative min offset

		`num32[:1]`,  // with max offset
		`num32[:-1]`, // with negative max offset

		`num128[-3]`, // with negative precision
		`num256[+3]`, // with positive precision

		`num[-100:100]`, // with full range
		`num[100:200]`,  // with positive full range
		`num[-10:-5]`,   // with negative full range

		`num[-100:100,-3]`, // with full range and precision
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			n, err := num.Parse(test)
			if err != nil || n == nil {
				t.Error(err)
			}
		})
	}
}

func TestParseFailure(t *testing.T) {
	tests := []string{
		``,      // empty
		`blah8`, // not num

		`num`,    // no width or range
		`number`, // width is not number
		`num8aa`, // width with extra characters
		`num0`,   // zero width
		`num08`,  // zero prefixed width
		`num-8`,  // negative width
		`num3`,   // non multiple of 8 width

		`num8[]`,  // empty range
		`num8[`,   // missing close bracket
		`num8]`,   // missing open bracket
		`num8][`,  // missmatched brackets
		`num8[]x`, // extra values after brackets

		`num8[-1:1,]`, // comma but no precision
		`num8[1:-1]`,  // inverted range
		`num8[1:1]`,   // single value range
		`num8[1:2]`,   // two value range (use bit)

		`num[1:]`,   // min with no width
		`num8[aa:]`, // min not a number
		`num[:1]`,   // max with no width
		`num8[:aa]`, // max not a number

		`num8[,-3]`, // comma but no range
		`num8[123]`, // precision doesnt start with + or -
		`num8[+0]`,  // zero precision
		`num8[+aa]`, // precision not a number
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			n, err := num.Parse(test)
			if err == nil || n != nil {
				t.Errorf(`expected failure but got "%+v"`, n)
			}
		})
	}
}
