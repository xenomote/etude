package lexer

import (
	"unicode"

	"github.com/xenomote/etude/internal/token"
)

type Error string

func (l Error) Error() string {
	return string(l)
}

const (
	ErrNotFound                = Error("no matching token found")
	ErrUnexpectedToken         = Error("unexpected token type found")
	ErrUnexpectedStringEOF     = Error("unexpected end of input before string close")
	ErrUnexpectedStringNewline = Error("unexpected newline before string close")
	ErrBadNumber               = Error("malformed number token")
)

type Position struct {
	l, c int
}

type lexer struct {
	src []byte
	a   int
	b   int
	pos Position
}

func New() *lexer {
	return &lexer{pos: Position{1, 1}}
}

func (l *lexer) Write(b []byte) (int, error) {
	l.src = append(l.src, b...)
	return len(b), nil
}

func (l *lexer) Any(tokens ...token.Kind) (token.Token, error) {
	got, err := l.Next()

	for _, want := range tokens {
		if want == got.Kind {
			return got, nil
		}
	}

	if err != nil {
		return got, err
	}

	return got, ErrUnexpectedToken
}

func (l *lexer) Next() (token.Token, error) {
	var c rune
	for {
		c = rune(l.read())

		if c == EOF || !unicode.IsSpace(c) {
			break
		}
	}

	if c == EOF {
		return l.emit(token.END)
	}

	switch c {
	case '{', '}', '[', ']', '(', ')', '?', '@', '#', '.', ',', '+', '-', '*', '/', '^', '%':
		return l.emit(token.Kind(c))

	case '=':
		return l.either(token.EQUALS, token.DOUBLE_EQUALS)

	case ':':
		return l.either(token.COLON, token.COLON_EQUALS)

	case '!':
		return l.either(token.EXCLAIM, token.EXCLAIM_EQUALS)

	case '<':
		return l.either(token.LESS, token.LESS_EQUALS)

	case '>':
		return l.either(token.GREATER, token.GREATER_EQUALS)

	case '"':
		return l.string()
	}

	if unicode.IsDigit(rune(c)) {
		return l.number()
	}

	if unicode.IsLetter(c) {
		return l.identifier()
	}

	return l.fail(ErrNotFound)
}

func (l *lexer) emit(kind token.Kind) (token.Token, error) {
	text := l.src[l.a:l.b]
	l.a = l.b

	return token.Token{Kind: kind, Text: text}, nil
}

func (l *lexer) fail(err error) (token.Token, error) {
	t, _ := l.emit(token.ERROR)

	return t, err
}

func (l *lexer) either(a, b token.Kind) (token.Token, error) {
	c := l.read()

	if c == '=' {
		return l.emit(b)
	}

	l.unread()
	return l.emit(a)
}

func (l *lexer) string() (token.Token, error) {
	for {
		switch l.readEscape() {
		case EOF:
			return l.fail(ErrUnexpectedStringEOF)

		case '\n':
			return l.fail(ErrUnexpectedStringNewline)

		case '"':
			return l.emit(token.STRING)

		default:
		}
	}
}

func (l *lexer) number() (token.Token, error) {
	var c rune
	for {
		c = rune(l.read())

		if !unicode.IsNumber(c) {
			l.unread()
			break
		}
	}

	if c == EOF || !unicode.IsLetter(c) {
		return l.emit(token.NUMBER)
	}

	return l.fail(ErrBadNumber)
}

func (l *lexer) identifier() (token.Token, error) {
	var c rune
	for {
		c = rune(l.read())

		if !(unicode.IsLetter(c) || unicode.IsNumber(c)) {
			l.unread()
			return l.emit(token.IDENTIFIER)
		}
	}
}

func (l *lexer) readEscape() byte {
	c := l.read()
	if c != '\\' {
		return c
	}

	l.read()
	return '\\' // anything that is not the escaped character
}

const EOF = 0

func (l *lexer) read() byte {
	if !(l.b < len(l.src)) {
		return EOF
	}

	l.b++
	c := l.src[l.b-1]

	if c == '\n' {
		l.pos.l += 1
		l.pos.c = 1
	} else {
		l.pos.c += 1
	}

	return c
}

func (l *lexer) unread() {
	if !(l.b > l.a) {
		panic("unread into previous token")
	}

	l.b--
	c := l.src[l.b]

	if c == '\n' {
		panic("unread into previous line")
	}

	l.pos.c -= 1
}
