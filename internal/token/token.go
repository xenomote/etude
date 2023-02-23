package token

//go:generate stringer -type Token
type Token uint

const (
	ERROR Token = iota
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

	CURLY_LEFT  Token = '{'
	CURLY_RIGHT Token = '}'

	ROUND_LEFT  Token = '('
	ROUND_RIGHT Token = ')'

	SQUARE_LEFT  Token = '['
	SQUARE_RIGHT Token = ']'

	QUESTION Token = '?'
	ADDRESS  Token = '@'
	HASH     Token = '#'

	PERIOD Token = '.'
	COMMA  Token = ','

	PLUS    Token = '+'
	MINUS   Token = '-'
	STAR    Token = '*'
	SLASH   Token = '/'
	CARRET  Token = '^'
	PERCENT Token = '%'
)