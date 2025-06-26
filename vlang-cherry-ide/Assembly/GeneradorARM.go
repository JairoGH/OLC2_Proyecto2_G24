package assembly

import (
	"fmt"
	"strings"
)

// ARMGenerator genera instrucciones ARM en formato texto
type ARMGenerator struct {
	Instructions        []string
	FloatConstants      map[string]string
	StringConstants     map[string]string
	FuncionesUsadas     map[string]bool
	funcionesPendientes []string // 🔴 NUEVO: para guardar funciones al final
}

// Constructor
func NewARMGenerator() *ARMGenerator {
	return &ARMGenerator{
		Instructions:        []string{},
		FloatConstants:      make(map[string]string),
		StringConstants:     make(map[string]string),
		FuncionesUsadas:     make(map[string]bool),
		funcionesPendientes: []string{}, // 🔴 Inicializar
	}
}

func (g *ARMGenerator) AgregarFuncion(nombre string, cuerpo []string) {
	codigo := append([]string{fmt.Sprintf("%s:", nombre)}, cuerpo...)
	codigo = append(codigo, "    ret") // siempre cerrar con ret
	g.funcionesPendientes = append(g.funcionesPendientes, strings.Join(codigo, "\n"))
}

// ============= FUNCIONES AUXILIARES A USAR =============
func (g *ARMGenerator) UsarFuncion(nombreFuncion string) {
	g.FuncionesUsadas[nombreFuncion] = true

	// Registrar dependencias automáticamente
	switch nombreFuncion {
	case "print_int":
		g.FuncionesUsadas["print_char"] = true
	case "print_float":
		g.FuncionesUsadas["print_int"] = true
		g.FuncionesUsadas["print_char"] = true
	case "print_string":
		g.FuncionesUsadas["strlen"] = true
	case "concat_strings":
		// No tiene dependencias adicionales
	case "strcmp":
		// No tiene dependencias adicionales
	case "print_bool":
		g.FuncionesUsadas["print_char"] = true
	}
}

func (g *ARMGenerator) Escribir(linea string) {
	g.Instructions = append(g.Instructions, linea)
}

// ============= OPERACIONES ENTERAS =============
func (g *ARMGenerator) Add(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("add %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) Sub(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("sub %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) Mul(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("mul %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) Div(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("udiv %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) Msub(rd, rn, rm, ra string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("msub %s, %s, %s, %s", strings.ToLower(rd), strings.ToLower(rn), strings.ToLower(rm), strings.ToLower(ra)))
}

// ============= OPERACIONES FLOAT =============
func (g *ARMGenerator) FAdd(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fadd %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) FSub(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fsub %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) FMul(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fmul %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

func (g *ARMGenerator) FDiv(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fdiv %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

// ============= OPERACIONES STRING =============

func (g *ARMGenerator) LoadString(rd, stringLabel string) {
	g.Instructions = append(g.Instructions,
		fmt.Sprintf("adr %s, %s", strings.ToLower(rd), stringLabel))
}

func (g *ARMGenerator) ConcatStrings(destReg, src1Reg, src2Reg string) {
	g.UsarFuncion("concat_strings") //  REGISTRAR USO
	g.Instructions = append(g.Instructions,
		fmt.Sprintf("// Concatenar strings: %s = %s + %s", destReg, src1Reg, src2Reg),
		fmt.Sprintf("mov x0, %s", strings.ToLower(src1Reg)),
		fmt.Sprintf("mov x1, %s", strings.ToLower(src2Reg)),
		fmt.Sprintf("mov x2, %s", strings.ToLower(destReg)),
		"bl concat_strings")
}

func (g *ARMGenerator) AddStringConstant(value string) string {
	constName := fmt.Sprintf("str_const_%d", len(g.StringConstants))
	g.StringConstants[constName] = value
	return constName
}

func (g *ARMGenerator) StrLen(destReg, srcReg string) {
	g.UsarFuncion("strlen") //  REGISTRAR USO
	g.Instructions = append(g.Instructions,
		fmt.Sprintf("mov x0, %s", strings.ToLower(srcReg)),
		"bl strlen",
		fmt.Sprintf("mov %s, x0", strings.ToLower(destReg)))
}

// ============= OPERACIONES DE COMPARACIÓN =============

// Comparaciones enteras
func (g *ARMGenerator) Cmp(rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("cmp %s, %s", strings.ToLower(rs1), strings.ToLower(rs2)))
}

// Establecer registro basado en condición
func (g *ARMGenerator) CSet(rd, condition string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("cset %s, %s", strings.ToLower(rd), condition))
}

// Comparaciones float
func (g *ARMGenerator) FCmp(rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fcmp %s, %s", strings.ToLower(rs1), strings.ToLower(rs2)))
}

// Comparar strings (llamada a función auxiliar)
func (g *ARMGenerator) StrCmp(rs1, rs2, resultReg string) {
	g.UsarFuncion("strcmp") //  REGISTRAR USO
	g.Instructions = append(g.Instructions,
		fmt.Sprintf("mov x0, %s", strings.ToLower(rs1)),
		fmt.Sprintf("mov x1, %s", strings.ToLower(rs2)),
		"bl strcmp",
		fmt.Sprintf("mov %s, x0", strings.ToLower(resultReg)))
}

// NUEVO: Método para agregar llamada a función auxiliar manualmente
func (g *ARMGenerator) LlamarFuncion(nombreFuncion string) {
	g.UsarFuncion(nombreFuncion)
	g.Instructions = append(g.Instructions, fmt.Sprintf("bl %s", nombreFuncion))
}

// ============= OPERACIONES LOGICAS =============
// AND lógico bit a bit
func (g *ARMGenerator) And(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("and %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

// OR lógico bit a bit
func (g *ARMGenerator) Orr(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("orr %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

// XOR (para negación lógica)
func (g *ARMGenerator) Eor(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("eor %s, %s, %s", strings.ToLower(rd), strings.ToLower(rs1), strings.ToLower(rs2)))
}

// XOR con inmediato (para negación lógica con constante)
func (g *ARMGenerator) EorImm(rd, rs string, imm int) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("eor %s, %s, #%d", strings.ToLower(rd), strings.ToLower(rs), imm))
}

// Salto condicional - Branch if Equal
func (g *ARMGenerator) Beq(label string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("beq %s", label))
}

// Salto incondicional - Branch
func (g *ARMGenerator) B(label string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("b %s", label))
}

// ============= MOVIMIENTOS =============
func (g *ARMGenerator) Mov(rd string, imm int) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("mov %s, #%d", strings.ToLower(rd), imm))
}

func (g *ARMGenerator) MovReg(rd, rs string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("mov %s, %s", strings.ToLower(rd), strings.ToLower(rs)))
}

// ============= MOVIMIENTOS FLOAT =============
func (g *ARMGenerator) FMovImm(rd string, value float64) string {
	// Crear nombre único para la constante
	constName := fmt.Sprintf("float_const_%d", len(g.FloatConstants))

	// Guardar la constante para la sección de datos
	g.FloatConstants[constName] = fmt.Sprintf("%.6f", value)

	// Cargar desde memoria
	g.Instructions = append(g.Instructions, fmt.Sprintf("ldr %s, =%s", strings.ToLower(rd), constName))

	return constName
}

func (g *ARMGenerator) FMov(rd, rs string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fmov %s, %s", strings.ToLower(rd), strings.ToLower(rs)))
}

// ============= CONVERSIONES =============
func (g *ARMGenerator) ScvtfIntToFloat(rd, rs string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("scvtf %s, %s", strings.ToLower(rd), strings.ToLower(rs)))
}

func (g *ARMGenerator) FcvtzsFloatToInt(rd, rs string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("fcvtzs %s, %s", strings.ToLower(rd), strings.ToLower(rs)))
}

// ============= UTILIDADES =============
func (g *ARMGenerator) Comment(comment string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("// %s", comment))
}

func (g *ARMGenerator) String() string {
	var sb strings.Builder

	for _, instr := range g.Instructions {
		if strings.HasPrefix(instr, "//") {
			sb.WriteString("    " + instr + "\n")
		} else {
			sb.WriteString("    " + instr + "\n")
		}
	}

	return sb.String()
}

func (gen *ARMGenerator) ExisteEtiqueta(etiqueta string) bool {
	for _, instr := range gen.Instructions {
		if strings.HasPrefix(instr, etiqueta+":") {
			return true
		}
	}
	return false
}

/*
func (g *ARMGenerator) FinalizarPrograma() {
	g.Escribir("\n// === FINALIZACIÓN DEL PROGRAMA ===")
	g.Escribir("mov x8, #93")
	g.Escribir("mov x0, #0")
	g.Escribir("svc 0")

	// Escribir funciones definidas por el usuario
	g.Escribir("\n// === FUNCIONES DEFINIDAS POR EL USUARIO ===")
	for _, f := range g.funcionesPendientes {
		g.Escribir("\n" + f)
	}
}
*/

func (g *ARMGenerator) FinalizarPrograma() {
	// NO agregar código de salida aquí - eso debe estar en _start

	// Solo escribir funciones definidas por el usuario
	g.Escribir("\n// === FUNCIONES DEFINIDAS POR EL USUARIO ===")
	for _, f := range g.funcionesPendientes {
		g.Escribir("\n" + f)
	}
}
