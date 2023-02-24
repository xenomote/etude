```ebnf
program
    = [ assign | func ]+

assign
    = expression "=" expression

func
    = "func" field [ ":" type ]? [ object_type ]? block
    = "func" field object_type block
    = "func" field "[" "->" type "]" block
    = "func" field "[" object_type "->" type "]"

block
    = "{" [ statement ]+ "}"

statement
    = if
    = on
    = for
    = assign
    = "return" [ expression ]?

if
    = "if" expression block [ "or" expression block ]* [ "or" block ]?

on
    = "on" expression "{" [ expression [ "," expression ]* [ "," ]? block ]+ [ "or" block ]? "}"

for
    = "for" [ [ assign "," ]? expression [ "," assign ]? ]? [ block ]?

expression
    = literal
    = expression constructor
    = op_prefix expression
    = expression op_suffix
    = expression [ op_infix expression ]+
   
op_prefix
    = [ "-" | "!" | "@" ]
    
op_suffix
    = [ "..." ]

op_infix
    = [ "+" | "-" | "*" | "/" | "^" | "%" ] // arithmetic
    = [ "==" | "!=" | "&&" | "||" ]         // boolean
    = [ "<<" | ">>"]                        // object/map merge

literal
    = reference
    = number
    = string
    = boolean
    = constructor

reference
    = [ "~" ]? [ "#" ]? path [ "?" ]?

path
    =  name [ "." name ]*

constructor
    = "(" expression_list ")"   // list literal/tuple literal/expression brackets
    = "(" expression_fields ")" // object literal/map literal

expression_list
    = expression [ "," expression ]* [ "," ]?

expression_fields
    = expression_field [ "," expression_field ]* [ "," ]?

expression_field
    = field ":" expression
    = ":" field

type
    = path                  // named type
    = object_type
    = "[" type_list "]"     // list type/tuple type
    = "[" type ":" type "]" // map type/function type

object_type
    = "[" type_fields "]"

type_list
    = type [ "," type ]* [ "," ]?

type_fields
    = field ":" type [ "," field ":" type ]* [ "," ]?

field
    = [ "~" ]? [ "#" ]? name [ "?" ]?
```