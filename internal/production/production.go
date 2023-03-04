package production

import "fmt"

type Production struct {
	Kind Kind
	Productions []any
}

func (p Production) String() string {
	return fmt.Sprint(p.Productions...)
}

//go:generate stringer -type Kind
type Kind uint

const (
	ERROR Kind = iota

	PROGRAM
	BLOCK
	STATEMENT

	FUNC
	COMP
	TYPEDEF

	IF
	ON
	FOR
	ASSIGN
	RETURN

	EXPRESSION
	EXPRESSION_OPERATOR
	EXPRESSION_CONSTRUCTOR
	EXPRESSION_FIELD

	TYPE
	TYPE_CONSTRUCTOR
	TYPE_FIELD
	TYPE_MAP

	OPERAND
	OP_PREFIX
	OP_SUFFIX
	OP_INFIX

	LITERAL

	REF_NAME
	REF_PATH
	PATH
)
