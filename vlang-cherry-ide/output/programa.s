.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "¡Hola, mundo!"
    msg_2: .asciz "¡Hola,"
    msg_3: .asciz "!"
    msg_4: .asciz "=== PRUEBA COMPLETA DE FUNCIONES PARA 20 PUNTOS ==="
    msg_5: .asciz "\n--- FUNCIONES SIN PARÁMETROS ---"
    msg_6: .asciz "Número obtenido:"
    msg_7: .asciz "✅ obtener_numero() CORRECTO"
    msg_8: .asciz "❌ obtener_numero() INCORRECTO"
    msg_9: .asciz "\n--- FUNCIONES CON PARÁMETROS ---"
    msg_10: .asciz "Juan"
    msg_11: .asciz "Resultado de suma:"
    msg_12: .asciz "✅ sumar() CORRECTO"
    msg_13: .asciz "❌ sumar() INCORRECTO"
    msg_14: .asciz "\n--- ATOI ---"
    msg_15: .asciz "Resultado atoi:"
    msg_16: .asciz "✅ atoi() CORRECTO"
    msg_17: .asciz "❌ atoi() INCORRECTO"
    msg_18: .asciz "\n--- PARSEFLOAT ---"
    float_const_19: .double 123.450000
    msg_20: .asciz "Resultado parseFloat 1:"
    float_const_21: .double 123.000000
    msg_22: .asciz "Resultado parseFloat 2:"
    float_const_23: .double 123.450000
    msg_24: .asciz "✅ parseFloat(\"123.45\") CORRECTO"
    msg_25: .asciz "❌ parseFloat(\"123.45\") INCORRECTO"
    float_const_26: .double 123.000000
    msg_27: .asciz "✅ parseFloat(\"123\") CORRECTO"
    msg_28: .asciz "❌ parseFloat(\"123\") INCORRECTO"
    msg_29: .asciz "\n--- RESUMEN ---"
    msg_30: .asciz "Puntos totales:"
    msg_31: .asciz "/ 20"
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
// Total de sentencias encontradas: 4
// Procesando sentencia 1
// === DECLARACIÓN DE FUNCIÓN ===
// Procesando función: saludar
// DEBUG: Tipo de sentencias: []parser.IStmtContext
// DEBUG: Número de sentencias: 1
// DEBUG: Tipo de sentencia[0]: *parser.StmtContext
// Función 'saludar' registrada exitosamente
// - Parámetros: 0
// - Tipo retorno: void
// Procesando sentencia 2
// === DECLARACIÓN DE FUNCIÓN ===
// Procesando función: obtener_numero
// DEBUG: Tipo de sentencias: []parser.IStmtContext
// DEBUG: Número de sentencias: 1
// DEBUG: Tipo de sentencia[0]: *parser.StmtContext
// Función 'obtener_numero' registrada exitosamente
// - Parámetros: 0
// - Tipo retorno: int
// Procesando sentencia 3
// === DECLARACIÓN DE FUNCIÓN ===
// Procesando función: sumar
// Debug: Verificando métodos disponibles en Lista_parametros
// DEBUG: Tipo de sentencias: []parser.IStmtContext
// DEBUG: Número de sentencias: 1
// DEBUG: Tipo de sentencia[0]: *parser.StmtContext
// Función 'sumar' registrada exitosamente
// - Parámetros: 2
// - Tipo retorno: int
// Procesando sentencia 4
// === DECLARACIÓN DE FUNCIÓN ===
// Procesando función: saludar_persona
// Debug: Verificando métodos disponibles en Lista_parametros
// DEBUG: Tipo de sentencias: []parser.IStmtContext
// DEBUG: Número de sentencias: 1
// DEBUG: Tipo de sentencia[0]: *parser.StmtContext
// Función 'saludar_persona' registrada exitosamente
// - Parámetros: 1
// - Tipo retorno: void
// === GENERANDO TODAS LAS FUNCIONES DE USUARIO ===
// Generando función: saludar
// === GENERANDO CÓDIGO PARA FUNCIÓN: fn_saludar ===
// DEBUG: En generarCodigoFuncion, tipo de fn.Cuerpo: []parser.IStmtContext
fn_saludar:
// Push scope: function
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp                 // Configurar frame pointer
    sub   sp, sp, #64             // Reservar espacio para variables locales
// >>> EJECUTANDO CUERPO DE LA FUNCIÓN <<<
// DEBUG: Kind del cuerpo: slice
// DEBUG: Longitud del slice: 1
// Ejecutando sentencia 1 de la función (tipo: *parser.StmtContext)
// === IMPRIMIR STRING  ===
    adr x9, msg_1
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// No hubo return explícito, agregando return automático
    mov   sp, x29                 // Restaurar stack usando frame pointer
    ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
    ret
// Pop scope
// === FIN CÓDIGO FUNCIÓN: fn_saludar ===
// Generando función: obtener_numero
// === GENERANDO CÓDIGO PARA FUNCIÓN: fn_obtener_numero ===
// DEBUG: En generarCodigoFuncion, tipo de fn.Cuerpo: []parser.IStmtContext
fn_obtener_numero:
// Push scope: function
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp                 // Configurar frame pointer
    sub   sp, sp, #64             // Reservar espacio para variables locales
// >>> EJECUTANDO CUERPO DE LA FUNCIÓN <<<
// DEBUG: Kind del cuerpo: slice
// DEBUG: Longitud del slice: 1
// Ejecutando sentencia 1 de la función (tipo: *parser.StmtContext)
// === RETURN STATEMENT ===
    mov x0, #42
    mov   sp, x29                 // Restaurar stack usando frame pointer
    ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
    ret
// Ya hubo return explícito, no agregar return automático
// Pop scope
// === FIN CÓDIGO FUNCIÓN: fn_obtener_numero ===
// Generando función: sumar
// === GENERANDO CÓDIGO PARA FUNCIÓN: fn_sumar ===
// DEBUG: En generarCodigoFuncion, tipo de fn.Cuerpo: []parser.IStmtContext
fn_sumar:
// Push scope: function
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp                 // Configurar frame pointer
    sub   sp, sp, #64             // Reservar espacio para variables locales
    sub sp, sp, #8  // Reservar espacio para a
    str x0, [sp]
// a en [sp+0] 

    sub sp, sp, #8  // Reservar espacio para b
    str x1, [sp]
// b en [sp+0] 

// >>> EJECUTANDO CUERPO DE LA FUNCIÓN <<<
// DEBUG: Kind del cuerpo: slice
// DEBUG: Longitud del slice: 1
// Ejecutando sentencia 1 de la función (tipo: *parser.StmtContext)
// === RETURN STATEMENT ===
// ===  EXPRESIÓN BINARIA + ===
// === CARGAR VARIABLE: a (offset: 8) ===
    ldr x10, [sp, #8]
// === CARGAR VARIABLE: b (offset: 0) ===
    ldr x11, [sp, #0]
    mov x12, x10
    mov x13, x11
    add x14, x12, x13
// === FIN EXPRESIÓN BINARIA === 

    mov x0, x14
    mov   sp, x29                 // Restaurar stack usando frame pointer
    ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
    ret
// Ya hubo return explícito, no agregar return automático
// Pop scope
// === FIN CÓDIGO FUNCIÓN: fn_sumar ===
// Generando función: saludar_persona
// === GENERANDO CÓDIGO PARA FUNCIÓN: fn_saludar_persona ===
// DEBUG: En generarCodigoFuncion, tipo de fn.Cuerpo: []parser.IStmtContext
fn_saludar_persona:
// Push scope: function
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp                 // Configurar frame pointer
    sub   sp, sp, #64             // Reservar espacio para variables locales
    sub sp, sp, #8  // Reservar espacio para nombre
    mov x15, x0
    str x15, [sp]
// nombre en [sp+0] 

// >>> EJECUTANDO CUERPO DE LA FUNCIÓN <<<
// DEBUG: Kind del cuerpo: slice
// DEBUG: Longitud del slice: 1
// Ejecutando sentencia 1 de la función (tipo: *parser.StmtContext)
// === CARGAR VARIABLE: nombre (offset: 0) ===
    ldr x16, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x17, msg_2
    mov x0, x17
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x16
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x18, msg_3
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// No hubo return explícito, agregando return automático
    mov   sp, x29                 // Restaurar stack usando frame pointer
    ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
    ret
// Pop scope
// === FIN CÓDIGO FUNCIÓN: fn_saludar_persona ===
// Procesando función main
// === FUNCIÓN MAIN ===
fn_main:
        stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
        mov   x29, sp                 // Configurar frame pointer
// === IMPRIMIR STRING  ===
    adr x9, msg_4
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_5
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === LLAMADA A FUNCIÓN USUARIO: saludar ===
// >>> PREPARANDO LLAMADA A FUNCIÓN <<<
// >>> PREPARANDO ARGUMENTOS PARA LLAMADA <<<
    bl fn_saludar
// === DECLARAR VARIABLE: numero ===
// === LLAMADA A FUNCIÓN USUARIO: obtener_numero ===
// >>> PREPARANDO LLAMADA A FUNCIÓN <<<
// >>> PREPARANDO ARGUMENTOS PARA LLAMADA <<<
    bl fn_obtener_numero
    mov x9, x0
    sub sp, sp, #8  // Reservar espacio para numero
    str x9, [sp]
// numero en [sp+0] 

// === CARGAR VARIABLE: numero (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_6
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: numero (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #42
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_1
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_7
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
    adr x14, msg_8
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_0:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_9
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === LLAMADA A FUNCIÓN USUARIO: saludar_persona ===
// >>> PREPARANDO LLAMADA A FUNCIÓN <<<
// >>> PREPARANDO ARGUMENTOS PARA LLAMADA <<<
    adr x0, msg_10
    bl fn_saludar_persona
// === DECLARAR VARIABLE: resultado ===
// === LLAMADA A FUNCIÓN USUARIO: sumar ===
// >>> PREPARANDO LLAMADA A FUNCIÓN <<<
// >>> PREPARANDO ARGUMENTOS PARA LLAMADA <<<
    mov x0, #10
    mov x1, #20
    bl fn_sumar
    mov x9, x0
    sub sp, sp, #8  // Reservar espacio para resultado
    str x9, [sp]
// resultado en [sp+0] 

// === CARGAR VARIABLE: resultado (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_11
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: resultado (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #30
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_3
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_12
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_2
.Lskip_branch_3:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x14, msg_13
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_2:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_14
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: val1 ===
    mov x9, #123
    sub sp, sp, #8  // Reservar espacio para val1
    str x9, [sp]
// val1 en [sp+0] 

// === CARGAR VARIABLE: val1 (offset: 0) ===
    ldr x9, [sp, #0]
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
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: val1 (offset: 0) ===
    ldr x9, [sp, #0]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #123
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_5
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_16
    mov x0, x13
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
    adr x14, msg_17
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_4:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_18
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: f1 ===
    sub sp, sp, #8  // Reservar espacio para f1
    adr x9, float_const_19
    ldr d0, [x9]
    str d0, [sp]
// f1 en [sp+0] 

// === CARGAR VARIABLE: f1 (offset: 0) ===
    ldr d0, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x9, msg_20
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR FLOAT  ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: f2 ===
    sub sp, sp, #8  // Reservar espacio para f2
    adr x9, float_const_21
    ldr d0, [x9]
    str d0, [sp]
// f2 en [sp+0] 

// === CARGAR VARIABLE: f2 (offset: 0) ===
    ldr d0, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x9, msg_22
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR FLOAT  ===
    fmov d0, d0
    bl print_float
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: ok ===
    sub sp, sp, #8  // Reservar espacio para ok
    mov x9, #0
    str x9, [sp]
// ok en [sp+0] 

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: f1 (offset: 16) ===
    ldr d0, [sp, #16]
// === OPERACIÓN RELACIONAL == ===
    fmov d1, d0
    adr x10, float_const_23
    ldr d2, [x10]
    fcmp d1, d2
    cset x9, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x9, #0
    beq .Lskip_branch_7
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x11, msg_24
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: ok += ===
// === CARGAR VARIABLE: ok (offset: 0) ===
    ldr x12, [sp, #0]
    mov x13, x12
    mov x14, #1
    add x15, x13, x14
// === ASIGNAR VARIABLE: ok = int (offset: 0) ===
    str x15, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: ok += === 

// Pop scope
    b .Lif_final_6
.Lskip_branch_7:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x16, msg_25
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_6:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: f2 (offset: 8) ===
    ldr d0, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    fmov d1, d0
    adr x10, float_const_26
    ldr d2, [x10]
    fcmp d1, d2
    cset x9, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x9, #0
    beq .Lskip_branch_9
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x11, msg_27
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNACIÓN ARITMÉTICA: ok += ===
// === CARGAR VARIABLE: ok (offset: 0) ===
    ldr x12, [sp, #0]
    mov x13, x12
    mov x14, #1
    add x15, x13, x14
// === ASIGNAR VARIABLE: ok = int (offset: 0) ===
    str x15, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: ok += === 

// Pop scope
    b .Lif_final_8
.Lskip_branch_9:
// === FIN PROCESAMIENTO BRANCH ===
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x16, msg_28
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_8:
// === FIN ESTRUCTURA IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_29
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: puntos ===
    sub sp, sp, #8  // Reservar espacio para puntos
    mov x9, #0
    str x9, [sp]
// puntos en [sp+0] 

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: numero (offset: 48) ===
    ldr x9, [sp, #48]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #42
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_11
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos += ===
// === CARGAR VARIABLE: puntos (offset: 0) ===
    ldr x13, [sp, #0]
    mov x14, x13
    mov x15, #8
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos = int (offset: 0) ===
    str x16, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos += === 

// Pop scope
    b .Lif_final_10
.Lskip_branch_11:
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_10:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: resultado (offset: 40) ===
    ldr x9, [sp, #40]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #30
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_13
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos += ===
// === CARGAR VARIABLE: puntos (offset: 0) ===
    ldr x13, [sp, #0]
    mov x14, x13
    mov x15, #7
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos = int (offset: 0) ===
    str x16, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos += === 

// Pop scope
    b .Lif_final_12
.Lskip_branch_13:
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_12:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: val1 (offset: 32) ===
    ldr x9, [sp, #32]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #123
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_15
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos += ===
// === CARGAR VARIABLE: puntos (offset: 0) ===
    ldr x13, [sp, #0]
    mov x14, x13
    mov x15, #2
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos = int (offset: 0) ===
    str x16, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos += === 

// Pop scope
    b .Lif_final_14
.Lskip_branch_15:
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_14:
// === FIN ESTRUCTURA IF-ELSE ===

// === INICIO ESTRUCTURA IF-ELSE ===
// ===  EXPRESIÓN BINARIA == ===
// === CARGAR VARIABLE: ok (offset: 8) ===
    ldr x9, [sp, #8]
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    mov x12, #2
    cmp x11, x12
    cset x10, eq
// === FIN EXPRESIÓN BINARIA === 

    cmp x10, #0
    beq .Lskip_branch_17
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === ASIGNACIÓN ARITMÉTICA: puntos += ===
// === CARGAR VARIABLE: puntos (offset: 0) ===
    ldr x13, [sp, #0]
    mov x14, x13
    mov x15, #3
    add x16, x14, x15
// === ASIGNAR VARIABLE: puntos = int (offset: 0) ===
    str x16, [sp, #0]
// === FIN ASIGNACIÓN ARITMÉTICA: puntos += === 

// Pop scope
    b .Lif_final_16
.Lskip_branch_17:
// === FIN PROCESAMIENTO BRANCH ===
.Lif_final_16:
// === FIN ESTRUCTURA IF-ELSE ===

// === CARGAR VARIABLE: puntos (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x10, msg_30
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
    adr x11, msg_31
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
        mov   sp, x29                 // Restaurar stack usando frame pointer
        ldp   x29, x30, [sp], #16     // Restaurar frame pointer y link register
        ret                           // Retornar a _start
// === FIN FUNCIÓN MAIN ===
    
// === FUNCIONES DEFINIDAS POR EL USUARIO ===
// Funciones registradas: 4
// - Función: saludar
// - Función: obtener_numero
// - Función: sumar
// - Función: saludar_persona

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

