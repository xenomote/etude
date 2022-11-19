package parser

type Program struct {
	funcs   []Func
	defines []Define
}

type Define struct {
	name string
	typ  Type
	exp  Expression
}

type Assign struct {
	to   Expression
	from Expression
}

type Func struct {
	name string
	reciever Type
	argument Type
	statements []Statement
}

type If struct {
	cases []Case
	alternative []Statement
}

type On struct {
	condition Condition
	cases []Case
	alternative []Statement
}

type Case struct {
	condition Condition
	statements []Statement
}

type For struct {
	definitions []Define
	expression Expression
	assignments []Assign
	statements []Statement
}

type Statement interface{}

type Condition interface{}

type Expression interface{}

type Literal interface{}

type Type interface{}

type Token interface {
	String() string
}

type Lexer interface {
	Want(...Token) (Token, error)
	Next() (Token, error)
}

func Parse(l Lexer) (*Program, error) {
	return nil, nil
}
