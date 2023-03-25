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

	ELLIPSIS
	DOUBLE_PLUS
	DOUBLE_MINUS

	DOUBLE_EQUALS
	EXCLAIM_EQUALS
	LESS_EQUALS
	GREATER_EQUALS

	DOUBLE_AMPERSAND
	DOUBLE_PIPE

	SHIFT_LEFT
	SHIFT_RIGHT

	NUMBER
	STRING
	BOOLEAN
	IDENTIFIER

	EXCLAIM Kind = '!'
	EQUALS  Kind = '='
	LESS    Kind = '<'
	GREATER Kind = '>'

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
