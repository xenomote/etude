package lexer_test

import (
	"strings"
	"testing"

	"github.com/xenomote/etude/internal/lexer"
	"github.com/xenomote/etude/internal/token"
)

func TestWant(t *testing.T) {
	s := `   ===@?><>=!=:=.,"abc \n"1c`
	l := lexer.New(strings.NewReader(s))

	toks := []token.Token{
		token.DOUBLE_EQUALS,
		token.EQUALS,
		token.ADDRESS,
		token.QUESTION,
		token.GREATER,
		token.LESS,
		token.GREATER_EQUALS,
		token.EXCLAIM_EQUALS,
		token.COLON_EQUALS,
		token.PERIOD,
		token.COMMA,
		token.STRING,
		token.NUMBER,
		token.IDENTIFIER,
	}

	for _, tok := range toks {
		got, err := l.Want(tok)
		if err != nil {
			t.Fatal(tok, err, got)
		}
	}
}

func TestNextFail(t *testing.T) {
	for _, test := range []struct {
		input  string
		output lexer.Error
	}{
		{
			"~",
			lexer.ErrNotFound,
		},
		{
			"",
			lexer.ErrUnexpectedEOF,
		},
		{
			"\"hello",
			lexer.ErrUnexpectedStringEOF,
		},
		{
			"\"hello\n",
			lexer.ErrUnexpectedStringNewline,
		},
	} {
		t.Run(test.output.Error(), func(t *testing.T) {
			l := lexer.New(strings.NewReader(test.input))
			_, err := l.Next()
			if err == nil {
				t.Fatal("expected an error but got none")
			}

			if err != test.output {
				t.Fatal("unexpected error", err)
			}
		})
	}
}
