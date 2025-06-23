.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "🔹 DECLARACIÓN DE SLICES:\n"
    msg_2: .asciz "hola"
    msg_3: .asciz "mundo"
    msg_4: .asciz "Numeros originales: "
    msg_5: .asciz "Palabras originales: "
    msg_6: .asciz "🔹 FUNCIÓN LEN (longitud):\n"
    msg_7: .asciz "No. Numeros:"
    msg_8: .asciz "No. Palabras: "
    msg_9: .asciz "🔹 ACCESO POR ÍNDICE:\n"
    msg_10: .asciz "Elemento en Posicion: "
    msg_11: .asciz "🔹 ASIGNACIÓN POR ÍNDICE:\n"
    msg_12: .asciz "Después de numeros[1] = 99: "
    msg_13: .asciz "🔹 FUNCIÓN APPEND (agregar elementos):\n"
    msg_14: .asciz "Después de append(numeros, 77): "
    msg_15: .asciz " Cantidas de Elementos "
    msg_16: .asciz "🔹 FUNCIÓN INDEXOF (buscar elementos):\n"
    msg_17: .asciz "Posicion de No. 99 : "
    msg_18: .asciz "Posicion de No. 999 : "
    msg_19: .asciz "🔹 FUNCIÓN JOIN (concatenar con separador):\n"
    msg_20: .asciz " "
    msg_21: .asciz ""
    msg_22: .asciz "JUNTAR PALABRAS"
    msg_23: .asciz " $ "
    msg_24: .asciz ""
    msg_25: .asciz "SEPARAR PALABRAS POR '$' : "
    msg_26: .asciz " * "
    msg_27: .asciz ""
    msg_28: .asciz "SEPARAR PALABRAS POR '*': "
    msg_29: .asciz "\n🔹 RESUMEN FINAL:\n"
    msg_30: .asciz "Numeros finales: "
    msg_31: .asciz " (longitud: "
    msg_32: .asciz ")"
    msg_33: .asciz "Palabras finales: "
    msg_34: .asciz " (longitud: "
    msg_35: .asciz ")"
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
// === DECLARACIÓN SLICE numeros: []int ===
// Ajustar stack para slices: 256 bytes
    sub sp, sp, #256
// Elemento 0 (literal): 10
    mov x0, #10
    str x0, [sp, #0]
// Elemento 1 (literal): 20
    mov x0, #20
    str x0, [sp, #8]
// Elemento 2 (literal): 30
    mov x0, #30
    str x0, [sp, #16]
// === DECLARACIÓN SLICE palabras: []string ===
// Elemento 0 (literal): hola
    adr x0, msg_2
    str x0, [sp, #24]
// Elemento 1 (literal): mundo
    adr x0, msg_3
    str x0, [sp, #32]
// === IMPRIMIR STRING  ===
    adr x9, msg_4
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #0]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #8]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #16]    // cargar int[2]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_5
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #24]    // cargar string[0]
    bl print_string
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #32]    // cargar string[1]
    bl print_string
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_6
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 3 elementos
    mov x9, #3
// === IMPRIMIR STRING  ===
    adr x10, msg_7
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === FUNCIÓN len(palabras) ===
// Slice palabras tiene 2 elementos
    mov x9, #2
// === IMPRIMIR STRING  ===
    adr x10, msg_8
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_9
    mov x0, x9
    bl print_string
    bl print_newline
// === ACCESO POR ÍNDICE numeros[1] ===
    ldr x9, [sp, #8]    // cargar numeros[1]
// === IMPRIMIR STRING  ===
    adr x10, msg_10
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_11
    mov x0, x9
    bl print_string
    bl print_newline
// === ACCESO POR ÍNDICE numeros[1] ===
    ldr x9, [sp, #8]    // cargar numeros[1]
// === ASIGNACIÓN POR ÍNDICE numeros[1] = 99 ===
// Asignación directa al elemento 1
    mov x10, #99
    str x10, [sp, #8]    // numeros[1] = 99
// Asignación a slice completada exitosamente
// === ACCESO POR ÍNDICE numeros[0] ===
    ldr x9, [sp, #0]    // cargar numeros[0]
// === IMPRIMIR STRING  ===
    adr x10, msg_12
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_13
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN append(numeros, 77) ===
// Expandiendo slice numeros de 3 a 4 elementos
// Copiando 3 elementos existentes
    ldr x18, [sp, #0]    // cargar elemento 0 original
    str x18, [sp, #40]    // guardar elemento 0 en nueva posición
    ldr x18, [sp, #8]    // cargar elemento 1 original
    str x18, [sp, #48]    // guardar elemento 1 en nueva posición
    ldr x18, [sp, #16]    // cargar elemento 2 original
    str x18, [sp, #56]    // guardar elemento 2 en nueva posición
// Agregando nuevo elemento en offset 64
    mov x0, #77
    str x0, [sp, #64]
// append completado: numeros ahora tiene 4 elementos
// === IMPRIMIR STRING  ===
    adr x9, msg_14
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #40]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #48]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #56]    // cargar int[2]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #64]    // cargar int[3]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 4 elementos
    mov x9, #4
// === IMPRIMIR STRING  ===
    adr x10, msg_15
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_16
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN indexOf(numeros, 99) ===
// Buscando 99 en slice numeros (tamaño: 4)
    mov x12, #99
    mov x9, #0
indexOf_loop_numeros_1:
    cmp x9, #4
    bge indexOf_not_found_numeros_3
    add x10, sp, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros_2
    add x9, x9, #1
    b indexOf_loop_numeros_1
indexOf_found_numeros_2:
    mov x11, x9
    b indexOf_end_numeros_4
indexOf_not_found_numeros_3:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros_4:
// indexOf completado, resultado en x11
// === IMPRIMIR STRING  ===
    adr x9, msg_17
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === FUNCIÓN indexOf(numeros, 999) ===
// Buscando 999 en slice numeros (tamaño: 4)
    mov x12, #999
    mov x9, #0
indexOf_loop_numeros_5:
    cmp x9, #4
    bge indexOf_not_found_numeros_7
    add x10, sp, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros_6
    add x9, x9, #1
    b indexOf_loop_numeros_5
indexOf_found_numeros_6:
    mov x11, x9
    b indexOf_end_numeros_8
indexOf_not_found_numeros_7:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros_8:
// indexOf completado, resultado en x11
// === IMPRIMIR STRING  ===
    adr x9, msg_18
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_19
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, " ") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_20
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_21
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_9:
    cmp x9, #2
    bge join_end_palabras_10
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_9
join_end_palabras_10:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_22
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, " $ ") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_23
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_24
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_11:
    cmp x9, #2
    bge join_end_palabras_12
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_11
join_end_palabras_12:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_25
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, " * ") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_26
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_27
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_13:
    cmp x9, #2
    bge join_end_palabras_14
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_13
join_end_palabras_14:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_28
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_29
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_30
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #40]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #48]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #56]    // cargar int[2]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #64]    // cargar int[3]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 4 elementos
    mov x9, #4
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
// === IMPRIMIR STRING  ===
    adr x9, msg_33
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #24]    // cargar string[0]
    bl print_string
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #32]    // cargar string[1]
    bl print_string
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(palabras) ===
// Slice palabras tiene 2 elementos
    mov x9, #2
// === IMPRIMIR STRING  ===
    adr x10, msg_34
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
    adr x11, msg_35
    mov x0, x11
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

