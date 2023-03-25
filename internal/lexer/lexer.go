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

type Lexer interface {
	Write(b []byte) (int, error)
	Next() (token.Token, error)
}

func ReadAll(l lexer) ([]token.Token, error) {
	var tokens []token.Token

	for {
		t, err := l.Next()
		if err != nil {
			return tokens, err
		}

		tokens = append(tokens, t)

		if t.Kind == token.END {
			return tokens, nil
		}
	}
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

func (l *lexer) Next() (token.Token, error) {
	var c rune
	for {
		c = rune(l.read())

		if c == EOF || !unicode.IsSpace(c) {
			l.a = l.b - 1
			break
		}
	}

	if c == EOF {
		return l.emit(token.END)
	}

	switch c {
	case '{', '}', '[', ']', '(', ')', '?', '@', '~', '#', ',', '*', '/', '^', '%', ':':
		return l.emit(token.Kind(c))
	
	case '.':
		for l.read() == '.' {}
		
		switch string(l.text()) {
		case ".":
			return l.emit(token.PERIOD)

		case "...":
			return l.emit(token.ELLIPSIS)

		default:
			return l.fail(ErrNotFound)
		}

	case '+':
		return l.ifPeek('+', token.PLUS, token.DOUBLE_PLUS)

	case '-':
		return l.ifPeek('-', token.MINUS, token.DOUBLE_MINUS)

	case '=':
		return l.ifPeek('=', token.EQUALS, token.DOUBLE_EQUALS)

	case '!':
		return l.ifPeek('=', token.EXCLAIM, token.EXCLAIM_EQUALS)

	case '<':
		return l.ifPeek('=', token.LESS, token.LESS_EQUALS)

	case '>':
		return l.ifPeek('=', token.GREATER, token.GREATER_EQUALS)

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

func (l *lexer) text() []byte {
	return l.src[l.a:l.b]
}

func (l *lexer) emit(kind token.Kind) (token.Token, error) {
	text := l.text()
	l.a = l.b

	return token.Token{Kind: kind, Text: text}, nil
}

func (l *lexer) fail(err error) (token.Token, error) {
	t, _ := l.emit(token.ERROR)

	return t, err
}

func (l *lexer) ifPeek(c byte, a, b token.Kind) (token.Token, error) {
	if l.peek() != c {
		return l.emit(a)
	}

	l.read()

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
		c = rune(l.peek())

		if c == EOF {
			break
		}

		if !unicode.IsNumber(c) {
			break
		}

		l.read()
	}

	if c == EOF || !unicode.IsLetter(c) {
		return l.emit(token.NUMBER)
	}

	return l.fail(ErrBadNumber)
}

var keywords = map[string]token.Kind{
	"comp": token.COMP,
	"type": token.TYPE,
	"func": token.FUNC,

	"if":     token.IF,
	"or":     token.OR,
	"for":    token.FOR,
	"copy":   token.COPY,
	"return": token.RETURN,

	"true":  token.BOOLEAN,
	"false": token.BOOLEAN,
}

func (l *lexer) identifier() (token.Token, error) {
	var c rune
	for {
		c = rune(l.peek())

		if c == EOF {
			break
		}

		if !(unicode.IsLetter(c) || unicode.IsNumber(c)) {
			break
		}

		l.read()
	}

	kind, exists := keywords[string(l.text())]
	if exists {
		return l.emit(kind)
	}

	return l.emit(token.IDENTIFIER)
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

	c := l.src[l.b]
	l.b++

	if c == '\n' {
		l.pos.l += 1
		l.pos.c = 1
	} else {
		l.pos.c += 1
	}

	return c
}

func (l *lexer) peek() byte {
	if !(l.b < len(l.src)) {
		return EOF
	}

	return l.src[l.b]
}
