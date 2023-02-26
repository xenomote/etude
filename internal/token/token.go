package token

type Token struct {
	Kind Kind
	Text []byte
}

//go:generate stringer -type Kind
type Kind uint

const (
	ERROR Kind = iota
	END

	FUNC
	RETURN
	FOR
	IF
	ON

	NUMBER
	STRING
	BOOLEAN
	IDENTIFIER

	EQUALS
	DOUBLE_EQUALS

	COLON
	COLON_EQUALS

	EXCLAIM
	EXCLAIM_EQUALS

	DOUBLE_AMPERSAND
	DOUBLE_PIPE

	LESS
	LESS_EQUALS

	GREATER
	GREATER_EQUALS

	CURLY_LEFT  Kind = '{'
	CURLY_RIGHT Kind = '}'

	ROUND_LEFT  Kind = '('
	ROUND_RIGHT Kind = ')'

	SQUARE_LEFT  Kind = '['
	SQUARE_RIGHT Kind = ']'

	QUESTION Kind = '?'
	ADDRESS  Kind = '@'
	HASH     Kind = '#'

	PERIOD Kind = '.'
	COMMA  Kind = ','

	PLUS    Kind = '+'
	MINUS   Kind = '-'
	STAR    Kind = '*'
	SLASH   Kind = '/'
	CARRET  Kind = '^'
	PERCENT Kind = '%'
)
