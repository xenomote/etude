package lexer_test

import (
	"io"
	"testing"

	"github.com/xenomote/etude/internal/lexer"
	"github.com/xenomote/etude/internal/token"
)

func TestWant(t *testing.T) {
	s := `   ===@?><>=!=comp.,"abc \n"1+cfunc#func`
	l := lexer.New()

	_, err := io.WriteString(l, s)
	if err != nil {
		t.Fatal(err)
	}

	toks := []token.Kind{
		token.DOUBLE_EQUALS,
		token.EQUALS,
		token.ADDRESS,
		token.QUESTION,
		token.GREATER,
		token.LESS,
		token.GREATER_EQUALS,
		token.EXCLAIM_EQUALS,
		token.COMP,
		token.PERIOD,
		token.COMMA,
		token.STRING,
		token.NUMBER,
		token.PLUS,
		token.IDENTIFIER,
		token.HASH,
		token.FUNC,
	}

	for _, tok := range toks {
		got, err := l.Next()
		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		if tok != got.Kind {
			t.Fatal("expected:", tok, "got:", got.Kind, string(got.Text))
		}
	}
}

func TestNextFail(t *testing.T) {
	for _, test := range []struct {
		input  string
		output lexer.Error
	}{
		{
			"`",
			lexer.ErrNotFound,
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
			l := lexer.New()
			_, err := io.WriteString(l, test.input)
			if err != nil {
				t.Fatal(err)
			}

			got, err := l.Next()
			if err == nil {
				t.Fatal("got:", got.Kind, "but expected error:", test.output)
			}

			if err != test.output {
				t.Fatal("unexpected error", err)
			}
		})
	}
}
