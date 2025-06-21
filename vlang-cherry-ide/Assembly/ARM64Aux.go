package assembly

import (
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
    codigo += generarSeccionDatos()

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

// generarSeccionDatos genera la sección .data
func generarSeccionDatos() string {
    codigo := ".section .data\n"
    codigo += "buffer_int: .skip 32\n"
    codigo += "msg_nl: .asciz \"\\n\"\n"
    codigo += "msg_menos: .asciz \"-\"\n"
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

// generarFuncionesAuxiliares genera las funciones auxiliares de ARM64
func generarFuncionesAuxiliares() string {
    return `// --------------------------------------------------------
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
`
}