lexer grammar VLexer;
//  ══════════════════════════════════════════
//          COMENTARIOS Y ESPACIOS
//  ══════════════════════════════════════════
WS:                         [ \t\r\n]+ -> skip;
COMENTARIO:                 '//' .*? ('\n' | EOF) -> skip;
COMENTARIOMULT:             '/*' .*? '*/' -> skip;

//  ══════════════════════════════════════════
//      DECLARACIONES Y ASIGNACIONES
//  ══════════════════════════════════════════
RW_MUT:                    'mut';
OP_ASSIGN:                 ':=';
OP_INCREMENTO:             '+=';
OP_DECREMENTO:             '-=';
OP_DECLARACION:            '=';

//  ══════════════════════════════════════════
//              TIPOS DE DATOS
//  ══════════════════════════════════════════
RW_INT:                     'int';
RW_FLOAT64:                 'float64'| 'f64';
RW_STRING:                  'string';
RW_BOOL:                    'bool';

//  ══════════════════════════════════════════
//              LITERALES
//  ══════════════════════════════════════════
INT_LITERAL:                [0-9]+;
FLOAT_LITERAL:              [0-9]+ '.' [0-9]+;
STRING_LITERAL:             '"' (~["\r\n\\] | ESC_SEQ)* '"';
BOOL_LITERAL:               'true' | 'false';
NIL_LITERAL:                'nil';

//  ══════════════════════════════════════════
//             OPERADORES
//  ══════════════════════════════════════════
// Aritméticos
OP_SUMA:                   '+';
OP_RESTA:                  '-';
OP_MULT:                   '*';
OP_DIV:                    '/';
OP_MOD:                    '%';
// Relacionales
OP_IGUAL:                  '==';
OP_DIFERENTE:              '!=';
OP_MENORQ:                 '<';
OP_MAYORQ:                 '>';
OP_MENORIGUAL:             '<=';
OP_MAYORIGUAL:             '>=';
// Lógicos
OP_AND:                    '&&';
OP_OR:                     '||';
OP_NOT:                    '!';

//  ══════════════════════════════════════════
//             PALABRAS RESERVADAS
//  ══════════════════════════════════════════
// Estructura del programa
RW_MAIN:                'main';
RW_FN:                  'fn';
RW_STRUCT:              'struct';
// Control de flujo
RW_IF:                  'if';
RW_ELSE:                'else';
RW_SWITCH:              'switch';
RW_CASE:                'case';
RW_DEFAULT:             'default';
RW_FOR:                 'for';
RW_WHILE:               'while';
RW_BREAK:               'break';
RW_CONTINUE:            'continue';
RW_RETURN:              'return';
RW_IN:                  'in';


//  ══════════════════════════════════════════
//              IDENTIFICADORES
//  ══════════════════════════════════════════
ID:                 [_a-zA-Z] [_a-zA-Z0-9]*;

//  ══════════════════════════════════════════
//              AGRUPACIÓN Y DELIMITADORES
//  ══════════════════════════════════════════
PAR_IZQ:                '(';
PAR_DER:                ')';
LLAVE_IZQ:              '{';
LLAVE_DER:              '}';
CORCHETE_IZQ:           '[';
CORCHETE_DER:           ']';
PUNTO:                  '.';
COMA:                   ',';
PUNTO_Y_COMA:           ';';
DOS_PUNTOS:             ':';

//  ══════════════════════════════════════════
//                  ERRORES
//  ══════════════════════════════════════════
fragment ESC_SEQ:          '\\' [btnfr"'\\];