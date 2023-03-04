```ebnf
program
    = { "-"? ( comp / typedef / func ) }+

block
    = "{" { statement } "}"

statement
    = if
    = on
    = for
    = comp
    = typedef
    = func
    = assign
    = return



if
    = "if" expression block { "or" expression block } ( "or" block )?

on
    = "on" expression "{" { expression block }+ ( "or" block )? "}"

for
    = "for" ( ( assign "," )? expression ( "," assign )? )? block?

assign
    = expression "=" expression

return
    = "return" expression?

comp
    = "comp" name expression

typedef
    = "type" name type

func
    = "func" ref_name ( ":" type )? type_constructor? block



expression
    = "copy"? _expression

_expression
    = literal
    = ref_path
    = expression_constructor
    = expression_operator

expression_operator
    = operand { op_infix operand }

operand
    = op_prefix? expression op_suffix?
   
op_prefix
    = ("-" / "!" / "@")

op_suffix
    = ("...")

op_infix
	= ( "+" / "-" / "*" / "/" / "^" / "%" )     (* arithmetic       *)
    = ( "==" / "!=" / "<" / ">" / "<=" / ">=")  (* comparison       *)
    = ( "&&" / "||")                            (* boolean          *)
    = ( "<<" / ">>" )                           (* object/map merge *)

literal
    = number
    = string
    = boolean

expression_constructor
    = "(" expression_field { "," expression_field } ")"   (* list/tuple/object/map literal, brackets *)

expression_field
    = ":" ref_name 
    = expression ( ":" ref_name )? 



type
    = "comp"? _type

_type
    = path              (* named type               *)
    = type_constructor  (* list/tuple/object type   *)
    = type_map          (* map/function type        *)

type_map
    = [" type "/" type "]"

type_constructor
    = "[" type_field { "," type_field } "]"

type_field
    = "-"?  type ( ":" ref_name )?



ref_name
    = "~"? "#"? name "?"?

ref_path
    = "~"? "#"? path "?"?

path
    = name { "." name }
```
