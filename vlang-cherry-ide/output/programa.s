.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "=== Archivo de prueba de slices ==="
    msg_2: .asciz "==== Creación de slices ===="
    msg_3: .asciz "Creación con literales:"
    msg_4: .asciz "###Validacion Manual"
    msg_5: .asciz "numeros:"
    msg_6: .asciz "OK Creación con literales: correcto"
    msg_7: .asciz "\n==== Acceso de Elementos ===="
    msg_8: .asciz "Acceso por índice:"
    msg_9: .asciz "Primer elemento:"
    msg_10: .asciz "OK Acceso por índice: correcto"
    msg_11: .asciz "X Acceso por índice: incorrecto"
    msg_12: .asciz "\nModificación de elementos:"
    msg_13: .asciz "numeros después de modificar:"
    msg_14: .asciz "OK Modificación de elementos: correcto"
    msg_15: .asciz "X Modificación de elementos: incorrecto"
    msg_16: .asciz "\n==== Función indexOf ===="
    msg_17: .asciz "Búsqueda de elementos con indexOf:"
    msg_18: .asciz "Índice de 30:"
    msg_19: .asciz "Índice de 60:"
    msg_20: .asciz "OK indexOf: correcto"
    msg_21: .asciz "X indexOf: incorrecto"
    msg_22: .asciz "\n==== Función join ===="
    msg_23: .asciz "Unión de strings con join:"
    msg_24: .asciz "Hola"
    msg_25: .asciz "mundo"
    msg_26: .asciz "desde"
    msg_27: .asciz "Go"
    msg_28: .asciz " "
    msg_29: .asciz ", "
    msg_30: .asciz "Frase con espacios:"
    msg_31: .asciz "Frase con comas:"
    msg_32: .asciz "Hola mundo desde Go"
    msg_33: .asciz "Hola, mundo, desde, Go"
    msg_34: .asciz "OK join: correcto"
    msg_35: .asciz "X join: incorrecto"
    msg_36: .asciz "\n==== Función len ===="
    msg_37: .asciz "Longitud de slices con len:"
    msg_38: .asciz "Longitud de numeros:"
    msg_39: .asciz "Longitud de palabras:"
    msg_40: .asciz "OK len: correcto"
    msg_41: .asciz "X len: incorrecto"
    msg_42: .asciz "\n==== Función append ===="
    msg_43: .asciz "Agregar elementos con append:"
    msg_44: .asciz "numeros después de append(numeros3, 4):"
    msg_45: .asciz "OK Agregar un elemento: correcto"
    msg_46: .asciz "X Agregar un elemento: incorrecto"
    msg_47: .asciz "\n=== Tabla de Resultados ==="
    msg_48: .asciz "+----------------------------------+--------+-------+"
    msg_49: .asciz "| Característica                   | Puntos | Total |"
    msg_50: .asciz "+----------------------------------+--------+-------+"
    msg_51: .asciz "| Creación de slices               | "
    msg_52: .asciz "    | 6     |"
    msg_53: .asciz "| Acceso de Elementos              | "
    msg_54: .asciz "    | 6     |"
    msg_55: .asciz "| Función indexOf                  | "
    msg_56: .asciz "    | 1     |"
    msg_57: .asciz "| Función join                     | "
    msg_58: .asciz "    | 1     |"
    msg_59: .asciz "| Función len                      | "
    msg_60: .asciz "    | 3     |"
    msg_61: .asciz "| Función append                   | "
    msg_62: .asciz "    | 3     |"
    msg_63: .asciz "+----------------------------------+--------+-------+"
    msg_64: .asciz "| TOTAL                            | "
    msg_65: .asciz "   | 20    |"
    msg_66: .asciz "+----------------------------------+--------+-------+"
buffer_int: .skip 32
buffer_float: .skip 64
buffer_string: .skip 512
temp_buffer: .skip 256
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0
str_empty: .asciz ""
join_buffer_17: .skip 512
join_buffer_18: .skip 512

.section .text
.global _start

_start:
    bl fn_main      // Llamar a la función main
    mov x8, #93     // sys_exit como respaldo
    mov x0, #0      // exit status
    svc #0          // system call

// === PROCESANDO PROGRAMA ===
// === GENERANDO TODAS LAS FUNCIONES DE USUARIO ===
// Procesando función main
// === FUNCIÓN MAIN ===
fn_main:
        stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
        mov   x29, sp                 // Configurar frame pointer
// === DECLARAR VARIABLE MUT: puntos ===
    sub sp, sp, #8  // Reservar espacio para puntos
    mov x9, #0
    str x9, [sp]
// puntos en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_1
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_2
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosCreacion ===
    sub sp, sp, #8  // Reservar espacio para puntosCreacion
    mov x9, #0
    str x9, [sp]
// puntosCreacion en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_3
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARACIÓN SLICE numeros: []int ===
// Ajustar stack para slices: 256 bytes
    sub sp, sp, #256
    mov x28, x29
    sub x28, x28, #256
// Elemento 0 (literal): 1
    mov x0, #1
    str x0, [x28, #0]
// Elemento 1 (literal): 2
    mov x0, #2
    str x0, [x28, #8]
// Elemento 2 (literal): 3
    mov x0, #3
    str x0, [x28, #16]
// Elemento 3 (literal): 4
    mov x0, #4
    str x0, [x28, #24]
// Elemento 4 (literal): 5
    mov x0, #5
    str x0, [x28, #32]
// === IMPRIMIR STRING  ===
    adr x9, msg_4
    mov x0, x9
    bl print_string
    bl print_newline
// Verificando slice 'numeros': true
// Patrón ID 'numeros' encontrado como slice
// === IMPRIMIR STRING  ===
    adr x9, msg_5
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [x28, #0]    // cargar int[0] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #8]    // cargar int[1] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #16]    // cargar int[2] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #24]    // cargar int[3] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #32]    // cargar int[4] desde base fija
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_6
    mov x0, x9
    bl print_string
    bl print_newline
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosCreacion': false
// === CARGAR VARIABLE: puntosCreacion (offset: 0) ===
    ldr x9, [sp, #0]
    mov x10, x9
    mov x11, #6
    add x12, x10, x11
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosCreacion = int (offset: 0) ===
    str x12, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x9, msg_7
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosAcceso ===
    sub sp, sp, #8  // Reservar espacio para puntosAcceso
    mov x9, #0
    str x9, [sp]
// puntosAcceso en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_8
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: primerElemento ===
// Verificando slice 'numeros': true
// === ACCESO POR ÍNDICE numeros[0] ===
    ldr x9, [x28, #0]    // cargar numeros[0] desde base fija
    sub sp, sp, #8  // Reservar espacio para primerElemento
    str x9, [sp]
// primerElemento en [sp+0] 

// Verificando slice 'primerElemento': false
// === CARGAR VARIABLE: primerElemento (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_9
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'primerElemento': false
// === CARGAR VARIABLE: primerElemento (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #1
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_1
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosAcceso': false
// === CARGAR VARIABLE: puntosAcceso (offset: 8) ===
    ldr x13, [sp, #8]
    mov x14, x13
    mov x15, #2
    add x16, x14, x15
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosAcceso = int (offset: 8) ===
    str x16, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x17, msg_10
    mov x0, x17
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_0
.Lskip_branch_1:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x18, msg_11
    mov x0, x18
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_0:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_12
    mov x0, x9
    bl print_string
    bl print_newline
// Verificando slice 'numeros': true
// === ACCESO POR ÍNDICE numeros[0] ===
    ldr x9, [x28, #0]    // cargar numeros[0] desde base fija
// === ASIGNACIÓN POR ÍNDICE numeros[0] = 10 ===
// Asignación directa al elemento 0
    mov x10, #10
    str x10, [x28, #0]    // numeros[0] = 10
// Asignación a slice completada exitosamente
// Verificando slice 'numeros': true
// Patrón ID 'numeros' encontrado como slice
// === IMPRIMIR STRING  ===
    adr x9, msg_13
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [x28, #0]    // cargar int[0] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #8]    // cargar int[1] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #16]    // cargar int[2] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #24]    // cargar int[3] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #32]    // cargar int[4] desde base fija
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'numeros': true
// === ACCESO POR ÍNDICE numeros[0] ===
    ldr x9, [x28, #0]    // cargar numeros[0] desde base fija
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_3
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosAcceso': false
// === CARGAR VARIABLE: puntosAcceso (offset: 8) ===
    ldr x13, [sp, #8]
    mov x14, x13
    mov x15, #4
    add x16, x14, x15
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosAcceso = int (offset: 8) ===
    str x16, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x17, msg_14
    mov x0, x17
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_2
.Lskip_branch_3:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x18, msg_15
    mov x0, x18
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_2:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_16
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosIndex ===
    sub sp, sp, #8  // Reservar espacio para puntosIndex
    mov x9, #0
    str x9, [sp]
// puntosIndex en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_17
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARACIÓN SLICE numeros2: []int ===
// Elemento 0 (literal): 10
    mov x0, #10
    str x0, [x28, #40]
// Elemento 1 (literal): 20
    mov x0, #20
    str x0, [x28, #48]
// Elemento 2 (literal): 30
    mov x0, #30
    str x0, [x28, #56]
// Elemento 3 (literal): 40
    mov x0, #40
    str x0, [x28, #64]
// Elemento 4 (literal): 50
    mov x0, #50
    str x0, [x28, #72]
// === DECLARAR VARIABLE MUT: indice1 ===
// Verificando slice 'numeros2': true
// Patrón ID 'numeros2' encontrado como slice
// === FUNCIÓN indexOf(numeros2, 30) ===
// Buscando 30 en slice numeros2 (tamaño: 5)
    mov x12, #30
    mov x9, #0
indexOf_loop_numeros2_5:
    cmp x9, #5
    bge indexOf_not_found_numeros2_7
    add x10, x28, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros2_6
    add x9, x9, #1
    b indexOf_loop_numeros2_5
indexOf_found_numeros2_6:
    mov x11, x9
    b indexOf_end_numeros2_8
indexOf_not_found_numeros2_7:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros2_8:
// indexOf completado, resultado en x11
    sub sp, sp, #8  // Reservar espacio para indice1
    str x11, [sp]
// indice1 en [sp+0] 

// === DECLARAR VARIABLE MUT: indice2 ===
// Verificando slice 'numeros2': true
// Patrón ID 'numeros2' encontrado como slice
// === FUNCIÓN indexOf(numeros2, 60) ===
// Buscando 60 en slice numeros2 (tamaño: 5)
    mov x12, #60
    mov x9, #0
indexOf_loop_numeros2_9:
    cmp x9, #5
    bge indexOf_not_found_numeros2_11
    add x10, x28, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros2_10
    add x9, x9, #1
    b indexOf_loop_numeros2_9
indexOf_found_numeros2_10:
    mov x11, x9
    b indexOf_end_numeros2_12
indexOf_not_found_numeros2_11:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros2_12:
// indexOf completado, resultado en x11
    sub sp, sp, #8  // Reservar espacio para indice2
    str x11, [sp]
// indice2 en [sp+0] 

// Verificando slice 'indice1': false
// === CARGAR VARIABLE: indice1 (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x10, msg_18
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// Verificando slice 'indice2': false
// === CARGAR VARIABLE: indice2 (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_19
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA && ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'indice1': false
// === CARGAR VARIABLE: indice1 (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #2
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'indice2': false
// === CARGAR VARIABLE: indice2 (offset: 0) ===
    ldr x13, [sp, #0]
// === NEGACIÓN ARITMÉTICA - ===
    mov x14, #1
    mov x15, #0
    sub x15, x15, x14
// === OPERACIÓN RELACIONAL == ===
    mov x17, x13
    mov x18, x15
    cmp x17, x18
    cset x16, eq
// === FIN EXPRESIÓN BINARIA === 

// === OPERACIÓN LÓGICA && ===
    mov x10, x10
    cmp x10, #0
    beq .Land_false_13
    mov x11, x16
    and x9, x10, x11
    b .Land_end_14
.Land_false_13:
    mov x9, #0
.Land_end_14:
// === FIN EXPRESIÓN BINARIA === 

    cmp x9, #0
    beq .Lskip_branch_15
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosIndex': false
// === CARGAR VARIABLE: puntosIndex (offset: 16) ===
    ldr x12, [sp, #16]
    mov x13, x12
    mov x14, #1
    add x15, x13, x14
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosIndex = int (offset: 16) ===
    str x15, [sp, #16]
// === IMPRIMIR STRING  ===
    adr x16, msg_20
    mov x0, x16
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_12
.Lskip_branch_15:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x17, msg_21
    mov x0, x17
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_12:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_22
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosJoin ===
    sub sp, sp, #8  // Reservar espacio para puntosJoin
    mov x9, #0
    str x9, [sp]
// puntosJoin en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_23
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARACIÓN SLICE palabras: []string ===
// Elemento 0 (literal): Hola
    adr x0, msg_24
    str x0, [x28, #80]
// Elemento 1 (literal): mundo
    adr x0, msg_25
    str x0, [x28, #88]
// Elemento 2 (literal): desde
    adr x0, msg_26
    str x0, [x28, #96]
// Elemento 3 (literal): Go
    adr x0, msg_27
    str x0, [x28, #104]
// === DECLARAR VARIABLE MUT: frase ===
// Verificando slice 'palabras': true
// Patrón ID 'palabras' encontrado como slice
// === FUNCIÓN join(palabras, " ") ===
// Uniendo 4 elementos con separador
// === INICIO JOIN SIMPLIFICADO ===
    adr x11, msg_28
// Separador: ' '
    adr x12, join_buffer_17
// Usando buffer único: join_buffer_17
    mov x13, x12
    add x16, x28, #80
    ldr x14, [x16, #0]
copy1_loop_17:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy1_end_17
    strb w15, [x13], #1
    b copy1_loop_17
copy1_end_17:
    mov x9, #1
join_loop_palabras_17:
    cmp x9, #4
    bge join_end_palabras_17
    mov x14, x11
copy_sep_loop_17_x9:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy_sep_end_17_x9
    strb w15, [x13], #1
    b copy_sep_loop_17_x9
copy_sep_end_17_x9:
    add x16, x28, #80
    ldr x14, [x16, x9, lsl #3]
copy_elem_loop_17_x9:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy_elem_end_17_x9
    strb w15, [x13], #1
    b copy_elem_loop_17_x9
copy_elem_end_17_x9:
    add x9, x9, #1
    b join_loop_palabras_17
join_end_palabras_17:
    mov w15, #0
    strb w15, [x13]
// === FIN JOIN SIMPLIFICADO ===
    sub sp, sp, #8  // Reservar espacio para frase
    mov x9, x12
    str x9, [sp]
// frase en [sp+0] 

// === DECLARAR VARIABLE MUT: fraseConComas ===
// Verificando slice 'palabras': true
// Patrón ID 'palabras' encontrado como slice
// === FUNCIÓN join(palabras, ", ") ===
// Uniendo 4 elementos con separador
// === INICIO JOIN SIMPLIFICADO ===
    adr x11, msg_29
// Separador: ', '
    adr x12, join_buffer_18
// Usando buffer único: join_buffer_18
    mov x13, x12
    add x16, x28, #80
    ldr x14, [x16, #0]
copy1_loop_18:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy1_end_18
    strb w15, [x13], #1
    b copy1_loop_18
copy1_end_18:
    mov x9, #1
join_loop_palabras_18:
    cmp x9, #4
    bge join_end_palabras_18
    mov x14, x11
copy_sep_loop_18_x9:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy_sep_end_18_x9
    strb w15, [x13], #1
    b copy_sep_loop_18_x9
copy_sep_end_18_x9:
    add x16, x28, #80
    ldr x14, [x16, x9, lsl #3]
copy_elem_loop_18_x9:
    ldrb w15, [x14], #1
    cmp w15, #0
    beq copy_elem_end_18_x9
    strb w15, [x13], #1
    b copy_elem_loop_18_x9
copy_elem_end_18_x9:
    add x9, x9, #1
    b join_loop_palabras_18
join_end_palabras_18:
    mov w15, #0
    strb w15, [x13]
// === FIN JOIN SIMPLIFICADO ===
    sub sp, sp, #8  // Reservar espacio para fraseConComas
    mov x9, x12
    str x9, [sp]
// fraseConComas en [sp+0] 

// Verificando slice 'frase': false
// === CARGAR VARIABLE: frase (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x10, msg_30
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x9
    bl print_string
    bl print_newline
// Verificando slice 'fraseConComas': false
// === CARGAR VARIABLE: fraseConComas (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_31
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x9
    bl print_string
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA && ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'frase': false
// === CARGAR VARIABLE: frase (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    adr x12, msg_32
    mov x0, x11
    mov x1, x12
    bl strcmp
    mov x10, x0
    cmp x10, #0
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'fraseConComas': false
// === CARGAR VARIABLE: fraseConComas (offset: 0) ===
    ldr x13, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x15, x13
    adr x16, msg_33
    mov x0, x15
    mov x1, x16
    bl strcmp
    mov x14, x0
    cmp x14, #0
    cset x14, eq
// === FIN EXPRESIÓN BINARIA === 

// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_19
    mov x9, x14
    and x17, x18, x9
    b .Land_end_20
.Land_false_19:
    mov x17, #0
.Land_end_20:
// === FIN EXPRESIÓN BINARIA === 

    cmp x17, #0
    beq .Lskip_branch_21
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosJoin': false
// === CARGAR VARIABLE: puntosJoin (offset: 16) ===
    ldr x10, [sp, #16]
    mov x11, x10
    mov x12, #1
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosJoin = int (offset: 16) ===
    str x13, [sp, #16]
// === IMPRIMIR STRING  ===
    adr x14, msg_34
    mov x0, x14
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_18
.Lskip_branch_21:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x15, msg_35
    mov x0, x15
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_18:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_36
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosLen ===
    sub sp, sp, #8  // Reservar espacio para puntosLen
    mov x9, #0
    str x9, [sp]
// puntosLen en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_37
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: longitud1 ===
// Verificando slice 'numeros': true
// Patrón ID 'numeros' encontrado como slice
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 5 elementos
    mov x9, #5
    sub sp, sp, #8  // Reservar espacio para longitud1
    str x9, [sp]
// longitud1 en [sp+0] 

// === DECLARAR VARIABLE MUT: longitud2 ===
// Verificando slice 'palabras': true
// Patrón ID 'palabras' encontrado como slice
// === FUNCIÓN len(palabras) ===
// Slice palabras tiene 4 elementos
    mov x9, #4
    sub sp, sp, #8  // Reservar espacio para longitud2
    str x9, [sp]
// longitud2 en [sp+0] 

// Verificando slice 'longitud1': false
// === CARGAR VARIABLE: longitud1 (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x10, msg_38
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// Verificando slice 'longitud2': false
// === CARGAR VARIABLE: longitud2 (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_39
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA && ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'longitud1': false
// === CARGAR VARIABLE: longitud1 (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #5
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'longitud2': false
// === CARGAR VARIABLE: longitud2 (offset: 0) ===
    ldr x13, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x15, x13
    mov x16, #4
    cmp x15, x16
    cset x14, eq
// === FIN EXPRESIÓN BINARIA === 

// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_23
    mov x9, x14
    and x17, x18, x9
    b .Land_end_24
.Land_false_23:
    mov x17, #0
.Land_end_24:
// === FIN EXPRESIÓN BINARIA === 

    cmp x17, #0
    beq .Lskip_branch_25
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosLen': false
// === CARGAR VARIABLE: puntosLen (offset: 16) ===
    ldr x10, [sp, #16]
    mov x11, x10
    mov x12, #3
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosLen = int (offset: 16) ===
    str x13, [sp, #16]
// === IMPRIMIR STRING  ===
    adr x14, msg_40
    mov x0, x14
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_22
.Lskip_branch_25:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x15, msg_41
    mov x0, x15
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_22:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_42
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosAppend ===
    sub sp, sp, #8  // Reservar espacio para puntosAppend
    mov x9, #0
    str x9, [sp]
// puntosAppend en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_43
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARACIÓN SLICE numeros3: []int ===
// Elemento 0 (literal): 1
    mov x0, #1
    str x0, [x28, #112]
// Elemento 1 (literal): 2
    mov x0, #2
    str x0, [x28, #120]
// Elemento 2 (literal): 3
    mov x0, #3
    str x0, [x28, #128]
// Verificando slice 'numeros3': true
// Patrón ID 'numeros3' encontrado como slice
// === FUNCIÓN append(numeros3, 4) ===
// Expandiendo slice numeros3 de 3 a 4 elementos
// Copiando 3 elementos existentes
    ldr x18, [x28, #112]    // cargar elemento 0 original
    str x18, [x28, #136]    // guardar elemento 0 en nueva posición
    ldr x18, [x28, #120]    // cargar elemento 1 original
    str x18, [x28, #144]    // guardar elemento 1 en nueva posición
    ldr x18, [x28, #128]    // cargar elemento 2 original
    str x18, [x28, #152]    // guardar elemento 2 en nueva posición
// Agregando nuevo elemento en offset 160
    mov x0, #4
    str x0, [x28, #160]
// append completado: numeros3 ahora tiene 4 elementos
// Verificando slice 'numeros3': true
// Patrón ID 'numeros3' encontrado como slice
// === IMPRIMIR STRING  ===
    adr x9, msg_44
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [x28, #136]    // cargar int[0] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #144]    // cargar int[1] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #152]    // cargar int[2] desde base fija
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [x28, #160]    // cargar int[3] desde base fija
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA && ===
// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'numeros3': true
// Patrón ID 'numeros3' encontrado como slice
// === FUNCIÓN len(numeros3) ===
// Slice numeros3 tiene 4 elementos
    mov x9, #4
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #4
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

// ===  EXPRESIÓN BINARIA == ===
// Verificando slice 'numeros3': true
// === ACCESO POR ÍNDICE numeros3[3] ===
    ldr x13, [x28, #160]    // cargar numeros3[3] desde base fija
// === OPERACIÓN RELACIONAL == ===
    mov x15, x13
    mov x16, #4
    cmp x15, x16
    cset x14, eq
// === FIN EXPRESIÓN BINARIA === 

// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_27
    mov x9, x14
    and x17, x18, x9
    b .Land_end_28
.Land_false_27:
    mov x17, #0
.Land_end_28:
// === FIN EXPRESIÓN BINARIA === 

    cmp x17, #0
    beq .Lskip_branch_29
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosAppend': false
// === CARGAR VARIABLE: puntosAppend (offset: 0) ===
    ldr x10, [sp, #0]
    mov x11, x10
    mov x12, #3
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntosAppend = int (offset: 0) ===
    str x13, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x14, msg_45
    mov x0, x14
    bl print_string
    bl print_newline
// Pop scope
    b .Lif_final_26
.Lskip_branch_29:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x15, msg_46
    mov x0, x15
    bl print_string
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_26:
// === FIN ESTRUCTURA IF-ELSE ===

// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// Verificando slice 'puntosCreacion': false
// === CARGAR VARIABLE: puntosCreacion (offset: 96) ===
    ldr x9, [sp, #96]
// Verificando slice 'puntosAcceso': false
// === CARGAR VARIABLE: puntosAcceso (offset: 88) ===
    ldr x10, [sp, #88]
    mov x11, x9
    mov x12, x10
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA === 

// Verificando slice 'puntosIndex': false
// === CARGAR VARIABLE: puntosIndex (offset: 72) ===
    ldr x14, [sp, #72]
    mov x15, x13
    mov x16, x14
    add x17, x15, x16
// === FIN EXPRESIÓN BINARIA === 

// Verificando slice 'puntosJoin': false
// === CARGAR VARIABLE: puntosJoin (offset: 48) ===
    ldr x18, [sp, #48]
    mov x9, x17
    mov x10, x18
    add x11, x9, x10
// === FIN EXPRESIÓN BINARIA === 

// Verificando slice 'puntosLen': false
// === CARGAR VARIABLE: puntosLen (offset: 24) ===
    ldr x12, [sp, #24]
    mov x13, x11
    mov x14, x12
    add x15, x13, x14
// === FIN EXPRESIÓN BINARIA === 

// Verificando slice 'puntosAppend': false
// === CARGAR VARIABLE: puntosAppend (offset: 0) ===
    ldr x16, [sp, #0]
    mov x17, x15
    mov x18, x16
    add x9, x17, x18
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntos = int (offset: 104) ===
    str x9, [sp, #104]
// === IMPRIMIR STRING  ===
    adr x9, msg_47
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_48
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_49
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_50
    mov x0, x9
    bl print_string
    bl print_newline
// Verificando slice 'puntosCreacion': false
// === CARGAR VARIABLE: puntosCreacion (offset: 96) ===
    ldr x9, [sp, #96]
// === IMPRIMIR STRING  ===
    adr x10, msg_51
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_52
    mov x0, x11
    bl print_string
    bl print_newline
// Verificando slice 'puntosAcceso': false
// === CARGAR VARIABLE: puntosAcceso (offset: 88) ===
    ldr x9, [sp, #88]
// === IMPRIMIR STRING  ===
    adr x10, msg_53
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_54
    mov x0, x11
    bl print_string
    bl print_newline
// Verificando slice 'puntosIndex': false
// === CARGAR VARIABLE: puntosIndex (offset: 72) ===
    ldr x9, [sp, #72]
// === IMPRIMIR STRING  ===
    adr x10, msg_55
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_56
    mov x0, x11
    bl print_string
    bl print_newline
// Verificando slice 'puntosJoin': false
// === CARGAR VARIABLE: puntosJoin (offset: 48) ===
    ldr x9, [sp, #48]
// === IMPRIMIR STRING  ===
    adr x10, msg_57
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_58
    mov x0, x11
    bl print_string
    bl print_newline
// Verificando slice 'puntosLen': false
// === CARGAR VARIABLE: puntosLen (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR STRING  ===
    adr x10, msg_59
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_60
    mov x0, x11
    bl print_string
    bl print_newline
// Verificando slice 'puntosAppend': false
// === CARGAR VARIABLE: puntosAppend (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_61
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_62
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_63
    mov x0, x9
    bl print_string
    bl print_newline
// Verificando slice 'puntos': false
// === CARGAR VARIABLE: puntos (offset: 104) ===
    ldr x9, [sp, #104]
// === IMPRIMIR STRING  ===
    adr x10, msg_64
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_65
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_66
    mov x0, x9
    bl print_string
    bl print_newline
// === LIMPIEZA FINAL ===
    mov x8, #93     // sys_exit
    mov x0, #0      // exit status 0
    svc #0          // llamada al sistema
        mov   sp, x29                 // Restaurar stack usando frame pointer
        ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
        ret                           // Retornar a _start
// === FIN FUNCIÓN MAIN ===
    
// === FUNCIONES DEFINIDAS POR EL USUARIO ===

// --------------------------------------------------------
//            FUNCIONES AUXILIARES ARM64
// --------------------------------------------------------

print_int:
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp
    stp   x1, x2, [sp, #-16]!     // Guardar registros que vamos a usar
    stp   x3, x4, [sp, #-16]!
    stp   x5, x6, [sp, #-16]!

    //  Usar comparación con registro zero, no inmediato
    cmp   x0, xzr                 // ¿Es negativo? (usar xzr en lugar de #0)
    bge   .Lpi_pos
    
    //  Manejar números negativos - GUARDAR x0 original
    stp   x7, x8, [sp, #-16]!     // Guardar más registros
    mov   x7, x0                  // GUARDAR valor original en x7
    
    // Imprimir signo menos
    mov   x8, #64                 // Syscall write
    ldr   x1, =msg_menos          // Imprimir "-"
    mov   x2, #1
    mov   x0, #1
    svc   0
    
    //  RESTAURAR y negar el valor original
    mov   x0, x7                  // Restaurar valor original
    neg   x0, x0                  // Hacer positivo
    ldp   x7, x8, [sp], #16       // Restaurar registros

.Lpi_pos:
    ldr   x2, =buffer_int         // Buffer para dígitos
    add   x2, x2, #32             // Empezar desde el final
    mov   x3, #0                  // Contador de dígitos
    mov   x6, #10                 // Divisor

.Lpi_loop:
    udiv  x4, x0, x6              // x4 = x0 / 10
    msub  x5, x4, x6, x0          // x5 = x0 % 10 (resto)
    add   x5, x5, #48             // Convertir a ASCII
    sub   x2, x2, #1              // Retroceder en buffer
    strb  w5, [x2]                // Guardar dígito
    mov   x0, x4                  // Siguiente iteración
    add   x3, x3, #1              // Incrementar contador
    cmp   x0, #0
    bne   .Lpi_loop

    mov   x0, #1                  // stdout
    mov   x1, x2                  // Buffer con dígitos
    mov   x2, x3                  // Cantidad de dígitos
    mov   x8, #64                 // Syscall write
    svc   0

    ldp   x5, x6, [sp], #16       // Restaurar registros
    ldp   x3, x4, [sp], #16
    ldp   x1, x2, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

print_string:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros adicionales
    
    mov   x19, x0                 // Guardar dirección del string
    mov   x20, #0                 // Contador manual de longitud

    // Verificar que el string no sea NULL
    cmp   x19, #0
    beq   .Lps_null_string

// Contar manualmente la longitud (más seguro que strlen separado)
.Lps_count_loop:
    ldrb  w1, [x19, x20]         // Cargar byte
    cmp   w1, #0                 // ¿Es '\0'?
    beq   .Lps_count_done
    add   x20, x20, #1           // Incrementar contador
    cmp   x20, #500              // Límite de seguridad
    bge   .Lps_count_done        // Evitar loops infinitos
    b     .Lps_count_loop

.Lps_count_done:
    // Verificar que tengamos algo que imprimir
    cmp   x20, #0
    beq   .Lps_empty_string

    // Imprimir el string
    mov   x0, #1                 // stdout
    mov   x1, x19                // dirección string
    mov   x2, x20                // longitud calculada
    mov   x8, #64                // syscall write
    svc   0
    b     .Lps_end

.Lps_null_string:
    // String NULL - no hacer nada
    b     .Lps_end

.Lps_empty_string:
    // String vacío - no hacer nada
    b     .Lps_end

.Lps_end:
    ldp   x19, x20, [sp], #16    // Restaurar registros
    ldp   x29, x30, [sp], #16
    ret

print_char:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp

    strb  w0, [sp, #-1]!          // Poner carácter en stack
    mov   x0, #1                  // stdout
    mov   x1, sp                  // Dirección del carácter
    mov   x2, #1                  // 1 byte
    mov   x8, #64                 // Syscall write
    svc   0
    add   sp, sp, #1              // Limpiar stack

    ldp   x29, x30, [sp], #16
    ret

print_newline:
    mov   x0, #1                  // stdout
    ldr   x1, =msg_nl             // "\n"
    mov   x2, #1                  // 1 byte
    mov   x8, #64                 // Syscall write
    svc   0
    ret

strlen:
    mov   x1, #0                  // contador = 0
.Lstrlen_loop:
    ldrb  w2, [x0, x1]           // Cargar byte
    cmp   w2, #0                 // ¿Es '\0'?
    beq   .Lstrlen_end
    add   x1, x1, #1             // Incrementar contador
    b     .Lstrlen_loop
.Lstrlen_end:
    mov   x0, x1                 // Retornar longitud
    ret

strcmp:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x2, x3, [sp, #-16]!

strcmp_loop:
    ldrb  w2, [x0], #1            // Cargar char de string1
    ldrb  w3, [x1], #1            // Cargar char de string2
    
    cmp   w2, w3                  // Comparar caracteres
    bne   strcmp_different
    
    cmp   w2, #0                  // ¿Final de string?
    beq   strcmp_equal
    
    b     strcmp_loop
    
strcmp_equal:
    mov   x0, #0                  // Iguales = 0
    b     strcmp_end
    
strcmp_different:
    mov   x0, #1                  // Diferentes = 1
    
strcmp_end:
    ldp   x2, x3, [sp], #16       // Restaurar registros
    ldp   x29, x30, [sp], #16
    ret

