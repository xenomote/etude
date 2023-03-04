package token

type Token struct {
	Kind Kind
	Text []byte
}

func (t Token) String() string {
	return string(t.Text)
}

//go:generate stringer -type Kind
type Kind uint

const (
	ERROR Kind = iota
	END

	COMP
	TYPE
	FUNC

	IF
	OR
	ON
	FOR
	COPY
	RETURN

	EQUALS
	EXCLAIM
	ELLIPSIS

	DOUBLE_EQUALS
	EXCLAIM_EQUALS
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	DOUBLE_AMPERSAND
	DOUBLE_PIPE

	SHIFT_LEFT
	SHIFT_RIGHT

	NUMBER
	STRING
	BOOLEAN
	IDENTIFIER

	CURLY_LEFT  Kind = '{'
	CURLY_RIGHT Kind = '}'

	ROUND_LEFT  Kind = '('
	ROUND_RIGHT Kind = ')'

	SQUARE_LEFT  Kind = '['
	SQUARE_RIGHT Kind = ']'

	MINUS   Kind = '-'
	PLUS    Kind = '+'
	STAR    Kind = '*'
	SLASH   Kind = '/'
	CARRET  Kind = '^'
	PERCENT Kind = '%'

	ADDRESS Kind = '@'

	QUESTION Kind = '?'
	HASH     Kind = '#'
	TILDE    Kind = '~'

	COLON  Kind = ':'
	COMMA  Kind = ','
	PERIOD Kind = '.'
)
