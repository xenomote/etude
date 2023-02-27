package parser

import (
	"github.com/xenomote/etude/internal/production"
	"github.com/xenomote/etude/internal/token"
)

func (p *parser) Program() error {
	p.start(production.PROGRAM)

	defs := []parse{
		p.Comp,
		p.TypeDef,
		p.Func,
	}

	loop:
	for {
		for _, def := range defs {
			err := def()
			if err != nil {
				break loop
			}
		}
	}
	
	if p.peek().Kind != token.END {
		return p.fail(nil)
	}

	return p.done()
}

func (p *parser) Comp() error {
	p.start(production.COMP)

	if p.peek().Kind != token.COMP {
		return p.fail(nil)
	}
	p.take()

	if p.peek().Kind != token.IDENTIFIER {
		return p.fail(nil)
	}
	p.take()

	err := p.Expression()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) TypeDef() error {
	p.start(production.TYPEDEF)

	if p.peek().Kind != token.TYPE {
		return p.fail(nil)
	}
	p.take()

	if p.peek().Kind != token.IDENTIFIER {
		return p.fail(nil)
	}
	p.take()

	err := p.Type()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) Func() error {
	p.start(production.FUNC)

	if p.peek().Kind != token.FUNC {
		return p.fail(nil)
	}
	p.take()

	err := p.RefName()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind == token.COLON {
		p.take()

		err := p.Type()
		if err != nil {
			return p.fail(err)
		}
	}

	err = p.TypeConstructor()
	if err != nil {
		return p.fail(err)
	}

	err = p.Block()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) Block() error {
	p.start(production.BLOCK)

	if p.peek().Kind != token.CURLY_LEFT {
		return p.fail(nil)
	}
	p.take()

	for {
		if p.peek().Kind == token.CURLY_RIGHT {
			p.take()

			return p.done()
		}

		err := p.Statement()
		if err != nil {
			return p.fail(err)
		}
	}
}

func (p *parser) Statement() error {
	p.start(production.STATEMENT)

	stmts := []parse{
		p.Comp,
		p.TypeDef,
		p.Func,

		p.If,
		p.On,
		p.For,
		p.Return,

		// last because it would require the most backtracking
		p.Assign,
	}

	for _, f := range stmts {
		err := f()
		if err == nil {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) If() error {
	p.start(production.IF)

	if p.peek().Kind != token.IF {
		return p.fail(nil)
	}
	p.take()

	err := p.Expression()
	if err != nil {
		return p.fail(err)
	}

	err = p.Block()
	if err != nil {
		return p.fail(err)
	}

	for {
		if p.peek().Kind != token.OR {
			break
		}
		p.take()

		err := p.Block()
		if err == nil {
			break
		}

		err = p.Expression()
		if err != nil {
			return p.fail(err)
		}

		err = p.Block()
		if err != nil {
			return p.fail(err)
		}
	}

	return p.done()
}

func (p *parser) On() error {
	p.start(production.ON)

	err := p.Expression()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind != token.CURLY_LEFT {
		return p.fail(nil)
	}
	p.take()

loop:
	for {
		switch p.peek().Kind {
		case token.OR, token.CURLY_RIGHT:
			p.take()

			break loop
		}

		err := p.Expression()
		if err != nil {
			return p.fail(err)
		}

		err = p.Block()
		if err != nil {
			return p.fail(err)
		}
	}

	if p.peek().Kind == token.OR {
		p.take()

		err := p.Block()
		if err != nil {
			return p.fail(err)
		}
	}

	if p.peek().Kind != token.CURLY_RIGHT {
		return p.fail(nil)
	}
	p.take()

	return p.done()
}

func (p *parser) For() error {
	p.start(production.FOR)

	if p.peek().Kind != token.FOR {
		return p.fail(nil)
	}
	p.take()

	err := p.Block()
	if err == nil {
		return p.done()
	}

	err = p.Assign()
	if err == nil {
		if p.peek().Kind != token.COMMA {
			return p.fail(nil)
		}
		p.take()
	}

	err = p.Expression()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind == token.COMMA {
		p.take()

		err := p.Assign()
		if err != nil {
			return p.fail(nil)
		}
	}

	err = p.Block()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) Return() error {
	p.start(production.RETURN)

	if p.peek().Kind != token.RETURN {
		return p.fail(nil)
	}
	p.take()

	p.Expression()

	return p.done()
}

func (p *parser) Assign() error {
	p.start(production.ASSIGN)

	err := p.Expression()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind != token.EQUALS {
		return p.fail(nil)
	}

	err = p.Expression()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) Expression() error {
	p.start(production.EXPRESSION)

	exprs := []func() error{
		p.Literal,
		p.RefPath,
		p.ExpressionConstructor,
		p.ExpressionOperator,
	}

	for _, expr := range exprs {
		err := expr()
		if err == nil {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) ExpressionOperator() error {
	p.start(production.EXPRESSION_OPERATOR)

	for {
		err := p.Operand()
		if err != nil {
			return p.fail(err)
		}

		err = p.OpInfix()
		if err != nil {
			return p.done()
		}
	}
}

func (p *parser) Operand() error {
	p.start(production.OPERAND)

	p.OpPrefix()

	oprs := []parse{
		p.RefPath,
		p.Literal,
		p.ExpressionConstructor,
		// no Expression or ExpressionOperator to avoid recursion
	}

	for _, opr := range oprs {
		err := opr()
		if err == nil {
			goto success
		}

	}
	return p.fail(nil)

success:
	p.OpSuffix()

	return p.done()
}

func (p *parser) OpPrefix() error {
	p.start(production.OP_PREFIX)

	prefs := []token.Kind{
		token.MINUS,
		token.EXCLAIM,
		token.ADDRESS,
	}

	for _, pref := range prefs {
		if p.peek().Kind == pref {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) OpSuffix() error {
	p.start(production.OP_SUFFIX)

	suffs := []token.Kind{
		token.ELLIPSIS,
	}

	for _, suff := range suffs {
		if p.peek().Kind == suff {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) OpInfix() error {
	p.start(production.OP_INFIX)

	infs := []token.Kind{
		token.PLUS,
		token.MINUS,
		token.STAR,
		token.SLASH,
		token.CARRET,
		token.PERCENT,

		token.DOUBLE_EQUALS,
		token.EXCLAIM_EQUALS,
		token.LESS,
		token.LESS_EQUALS,
		token.GREATER,
		token.GREATER_EQUALS,

		token.DOUBLE_AMPERSAND,
		token.DOUBLE_PIPE,

		token.SHIFT_LEFT,
		token.SHIFT_RIGHT,
	}

	for _, inf := range infs {
		if p.peek().Kind == inf {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) Literal() error {
	p.start(production.LITERAL)

	lits := []token.Kind{
		token.IDENTIFIER,
		token.STRING,
		token.NUMBER,
	}

	for _, lit := range lits {
		if p.peek().Kind == lit {
			p.take()
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) ExpressionConstructor() error {
	p.start(production.EXPRESSION_CONSTRUCTOR)

	if p.peek().Kind != token.ROUND_LEFT {
		return p.fail(nil)
	}
	p.take()

	for {
		err := p.ExpressionField()
		if err != nil {
			return p.fail(err)
		}

		if p.peek().Kind != token.COMMA {
			break
		}
		p.take()
	}

	if p.peek().Kind != token.ROUND_RIGHT {
		return p.fail(nil)
	}
	p.take()

	return p.done()
}

func (p *parser) ExpressionField() error {
	p.start(production.EXPRESSION_FIELD)

	err := p.RefName()
	if p.peek().Kind == token.COLON {
		if err != nil {
			return p.fail(nil) // should not have been a colon, ref is fine
		}

		p.take()
	}

	err = p.Expression()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) Type() error {
	p.start(production.TYPE)

	if p.peek().Kind == token.COMP {
		p.take()
	}

	typs := []parse{
		p.Path,
		p.TypeConstructor,
		p.TypeMap,
	}

	for _, typ := range typs {
		err := typ()
		if err == nil {
			return p.done()
		}
	}

	return p.fail(nil)
}

func (p *parser) TypeConstructor() error {
	p.start(production.TYPE_CONSTRUCTOR)

	if p.peek().Kind != token.SQUARE_LEFT {
		return p.fail(nil)
	}
	p.take()

	for {
		err := p.TypeField()
		if err != nil {
			return p.fail(err)
		}

		if p.peek().Kind != token.COMMA {
			break
		}
		p.take()
	}

	if p.peek().Kind != token.SQUARE_RIGHT {
		return p.fail(nil)
	}
	p.take()

	return p.done()
}

func (p *parser) TypeField() error {
	p.start(production.TYPE_FIELD)

	if p.peek().Kind == token.MINUS {
		p.take()
	}

	err := p.RefName()
	if err == nil {
		if p.peek().Kind != token.COLON {
			return p.fail(nil)
		}
		p.take()
	}

	err = p.Expression()
	if err != nil {
		return p.fail(err)
	}

	return p.done()
}

func (p *parser) TypeMap() error {
	p.start(production.TYPE_MAP)

	if p.peek().Kind != token.SQUARE_LEFT {
		return p.fail(nil)
	}
	p.take()

	err := p.Type()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind != token.SLASH {
		return p.fail(nil)
	}
	p.take()

	err = p.Type()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind != token.SQUARE_RIGHT {
		return p.fail(nil)
	}
	p.take()

	return p.done()
}

func (p *parser) RefName() error {
	p.start(production.REF_NAME)

	if p.peek().Kind == token.TILDE {
		p.take()
	}

	if p.peek().Kind == token.HASH {
		p.take()
	}

	if p.peek().Kind != token.IDENTIFIER {
		return p.fail(nil)
	}
	p.take()

	if p.peek().Kind == token.QUESTION {
		p.take()
	}

	return p.done()
}

func (p *parser) RefPath() error {
	p.start(production.REF_PATH)

	if p.peek().Kind == token.TILDE {
		p.take()
	}

	if p.peek().Kind == token.HASH {
		p.take()
	}

	err := p.Path()
	if err != nil {
		return p.fail(err)
	}

	if p.peek().Kind == token.QUESTION {
		p.take()
	}

	return p.done()
}

func (p *parser) Path() error {
	p.start(production.PATH)

	for {
		if p.peek().Kind != token.IDENTIFIER {
			return p.fail(nil)
		}
		p.take()

		if p.peek().Kind != token.PERIOD {
			return p.done()
		}
		p.take()
	}
}
