.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "============ Switch/Case ============"
    msg_2: .asciz "Switch simple"
    msg_3: .asciz "Lunes"
    msg_4: .asciz "Martes"
    msg_5: .asciz "Miércoles"
    msg_6: .asciz "Jueves"
    msg_7: .asciz "Viernes"
    msg_8: .asciz "Sábado"
    msg_9: .asciz "Domingo"
    msg_10: .asciz "Día inválido"
    msg_11: .asciz "\nSwitch con default"
    msg_12: .asciz "No se debería imprimir"
    msg_13: .asciz "No se debería imprimir"
    msg_14: .asciz "Número no reconocido, se ejecuta default"
    msg_15: .asciz "\nSwitch con break explícito"
    msg_16: .asciz "No se debería imprimir"
    msg_17: .asciz "Caso 2 - Se ejecuta este y debe detenerse"
    msg_18: .asciz "No debería ejecutarse si el break funciona"
    msg_19: .asciz "No se debería imprimir"
    msg_20: .asciz "TotalPuntos: "
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
// === DECLARAR VARIABLE MUT: puntosSwitch ===
// === DECLARAR VARIABLE: puntosSwitch int ===
    sub sp, sp, #8  // Reservar espacio para puntosSwitch
    mov x9, #0
    str x9, [sp]
// Variable puntosSwitch declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut puntosSwitch := ? (int)
// === IMPRIMIR STRING  ===
    adr x9, msg_2
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: dia ===
// === DECLARAR VARIABLE: dia int ===
    sub sp, sp, #8  // Reservar espacio para dia
    mov x9, #1
    str x9, [sp]
// Variable dia declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut dia := ? (int)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: dia ===
// === CARGAR VARIABLE: dia (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: dia
// Switch sobre expresión tipo: int
// Push control: switch
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_10_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_10_1
// Comparación case 2
    mov x13, #3
    cmp x10, x13
    beq .Lcase_10_2
// Comparación case 3
    mov x14, #4
    cmp x10, x14
    beq .Lcase_10_3
// Comparación case 4
    mov x15, #5
    cmp x10, x15
    beq .Lcase_10_4
// Comparación case 5
    mov x16, #6
    cmp x10, x16
    beq .Lcase_10_5
// Comparación case 6
    mov x17, #7
    cmp x10, x17
    beq .Lcase_10_6
    b .Lswitch_default_10
.Lcase_10_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x18, msg_3
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntosSwitch += ===
// === CARGAR VARIABLE: puntosSwitch (offset: 8) ===
    ldr x19, [sp, #8]
// Operador: += -> +
    mov x20, x19
    mov x21, #1
    add x22, x20, x21
// === ASIGNAR VARIABLE: puntosSwitch = int (offset: 8) ===
    str x22, [sp, #8]
// Asignación aritmética completada: puntosSwitch +=
    b .Lswitch_end_10
.Lcase_10_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x23, msg_4
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x24, msg_5
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x25, msg_6
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x26, msg_7
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x27, msg_8
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x28, msg_9
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lswitch_default_10:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x29, msg_10
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_10:
// === FIN SWITCH STATEMENT ===
// Pop control: switch
// === IMPRIMIR STRING  ===
    adr x9, msg_11
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: numero ===
// === DECLARAR VARIABLE: numero int ===
    sub sp, sp, #8  // Reservar espacio para numero
    mov x9, #100
    str x9, [sp]
// Variable numero declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut numero := ? (int)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: numero ===
// === CARGAR VARIABLE: numero (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: numero
// Switch sobre expresión tipo: int
// Push control: switch
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_11_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_11_1
    b .Lswitch_default_11
.Lcase_11_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x13, msg_12
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x14, msg_13
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lswitch_default_11:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x15, msg_14
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntosSwitch += ===
// === CARGAR VARIABLE: puntosSwitch (offset: 16) ===
    ldr x16, [sp, #16]
// Operador: += -> +
    mov x17, x16
    mov x18, #1
    add x19, x17, x18
// === ASIGNAR VARIABLE: puntosSwitch = int (offset: 16) ===
    str x19, [sp, #16]
// Asignación aritmética completada: puntosSwitch +=
.Lswitch_end_11:
// === FIN SWITCH STATEMENT ===
// Pop control: switch
// === IMPRIMIR STRING  ===
    adr x9, msg_15
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: numeroBreak ===
// === DECLARAR VARIABLE: numeroBreak int ===
    sub sp, sp, #8  // Reservar espacio para numeroBreak
    mov x9, #2
    str x9, [sp]
// Variable numeroBreak declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut numeroBreak := ? (int)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: numeroBreak ===
// === CARGAR VARIABLE: numeroBreak (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: numeroBreak
// Switch sobre expresión tipo: int
// Push control: switch
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_12_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_12_1
// Comparación case 2
    mov x13, #3
    cmp x10, x13
    beq .Lcase_12_2
    b .Lswitch_end_12
.Lcase_12_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x14, msg_16
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_12
.Lcase_12_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x15, msg_17
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntosSwitch += ===
// === CARGAR VARIABLE: puntosSwitch (offset: 24) ===
    ldr x16, [sp, #24]
// Operador: += -> +
    mov x17, x16
    mov x18, #1
    add x19, x17, x18
// === ASIGNAR VARIABLE: puntosSwitch = int (offset: 24) ===
    str x19, [sp, #24]
// Asignación aritmética completada: puntosSwitch +=
// === BREAK STATEMENT ===
// Break desde switch
    b .Lswitch_end_12
// === IMPRIMIR STRING  ===
    adr x20, msg_18
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: puntosSwitch -= ===
// === CARGAR VARIABLE: puntosSwitch (offset: 24) ===
    ldr x21, [sp, #24]
// Operador: -= -> -
    mov x22, x21
    mov x23, #1
    sub x24, x22, x23
// === ASIGNAR VARIABLE: puntosSwitch = int (offset: 24) ===
    str x24, [sp, #24]
// Asignación aritmética completada: puntosSwitch -=
    b .Lswitch_end_12
.Lcase_12_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x25, msg_19
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_12
.Lswitch_end_12:
// === FIN SWITCH STATEMENT ===
// Pop control: switch
// === CARGAR VARIABLE: puntosSwitch (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR STRING  ===
    adr x10, msg_20
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
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

