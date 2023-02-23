```ebnf
program
    = [ func | assign ]+
 
assign
    = expression "=" expression

func
    = "func" name [ type ]? ":" [ type ]? block

if
    = "if" expression block [ "or" expression block ]* [ "or" block ]?

on
    = "on" expression "{" [ expression [ "," expression ]* [ "," ]? block ]+ [ "or" block ]? "}"

for
    = "for" [ [ assign "," ]* expression [ "," assign ]* ]? [ block ]?

block
    = "{" [ statement ]+ "}"

statement
    = if
    = on
    = for
    = assign
    = "return" [ expression ]?

expression
    = literal
    = expression constructor
    = op_prefix expression
    = expression op_suffix
    = expression [ op_infix expression ]+
   
op_prefix
    = [ "-" | "!" | "@" ]
    
op_suffix
    = [ "++" | "--" | "..." ]

op_infix
    = [ "+" | "-" | "*" | "/" | "^" | "%" | "==" | "!=" | "&&" | "||" | "#" ]

literal
    = reference
    = number
    = string
    = boolean
    = constructor

reference
    = [ "~" ]? [ "#" ]? path [ "?" ]?

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
    = "[" type_list "]"     // list type/tuple type
    = "[" type_fields "]"   // object type
    = "[" type ":" type "]" // map type/function type
 
path
    =  name [ "." name ]*

type_list
    = type [ "," type ]* [ "," ]?

type_fields
    = field ":" type [ "," field ":" type ]* [ "," ]?

field
    = [ "~" ]? [ "#" ]? name [ "?" ]?
```
