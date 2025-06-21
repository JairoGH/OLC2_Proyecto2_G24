.section .data
buffer_int: .skip 32
msg_nl: .asciz "\n"
msg_menos: .asciz "-"

.section .text
.global _start

_start:
    // Simula: 10 * 10
    mov x0, #10
    mov x1, #10
    mul x2, x0, x1      // x2 = x0 * x1

    // Imprimir resultado
    mov x0, x2
    bl print_int

    // Imprimir salto de línea
    bl print_newline

    // Salir
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
// FUNCIONES: print_int y print_newline
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

print_newline:
    mov x0, #1
    ldr x1, =msg_nl
    mov x2, #1
    mov x8, #64
    svc 0
    ret

