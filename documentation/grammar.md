```ebnf
program :=
    import_declaration_list toplevel_definition_list

import_declaration_list :=
    import_declaration  import_declaration_list
    ~

import_declaration :=
    'import' LITERAL_STRING

toplevel_definition_list :=
    toplevel_definition toplevel_definition_list
    ~

toplevel_definition :=
    constant_definition
    function_definition

constant_definition :=
    IDENTIFIER ':=' expression

expression :=
    literal
    structure_value
    prefix_operator expression
    expression binary_operator expression

literal :=
    LITERAL_BOOLEAN
    LITERAL_NUMBER
    LITERAL_STRING

function_definition :=
    'func' structure_type structure_type '{' statement_list '}'
```