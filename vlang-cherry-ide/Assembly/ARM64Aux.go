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

	// Funciones auxiliares
	codigo += generarFuncionesAuxiliares()

	return codigo
}

// generarSeccionDatos genera la sección .data con constantes float
func generarSeccionDatos(armGen *ARMGenerator) string {
	codigo := ".section .data\n"
	codigo += "buffer_int: .skip 32\n"
	codigo += "buffer_float: .skip 64\n"
	codigo += "msg_nl: .asciz \"\\n\"\n"
	codigo += "msg_menos: .asciz \"-\"\n"
	codigo += "msg_punto: .asciz \".\"\n"
	codigo += "const_100: .double 100.0\n" // ✅ AGREGAR constante 100.0

	// ✅ AGREGAR CONSTANTES FLOAT
	for constName, value := range armGen.FloatConstants {
		codigo += fmt.Sprintf("%s: .double %s\n", constName, value)
	}

	codigo += "\n"
	return codigo
}

// generarSeccionTexto genera la sección .text
func generarSeccionTexto() string {
	codigo := ".section .text\n"
	codigo += ".global _start\n\n"
	codigo += "_start:\n"
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
	codigo := "\n    // Salir del programa\n"
	codigo += "    mov x8, #93\n"
	codigo += "    mov x0, #0\n"
	codigo += "    svc 0\n\n"
	return codigo
}

// ✅ CORREGIR función print_float
func generarFuncionesAuxiliares() string {
	return `// --------------------------------------------------------
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
`
}
