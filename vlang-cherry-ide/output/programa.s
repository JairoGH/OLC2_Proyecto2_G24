.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "\n==== For Range ===="
    msg_2: .asciz "For range con slice"
    msg_3: .asciz "Índice"
    msg_4: .asciz "="
    msg_5: .asciz "OK For range con slice: correcto"
    msg_6: .asciz "X For range con slice: incorrecto"
    msg_7: .asciz "OK For range con índices: correcto"
    msg_8: .asciz "X For range con índices: incorrecto"
    msg_9: .asciz "PUNTOS POR FOR RANGE:"
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
// === GENERANDO FUNCIÓN MAIN CON CÓDIGO PRINCIPAL ===
// Procesando función main desde Main_func
// === FUNCIÓN MAIN ===
fn_main:
        stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
        mov   x29, sp                 // Configurar frame pointer
// === IMPRIMIR STRING  ===
    adr x9, msg_1
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: puntosForRange ===
    sub sp, sp, #8  // Reservar espacio para puntosForRange
    mov x9, #0
    str x9, [sp]
// puntosForRange en [sp+0] 

// === IMPRIMIR STRING  ===
    adr x9, msg_2
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARACIÓN SLICE numeros: []int ===
// Ajustar stack para slices: 256 bytes
    sub sp, sp, #256
    mov x28, x29
    sub x28, x28, #256
// Elemento 0 (literal): 10
    mov x0, #10
    str x0, [x28, #0]
// Elemento 1 (literal): 20
    mov x0, #20
    str x0, [x28, #8]
// Elemento 2 (literal): 30
    mov x0, #30
    str x0, [x28, #16]
// Elemento 3 (literal): 40
    mov x0, #40
    str x0, [x28, #24]
// Elemento 4 (literal): 50
    mov x0, #50
    str x0, [x28, #32]
// === DECLARAR VARIABLE: suma ===
    sub sp, sp, #8  // Reservar espacio para suma
    mov x9, #0
    str x9, [sp]
// suma en [sp+0] 

// === DECLARAR VARIABLE MUT: sumaIndices ===
    sub sp, sp, #8  // Reservar espacio para sumaIndices
    mov x9, #0
    str x9, [sp]
// sumaIndices en [sp+0] 

// === INICIO FOR TIPO RANGE ===
// Verificando slice 'numeros': true
// Verificando slice 'numeros': true
// Verificando slice 'numeros': true
// Slice 'numeros' encontrado exitosamente
// Verificando slice 'numeros': true
// === GENERANDO FOR RANGE ARM64 ===
// Verificando slice 'numeros': true
// Push control: for_range
    mov x25, #0
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 5 elementos
    mov x9, #5
    mov x26, x9
// Push scope: for_range_1
    sub sp, sp, #8  // Reservar espacio para idx
    str x25, [sp]
// idx en [sp+0] 

    sub sp, sp, #8  // Reservar espacio para val
    mov x10, #0
    str x10, [sp]
// val en [sp+0] 

.Lfor_range_start_1:
    cmp x25, x26
    bge .Lfor_range_end_1
    mov x27, x25
// === ASIGNAR VARIABLE: idx = int (offset: 8) ===
    str x27, [sp, #8]
// === ACCESO POR ÍNDICE numeros[<nil>] ===
// Acceso con verificación de rango en tiempo de ejecución
    mov x9, x27
    cmp x9, #0
    blt range_invalid_numeros_3
    cmp x9, #5
    bge range_invalid_numeros_3
range_valid_numeros_2:
    add x10, x28, #0
    ldr x11, [x10, x9, lsl #3]
    b access_end_numeros_4
range_invalid_numeros_3:
    mov x11, #0
    sub x11, x11, #1
access_end_numeros_4:
// Acceso completado, resultado en x11
// === ASIGNAR VARIABLE: val = int (offset: 0) ===
    str x11, [sp, #0]
// Ejecutando sentencia 1 del for range
// Verificando slice 'idx': false
// === CARGAR VARIABLE: idx (offset: 8) ===
    ldr x11, [sp, #8]
// Verificando slice 'val': false
// === CARGAR VARIABLE: val (offset: 0) ===
    ldr x12, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x13, msg_3
    mov x0, x13
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x14, msg_4
    mov x0, x14
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x12
    bl print_int
    bl print_newline
    bl print_newline
// Ejecutando sentencia 2 del for range
// === ASIGNACIÓN ARITMÉTICA: suma += ===
// === CARGAR VARIABLE: suma (offset: 24) ===
    ldr x15, [sp, #24]
// === CARGAR VARIABLE: val (offset: 0) ===
    ldr x16, [sp, #0]
    mov x17, x15
    mov x18, x16
    add x9, x17, x18
// === ASIGNAR VARIABLE: suma = int (offset: 24) ===
    str x9, [sp, #24]
// === FIN ASIGNACIÓN ARITMÉTICA: suma += === 

// Ejecutando sentencia 3 del for range
// === ASIGNACIÓN ARITMÉTICA: sumaIndices += ===
// === CARGAR VARIABLE: sumaIndices (offset: 16) ===
    ldr x10, [sp, #16]
// === CARGAR VARIABLE: idx (offset: 8) ===
    ldr x11, [sp, #8]
    mov x12, x10
    mov x13, x11
    add x14, x12, x13
// === ASIGNAR VARIABLE: sumaIndices = int (offset: 16) ===
    str x14, [sp, #16]
// === FIN ASIGNACIÓN ARITMÉTICA: sumaIndices += === 

.Lfor_range_continue_1:
    add x25, x25, #1
    b .Lfor_range_start_1
.Lfor_range_end_1:
// Pop scope
// === FIN GENERACIÓN FOR RANGE ARM64 ===
// Pop control: for_range
// === FIN FOR TIPO RANGE ===
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: suma (offset: 24) ===
    ldr x9, [sp, #24]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #150
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_5
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntosForRange += ===
// === CARGAR VARIABLE: puntosForRange (offset: 32) ===
    ldr x13, [sp, #32]
    mov x14, x13
    mov x15, #2
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntosForRange = int (offset: 32) ===
    str x16, [sp, #32]
// === FIN ASIGNACIÓN ARITMÉTICA: puntosForRange += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_5
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_4
.Lskip_branch_5:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x18, msg_6
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_4:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: sumaIndices (offset: 16) ===
    ldr x9, [sp, #16]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_7
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntosForRange += ===
// === CARGAR VARIABLE: puntosForRange (offset: 32) ===
    ldr x13, [sp, #32]
    mov x14, x13
    mov x15, #1
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntosForRange = int (offset: 32) ===
    str x16, [sp, #32]
// === FIN ASIGNACIÓN ARITMÉTICA: puntosForRange += === 

// === IMPRIMIR STRING  ===
    adr x17, msg_7
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_6
.Lskip_branch_7:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x18, msg_8
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_6:
// === FIN ESTRUCTURA IF-ELSE ===

// Verificando slice 'puntosForRange': false
// === CARGAR VARIABLE: puntosForRange (offset: 32) ===
    ldr x9, [sp, #32]
// === IMPRIMIR STRING  ===
    adr x10, msg_9
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === SALIDA EXITOSA DEL PROGRAMA ===
    mov   x0, #0                  // Exit status: success (0)
    mov   x8, #93                 // sys_exit syscall number
    svc   #0                      // Llamada al sistema para terminar programa
// === FIN FUNCIÓN MAIN ===
// Programa termina aquí - no hay return a _start
    
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

