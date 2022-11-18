// Code generated by "stringer -type Token"; DO NOT EDIT.

package parser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERROR-0]
	_ = x[END-1]
	_ = x[FUNC-2]
	_ = x[RETURN-3]
	_ = x[FOR-4]
	_ = x[IF-5]
	_ = x[ON-6]
	_ = x[NUMBER-7]
	_ = x[STRING-8]
	_ = x[BOOLEAN-9]
	_ = x[EQUALS-10]
	_ = x[DOUBLE_EQUALS-11]
	_ = x[COLON-12]
	_ = x[COLON_EQUALS-13]
	_ = x[EXCLAIM-14]
	_ = x[EXCLAIM_EQUALS-15]
	_ = x[DOUBLE_AMPERSAND-16]
	_ = x[DOUBLE_PIPE-17]
	_ = x[CURLY_LEFT-123]
	_ = x[CURLY_RIGHT-125]
	_ = x[ROUND_LEFT-40]
	_ = x[ROUND_RIGHT-41]
	_ = x[SQUARE_LEFT-91]
	_ = x[SQUARE_RIGHT-93]
	_ = x[QUESTION-63]
	_ = x[ADDRESS-64]
	_ = x[HASH-35]
	_ = x[PERIOD-46]
	_ = x[COMMA-44]
	_ = x[PLUS-43]
	_ = x[MINUS-45]
	_ = x[STAR-42]
	_ = x[SLASH-47]
	_ = x[CARRET-94]
	_ = x[PERCENT-37]
}

const (
	_Token_name_0 = "ERRORENDFUNCRETURNFORIFONNUMBERSTRINGBOOLEANEQUALSDOUBLE_EQUALSCOLONCOLON_EQUALSEXCLAIMEXCLAIM_EQUALSDOUBLE_AMPERSANDDOUBLE_PIPE"
	_Token_name_1 = "HASH"
	_Token_name_2 = "PERCENT"
	_Token_name_3 = "ROUND_LEFTROUND_RIGHTSTARPLUSCOMMAMINUSPERIODSLASH"
	_Token_name_4 = "QUESTIONADDRESS"
	_Token_name_5 = "SQUARE_LEFT"
	_Token_name_6 = "SQUARE_RIGHTCARRET"
	_Token_name_7 = "CURLY_LEFT"
	_Token_name_8 = "CURLY_RIGHT"
)

var (
	_Token_index_0 = [...]uint8{0, 5, 8, 12, 18, 21, 23, 25, 31, 37, 44, 50, 63, 68, 80, 87, 101, 117, 128}
	_Token_index_3 = [...]uint8{0, 10, 21, 25, 29, 34, 39, 45, 50}
	_Token_index_4 = [...]uint8{0, 8, 15}
	_Token_index_6 = [...]uint8{0, 12, 18}
)

func (i Token) String() string {
	switch {
	case i <= 17:
		return _Token_name_0[_Token_index_0[i]:_Token_index_0[i+1]]
	case i == 35:
		return _Token_name_1
	case i == 37:
		return _Token_name_2
	case 40 <= i && i <= 47:
		i -= 40
		return _Token_name_3[_Token_index_3[i]:_Token_index_3[i+1]]
	case 63 <= i && i <= 64:
		i -= 63
		return _Token_name_4[_Token_index_4[i]:_Token_index_4[i+1]]
	case i == 91:
		return _Token_name_5
	case 93 <= i && i <= 94:
		i -= 93
		return _Token_name_6[_Token_index_6[i]:_Token_index_6[i+1]]
	case i == 123:
		return _Token_name_7
	case i == 125:
		return _Token_name_8
	default:
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
