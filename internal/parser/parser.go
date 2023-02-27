package parser

import (
	"github.com/xenomote/etude/internal/production"
	"github.com/xenomote/etude/internal/token"
)

func (p *Parser) Write(tokens []token.Token)  {
	p.tokens = append(p.tokens, tokens...)
}

type Parser struct {
	tokens []token.Token
	states []state

	output *production.Production
}

type state struct {
	prod production.Production
	offset int
}

func (p *Parser) top() *state {
	return &p.states[len(p.states)-1]
}

func (s *state) push(prod any) {
	s.prod.Productions = append(s.prod.Productions, prod)
}

/*
	state management
*/

type parse func()error

func (p *Parser) start(kind production.Kind) {
	offset := p.top().offset

	p.states = append(p.states, state{})

	top := p.top()
	top.prod.Kind = kind
	top.offset = offset
}

func (p *Parser) done() error {
	prod := p.top().prod
	p.states = p.states[:len(p.states)-1]

	if len(p.states) < 1 {
		p.output = &prod
		return nil
	}

	p.top().push(prod)

	return nil
}

func (p *Parser) fail(err error) error {
	p.states = p.states[:len(p.states)-1]
	return err
}

/*
	stream management
*/

func (p *Parser) peek() token.Token {
	return p.tokens[p.top().offset]
}

func (p *Parser) take() {
	top := p.top()

	top.push(p.peek())
	top.offset++
}