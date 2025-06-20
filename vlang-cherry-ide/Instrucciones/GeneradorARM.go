package instrucciones

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

func (g *ARMGenerator) Add(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Sub(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Mul(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("MUL %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Div(rd, rs1, rs2 string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("UDIV %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Mov(rd string, imm int) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("MOV %s, #%d", rd, imm))
}

func (g *ARMGenerator) MovReg(rd, rs string) {
	g.Instructions = append(g.Instructions, fmt.Sprintf("MOV %s, %s", rd, rs))
}

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
