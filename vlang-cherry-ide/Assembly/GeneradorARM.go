package assembly

import (
	"fmt"
	"strings"
)

// ARMGenerator genera instrucciones ARM en formato texto
type ARMGenerator struct {
	Instructions   []string
	FloatConstants map[string]string // Para almacenar constantes float
}

// Constructor
func NewARMGenerator() *ARMGenerator {
	return &ARMGenerator{
		Instructions:   []string{},
		FloatConstants: make(map[string]string),
	}
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
