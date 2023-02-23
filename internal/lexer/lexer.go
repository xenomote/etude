package lexer

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	"github.com/xenomote/etude/internal/token"
)

type Error string

func (l Error) Error() string {
	return string(l)
}

const (
	ErrNotFound                = Error("no matching token found")
	ErrUnexpectedEOF           = Error("unexpected end of input")
	ErrUnexpectedStringEOF     = Error("unexpected end of input before string close")
	ErrUnexpectedStringNewline = Error("unexpected newline before string close")
	ErrUnexpectedToken         = Error("unexpected token type found")
)

type Position struct {
	l, c int
}

type Lexer struct {
	src  *bufio.Reader
	buf  *strings.Builder
	pos  Position
	got  token.Token
	end  bool
	last rune
}

func New(r io.Reader) Lexer {
	l := Lexer{
		src: bufio.NewReader(r),
		pos: Position{1, 1},
		buf: &strings.Builder{},
	}
	l.read()

	return l
}

func (l *Lexer) Want(tokens ...token.Token) (token.Token, error) {
	got, err := l.Next()

	for _, want := range tokens {
		if want == got {
			return want, nil
		}
	}

	if err != nil {
		return got, err
	}

	return got, ErrUnexpectedToken
}

func (l *Lexer) Next() (token.Token, error) {
	for !l.end && unicode.IsSpace(l.last) {
		l.read()
	}
	l.Clear()

	if l.end {
		return l.fail(ErrUnexpectedEOF)
	}

	switch l.last {
	case '{', '}', '[', ']', '(', ')', '?', '@', '#', '.', ',', '+', '-', '*', '/', '^', '%':
		return l.emit(token.Token(l.last))

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

	if unicode.IsDigit(l.last) {
		return l.number()
	}

	if unicode.IsLetter(l.last) {
		return l.identifier()
	}

	return l.fail(ErrNotFound)
}

func (l *Lexer) Text() string {
	s := l.buf.String()
	if !l.end {
		s = s[:len(s)-1]
	}

	return s
}

func (l *Lexer) Clear() {
	l.buf.Reset()
	if !l.end {
		l.buf.WriteRune(l.last)
	}
}

func (l *Lexer) emit(t token.Token) (token.Token, error) {
	l.got = t
	l.read()
	return t, nil
}

func (l *Lexer) fail(err error) (token.Token, error) {
	l.got = token.ERROR
	return token.ERROR, err
}

func (l *Lexer) either(a, b token.Token) (token.Token, error) {
	l.read()

	if l.last == '=' {
		return l.emit(b)
	}

	l.got = a
	return a, nil
}

func (l *Lexer) string() (token.Token, error) {
	for {
		l.readEscape()

		if l.end {
			return l.fail(ErrUnexpectedStringEOF)
		}

		if l.last == '\n' {
			return l.fail(ErrUnexpectedStringNewline)
		}

		if l.last == '"' {
			return l.emit(token.STRING)
		}
	}
}

func (l *Lexer) number() (token.Token, error) {
	l.read()
	return token.NUMBER, nil
}

func (l *Lexer) identifier() (token.Token, error) {
	l.read()
	return token.IDENTIFIER, nil
}

func (l *Lexer) readEscape() {
	l.read()
	if l.last == '\\' {
		l.read()
		l.read()
	}
}

func (l *Lexer) read() {
	c, _, err := l.src.ReadRune()
	if err == io.EOF {
		l.end = true
		l.last = 0
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

	_, err = l.buf.WriteRune(c)
	if err != nil {
		panic(err)
	}

	l.last = c
}
