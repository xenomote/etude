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
	l.Read()

	return l
}

func (l *Lexer) Read() {
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

func (l *Lexer) Next() (Token, error) {
	for !l.end && unicode.IsSpace(l.last) {
		l.Read()
	}
	l.Clear()

	if l.end {
		return l.Fail(UnexpectedEOF)
	}

	switch l.last {
	case '{', '}', '[', ']', '(', ')', '?', '@', '#', '.', ',', '+', '-', '*', '/', '^', '%':
		return l.Emit(Token(l.last))

	case '=':
		return l.Either(EQUALS, DOUBLE_EQUALS)

	case ':':
		return l.Either(COLON, COLON_EQUALS)

	case '!':
		return l.Either(EXCLAIM, EXCLAIM_EQUALS)

	case '<':
		return l.Either(LESS, LESS_EQUALS)

	case '>':
		return l.Either(GREATER, GREATER_EQUALS)

	case '"':
		return l.String()
	}

	if unicode.IsDigit(l.last) {
		return l.Number()
	}

	if unicode.IsLetter(l.last) {
		return l.Identifier()
	}

	return l.Fail(NotFound)
}

func (l *Lexer) Emit(t Token) (Token, error) {
	l.got = t
	l.Read()
	return t, nil
}

func (l *Lexer) Fail(err error) (Token, error) {
	l.got = ERROR
	return ERROR, err
}

func (l *Lexer) Either(a, b Token) (Token, error) {
	l.Read()

	if l.last == '=' {
		return l.Emit(b)
	}

	l.got = a
	return a, nil
}

func (l *Lexer) String() (Token, error) {
	for {
		l.ReadEscape()

		if l.end {
			return l.Fail(UnexpectedStringEOF)
		}

		if l.last == '\n' {
			return l.Fail(UnexpectedStringNewline)
		}

		if l.last == '"' {
			return l.Emit(STRING)
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
	l.Read()
	return NUMBER, nil
}

func (l *Lexer) Identifier() (Token, error) {
	l.Read()
	return IDENTIFIER, nil
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
