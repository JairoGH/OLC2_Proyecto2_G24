.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "a ="
    msg_2: .asciz "OK a = 10"
    msg_3: .asciz "b ="
    msg_4: .asciz "OK b = 20"
    msg_5: .asciz "c ="
    msg_6: .asciz "d ="
    msg_7: .asciz "OK c = 30"
    msg_8: .asciz "OK true"
    msg_9: .asciz "OK 1 == 1"
    msg_10: .asciz "OK 2 > 1"
    msg_11: .asciz "OK suma1 == 10"
    msg_12: .asciz "OK i == 5"
    msg_13: .asciz "OK suma2 == 10"
    msg_14: .asciz "Lunes"
    msg_15: .asciz "Martes"
    msg_16: .asciz "Miércoles"
    msg_17: .asciz "Jueves"
    msg_18: .asciz "Viernes"
    msg_19: .asciz "Sábado"
    msg_20: .asciz "Domingo"
    msg_21: .asciz "Día inválido"
    msg_22: .asciz "OK suma3 == 10"
    msg_23: .asciz "OK suma_pares == 20"
    msg_24: .asciz "|══════════════════════════════════════|══════|══════|"
    msg_25: .asciz "| Característica                       |Puntos| Total|"
    msg_26: .asciz "|══════════════════════════════════════|══════|══════|"
    msg_27: .asciz "| Declaración y Asignación Variables   |"
    msg_28: .asciz "   |  3   |"
    msg_29: .asciz "| Estructuras Condicionales IF         |"
    msg_30: .asciz "   |  3   |"
    msg_31: .asciz "| Bucles FOR tipo WHILE                |"
    msg_32: .asciz "   |  2   |"
    msg_33: .asciz "| Bucles FOR Clásico                   |"
    msg_34: .asciz "   |  3   |"
    msg_35: .asciz "| Estructura SWITCH-CASE               |"
    msg_36: .asciz "   |  1   |"
    msg_37: .asciz "| Sentencia BREAK                      |"
    msg_38: .asciz "   |  3   |"
    msg_39: .asciz "| Sentencia CONTINUE                   |"
    msg_40: .asciz "   |  3   |"
    msg_41: .asciz "|══════════════════════════════════════|══════|══════|"
    msg_42: .asciz "| TOTAL PUNTOS OBTENIDOS               |"
    msg_43: .asciz "  | 18   |"
    msg_44: .asciz "|══════════════════════════════════════|══════|══════|"
buffer_int: .skip 32
buffer_float: .skip 64
buffer_string: .skip 512
temp_buffer: .skip 256
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0
str_empty: .asciz ""

.section .text
.global _start

_start:
// === DECLARAR VARIABLE MUT: puntos ===
    sub sp, sp, #8  // Reservar espacio para puntos
    mov x9, #0
    str x9, [sp]
// puntos en [sp+0] 

// === DECLARAR VARIABLE MUT: puntos_entornos ===
    sub sp, sp, #8  // Reservar espacio para puntos_entornos
    mov x9, #0
    str x9, [sp]
// puntos_entornos en [sp+0] 

// === DECLARAR VARIABLE MUT: a ===
    sub sp, sp, #8  // Reservar espacio para a
    mov x9, #10
    str x9, [sp]
// a en [sp+0] 

// === CARGAR VARIABLE: a (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_1
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: a (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_1
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_entornos += ===
// === CARGAR VARIABLE: puntos_entornos (offset: 8) ===
    ldr x13, [sp, #8]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_entornos = int (offset: 8) ===
    str x16, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_entornos += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_2
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_0
.Lskip_branch_1:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_0:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: b ===
    sub sp, sp, #8  // Reservar espacio para b
    mov x9, #10
    str x9, [sp]
// b en [sp+0] 

// === ASIGNAR VARIABLE: b = int (offset: 0) ===
    mov x9, #20
    str x9, [sp, #0]
// === CARGAR VARIABLE: b (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_3
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: b (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #20
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_3
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_entornos += ===
// === CARGAR VARIABLE: puntos_entornos (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_entornos = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_entornos += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_4
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_2
.Lskip_branch_3:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_2:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: c ===
    sub sp, sp, #8  // Reservar espacio para c
    mov x9, #10
    str x9, [sp]
// c en [sp+0] 

// === DECLARAR VARIABLE MUT: d ===
    sub sp, sp, #8  // Reservar espacio para d
    mov x9, #10
    str x9, [sp]
// d en [sp+0] 

// === ASIGNAR VARIABLE: c = int (offset: 8) ===
    mov x9, #30
    str x9, [sp, #8]
// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING  ===
    adr x10, msg_5
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: d (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_6
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #30
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_5
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_entornos += ===
// === CARGAR VARIABLE: puntos_entornos (offset: 32) ===
    ldr x13, [sp, #32]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_entornos = int (offset: 32) ===
    str x16, [sp, #32]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_entornos += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_7
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_4
.Lskip_branch_5:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_4:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: puntos_if ===
    sub sp, sp, #8  // Reservar espacio para puntos_if
    mov x9, #0
    str x9, [sp]
// puntos_if en [sp+0] 

// === INICIO ESTRUCTURA IF-ELSE ===
// === LITERAL BOOL: true ===
// Condición literal: true
// === EJECUTANDO BLOQUE (literal true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_if += ===
// === CARGAR VARIABLE: puntos_if (offset: 0) ===
    ldr x9, [sp, #0]
    mov x10, x9
    mov x11, #1
    add x12, x10, x11
// === ASIGNAR VARIABLE: puntos_if = int (offset: 0) ===
    str x12, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_if += === 

// === IMPRIMIR STRING  ===
    adr x13, msg_8
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_6
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_6:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === OPERACIÓN RELACIONAL == ===
    mov x10, #1
    mov x11, #1
    cmp x10, x11
    cset x9, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x9, #0
    beq .Lskip_branch_9
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_if += ===
// === CARGAR VARIABLE: puntos_if (offset: 0) ===
    ldr x12, [sp, #0]
    mov x13, x12
    mov x14, #1
    add x15, x13, x14
// === ASIGNAR VARIABLE: puntos_if = int (offset: 0) ===
    str x15, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_if += === 

// === IMPRIMIR STRING  ===
    adr x16, msg_9
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_8
.Lskip_branch_9:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_8:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA > ===
// === OPERACIÓN RELACIONAL > ===
    mov x10, #2
    mov x11, #1
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA === 

    cmp x9, #0
    beq .Lskip_branch_11
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_if += ===
// === CARGAR VARIABLE: puntos_if (offset: 0) ===
    ldr x12, [sp, #0]
    mov x13, x12
    mov x14, #1
    add x15, x13, x14
// === ASIGNAR VARIABLE: puntos_if = int (offset: 0) ===
    str x15, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_if += === 

// === IMPRIMIR STRING  ===
    adr x16, msg_10
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_10
.Lskip_branch_11:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_10:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: puntos_while ===
    sub sp, sp, #8  // Reservar espacio para puntos_while
    mov x9, #0
    str x9, [sp]
// puntos_while en [sp+0] 

// === DECLARAR VARIABLE MUT: i ===
    sub sp, sp, #8  // Reservar espacio para i
    mov x9, #0
    str x9, [sp]
// i en [sp+0] 

// === DECLARAR VARIABLE MUT: suma1 ===
    sub sp, sp, #8  // Reservar espacio para suma1
    mov x9, #0
    str x9, [sp]
// suma1 en [sp+0] 

// === INICIO WHILE ARM64 ===
// Push control: while
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL < ===
    mov x11, x9
    mov x12, #5
    cmp x11, x12
    cset x10, lt
// === FIN EXPRESIÓN BINARIA === 

.Lwhile_start_13:
// Evaluando condición del while...
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x13, [sp, #8]
// === OPERACIÓN RELACIONAL < ===
    mov x15, x13
    mov x16, #5
    cmp x15, x16
    cset x14, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x14, #0
    beq .Lwhile_end_13
// Ejecutando sentencia 1 del while
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x17, [sp, #8]
// === IMPRIMIR INT  ===
    mov x0, x17
    bl print_int
    bl print_newline
    bl print_newline
// Ejecutando sentencia 2 del while
// === ASIGNACIÓN ARITMÉTICA: suma1 += ===
// === CARGAR VARIABLE: suma1 (offset: 0) ===
    ldr x18, [sp, #0]
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x19, [sp, #8]
    mov x20, x18
    mov x21, x19
    add x22, x20, x21
// === ASIGNAR VARIABLE: suma1 = int (offset: 0) ===
    str x22, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: suma1 += === 

// Ejecutando sentencia 3 del while
// === ASIGNACIÓN ARITMÉTICA: i += ===
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x23, [sp, #8]
    mov x24, x23
    mov x25, #1
    add x26, x24, x25
// === ASIGNAR VARIABLE: i = int (offset: 8) ===
    str x26, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: i += === 

.Lwhile_continue_13:
    b .Lwhile_start_13
.Lwhile_end_13:
// >>> FIN DEL BUCLE WHILE <<<
// === FIN WHILE ARM64 ===
// Pop control: while
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: suma1 (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_14
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_while += ===
// === CARGAR VARIABLE: puntos_while (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_while = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_while += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_11
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_13
.Lskip_branch_14:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_13:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: i (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #5
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_16
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_while += ===
// === CARGAR VARIABLE: puntos_while (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_while = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_while += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_12
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_15
.Lskip_branch_16:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_15:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: j ===
    sub sp, sp, #8  // Reservar espacio para j
    mov x9, #3
    str x9, [sp]
// j en [sp+0] 

// === INICIO WHILE ARM64 ===
// Push control: while
// ===  EXPRESIÓN BINARIA > ===
// === CARGAR VARIABLE: j (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL > ===
    mov x11, x9
    mov x12, #0
    cmp x11, x12
    cset x10, gt
// === FIN EXPRESIÓN BINARIA === 

.Lwhile_start_18:
// Evaluando condición del while...
// ===  EXPRESIÓN BINARIA > ===
// === CARGAR VARIABLE: j (offset: 0) ===
    ldr x13, [sp, #0]
// === OPERACIÓN RELACIONAL > ===
    mov x15, x13
    mov x16, #0
    cmp x15, x16
    cset x14, gt
// === FIN EXPRESIÓN BINARIA === 

    cmp x14, #0
    beq .Lwhile_end_18
// Ejecutando sentencia 1 del while
// === CARGAR VARIABLE: j (offset: 0) ===
    ldr x17, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x17
    bl print_int
    bl print_newline
    bl print_newline
// Ejecutando sentencia 2 del while
// === ASIGNACIÓN ARITMÉTICA: j -= ===
// === CARGAR VARIABLE: j (offset: 0) ===
    ldr x18, [sp, #0]
    mov x19, x18
    mov x20, #1
    sub x21, x19, x20
// === ASIGNAR VARIABLE: j = int (offset: 0) ===
    str x21, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: j -= === 

.Lwhile_continue_18:
    b .Lwhile_start_18
.Lwhile_end_18:
// >>> FIN DEL BUCLE WHILE <<<
// === FIN WHILE ARM64 ===
// Pop control: while
// === DECLARAR VARIABLE MUT: k ===
    sub sp, sp, #8  // Reservar espacio para k
    mov x9, #0
    str x9, [sp]
// k en [sp+0] 

// === INICIO WHILE ARM64 ===
// Push control: while
// ===  EXPRESIÓN BINARIA <= ===
// === CARGAR VARIABLE: k (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL <= ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, le
// === FIN EXPRESIÓN BINARIA === 

.Lwhile_start_19:
// Evaluando condición del while...
// ===  EXPRESIÓN BINARIA <= ===
// === CARGAR VARIABLE: k (offset: 0) ===
    ldr x13, [sp, #0]
// === OPERACIÓN RELACIONAL <= ===
    mov x15, x13
    mov x16, #10
    cmp x15, x16
    cset x14, le
// === FIN EXPRESIÓN BINARIA === 

    cmp x14, #0
    beq .Lwhile_end_19
// Ejecutando sentencia 1 del while
// === CARGAR VARIABLE: k (offset: 0) ===
    ldr x17, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x17
    bl print_int
    bl print_newline
    bl print_newline
// Ejecutando sentencia 2 del while
// === ASIGNACIÓN ARITMÉTICA: k += ===
// === CARGAR VARIABLE: k (offset: 0) ===
    ldr x18, [sp, #0]
    mov x19, x18
    mov x20, #2
    add x21, x19, x20
// === ASIGNAR VARIABLE: k = int (offset: 0) ===
    str x21, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: k += === 

.Lwhile_continue_19:
    b .Lwhile_start_19
.Lwhile_end_19:
// >>> FIN DEL BUCLE WHILE <<<
// === FIN WHILE ARM64 ===
// Pop control: while
// === DECLARAR VARIABLE MUT: puntos_for ===
    sub sp, sp, #8  // Reservar espacio para puntos_for
    mov x9, #0
    str x9, [sp]
// puntos_for en [sp+0] 

// === DECLARAR VARIABLE MUT: suma2 ===
    sub sp, sp, #8  // Reservar espacio para suma2
    mov x9, #0
    str x9, [sp]
// suma2 en [sp+0] 

// === DECLARAR VARIABLE MUT: x ===
    sub sp, sp, #8  // Reservar espacio para x
    mov x9, #0
    str x9, [sp]
// x en [sp+0] 

// === INICIO FOR CLÁSICO ARM64 ===
// Push control: for_clasico
// Inicialización del for clásico
// === ASIGNAR VARIABLE: x = int (offset: 0) ===
    mov x9, #0
    str x9, [sp, #0]
.Lforc_start_20:
// Evaluar condición del for clásico
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: x (offset: 0) ===
    ldr x10, [sp, #0]
// === OPERACIÓN RELACIONAL < ===
    mov x12, x10
    mov x13, #5
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x11, #0
    beq .Lforc_end_20
// Cuerpo del for clásico
// Sentencia 1 del for clásico
// === CARGAR VARIABLE: x (offset: 0) ===
    ldr x14, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x14
    bl print_int
    bl print_newline
    bl print_newline
// Sentencia 2 del for clásico
// === ASIGNACIÓN ARITMÉTICA: suma2 += ===
// === CARGAR VARIABLE: suma2 (offset: 8) ===
    ldr x15, [sp, #8]
// === CARGAR VARIABLE: x (offset: 0) ===
    ldr x16, [sp, #0]
    mov x17, x15
    mov x18, x16
    add x19, x17, x18
// === ASIGNAR VARIABLE: suma2 = int (offset: 8) ===
    str x19, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: suma2 += === 

.Lforc_continue_20:
// Post-iteración del for clásico
// === ASIGNACIÓN ARITMÉTICA: x += ===
// === CARGAR VARIABLE: x (offset: 0) ===
    ldr x20, [sp, #0]
    mov x21, x20
    mov x22, #1
    add x23, x21, x22
// === ASIGNAR VARIABLE: x = int (offset: 0) ===
    str x23, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: x += === 

    b .Lforc_start_20
.Lforc_end_20:
// === FIN FOR CLÁSICO ARM64 INTERNO ===
// === FIN FOR CLÁSICO ARM64 ===
// Pop control: for_clasico
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: suma2 (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_21
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_for += ===
// === CARGAR VARIABLE: puntos_for (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_for = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_for += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_13
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_20
.Lskip_branch_21:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_20:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: y ===
    sub sp, sp, #8  // Reservar espacio para y
    mov x9, #0
    str x9, [sp]
// y en [sp+0] 

// === INICIO FOR CLÁSICO ARM64 ===
// Push control: for_clasico
// Inicialización del for clásico
// === ASIGNAR VARIABLE: y = int (offset: 0) ===
    mov x9, #0
    str x9, [sp, #0]
.Lforc_start_23:
// Evaluar condición del for clásico
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: y (offset: 0) ===
    ldr x10, [sp, #0]
// === OPERACIÓN RELACIONAL < ===
    mov x12, x10
    mov x13, #3
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x11, #0
    beq .Lforc_end_23
// Cuerpo del for clásico
// Sentencia 1 del for clásico
// === CARGAR VARIABLE: y (offset: 0) ===
    ldr x14, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x14
    bl print_int
    bl print_newline
    bl print_newline
.Lforc_continue_23:
// Post-iteración del for clásico
// === ASIGNACIÓN ARITMÉTICA: y += ===
// === CARGAR VARIABLE: y (offset: 0) ===
    ldr x15, [sp, #0]
    mov x16, x15
    mov x17, #1
    add x18, x16, x17
// === ASIGNAR VARIABLE: y = int (offset: 0) ===
    str x18, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: y += === 

    b .Lforc_start_23
.Lforc_end_23:
// === FIN FOR CLÁSICO ARM64 INTERNO ===
// === FIN FOR CLÁSICO ARM64 ===
// Pop control: for_clasico
// === DECLARAR VARIABLE MUT: z ===
    sub sp, sp, #8  // Reservar espacio para z
    mov x9, #0
    str x9, [sp]
// z en [sp+0] 

// === INICIO FOR CLÁSICO ARM64 ===
// Push control: for_clasico
// Inicialización del for clásico
// === ASIGNAR VARIABLE: z = int (offset: 0) ===
    mov x9, #0
    str x9, [sp, #0]
.Lforc_start_24:
// Evaluar condición del for clásico
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x10, [sp, #0]
// === OPERACIÓN RELACIONAL < ===
    mov x12, x10
    mov x13, #2
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x11, #0
    beq .Lforc_end_24
// Cuerpo del for clásico
// Sentencia 1 del for clásico
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x14, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x14
    bl print_int
    bl print_newline
    bl print_newline
.Lforc_continue_24:
// Post-iteración del for clásico
// === ASIGNACIÓN ARITMÉTICA: z += ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x15, [sp, #0]
    mov x16, x15
    mov x17, #1
    add x18, x16, x17
// === ASIGNAR VARIABLE: z = int (offset: 0) ===
    str x18, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: z += === 

    b .Lforc_start_24
.Lforc_end_24:
// === FIN FOR CLÁSICO ARM64 INTERNO ===
// === FIN FOR CLÁSICO ARM64 ===
// Pop control: for_clasico
// === ASIGNACIÓN ARITMÉTICA: puntos_for += ===
// === CARGAR VARIABLE: puntos_for (offset: 32) ===
    ldr x9, [sp, #32]
    mov x10, x9
    mov x11, #2
    add x12, x10, x11
// === ASIGNAR VARIABLE: puntos_for = int (offset: 32) ===
    str x12, [sp, #32]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_for += === 

// === DECLARAR VARIABLE MUT: puntos_case ===
    sub sp, sp, #8  // Reservar espacio para puntos_case
    mov x9, #0
    str x9, [sp]
// puntos_case en [sp+0] 

// === DECLARAR VARIABLE MUT: dia ===
    sub sp, sp, #8  // Reservar espacio para dia
    mov x9, #3
    str x9, [sp]
// dia en [sp+0] 

// === INICIO SWITCH STATEMENT ===
// === CARGAR VARIABLE: dia (offset: 0) ===
    ldr x9, [sp, #0]
// Switch sobre expresión tipo: int
// Push control: switch
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_7_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_7_1
// Comparación case 2
    mov x13, #3
    cmp x10, x13
    beq .Lcase_7_2
// Comparación case 3
    mov x14, #4
    cmp x10, x14
    beq .Lcase_7_3
// Comparación case 4
    mov x15, #5
    cmp x10, x15
    beq .Lcase_7_4
// Comparación case 5
    mov x16, #6
    cmp x10, x16
    beq .Lcase_7_5
// Comparación case 6
    mov x17, #7
    cmp x10, x17
    beq .Lcase_7_6
    b .Lswitch_default_7
.Lcase_7_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x18, msg_14
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntos_case += ===
// === CARGAR VARIABLE: puntos_case (offset: 8) ===
    ldr x19, [sp, #8]
    mov x20, x19
    mov x21, #1
    add x22, x20, x21
// === ASIGNAR VARIABLE: puntos_case = int (offset: 8) ===
    str x22, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_case += === 

    b .Lswitch_end_7
.Lcase_7_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x23, msg_15
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x24, msg_16
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntos_case += ===
// === CARGAR VARIABLE: puntos_case (offset: 8) ===
    ldr x25, [sp, #8]
    mov x26, x25
    mov x27, #1
    add x28, x26, x27
// === ASIGNAR VARIABLE: puntos_case = int (offset: 8) ===
    str x28, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_case += === 

    b .Lswitch_end_7
.Lcase_7_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x29, msg_17
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x30, msg_18
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_19
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x10, msg_20
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lswitch_default_7:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x11, msg_21
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_7:
// === FIN SWITCH STATEMENT ===
// Pop control: switch
// === DECLARAR VARIABLE MUT: puntos_break ===
    sub sp, sp, #8  // Reservar espacio para puntos_break
    mov x9, #0
    str x9, [sp]
// puntos_break en [sp+0] 

// === DECLARAR VARIABLE MUT: suma3 ===
    sub sp, sp, #8  // Reservar espacio para suma3
    mov x9, #0
    str x9, [sp]
// suma3 en [sp+0] 

// === DECLARAR VARIABLE MUT: n ===
    sub sp, sp, #8  // Reservar espacio para n
    mov x9, #0
    str x9, [sp]
// n en [sp+0] 

// === INICIO FOR CLÁSICO ARM64 ===
// Push control: for_clasico
// Inicialización del for clásico
// === ASIGNAR VARIABLE: n = int (offset: 0) ===
    mov x9, #0
    str x9, [sp, #0]
.Lforc_start_25:
// Evaluar condición del for clásico
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: n (offset: 0) ===
    ldr x10, [sp, #0]
// === OPERACIÓN RELACIONAL < ===
    mov x12, x10
    mov x13, #10
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x11, #0
    beq .Lforc_end_25
// Cuerpo del for clásico
// Sentencia 1 del for clásico
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: n (offset: 0) ===
    ldr x14, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x16, x14
    mov x17, #5
    cmp x16, x17
    cset x15, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x15, #0
    beq .Lskip_branch_26
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === BREAK STATEMENT ===
// Break desde for_clasico
    b .Lforc_end_25
// Pop scope
    b .Lif_final_25
.Lskip_branch_26:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_25:
// === FIN ESTRUCTURA IF-ELSE ===

// Sentencia 2 del for clásico
// === CARGAR VARIABLE: n (offset: 0) ===
    ldr x18, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x18
    bl print_int
    bl print_newline
    bl print_newline
// Sentencia 3 del for clásico
// === ASIGNACIÓN ARITMÉTICA: suma3 += ===
// === CARGAR VARIABLE: suma3 (offset: 8) ===
    ldr x19, [sp, #8]
// === CARGAR VARIABLE: n (offset: 0) ===
    ldr x20, [sp, #0]
    mov x21, x19
    mov x22, x20
    add x23, x21, x22
// === ASIGNAR VARIABLE: suma3 = int (offset: 8) ===
    str x23, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: suma3 += === 

.Lforc_continue_25:
// Post-iteración del for clásico
// === ASIGNACIÓN ARITMÉTICA: n += ===
// === CARGAR VARIABLE: n (offset: 0) ===
    ldr x24, [sp, #0]
    mov x25, x24
    mov x26, #1
    add x27, x25, x26
// === ASIGNAR VARIABLE: n = int (offset: 0) ===
    str x27, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: n += === 

    b .Lforc_start_25
.Lforc_end_25:
// === FIN FOR CLÁSICO ARM64 INTERNO ===
// === FIN FOR CLÁSICO ARM64 ===
// Pop control: for_clasico
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: suma3 (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_28
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_break += ===
// === CARGAR VARIABLE: puntos_break (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #3
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_break = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_break += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_22
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_27
.Lskip_branch_28:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_27:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: puntos_continue ===
    sub sp, sp, #8  // Reservar espacio para puntos_continue
    mov x9, #0
    str x9, [sp]
// puntos_continue en [sp+0] 

// === DECLARAR VARIABLE MUT: suma_pares ===
    sub sp, sp, #8  // Reservar espacio para suma_pares
    mov x9, #0
    str x9, [sp]
// suma_pares en [sp+0] 

// === DECLARAR VARIABLE MUT: m ===
    sub sp, sp, #8  // Reservar espacio para m
    mov x9, #0
    str x9, [sp]
// m en [sp+0] 

// === INICIO FOR CLÁSICO ARM64 ===
// Push control: for_clasico
// Inicialización del for clásico
// === ASIGNAR VARIABLE: m = int (offset: 0) ===
    mov x9, #0
    str x9, [sp, #0]
.Lforc_start_30:
// Evaluar condición del for clásico
// ===  EXPRESIÓN BINARIA < ===
// === CARGAR VARIABLE: m (offset: 0) ===
    ldr x10, [sp, #0]
// === OPERACIÓN RELACIONAL < ===
    mov x12, x10
    mov x13, #10
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA === 

    cmp x11, #0
    beq .Lforc_end_30
// Cuerpo del for clásico
// Sentencia 1 del for clásico
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA != ===
// ===  EXPRESIÓN BINARIA % ===
// === CARGAR VARIABLE: m (offset: 0) ===
    ldr x14, [sp, #0]
    mov x15, x14
    mov x16, #2
    udiv x17, x15, x16
    msub x17, x17, x16, x15
// === FIN EXPRESIÓN BINARIA === 

// === OPERACIÓN RELACIONAL != ===
    mov x19, x17
    mov x20, #0
    cmp x19, x20
    cset x18, ne
// === FIN EXPRESIÓN BINARIA === 

    cmp x18, #0
    beq .Lskip_branch_31
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === CONTINUE STATEMENT ===
// Continue en for_clasico
    b .Lforc_continue_30
// Pop scope
    b .Lif_final_30
.Lskip_branch_31:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_30:
// === FIN ESTRUCTURA IF-ELSE ===

// Sentencia 2 del for clásico
// === CARGAR VARIABLE: m (offset: 0) ===
    ldr x21, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x21
    bl print_int
    bl print_newline
    bl print_newline
// Sentencia 3 del for clásico
// === ASIGNACIÓN ARITMÉTICA: suma_pares += ===
// === CARGAR VARIABLE: suma_pares (offset: 8) ===
    ldr x22, [sp, #8]
// === CARGAR VARIABLE: m (offset: 0) ===
    ldr x23, [sp, #0]
    mov x24, x22
    mov x25, x23
    add x26, x24, x25
// === ASIGNAR VARIABLE: suma_pares = int (offset: 8) ===
    str x26, [sp, #8]
// === FIN ASIGNACIÓN ARITMÉTICA: suma_pares += === 

.Lforc_continue_30:
// Post-iteración del for clásico
// === ASIGNACIÓN ARITMÉTICA: m += ===
// === CARGAR VARIABLE: m (offset: 0) ===
    ldr x27, [sp, #0]
    mov x28, x27
    mov x29, #1
    add x30, x28, x29
// === ASIGNAR VARIABLE: m = int (offset: 0) ===
    str x30, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: m += === 

    b .Lforc_start_30
.Lforc_end_30:
// === FIN FOR CLÁSICO ARM64 INTERNO ===
// === FIN FOR CLÁSICO ARM64 ===
// Pop control: for_clasico
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: suma_pares (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #20
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_33
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos_continue += ===
// === CARGAR VARIABLE: puntos_continue (offset: 16) ===
    ldr x13, [sp, #16]
    mov x14, x13
    mov x15, #3
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos_continue = int (offset: 16) ===
    str x16, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos_continue += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_23
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_32
.Lskip_branch_33:
// Condición falsa - continuar al siguiente branch
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_32:
// === FIN ESTRUCTURA IF-ELSE ===

// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// ===  EXPRESIÓN BINARIA + ===
// === CARGAR VARIABLE: puntos_entornos (offset: 184) ===
    ldr x9, [sp, #184]
// === CARGAR VARIABLE: puntos_if (offset: 144) ===
    ldr x10, [sp, #144]
    mov x11, x9
    mov x12, x10
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA === 

// === CARGAR VARIABLE: puntos_while (offset: 136) ===
    ldr x14, [sp, #136]
    mov x15, x13
    mov x16, x14
    add x17, x15, x16
// === FIN EXPRESIÓN BINARIA === 

// === CARGAR VARIABLE: puntos_for (offset: 96) ===
    ldr x18, [sp, #96]
    mov x19, x17
    mov x20, x18
    add x21, x19, x20
// === FIN EXPRESIÓN BINARIA === 

// === CARGAR VARIABLE: puntos_case (offset: 56) ===
    ldr x22, [sp, #56]
    mov x23, x21
    mov x24, x22
    add x25, x23, x24
// === FIN EXPRESIÓN BINARIA === 

// === CARGAR VARIABLE: puntos_break (offset: 40) ===
    ldr x26, [sp, #40]
    mov x27, x25
    mov x28, x26
    add x29, x27, x28
// === FIN EXPRESIÓN BINARIA === 

// === CARGAR VARIABLE: puntos_continue (offset: 16) ===
    ldr x30, [sp, #16]
// ADVERTENCIA: Reciclando registros temporales
    mov x9, x29
    mov x10, x30
    add x11, x9, x10
// === FIN EXPRESIÓN BINARIA === 

// === ASIGNAR VARIABLE: puntos = int (offset: 192) ===
    str x11, [sp, #192]
// === IMPRIMIR STRING  ===
    adr x9, msg_24
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_25
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_26
    mov x0, x9
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_entornos (offset: 184) ===
    ldr x9, [sp, #184]
// === IMPRIMIR STRING  ===
    adr x10, msg_27
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
    adr x11, msg_28
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_if (offset: 144) ===
    ldr x9, [sp, #144]
// === IMPRIMIR STRING  ===
    adr x10, msg_29
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
    adr x11, msg_30
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_while (offset: 136) ===
    ldr x9, [sp, #136]
// === IMPRIMIR STRING  ===
    adr x10, msg_31
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
    adr x11, msg_32
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_for (offset: 96) ===
    ldr x9, [sp, #96]
// === IMPRIMIR STRING  ===
    adr x10, msg_33
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
    adr x11, msg_34
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_case (offset: 56) ===
    ldr x9, [sp, #56]
// === IMPRIMIR STRING  ===
    adr x10, msg_35
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
    adr x11, msg_36
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_break (offset: 40) ===
    ldr x9, [sp, #40]
// === IMPRIMIR STRING  ===
    adr x10, msg_37
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
    adr x11, msg_38
    mov x0, x11
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos_continue (offset: 16) ===
    ldr x9, [sp, #16]
// === IMPRIMIR STRING  ===
    adr x10, msg_39
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
    adr x11, msg_40
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_41
    mov x0, x9
    bl print_string
    bl print_newline
// === CARGAR VARIABLE: puntos (offset: 192) ===
    ldr x9, [sp, #192]
// === IMPRIMIR STRING  ===
    adr x10, msg_42
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
    adr x11, msg_43
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_44
    mov x0, x9
    bl print_string
    bl print_newline

    // Limpiar stack de slices antes de salir
    add sp, sp, #256
    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

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
    mov   x19, x0                 // Guardar dirección del string

    bl    strlen                  // Calcular longitud
    mov   x2, x0                  // x2 = longitud
    mov   x1, x19                 // x1 = dirección string
    mov   x0, #1                  // stdout
    mov   x8, #64                 // Syscall write
    svc   0

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

