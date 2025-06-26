.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "Eres mayor de edad"
    msg_2: .asciz "Eres menor de edad"
    msg_3: .asciz "Excelente"
    msg_4: .asciz "Bueno"
    msg_5: .asciz "Necesitas mejorar"
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
// === DECLARAR VARIABLE MUT: edad ===
    sub sp, sp, #8  // Reservar espacio para edad
    mov x9, #18
    str x9, [sp]
// edad en [sp+0] 

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA >= ===
// === CARGAR VARIABLE: edad (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL >= ===
    mov x11, x9
    mov x12, #18
    cmp x11, x12
    cset x10, ge
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_1
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_1
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_0
.Lskip_branch_1:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x14, msg_2
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_0:
// === FIN ESTRUCTURA IF-ELSE ===

// === DECLARAR VARIABLE MUT: nota ===
    sub sp, sp, #8  // Reservar espacio para nota
    mov x9, #85
    str x9, [sp]
// nota en [sp+0] 

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA >= ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL >= ===
    mov x11, x9
    mov x12, #90
    cmp x11, x12
    cset x10, ge
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_3
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_3
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_2
.Lskip_branch_3:
// === FIN PROCESAMIENTO BRANCH ===
// ===  EXPRESIÓN BINARIA >= ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x14, [sp, #0]
// === OPERACIÓN RELACIONAL >= ===
    mov x16, x14
    mov x17, #70
    cmp x16, x17
    cset x15, ge
// === FIN EXPRESIÓN BINARIA === 

    cmp x15, #0
    beq .Lskip_branch_4
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x18, msg_4
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_2
.Lskip_branch_4:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x19, msg_5
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_2:
// === FIN ESTRUCTURA IF-ELSE ===

        mov   sp, x29                 // Restaurar stack usando frame pointer
        ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
        ret                           // Retornar a _start
// === FIN FUNCIÓN MAIN ===
    
// === FUNCIONES DEFINIDAS POR EL USUARIO ===
// Funciones registradas: 0

// --------------------------------------------------------
//            FUNCIONES AUXILIARES ARM64
// --------------------------------------------------------

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

