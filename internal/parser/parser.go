package parser

import (
	"fmt"

	"github.com/xenomote/etude/internal/production"
	"github.com/xenomote/etude/internal/token"
)

func New() parser {
	return parser{}
}

func (p *parser) Write(tokens []token.Token) {
	p.Tokens = append(p.Tokens, tokens...)
}

type parser struct {
	Tokens []token.Token
	States []state

	Output *production.Production
}

type state struct {
	prod   production.Production
	offset int
}

func (p *parser) top() *state {
	return &p.States[len(p.States)-1]
}

func (s *state) push(prod any) {
	s.prod.Productions = append(s.prod.Productions, prod)
}

/*
	state management
*/

type parse func() error

func (p *parser) start(kind production.Kind) {
	offset := p.right()

	p.States = append(p.States, state{})

	top := p.top()
	top.prod.Kind = kind
	top.offset = offset
}

func (p *parser) done() error {
	p.print("+++")

	done := p.top()
	p.States = p.States[:len(p.States)-1]

	if len(p.States) < 1 {
		p.Output = &done.prod
		return nil
	}

	next := p.top()

	next.offset = done.offset
	next.push(done.prod)

	return nil
}

type ParseError struct {
	error
}

func (p ParseError) Error() string {
	if p.error != nil {
		return p.error.Error()
	}

	return "failed to parse"
}

func (p *parser) fail(err error) error {
	p.print("---")

	p.States = p.States[:len(p.States)-1]

	if pe, ok := err.(ParseError); ok {
		return pe
	}

	return ParseError{err}
}

/*
	stream management
*/

func (p *parser) peek() token.Token {
	return p.Tokens[p.top().offset]
}

func (p *parser) take() {
	top := p.top()

	top.push(p.peek())
	top.offset++
}

func (p *parser) left() int {
	if len(p.States) < 2 {
		return 0
	}

	return p.States[len(p.States)-2].offset
}

func (p *parser) right() int {
	if len(p.States) < 1 {
		return 0
	}

	return p.top().offset
}

/*
	debug
*/

func (p *parser) print(msg string) {
	for _, s := range p.States {
		fmt.Print(s.prod.Kind, " ")
	}

	fmt.Print(msg, " ")

	l := p.left()
	r := p.right()
	
	for i := l; i < r; i++ {
		fmt.Print(p.Tokens[i].Kind, " ")
	}

	fmt.Println()
}