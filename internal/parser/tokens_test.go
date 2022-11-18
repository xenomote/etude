package parser_test

import (
	"strings"
	"testing"

	"github.com/xenomote/etude/internal/parser"
)

func TestLexer(t *testing.T) {
	s := "====?><>!=:=.,"
	l := parser.New(strings.NewReader(s))

	for {
		tok, err := l.Next()
		if err != nil {
			t.Log(err)
			break
		}
		t.Log(tok, l.Text())
	}
}