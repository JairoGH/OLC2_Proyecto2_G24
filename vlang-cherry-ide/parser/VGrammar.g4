parser grammar VGrammar;

options {
    tokenVocab = VLexer;	
}

// ═══════════════════════════ PROGRAMA PRINCIPAL ═══════════════════════════
program: (stmt)* main_func? EOF?;

main_func: 
    RW_FN RW_MAIN PAR_IZQ PAR_DER LLAVE_IZQ stmt* LLAVE_DER # FuncionMain;

// ═══════════════════════════ SENTENCIAS GENERALES ═══════════════════════════
stmt:
      stmt_declaracion
    | stmt_asignar
    | stmt_transferencia
    | if_stmt
    | switch_stmt
    | for_clasico_stmt
    | while_stmt
    | for_stmt
    | llamar_funcion
    | fun_slice
    | declarar_funcion
 	| declarar_struct;

// ═══════════════════════════ DECLARACIONES DE VARIABLES ═══════════════════════════
stmt_declaracion: 
      RW_MUT ID OP_ASSIGN expr                 # DeclararInferenciaMut
  	| ID OP_ASSIGN expr                        # DeclararInferencia
  	| RW_MUT ID tipo OP_DECLARACION expr       # DeclaraTipoValor
  	| RW_MUT ID tipo                           # DeclararTipo
  	| ID tipo OP_DECLARACION expr              # DeclararSinMutValor
  	| RW_MUT? ID OP_DECLARACION  tipo expr?    # DeclararSlice;

// ═══════════════════════════ ASIGNACIONES ═══════════════════════════
stmt_asignar: 
      item_slice OP_DECLARACION expr    									# AsignacionSliceItem
    | patronId OP_DECLARACION expr											# AsignacionDirecta
    | patronId op = (OP_INCREMENTO | OP_DECREMENTO) expr				    # AsignacionAritmetica
    | item_slice op = (OP_INCREMENTO | OP_DECREMENTO | OP_DECLARACION) expr	# AsignacionSlice;

// ═══════════════════════════ SENTENCIAS DE CONTROL ═══════════════════════════
stmt_transferencia:
      RW_RETURN expr?	# ReturnStmt
    | RW_BREAK		# BreakStmt
    | RW_CONTINUE	# ContinueStmt;

// ═══════════════════════════ ESTRUCTURAS CONDICIONALES ═══════════════════════════
if_stmt: 
    if_chain (RW_ELSE if_chain)* else_stmt? # IFstmt;

if_chain: 
    RW_IF expr LLAVE_IZQ stmt* LLAVE_DER # IFcadena;

else_stmt: 
    RW_ELSE LLAVE_IZQ stmt* LLAVE_DER # ElseStmt; 

switch_stmt:
    RW_SWITCH expr LLAVE_IZQ switch_case* default_case? LLAVE_DER # SwitchStmt;

switch_case: 
    RW_CASE expr DOS_PUNTOS stmt* # SwitchCase;

default_case: 
    RW_DEFAULT DOS_PUNTOS stmt* # DefaultCase;

// ═══════════════════════════ ESTRUCTURAS DE REPETICIÓN ═══════════════════════════
while_stmt: 
    RW_FOR expr LLAVE_IZQ stmt* LLAVE_DER # WhileStmt;

for_clasico_stmt:
    RW_FOR stmt_asignar? PUNTO_Y_COMA expr? PUNTO_Y_COMA stmt_asignar? LLAVE_IZQ stmt* LLAVE_DER # ForClasicoStmt;

for_stmt:
    RW_FOR ID (COMA ID)? RW_IN (expr | range) LLAVE_IZQ stmt* LLAVE_DER # ForStmt;

range: expr PUNTO PUNTO PUNTO expr # RangoNum;

// ═══════════════════════════ FUNCIONES ═══════════════════════════
declarar_funcion:
    RW_FN ID PAR_IZQ lista_parametros? PAR_DER tipo? LLAVE_IZQ stmt* LLAVE_DER    # FuncionDeclerada;

llamar_funcion: patronId PAR_IZQ lista_argumentos? PAR_DER # LlamarFuncion;

lista_parametros: 
    parametro_fun (COMA parametro_fun)* COMA? # ListaParametros;

parametro_fun: 
    ID? ID tipo # ParametroFun;

lista_argumentos: 
    argumento_fun (COMA argumento_fun)* COMA? # ListaArgumentos;

argumento_fun: 
    (ID DOS_PUNTOS)? (patronId | expr) # FuncionArg; 

// ═══════════════════════════ ESTRUCTURAS (STRUCTS) ═══════════════════════════
declarar_struct: 
    RW_STRUCT ID LLAVE_IZQ propiedad_struct* LLAVE_DER # DeclararStruct;

propiedad_struct:
    tipo ID (OP_DECLARACION expr)?              # PropiedadStruct;

crear_struct: 
    ID LLAVE_IZQ lista_parametros_init? LLAVE_DER # CrearStruct;

lista_parametros_init: 
    parametros_init_struct (COMA parametros_init_struct)* COMA? # ListaParametrosInit;

parametros_init_struct: 
    ID DOS_PUNTOS expr # ParametrosInitStruct;

// ═══════════════════════════ SLICES/VECTORES ═══════════════════════════
lista_slice:
    LLAVE_IZQ (expr (COMA expr)*)? COMA? LLAVE_DER # ListaSlice;

item_slice: 
    patronId (CORCHETE_IZQ expr CORCHETE_DER)+ # ItemSlice;

prop_slice: 
    item_slice PUNTO patronId # PropSlice;

fun_slice: 
    item_slice PUNTO llamar_funcion # FuncionSlice;

// ═══════════════════════════ TIPOS DE DATOS ═══════════════════════════
tipo:
    ID | RW_INT | RW_FLOAT64 | RW_STRING | RW_BOOL | tipos_slices ;

tipos_slices: 
      CORCHETE_IZQ CORCHETE_DER tipo                           # VectorSimple   
    | CORCHETE_IZQ CORCHETE_DER CORCHETE_IZQ CORCHETE_DER tipo # MatrizDoble;

// ═══════════════════════════ EXPRESIONES ═══════════════════════════
expr:
     PAR_IZQ expr PAR_DER								# ParentecisExp 
    | llamar_funcion								    # LlamarFuncionExp 
    | patronId											# IdExp 
    | item_slice										# ItemSliceExp 
    | prop_slice										# PropSliceExp 
    | fun_slice										    # VectorFuncExp 
    | literal											# LiteralExp 
    | lista_slice										# SliceExp 
    | crear_struct                                       # CrearStructExp
    | op = (OP_NOT | OP_RESTA) expr							 # UnarioExp 	
    | left = expr op = (OP_MULT | OP_DIV | OP_MOD) right = expr	# BinarioExp
    | left = expr op = (OP_SUMA | OP_RESTA) right = expr		# BinarioExp 
    | left = expr op = (OP_MENORQ | OP_MENORIGUAL | OP_MAYORQ | OP_MAYORIGUAL ) right = expr# BinarioExp 
    | left = expr op = (OP_IGUAL | OP_DIFERENTE) right = expr	# BinarioExp 
    | left = expr op = OP_AND right = expr						# BinarioExp 
    | left = expr op = OP_OR right = expr						# BinarioExp;

// ═══════════════════════════ LITERALES Y PATRONES ═══════════════════════════
literal:
     INT_LITERAL		    # IntLiteral 
    | FLOAT_LITERAL		# FloatLiteral
    | STRING_LITERAL	# StringLiteral 
    | BOOL_LITERAL		# BoolLiteral   
    | NIL_LITERAL		# NilLiteral; 

patronId: 
    ID (PUNTO ID)* # ID_Patron;