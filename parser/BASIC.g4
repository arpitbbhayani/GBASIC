grammar BASIC;

program
    : statement* EOF
    ;

statement
    : printStatement
    ;

printStatement
    : NUMBER SPACES PRINT SPACES expression
    ;

expression
    : NUMBER
    | STRING
    ;


fragment ESC_SEQ
    : '\\' ('b'|'t'|'n'|'f'|'r'|'"'|'\''|'\\')
    ;

fragment DIGIT
    : [0-9]
    ;

fragment SPACE
    : ' '
    ;


PRINT: 'PRINT';

NUMBER
    : DIGIT+
    ;

STRING
    : '"' (ESC_SEQ | ~["\\])* '"'
    ;

SPACES
    : SPACE+
    ;

WHITESPACE
    : [ \t\r\n] -> skip
    ;
