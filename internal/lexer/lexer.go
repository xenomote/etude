package lexer

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	. "github.com/xenomote/etude/internal/tokens"
)

type Error string

func (l Error) Error() string {
	return string(l)
}

const (
	NotFound                = Error("no matching token found")
	UnexpectedEOF           = Error("unexpected end of input")
	UnexpectedStringEOF     = Error("unexpected end of input before string close")
	UnexpectedStringNewline = Error("unexpected newline before string close")
	UnexpectedToken         = Error("unexpected token type found")
)

type Position struct {
	l, c int
}

type Lexer struct {
	src  *bufio.Reader
	buf  *strings.Builder
	pos  Position
	got  Token
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

func (l *Lexer) Want(tokens ...Token) (Token, error) {
	got, err := l.Next()

	for _, want := range tokens {
		if want == got {
			return want, nil
		}
	}

	if err != nil {
		return got, err
	}

	return got, UnexpectedToken
}

func (l *Lexer) Next() (Token, error) {
	for !l.end && unicode.IsSpace(l.last) {
		l.read()
	}
	l.Clear()

	if l.end {
		return l.fail(UnexpectedEOF)
	}

	switch l.last {
	case '{', '}', '[', ']', '(', ')', '?', '@', '#', '.', ',', '+', '-', '*', '/', '^', '%':
		return l.emit(Token(l.last))

	case '=':
		return l.either(EQUALS, DOUBLE_EQUALS)

	case ':':
		return l.either(COLON, COLON_EQUALS)

	case '!':
		return l.either(EXCLAIM, EXCLAIM_EQUALS)

	case '<':
		return l.either(LESS, LESS_EQUALS)

	case '>':
		return l.either(GREATER, GREATER_EQUALS)

	case '"':
		return l.string()
	}

	if unicode.IsDigit(l.last) {
		return l.number()
	}

	if unicode.IsLetter(l.last) {
		return l.identifier()
	}

	return l.fail(NotFound)
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

func (l *Lexer) emit(t Token) (Token, error) {
	l.got = t
	l.read()
	return t, nil
}

func (l *Lexer) fail(err error) (Token, error) {
	l.got = ERROR
	return ERROR, err
}

func (l *Lexer) either(a, b Token) (Token, error) {
	l.read()

	if l.last == '=' {
		return l.emit(b)
	}

	l.got = a
	return a, nil
}

func (l *Lexer) string() (Token, error) {
	for {
		l.readEscape()

		if l.end {
			return l.fail(UnexpectedStringEOF)
		}

		if l.last == '\n' {
			return l.fail(UnexpectedStringNewline)
		}

		if l.last == '"' {
			return l.emit(STRING)
		}
	}
}

func (l *Lexer) number() (Token, error) {
	l.read()
	return NUMBER, nil
}

func (l *Lexer) identifier() (Token, error) {
	l.read()
	return IDENTIFIER, nil
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
