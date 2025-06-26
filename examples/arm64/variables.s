.section .data
    .align 3    // alinea dobles a 8 bytes
    float_const_1: .double 3.141590
    msg_2: .asciz "Hola, Cherry!"
    float_const_3: .double 3.140000
    msg_4: .asciz "V-Lang Cherry"
    msg_5: .asciz "entero:"
    msg_6: .asciz "decimal:"
    msg_7: .asciz "texto:"
    msg_8: .asciz "booleano:"
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
    bl fn_main      // Llamar a la función main
    mov x8, #93     // sys_exit como respaldo
    mov x0, #0      // exit status
    svc #0          // system call

// === PROCESANDO PROGRAMA ===
// Total de sentencias encontradas: 0
// === GENERANDO TODAS LAS FUNCIONES DE USUARIO ===
// Procesando función main
// === FUNCIÓN MAIN ===
fn_main:
        stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
        mov   x29, sp                 // Configurar frame pointer
// === DECLARAR VARIABLE MUT: entero ===
    sub sp, sp, #8  // Reservar espacio para entero
    mov x9, #42
    str x9, [sp]
// entero en [sp+0] 

// === DECLARAR VARIABLE MUT: decimal ===
    sub sp, sp, #8  // Reservar espacio para decimal
    adr x9, float_const_1
    ldr d0, [x9]
    str d0, [sp]
// decimal en [sp+0] 

// === DECLARAR VARIABLE MUT: texto ===
    sub sp, sp, #8  // Reservar espacio para texto
    adr x9, msg_2
    str x9, [sp]
// texto en [sp+0] 

// === DECLARAR VARIABLE MUT: booleano ===
// === LITERAL BOOL: true ===
    sub sp, sp, #8  // Reservar espacio para booleano
    mov x9, #1
    str x9, [sp]
// booleano en [sp+0] 

// === DECLARAR VARIABLE MUT: edad ===
    sub sp, sp, #8  // Reservar espacio para edad
    mov x9, #25
    str x9, [sp]
// edad en [sp+0] 

// === DECLARAR VARIABLE MUT: pi ===
    sub sp, sp, #8  // Reservar espacio para pi
    adr x9, float_const_3
    ldr d0, [x9]
    str d0, [sp]
// pi en [sp+0] 

// === DECLARAR VARIABLE MUT: nombre ===
    sub sp, sp, #8  // Reservar espacio para nombre
    adr x9, msg_4
    str x9, [sp]
// nombre en [sp+0] 

// === CARGAR VARIABLE: entero (offset: 48) ===
    ldr x9, [sp, #48]
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
// === CARGAR VARIABLE: decimal (offset: 40) ===
    ldr d0, [sp, #40]
// === IMPRIMIR STRING  ===
    adr x9, msg_6
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR FLOAT  ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: texto (offset: 32) ===
    ldr x9, [sp, #32]
// === IMPRIMIR STRING  ===
    adr x10, msg_7
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: booleano (offset: 24) ===
    ldr x9, [sp, #24]
// === IMPRIMIR STRING  ===
    adr x10, msg_8
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
        mov   sp, x29                 // Restaurar stack usando frame pointer
        ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
        ret                           // Retornar a _start
// === FIN FUNCIÓN MAIN ===
    
// === FUNCIONES DEFINIDAS POR EL USUARIO ===
// Funciones registradas: 0

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

