package assembly

import (
	"fmt"
	"strings"
)

// VisitorARM64Interface define los métodos que necesita el generador de código
type VisitorARM64Interface interface {
	GetARMGenerator() *ARMGenerator
}

// GenerateCode genera el código ARM64 completo
func GenerateCode(visitor VisitorARM64Interface) string {
	codigo := ""

	// Sección de datos
	codigo += generarSeccionDatos(visitor.GetARMGenerator())

	// Sección de texto
	codigo += generarSeccionTexto()

	// Instrucciones del programa
	codigo += generarInstruccionesPrograma(visitor.GetARMGenerator())

	// Salida del programa
	codigo += generarSalidaPrograma()

	// Solo funciones auxiliares necesarias
	codigo += generarFuncionesAuxiliaresDinamicas(visitor.GetARMGenerator())

	return codigo
}

// generarSeccionDatos genera la sección .data con constantes float y string
func generarSeccionDatos(armGen *ARMGenerator) string {
    codigo := ".section .data\n"
    codigo += "buffer_int: .skip 32\n"
    codigo += "buffer_float: .skip 64\n"
    codigo += "buffer_string: .skip 512\n"
    codigo += "temp_buffer: .skip 256\n"
    codigo += "msg_nl: .asciz \"\\n\"\n"
    codigo += "msg_menos: .asciz \"-\"\n"
    codigo += "msg_punto: .asciz \".\"\n"
    codigo += "const_100: .double 100.0\n"
    codigo += "str_empty: .asciz \"\"\n"

    // CONSTANTES FLOAT
    for constName, value := range armGen.FloatConstants {
        codigo += fmt.Sprintf("%s: .double %s\n", constName, value)
    }

    // CONSTANTES STRING (incluyendo buffers dinámicos)
    for constName, value := range armGen.StringConstants {
        // 🔥 NUEVO: Detectar buffers de join
        if strings.HasPrefix(constName, "join_buffer_") {
            // Crear buffer de 512 bytes lleno de ceros
            codigo += fmt.Sprintf("%s: .skip 512\n", constName)
        } else {
            codigo += fmt.Sprintf("%s: .asciz \"%s\"\n", constName, value)
        }
    }

    codigo += "\n"
    return codigo
}

// generarSeccionTexto genera la sección .text
func generarSeccionTexto() string {
	codigo := ".section .text\n"
	codigo += ".global _start\n\n"
	codigo += "_start:\n"
	codigo += "    bl fn_main      // Llamar a la función main\n"
	codigo += "    mov x8, #93     // sys_exit como respaldo\n"
	codigo += "    mov x0, #0      // exit status\n"
	codigo += "    svc #0          // system call\n\n"
	return codigo
}

// generarInstruccionesPrograma genera las instrucciones del programa
func generarInstruccionesPrograma(armGen *ARMGenerator) string {
	codigo := ""

	for _, instr := range armGen.Instructions {
		if instr == "" {
			codigo += "\n"
			continue
		}

		if !strings.HasPrefix(instr, "//") && !strings.HasSuffix(instr, ":") {
			codigo += "    " + instr + "\n"
		} else {
			codigo += instr + "\n"
		}
	}

	return codigo
}

// generarSalidaPrograma genera el código de salida del programa
func generarSalidaPrograma() string {
	// Ya no generar código de salida aquí porque lo manejamos en _start
	return "\n"
}

// Generar solo las funciones auxiliares que se usan
func generarFuncionesAuxiliaresDinamicas(armGen *ARMGenerator) string {
	var codigo strings.Builder

	funcionesUsadas := armGen.FuncionesUsadas

	// Si no se usa ninguna función auxiliar, no generar nada
	if len(funcionesUsadas) == 0 {
		return ""
	}

	codigo.WriteString("// --------------------------------------------------------\n")
	codigo.WriteString("//            FUNCIONES AUXILIARES ARM64\n")
	codigo.WriteString("// --------------------------------------------------------\n\n")

	//  GENERAR SOLO LAS FUNCIONES USADAS

	if funcionesUsadas["print_int"] {
		codigo.WriteString(getFuncionPrintInt())
	}

	if funcionesUsadas["print_float"] {
		codigo.WriteString(getFuncionPrintFloat())
	}

	if funcionesUsadas["print_string"] {
		codigo.WriteString(getFuncionPrintString())
	}

	if funcionesUsadas["print_bool"] {
		codigo.WriteString(getFuncionPrintBool())
	}

	if funcionesUsadas["print_char"] {
		codigo.WriteString(getFuncionPrintChar())
	}

	if funcionesUsadas["print_newline"] {
		codigo.WriteString(getFuncionPrintNewline())
	}

	if funcionesUsadas["strlen"] {
		codigo.WriteString(getFuncionStrlen())
	}

	if funcionesUsadas["strcmp"] {
		codigo.WriteString(getFuncionStrcmp())
	}

	if funcionesUsadas["concat_strings"] {
		codigo.WriteString(getFuncionConcatStrings())
	}

	return codigo.String()
}

// Funciones individuales para cada función auxiliar

func getFuncionPrintInt() string {
	return `print_int:
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

`
}

func getFuncionPrintFloat() string {
	return `print_float:
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

`
}

/*
func getFuncionPrintString() string {
	return `print_string:
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

`
}*/

// REEMPLAZA la función getFuncionPrintString() en tu assembly/code_generator.go:

func getFuncionPrintString() string {
	return `print_string:
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

`
}

func getFuncionPrintBool() string {
	return `print_bool:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp

    add   w0, w0, #48             // 0→'0', 1→'1' (ASCII)
    bl    print_char              // Imprimir carácter

    ldp   x29, x30, [sp], #16
    ret

`
}

func getFuncionPrintChar() string {
	return `print_char:
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

`
}

func getFuncionPrintNewline() string {
	return `print_newline:
    mov   x0, #1                  // stdout
    ldr   x1, =msg_nl             // "\n"
    mov   x2, #1                  // 1 byte
    mov   x8, #64                 // Syscall write
    svc   0
    ret

`
}

func getFuncionStrlen() string {
	return `strlen:
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

`
}

func getFuncionStrcmp() string {
	return `strcmp:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x2, x3, [sp, #-16]!

strcmp_loop:
    ldrb  w2, [x0], #1            // Cargar char de string1
    ldrb  w3, [x1], #1            // Cargar char de string2
    
    cmp   w2, w3                  // Comparar caracteres
    bne   strcmp_different
    
    cmp   w2, #0                  // ¿Final de string?
    beq   strcmp_equal
    
    b     strcmp_loop
    
strcmp_equal:
    mov   x0, #0                  // Iguales = 0
    b     strcmp_end
    
strcmp_different:
    mov   x0, #1                  // Diferentes = 1
    
strcmp_end:
    ldp   x2, x3, [sp], #16       // Restaurar registros
    ldp   x29, x30, [sp], #16
    ret

`
}

func getFuncionConcatStrings() string {
    return `concat_strings:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros
    stp   x21, x22, [sp, #-16]!

    mov   x19, x0                 // string1
    mov   x20, x1                 // string2  
    mov   x21, x2                 // buffer destino
    mov   x22, #0                 // índice destino

    // VERIFICAR SI STRING1 ES VÁLIDO
    cmp   x19, #0
    beq   .Lconcat_copy2_start    // Si string1 es NULL, ir directo a string2

.Lconcat_copy1:                   // Copiar string1
    ldrb  w3, [x19], #1           // Cargar byte de string1
    cmp   w3, #0                  // ¿Es '\0'?
    beq   .Lconcat_copy2_start    // Terminar string1, ir a string2
    strb  w3, [x21, x22]          // Guardar en destino
    add   x22, x22, #1            // Avanzar índice
    cmp   x22, #500               // Límite de seguridad
    bge   .Lconcat_end_null       // Evitar overflow
    b     .Lconcat_copy1

.Lconcat_copy2_start:
    // VERIFICAR SI STRING2 ES VÁLIDO
    cmp   x20, #0
    beq   .Lconcat_end_null       // Si string2 es NULL, terminar

.Lconcat_copy2:                   // Copiar string2
    ldrb  w3, [x20], #1           // Cargar byte de string2
    strb  w3, [x21, x22]          // Guardar en destino
    cmp   w3, #0                  // ¿Era '\0'?
    beq   .Lconcat_end            // Si era '\0', terminar
    add   x22, x22, #1            // Avanzar índice
    cmp   x22, #500               // Límite de seguridad
    bge   .Lconcat_end_null       // Evitar overflow
    b     .Lconcat_copy2

.Lconcat_end_null:
    // Terminar con '\0' si hubo algún problema
    mov   w3, #0
    strb  w3, [x21, x22]

.Lconcat_end:
    mov   x0, x21                 // Retornar buffer resultado

    ldp   x21, x22, [sp], #16     // Restaurar registros
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

`
}