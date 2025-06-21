.section .data
buffer_int: .skip 32
buffer_float: .skip 64
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0
float_const_0: .double 2.500000
float_const_1: .double 3.700000
float_const_2: .double 3.000000
float_const_3: .double 3.500000

.section .text
.global _start

_start:
    ldr d0, =float_const_0
    ldr d1, =float_const_1
    fadd d2, d0, d1
// Imprimir float
    fmov d0, d2
    bl print_float
// println: Agregar dos saltos de línea
    bl print_newline
    bl print_newline

    mov x9, #10
    scvtf d0, x9
    ldr d1, =float_const_2
    fdiv d2, d0, d1
// Imprimir float
    fmov d0, d2
    bl print_float
// println: Agregar dos saltos de línea
    bl print_newline
    bl print_newline

    mov x9, #2
    scvtf d0, x9
    ldr d1, =float_const_3
    fmul d2, d0, d1
// Imprimir float
    fmov d0, d2
    bl print_float
// println: Agregar dos saltos de línea
    bl print_newline
    bl print_newline


    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
// FUNCIONES: print_int, print_float y print_newline
// --------------------------------------------------------

print_int:
    stp x29, x30, [sp, #-16]!
    mov x29, sp
    stp x1, x2, [sp, #-16]!
    stp x3, x4, [sp, #-16]!
    stp x5, x6, [sp, #-16]!

    cmp x0, #0
    bge .Lpositive_int
    mov x8, #64
    ldr x1, =msg_menos
    mov x2, #1
    mov x0, #1
    svc 0
    neg x0, x0

.Lpositive_int:
    ldr x2, =buffer_int
    add x2, x2, #32
    mov x3, #0
    mov x6, #10

.Lloop:
    udiv x4, x0, x6
    msub x5, x4, x6, x0
    add x5, x5, #48
    sub x2, x2, #1
    strb w5, [x2]
    mov x0, x4
    add x3, x3, #1
    cmp x0, #0
    bne .Lloop

    mov x0, #1
    mov x1, x2
    mov x2, x3
    mov x8, #64
    svc 0

    ldp x5, x6, [sp], #16
    ldp x3, x4, [sp], #16
    ldp x1, x2, [sp], #16
    ldp x29, x30, [sp], #16
    ret

print_float:
    stp x29, x30, [sp, #-16]!
    mov x29, sp
    stp x1, x2, [sp, #-16]!
    stp x3, x4, [sp, #-16]!
    stp x5, x6, [sp, #-16]!
    stp d1, d2, [sp, #-16]!

    // Convertir float a entero para la parte entera
    fcvtzs x1, d0
    
    // Imprimir parte entera
    mov x0, x1
    bl print_int
    
    // Imprimir punto decimal
    mov x8, #64
    ldr x1, =msg_punto
    mov x2, #1
    mov x0, #1
    svc 0
    
    // Para simplificar, imprimir solo 2 decimales
    // Obtener parte fraccionaria: (float - int) * 100
    scvtf d1, x1           // Convertir entero de vuelta a float
    fsub d2, d0, d1        // d2 = parte fraccionaria
    ldr d1, =const_100     // ✅ CORREGIDO: Cargar 100.0 desde memoria
    fmul d2, d2, d1        // d2 = fraccionaria * 100
    fcvtzs x1, d2          // Convertir a entero
    
    // Asegurar que sea valor absoluto
    cmp x1, #0
    bge .Lpositive_frac
    neg x1, x1
.Lpositive_frac:
    
    // Imprimir parte fraccionaria (siempre 2 dígitos)
    mov x2, #10
    udiv x3, x1, x2        // Primer dígito
    msub x4, x3, x2, x1    // Segundo dígito
    
    add x3, x3, #48        // Convertir a ASCII
    add x4, x4, #48        // Convertir a ASCII
    
    // Imprimir primer dígito
    strb w3, [sp, #-1]!
    mov x8, #64
    mov x1, sp
    mov x2, #1
    mov x0, #1
    svc 0
    add sp, sp, #1
    
    // Imprimir segundo dígito
    strb w4, [sp, #-1]!
    mov x8, #64
    mov x1, sp
    mov x2, #1
    mov x0, #1
    svc 0
    add sp, sp, #1

    ldp d1, d2, [sp], #16
    ldp x5, x6, [sp], #16
    ldp x3, x4, [sp], #16
    ldp x1, x2, [sp], #16
    ldp x29, x30, [sp], #16
    ret

print_newline:
    mov x0, #1
    ldr x1, =msg_nl
    mov x2, #1
    mov x8, #64
    svc 0
    ret
