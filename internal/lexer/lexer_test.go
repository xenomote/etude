package lexer_test

import (
	"strings"
	"testing"

	"github.com/xenomote/etude/internal/lexer"
	. "github.com/xenomote/etude/internal/tokens"
)

func TestWant(t *testing.T) {
	s := `   ===@?><>=!=:=.,"abc \n"1c`
	l := lexer.New(strings.NewReader(s))

	toks := []Token{
		DOUBLE_EQUALS,
		EQUALS,
		ADDRESS,
		QUESTION,
		GREATER,
		LESS,
		GREATER_EQUALS,
		EXCLAIM_EQUALS,
		COLON_EQUALS,
		PERIOD,
		COMMA,
		STRING,
		NUMBER,
		IDENTIFIER,
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
			lexer.NotFound,
		},
		{
			"",
			lexer.UnexpectedEOF,
		},
		{
			"\"hello",
			lexer.UnexpectedStringEOF,
		},
		{
			"\"hello\n",
			lexer.UnexpectedStringNewline,
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
