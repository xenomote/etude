```ebnf
program
    = [ [ "-" ]? [ comp | typedef | func ] ]+

comp
    = "comp" name expression

typedef
    = "type" name type

func
    = "func" field [ ":" type ]? [ type_constructor ]? block

block
    = "{" [ statement ]+ "}"

statement
    = if
    = on
    = for
    = comp
    = assign
    = "return" [ expression ]?

if
    = "if" expression block [ "or" expression block ]* [ "or" block ]?

on
    = "on" expression "{" [ expression [ "," expression ]* [ "," ]? block ]+ [ "or" block ]? "}"

for
    = "for" [ [ assign "," ]? expression [ "," assign ]? ]? [ block ]?

assign
    = expression "=" expression

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
    = expression_constructor

reference
    = [ "~" ]? [ "#" ]? path [ "?" ]?

path
    =  name [ "." name ]*

expression_constructor
    = "(" expression_fields ")"   // list/tuple/object/map literal, brackets

expression_fields
    = expression_field [ "," expression_field ]* [ "," ]?

expression_field
    = [ [ field ]? ":" ]? expression

type
    = [ "comp" ]? inner_type

inner_type
    = path                  // named type
    = type_constructor      // list/tuple/object type 
    = "[" type ":" type "]" // map/function type

type_constructor
    = "[" type_fields "]"

type_fields
    = type_field [ "," type_field ]* [ "," ]?

type_field
    = [ "-" ]? [ field ":" ]? type

field
    = [ "~" ]? [ "#" ]? name [ "?" ]?
```