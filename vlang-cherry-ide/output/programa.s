.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "\nFor como while anidado (patrón X)"
    msg_2: .asciz ""
    msg_3: .asciz ""
    msg_4: .asciz "100"
    msg_5: .asciz " "
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
// === IMPRIMIR STRING  ===
    adr x9, msg_1
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: n ===
// === DECLARAR VARIABLE: n int ===
    sub sp, sp, #8  // Reservar espacio para n
    mov x9, #5
    str x9, [sp]
// Variable n declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut n := ? (int)
// === DECLARAR VARIABLE MUT: x ===
// === DECLARAR VARIABLE: x int ===
    sub sp, sp, #8  // Reservar espacio para x
    mov x9, #0
    str x9, [sp]
// Variable x declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut x := ? (int)
// === DECLARAR VARIABLE MUT: j ===
// === DECLARAR VARIABLE: j int ===
    sub sp, sp, #8  // Reservar espacio para j
    mov x9, #0
    str x9, [sp]
// Variable j declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut j := ? (int)
// === DECLARAR VARIABLE MUT: fila ===
// === DECLARAR VARIABLE: fila string ===
    sub sp, sp, #8  // Reservar espacio para fila
    adr x9, msg_2
    str x9, [sp]
// Variable fila declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut fila := ? (string)
// === INICIO FOR TIPO WHILE (WhileStmt) ===
// Push scope: while
// ===  EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x9, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: n ===
// === CARGAR VARIABLE: n (offset: 24) ===
    ldr x10, [sp, #24]
// Variable cargada exitosamente: n
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x12, x9
    mov x13, x10
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA ===
// === GENERANDO BUCLE WHILE ARM64 ===
.Lwhile_start_1:
// >>> INICIO DEL BUCLE WHILE <<<
// Evaluando condición del while...
// ===  EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x14, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: n ===
// === CARGAR VARIABLE: n (offset: 24) ===
    ldr x15, [sp, #24]
// Variable cargada exitosamente: n
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x17, x14
    mov x18, x15
    cmp x17, x18
    cset x16, lt
// === FIN EXPRESIÓN BINARIA ===
    cmp x16, #0
    beq .Lwhile_end_1
// Ejecutando sentencia 1 del while
// === ASIGNAR VARIABLE: j = int (offset: 8) ===
    mov x19, #0
    str x19, [sp, #8]
// Ejecutando sentencia 2 del while
// === ASIGNAR VARIABLE: fila = string (offset: 0) ===
    adr x20, msg_3
    str x20, [sp, #0]
// Ejecutando sentencia 3 del while
// === INICIO FOR TIPO WHILE (WhileStmt) ===
// Push scope: while
// ===  EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: j ===
// === CARGAR VARIABLE: j (offset: 8) ===
    ldr x21, [sp, #8]
// Variable cargada exitosamente: j
// === ACCESO A VARIABLE: n ===
// === CARGAR VARIABLE: n (offset: 24) ===
    ldr x22, [sp, #24]
// Variable cargada exitosamente: n
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x24, x21
    mov x25, x22
    cmp x24, x25
    cset x23, lt
// === FIN EXPRESIÓN BINARIA ===
// === GENERANDO BUCLE WHILE ARM64 ===
.Lwhile_start_2:
// >>> INICIO DEL BUCLE WHILE <<<
// Evaluando condición del while...
// ===  EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: j ===
// === CARGAR VARIABLE: j (offset: 8) ===
    ldr x26, [sp, #8]
// Variable cargada exitosamente: j
// === ACCESO A VARIABLE: n ===
// === CARGAR VARIABLE: n (offset: 24) ===
    ldr x27, [sp, #24]
// Variable cargada exitosamente: n
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x29, x26
    mov x30, x27
    cmp x29, x30
    cset x28, lt
// === FIN EXPRESIÓN BINARIA ===
    cmp x28, #0
    beq .Lwhile_end_2
// Ejecutando sentencia 1 del while
// Reseteando contadores de registros por límite
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// ===  EXPRESIÓN BINARIA || ===
// ===  EXPRESIÓN BINARIA == ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x9, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: j ===
// === CARGAR VARIABLE: j (offset: 8) ===
    ldr x10, [sp, #8]
// Variable cargada exitosamente: j
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x12, x9
    mov x13, x10
    cmp x12, x13
    cset x11, eq
// === FIN EXPRESIÓN BINARIA ===
// ===  EXPRESIÓN BINARIA == ===
// ===  EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x14, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: j ===
// === CARGAR VARIABLE: j (offset: 8) ===
    ldr x15, [sp, #8]
// Variable cargada exitosamente: j
// Operador: +, Tipos: int + int
    mov x16, x14
    mov x17, x15
    add x18, x16, x17
// === FIN EXPRESIÓN BINARIA ===
// ===  EXPRESIÓN BINARIA - ===
// === ACCESO A VARIABLE: n ===
// === CARGAR VARIABLE: n (offset: 24) ===
    ldr x19, [sp, #24]
// Variable cargada exitosamente: n
// Operador: -, Tipos: int - int
    mov x20, x19
    mov x21, #1
    sub x22, x20, x21
// === FIN EXPRESIÓN BINARIA ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x24, x18
    mov x25, x22
    cmp x24, x25
    cset x23, eq
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x27, x11
    cmp x27, #1
    beq .Lor_true_3
    mov x28, x23
    orr x26, x27, x28
    b .Lor_end_4
.Lor_true_3:
    mov x26, #1
.Lor_end_4:
// === FIN EXPRESIÓN BINARIA ===
    cmp x26, #0
    beq .Lskip_branch_5
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: fila += ===
// === CARGAR VARIABLE: fila (offset: 0) ===
    ldr x29, [sp, #0]
// Operador: += -> +
// === CONCATENACIÓN DE STRINGS ===
// ADVERTENCIA: Reciclando registros temporales
    mov x30, x29
    adr x9, msg_4
    adr x10, buffer_string
// Concatenar strings: x10 = x30 + x9
    mov x0, x30
    mov x1, x9
    mov x2, x10
    bl concat_strings
// === ASIGNAR VARIABLE: fila = string (offset: 0) ===
    str x10, [sp, #0]
// Asignación aritmética completada: fila +=
// Pop scope
    b .Lif_final_2
.Lskip_branch_5:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === ASIGNACIÓN ARITMÉTICA: fila += ===
// === CARGAR VARIABLE: fila (offset: 0) ===
    ldr x11, [sp, #0]
// Operador: += -> +
// === CONCATENACIÓN DE STRINGS ===
    mov x12, x11
    adr x13, msg_5
    adr x14, buffer_string
// Concatenar strings: x14 = x12 + x13
    mov x0, x12
    mov x1, x13
    mov x2, x14
    bl concat_strings
// === ASIGNAR VARIABLE: fila = string (offset: 0) ===
    str x14, [sp, #0]
// Asignación aritmética completada: fila +=
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_2:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// Ejecutando sentencia 2 del while
// === ASIGNACIÓN ARITMÉTICA: j += ===
// === CARGAR VARIABLE: j (offset: 8) ===
    ldr x15, [sp, #8]
// Operador: += -> +
    mov x16, x15
    mov x17, #1
    add x18, x16, x17
// === ASIGNAR VARIABLE: j = int (offset: 8) ===
    str x18, [sp, #8]
// Asignación aritmética completada: j +=
.Lwhile_continue_2:
// >>> PUNTO DE CONTINUE <<<
    b .Lwhile_start_2
.Lwhile_end_2:
// >>> FIN DEL BUCLE WHILE <<<
// === FIN GENERACIÓN WHILE ARM64 ===
// === FIN FOR TIPO WHILE (WhileStmt) ===
// Pop scope
// Ejecutando sentencia 4 del while
// === CARGAR VARIABLE: fila (offset: 0) ===
    ldr x19, [sp, #0]
// === IMPRIMIR STRING  ===
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
// Ejecutando sentencia 5 del while
// === ASIGNACIÓN ARITMÉTICA: x += ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x20, [sp, #16]
// Operador: += -> +
    mov x21, x20
    mov x22, #1
    add x23, x21, x22
// === ASIGNAR VARIABLE: x = int (offset: 16) ===
    str x23, [sp, #16]
// Asignación aritmética completada: x +=
.Lwhile_continue_1:
// >>> PUNTO DE CONTINUE <<<
    b .Lwhile_start_1
.Lwhile_end_1:
// >>> FIN DEL BUCLE WHILE <<<
// === FIN GENERACIÓN WHILE ARM64 ===
// === FIN FOR TIPO WHILE (WhileStmt) ===
// Pop scope

    // Limpiar stack de slices antes de salir
    add sp, sp, #256
    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
//            FUNCIONES AUXILIARES ARM64
// --------------------------------------------------------

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

concat_strings:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros
    stp   x21, x22, [sp, #-16]!

    mov   x19, x0                 // string1
    mov   x20, x1                 // string2
    mov   x21, x2                 // buffer destino
    mov   x22, #0                 // índice destino

.Lconcat_copy1:                   // Copiar string1
    ldrb  w3, [x19], #1           // Cargar byte y avanzar
    cmp   w3, #0                  // ¿Es '\0'?
    beq   .Lconcat_copy2
    strb  w3, [x21, x22]          // Guardar en destino
    add   x22, x22, #1            // Avanzar índice
    b     .Lconcat_copy1

.Lconcat_copy2:                   // Copiar string2
    ldrb  w3, [x20], #1           // Cargar byte y avanzar
    strb  w3, [x21, x22]          // Guardar (incluye '\0')
    cmp   w3, #0                  // ¿Era '\0'?
    beq   .Lconcat_end
    add   x22, x22, #1            // Avanzar índice
    b     .Lconcat_copy2

.Lconcat_end:
    mov   x0, x21                 // Retornar resultado

    ldp   x21, x22, [sp], #16     // Restaurar registros
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

