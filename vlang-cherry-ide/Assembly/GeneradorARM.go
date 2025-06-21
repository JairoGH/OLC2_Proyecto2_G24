package assembly

import (
    "fmt"
    "strings"
)

// ARMGenerator genera instrucciones ARM en formato texto
type ARMGenerator struct {
    Instructions []string
}

// Constructor
func NewARMGenerator() *ARMGenerator {
    return &ARMGenerator{
        Instructions: []string{},
    }
}

// Operaciones aritméticas
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

// Movimientos
func (g *ARMGenerator) Mov(rd string, imm int) {
    g.Instructions = append(g.Instructions, fmt.Sprintf("mov %s, #%d", strings.ToLower(rd), imm))
}

func (g *ARMGenerator) MovReg(rd, rs string) {
    g.Instructions = append(g.Instructions, fmt.Sprintf("mov %s, %s", strings.ToLower(rd), strings.ToLower(rs)))
}

// Utilidades
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