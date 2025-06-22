.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "HOLA"
    msg_2: .asciz "MUNDO"
    msg_3: .asciz "V"
    msg_4: .asciz "lang"
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
// === CONCATENACIÓN DE STRINGS ===
    adr x9, msg_1
    adr x10, msg_2
    adr x11, buffer_string
// Concatenar strings: x11 = x9 + x10
    mov x0, x9
    mov x1, x10
    mov x2, x11
    bl concat_strings
// === IMPRIMIR STRING ===
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline

// === CONCATENACIÓN DE STRINGS ===
    adr x9, msg_3
    adr x10, msg_4
    adr x11, buffer_string
// Concatenar strings: x11 = x9 + x10
    mov x0, x9
    mov x1, x10
    mov x2, x11
    bl concat_strings
// === IMPRIMIR STRING ===
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline


    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
// FUNCIONES AUXILIARES ARM64 (Solo las necesarias)
// --------------------------------------------------------

print_string:
    stp   x29, x30, [sp, #-16]!
    mov   x29, sp
    mov   x19, x0               // GUARDAR dirección del string en x19

    // Calcular longitud
    bl    strlen                // x0 = longitud del string
    mov   x2, x0                // x2 = longitud
    mov   x1, x19               // RESTAURAR dirección del string
    mov   x0, #1                // stdout
    mov   x8, #64               // syscall write
    svc   0

    ldp   x29, x30, [sp], #16
    ret

print_newline:
    mov   x0, #1
    ldr   x1, =msg_nl
    mov   x2, #1
    mov   x8, #64
    svc   0
    ret

strlen:
    mov   x1, #0                // contador = 0
.Lstrlen_loop:
    ldrb  w2, [x0, x1]         // cargar byte en posición x1
    cmp   w2, #0               // ¿es null terminator?
    beq   .Lstrlen_end
    add   x1, x1, #1           // incrementar contador
    b     .Lstrlen_loop
.Lstrlen_end:
    mov   x0, x1               // retornar longitud en x0
    ret

concat_strings:
    stp   x29, x30, [sp, #-16]!
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!
    stp   x21, x22, [sp, #-16]!

    // x0 = string1, x1 = string2, x2 = buffer destino
    mov   x19, x0               // string1
    mov   x20, x1               // string2  
    mov   x21, x2               // buffer destino
    mov   x22, #0               // índice destino

    // Copiar string1 al destino
.Lconcat_copy1:
    ldrb  w3, [x19], #1        // cargar byte de string1 (post-increment)
    cmp   w3, #0               // ¿es null?
    beq   .Lconcat_copy2       // si es null, empezar con string2
    strb  w3, [x21, x22]       // guardar en destino
    add   x22, x22, #1         // incrementar índice destino
    b     .Lconcat_copy1

    // Copiar string2 al destino
.Lconcat_copy2:
    ldrb  w3, [x20], #1        // cargar byte de string2 (post-increment)
    strb  w3, [x21, x22]       // guardar en destino (incluye null terminator)
    cmp   w3, #0               // ¿era null terminator?
    beq   .Lconcat_end         // si era null, terminamos
    add   x22, x22, #1         // incrementar índice destino
    b     .Lconcat_copy2

.Lconcat_end:
    mov   x0, x21              // retornar dirección del resultado

    ldp   x21, x22, [sp], #16
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

