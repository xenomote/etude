package parser

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

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

	EQUALS
	DOUBLE_EQUALS

	COLON
	COLON_EQUALS

	EXCLAIM
	EXCLAIM_EQUALS

	DOUBLE_AMPERSAND
	DOUBLE_PIPE

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

type LexerError string

func (l LexerError) Error() string {
	return string(l)
}

const (
	NotFound          = LexerError("no matching token found")
	UnexpectedEOF     = LexerError("unexpected end of input")
	UnexpectedNewline = LexerError("unexpected newline")
	UnexpectedToken   = LexerError("unexpected token type found")
)

type Position struct {
	l, c int
}

type Lexer struct {
	src  bufio.Reader
	buf  strings.Builder
	pos  Position
	end  bool
	last rune
}

func New(r io.Reader) Lexer {
	l := Lexer{
		src: *bufio.NewReader(r),
		pos: Position{1, 1},
	}
	l.Read()

	return l
}

func (l *Lexer) Read() {
	c, _, err := l.src.ReadRune()
	if err == io.EOF {
		l.end = true
		return
	}
	if err != nil {
		panic(err)
	}

	if c == '\n' {
		l.pos.l += 1
		l.pos.c = 1
	} else {
		l.pos.c += 1
	}

	l.buf.WriteRune(c)
	l.last = c
}

func (l *Lexer) Next() (Token, error) {
	for !l.end && unicode.IsSpace(l.last) {
		l.Read()
	}
	l.Clear()

	if l.end {
		return END, nil
	}

	switch l.last {
	case '{', '}', '[', ']', '(', ')', '?', '@', '#', '.', ',', '+', '-', '*', '/', '^', '%':
		l.Read()
		return Token(l.last), nil

	case '=':
		l.Read()
		if l.last == '=' {
			l.Read()
			return DOUBLE_EQUALS, nil
		}

		return EQUALS, nil

	case ':':
		l.Read()
		if l.last == '=' {
			l.Read()
			return COLON_EQUALS, nil
		}

		return COLON, nil

	case '!':
		l.Read()
		if l.last == '=' {
			l.Read()
			return EXCLAIM_EQUALS, nil
		}

		return EXCLAIM, nil
	}

	if l.last == '"' {
		return l.String()
	}

	if unicode.IsDigit(l.last) {
		return l.Number()
	}

	if unicode.IsLetter(l.last) {
		return l.Identifier()
	}

	return ERROR, NotFound
}

func (l *Lexer) String() (Token, error) {
	for {
		l.ReadEscape()

		if l.end {
			return ERROR, UnexpectedEOF
		}

		if l.last == '\n' {
			return ERROR, UnexpectedNewline
		}

		if l.last == '"' {
			l.Read()
			return STRING, nil
		}
	}
}

func (l *Lexer) ReadEscape() {
	l.Read()
	if l.last == '\\' {
		l.Read()
		l.Read()
	}
}

func (l *Lexer) Number() (Token, error) {
	return END, nil
}

func (l *Lexer) Identifier() (Token, error) {
	return END, nil
}

func (l *Lexer) Want(want Token) error {
	got, err := l.Next()
	if err != nil {
		return err
	}

	if want != got {
		return UnexpectedToken
	}

	return nil
}

func (l *Lexer) Text() string {
	s := l.buf.String()
	return s[:len(s)-1]
}

func (l *Lexer) Clear() {
	l.buf.Reset()
	l.buf.WriteRune(l.last)
}
