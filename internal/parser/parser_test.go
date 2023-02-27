package parser_test

import (
	"fmt"
	"io"

	"github.com/xenomote/etude/internal/lexer"
	"github.com/xenomote/etude/internal/parser"
	"github.com/xenomote/etude/internal/token"
)

const s =
`
comp a (x: a)
`

func Example() {
	p := parser.New()
	l := lexer.New()

	_, err := io.WriteString(l, s)
	if err != nil {
		panic(err)
	}

	var ts []token.Token
	for {
		t, err := l.Next()
		if err != nil {
			panic(err)
		}
		ts = append(ts, t)

		if t.Kind == token.END {
			break
		}
	}
	fmt.Println(ts)
	
	p.Write(ts)
	err = p.Program()
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Output)
	// Output:
}