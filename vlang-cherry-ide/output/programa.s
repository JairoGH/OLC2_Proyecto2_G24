.section .data
    .align 3    // alinea dobles a 8 bytes
    float_const_1: .double 2.450000
    float_const_2: .double 1.550000
    float_const_3: .double 5.580000
    float_const_4: .double 1.250000
    float_const_5: .double 2.364200
    float_const_6: .double 5.623400
    float_const_7: .double 12.408000
    float_const_8: .double 4.207600
buffer_int: .skip 32
buffer_float: .skip 64
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0

.section .text
.global _start

_start:
    adr x9, float_const_1
    ldr d0, [x9]
    adr x10, float_const_2
    ldr d1, [x10]
    fadd d2, d0, d1
    fmov d0, d2
    bl print_float
    bl print_newline
    adr x11, float_const_3
    ldr d3, [x11]
    adr x12, float_const_4
    ldr d4, [x12]
    fsub d5, d3, d4
    fmov d0, d5
    bl print_float
    bl print_newline
    adr x13, float_const_5
    ldr d6, [x13]
    adr x14, float_const_6
    ldr d7, [x14]
    fmul d8, d6, d7
    fmov d0, d8
    bl print_float
    bl print_newline
    adr x15, float_const_7
    ldr d9, [x15]
    adr x16, float_const_8
    ldr d10, [x16]
    fdiv d11, d9, d10
    fmov d0, d11
    bl print_float
    bl print_newline

    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
// FUNCIONES AUXILIARES ARM64
// --------------------------------------------------------

print_int:
    stp   x29, x30, [sp, #-16]!
    mov   x29, sp
    stp   x1, x2, [sp, #-16]!
    stp   x3, x4, [sp, #-16]!
    stp   x5, x6, [sp, #-16]!

    cmp   x0, #0
    bge   .Lpi_pos
    // negativo
    mov   x8, #64
    ldr   x1, =msg_menos
    mov   x2, #1
    mov   x0, #1
    svc   0
    neg   x0, x0

.Lpi_pos:
    ldr   x2, =buffer_int
    add   x2, x2, #32
    mov   x3, #0
    mov   x6, #10

.Lpi_loop:
    udiv  x4, x0, x6
    msub  x5, x4, x6, x0
    add   x5, x5, #48
    sub   x2, x2, #1
    strb  w5, [x2]
    mov   x0, x4
    add   x3, x3, #1
    cmp   x0, #0
    bne   .Lpi_loop

    mov   x0, #1
    mov   x1, x2
    mov   x2, x3
    mov   x8, #64
    svc   0

    ldp   x5, x6, [sp], #16
    ldp   x3, x4, [sp], #16
    ldp   x1, x2, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

print_float:
    stp   x29, x30, [sp, #-16]!    // frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!    // salvar callee‐saved
    stp   x21, x22, [sp, #-16]!

    // 1) Parte entera → x20
    fcvtzs  x20, d0
    mov     x0, x20
    bl      print_int

    // 2) Punto decimal
    mov     x8, #64
    ldr     x1, =msg_punto
    mov     x2, #1
    mov     x0, #1
    svc     0

    // 3) Fracción * 100
    scvtf   d1, x20              // d1 = float(int(d0))
    fsub    d2, d0, d1           // d2 = fractional part
    mov     x2, #100
    scvtf   d1, x2               // d1 = 100.0
    fmul    d2, d2, d1           // d2 = fraction * 100
    fcvtzs  x20, d2              // x20 = trunc(d2)

    // 4) Valor absoluto
    cmp     x20, #0
    bge     .Lpf_pos
    neg     x20, x20
.Lpf_pos:

    // 5) Extraer dos dígitos
    mov     x19, #10
    udiv    x21, x20, x19        // x21 = tens digit
    msub    x22, x21, x19, x20   // x22 = ones digit
    add     w21, w21, #48        // ASCII tens
    add     w22, w22, #48        // ASCII ones

    // 6) Imprimir dígitos
    mov     x0, x21
    bl      print_char
    mov     x0, x22
    bl      print_char

    // Restaurar y return
    ldp   x21, x22, [sp], #16
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

print_char:
    stp   x29, x30, [sp, #-16]!
    mov   x29, sp

    strb  w0, [sp, #-1]!       // apilar carácter
    mov   x0, #1
    mov   x1, sp
    mov   x2, #1
    mov   x8, #64
    svc   0
    add   sp, sp, #1           // desapilar

    ldp   x29, x30, [sp], #16
    ret

print_newline:
    mov   x0, #1
    ldr   x1, =msg_nl
    mov   x2, #1
    mov   x8, #64
    svc   0
    ret
