grammar Calc;

start : expr EOF ;

expr : expr '+' expr # Addition
     | NUMBER        # Number ;

NUMBER     : [0-9]+ ;
WHITESPACE : [ \r\n\t]+ -> skip ;
