.section .data
    .align 3    // alinea dobles a 8 bytes
    float_const_1: .double 3.140000
    msg_2: .asciz "Texto"
    float_const_3: .double 2.718000
    msg_4: .asciz "Hola mundo"
    float_const_5: .double 0.000000
    msg_6: .asciz ""
buffer_int: .skip 32
buffer_float: .skip 64
buffer_string: .skip 512
temp_buffer: .skip 256
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0

.section .text
.global _start

_start:
// === DECLARAR VARIABLE: a int ===
    sub sp, sp, #8  // Reservar espacio para a
    mov x9, #10
    str x9, [sp]
// Variable a declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut a := ?
// === DECLARAR VARIABLE: b float ===
    sub sp, sp, #8  // Reservar espacio para b
    adr x9, float_const_1
    ldr d0, [x9]
    str d0, [sp]
// Variable b declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut b := ?
// === DECLARAR VARIABLE: c string ===
    sub sp, sp, #8  // Reservar espacio para c
    adr x9, msg_2
    str x9, [sp]
// Variable c declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut c := ?
// === DECLARAR VARIABLE: d bool ===
    sub sp, sp, #8  // Reservar espacio para d
    mov x9, #1
    str x9, [sp]
// Variable d declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut d := ?
// === CARGAR VARIABLE: a (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR INT ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: b (offset: 16) ===
    ldr d0, [sp, #16]
// === IMPRIMIR FLOAT ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING ===
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: d (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR BOOL ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline

// === DECLARAR VARIABLE: x int ===
    sub sp, sp, #8  // Reservar espacio para x
    mov x9, #42
    str x9, [sp]
// Variable x declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: y float ===
    sub sp, sp, #8  // Reservar espacio para y
    adr x9, float_const_3
    ldr d0, [x9]
    str d0, [sp]
// Variable y declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: z string ===
    sub sp, sp, #8  // Reservar espacio para z
    adr x9, msg_4
    str x9, [sp]
// Variable z declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: w bool ===
    sub sp, sp, #8  // Reservar espacio para w
    mov x9, #0
    str x9, [sp]
// Variable w declarada en [sp] (offset actual: 0)
// === CARGAR VARIABLE: x (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR INT ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: y (offset: 16) ===
    ldr d0, [sp, #16]
// === IMPRIMIR FLOAT ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: z (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING ===
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: w (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR BOOL ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline

// === DECLARAR VARIABLE: sin_entero int ===
    sub sp, sp, #8  // Reservar espacio para sin_entero
    mov x9, #0
    str x9, [sp]
// Variable sin_entero declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: sin_flotante float ===
    sub sp, sp, #8  // Reservar espacio para sin_flotante
    adr x10, float_const_5
    ldr d0, [x10]
    str d0, [sp]
// Variable sin_flotante declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: sin_texto string ===
    sub sp, sp, #8  // Reservar espacio para sin_texto
    adr x9, msg_6
    str x9, [sp]
// Variable sin_texto declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: sin_bool bool ===
    sub sp, sp, #8  // Reservar espacio para sin_bool
    mov x9, #0
    str x9, [sp]
// Variable sin_bool declarada en [sp] (offset actual: 0)
// === CARGAR VARIABLE: sin_entero (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR INT ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: sin_flotante (offset: 16) ===
    ldr d0, [sp, #16]
// === IMPRIMIR FLOAT ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: sin_texto (offset: 8) ===
    ldr x9, [sp, #8]
// === IMPRIMIR STRING ===
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline

// === CARGAR VARIABLE: sin_bool (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR BOOL ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline


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

    cmp   x0, #0                  // ¿Es negativo?
    bge   .Lpi_pos
    // negativo
    mov   x8, #64                 // Syscall write
    ldr   x1, =msg_menos          // Imprimir "-"
    mov   x2, #1
    mov   x0, #1
    svc   0
    neg   x0, x0                  // Hacer positivo

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

print_float:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros
    stp   x21, x22, [sp, #-16]!

    fcvtzs  x20, d0               // Convertir parte entera
    mov     x0, x20
    bl      print_int             // Imprimir parte entera

    mov     x8, #64               // Imprimir punto decimal
    ldr     x1, =msg_punto
    mov     x2, #1
    mov     x0, #1
    svc     0

    scvtf   d1, x20               // Convertir entero a float
    fsub    d2, d0, d1            // d2 = parte fraccionaria
    mov     x2, #100
    scvtf   d1, x2                // d1 = 100.0
    fmul    d2, d2, d1            // Multiplicar por 100
    fcvtzs  x20, d2               // Convertir a entero

    cmp     x20, #0               // Valor absoluto
    bge     .Lpf_pos
    neg     x20, x20
.Lpf_pos:

    mov     x19, #10              // Extraer dígitos
    udiv    x21, x20, x19         // Decenas
    msub    x22, x21, x19, x20    // Unidades
    add     w21, w21, #48         // Convertir a ASCII
    add     w22, w22, #48

    mov     x0, x21               // Imprimir dígitos
    bl      print_char
    mov     x0, x22
    bl      print_char

    ldp   x21, x22, [sp], #16     // Restaurar registros
    ldp   x19, x20, [sp], #16
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

print_bool:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp

    add   w0, w0, #48             // 0→'0', 1→'1' (ASCII)
    bl    print_char              // Imprimir carácter

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

