package instrucciones

import (
	"fmt"
	assembly "main/Assembly"
	"main/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorARM64 struct {
	parser.BaseVGrammarVisitor
	armGen               *assembly.ARMGenerator
	expresionesProcessor *assembly.ExpresionesProcessor
	TablaError           *TablaError
}

func NewVisitorARM64() *VisitorARM64 {
	armGen := assembly.NewARMGenerator()
	return &VisitorARM64{
		armGen:               armGen,
		expresionesProcessor: assembly.NewExpresionesProcessor(armGen),
	}
}

func (v *VisitorARM64) GetARMGenerator() *assembly.ARMGenerator {
	return v.armGen
}

func (v *VisitorARM64) GetCodigo() string {
	codigoBase := assembly.GenerateCode(v)
	mensajesDatos := v.expresionesProcessor.GetMensajesDatos()

	if len(mensajesDatos) == 0 {
		return codigoBase
	}

	lines := strings.Split(codigoBase, "\n")
	var out []string
	inserted := false

	for _, line := range lines {
		out = append(out, line)
		if !inserted && strings.Contains(line, ".section .data") {
			out = append(out, "    .align 3    // alinea dobles a 8 bytes")
			for _, c := range mensajesDatos {
				out = append(out, "    "+c)
			}
			inserted = true
		}
	}

	if !inserted {
		for i, line := range out {
			if strings.Contains(line, ".section .text") {
				prefix := append([]string{}, out[:i]...)
				prefix = append(prefix, "    .align 3    // alinea dobles a 8 bytes")
				for _, c := range mensajesDatos {
					prefix = append(prefix, "    "+c)
				}
				out = append(prefix, out[i:]...)
				break
			}
		}
	}

	return strings.Join(out, "\n")
}

func (v *VisitorARM64) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		token := val.GetSymbol()
		v.TablaError.NewErrorSintactico(token.GetLine(), token.GetColumn(), "Error de análisis: "+val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

// ============= MÉTODOS ANTLR - SOLO LÓGICA DE CONTEXTO =============

func (v *VisitorARM64) VisitProgram(ctx *parser.ProgramContext) interface{} {
	// Procesar sentencias ANTES de main
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Procesar función main si existe
	if ctx.Main_func() != nil {
		v.Visit(ctx.Main_func())
	}

	return nil
}

func (v *VisitorARM64) VisitFuncionMain(ctx *parser.FuncionMainContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		// Resetear contadores para evitar agotar registros
		v.expresionesProcessor.ResetearContadores()
		v.Visit(stmt)
	}
	return nil
}

func (v *VisitorARM64) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones
	if ds := ctx.Stmt_declaracion(); ds != nil {
		return nil
	}

	// 2) Asignaciones
	if ctx.Stmt_asignar() != nil {
		return nil
	}

	// 3) Resto de sentencias
	switch {
	case ctx.If_stmt() != nil:
		return nil
	case ctx.Switch_stmt() != nil:
		return nil
	case ctx.While_stmt() != nil:
		return nil
	case ctx.For_clasico_stmt() != nil:
		return nil
	case ctx.For_stmt() != nil:
		return nil
	case ctx.Stmt_transferencia() != nil:
		return nil
	case ctx.Llamar_funcion() != nil:
		return v.Visit(ctx.Llamar_funcion())
	case ctx.Declarar_funcion() != nil:
		return nil
	case ctx.Declarar_struct() != nil:
		return nil
	case ctx.Fun_slice() != nil:
		return nil
	default:
		return nil
	}
}

func (v *VisitorARM64) VisitLlamarFuncion(ctx *parser.LlamarFuncionContext) interface{} {
	nombreFuncion := ctx.PatronId().GetText()

	if (nombreFuncion == "print" || nombreFuncion == "println") && ctx.Lista_argumentos() != nil {
		// Procesar argumentos de print/println
		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			return nil
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok {
			return nil
		}

		for _, arg := range argumentos {
			if resultado, ok := arg.(*assembly.ResultadoExpresion); ok {
				v.generarCodigoImpresion(resultado, nombreFuncion == "println")
			}
		}
	}
	return nil
}

func (v *VisitorARM64) VisitListaArgumentos(ctx *parser.ListaArgumentosContext) interface{} {
	var argumentos []interface{}

	for _, argCtx := range ctx.AllArgumento_fun() {
		argResult := v.Visit(argCtx)
		if argResult != nil {
			argumentos = append(argumentos, argResult)
		}
	}

	return argumentos
}

func (v *VisitorARM64) VisitFuncionArg(ctx *parser.FuncionArgContext) interface{} {
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	if ctx.PatronId() != nil {
		return v.Visit(ctx.PatronId())
	}
	return nil
}

// ============= EXPRESIONES - DELEGANDO A ExpresionesProcessor =============

func (v *VisitorARM64) VisitBinarioExp(ctx *parser.BinarioExpContext) interface{} {
	leftResult := v.Visit(ctx.Expr(0))
	rightResult := v.Visit(ctx.Expr(1))

	if leftResult == nil || rightResult == nil {
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	left := leftResult.(*assembly.ResultadoExpresion)
	right := rightResult.(*assembly.ResultadoExpresion)
	operador := ctx.GetOp().GetText()

	// DELEGAR a ExpresionesProcessor
	return v.expresionesProcessor.ProcesarOperacionBinaria(left, right, operador)
}

func (v *VisitorARM64) VisitUnarioExp(ctx *parser.UnarioExpContext) interface{} {
	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	expr := exprResult.(*assembly.ResultadoExpresion)
	operador := ctx.GetOp().GetText()

	// DELEGAR a ExpresionesProcessor
	return v.expresionesProcessor.ProcesarOperacionUnaria(expr, operador)
}

func (v *VisitorARM64) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	nombreVariable := ctx.PatronId().GetText()
	v.armGen.Comment(fmt.Sprintf("Variable no soportada: %s", nombreVariable))

	return &assembly.ResultadoExpresion{
		Registro:  "",
		Tipo:      "int",
		EsLiteral: true,
		Valor:     0,
	}
}

func (v *VisitorARM64) VisitLlamarFuncionExp(ctx *parser.LlamarFuncionExpContext) interface{} {
	return v.Visit(ctx.Llamar_funcion())
}

// ============= LITERALES - DELEGANDO A ExpresionesProcessor =============

func (v *VisitorARM64) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	return v.expresionesProcessor.CrearLiteralInt(ctx.INT_LITERAL().GetText())
}

func (v *VisitorARM64) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	return v.expresionesProcessor.CrearLiteralFloat(ctx.FLOAT_LITERAL().GetText())
}

func (v *VisitorARM64) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return v.expresionesProcessor.CrearLiteralString(ctx.STRING_LITERAL().GetText())
}

func (v *VisitorARM64) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	return v.expresionesProcessor.CrearLiteralBool(ctx.BOOL_LITERAL().GetText())
}

func (v *VisitorARM64) VisitNilLiteral(ctx *parser.NilLiteralContext) interface{} {
	return &assembly.ResultadoExpresion{
		Registro:  "",
		Tipo:      "nil",
		EsLiteral: true,
		Valor:     nil,
	}
}

func (v *VisitorARM64) VisitLiteralExp(ctx *parser.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *VisitorARM64) VisitParentecisExp(ctx *parser.ParentecisExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *VisitorARM64) VisitID_Patron(ctx *parser.ID_PatronContext) interface{} {
	return ctx.GetText()
}

// ============= IMPRESIÓN - LÓGICA ESPECÍFICA DEL VISITOR =============

func (v *VisitorARM64) generarCodigoImpresion(resultado *assembly.ResultadoExpresion, esPrintln bool) {
    v.armGen.Comment(fmt.Sprintf("=== IMPRIMIR %s ===", strings.ToUpper(resultado.Tipo)))

    if resultado.Tipo == "string" {
        // Manejar strings
        if resultado.EsLiteral {
            etiqueta := v.expresionesProcessor.AgregarMensajeString(resultado.Valor.(string))
            regTmp := v.expresionesProcessor.NuevoRegistroTmp()
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
                fmt.Sprintf("mov x0, %s", regTmp))
            v.armGen.LlamarFuncion("print_string") 
        } else {
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("mov x0, %s", resultado.Registro))
            v.armGen.LlamarFuncion("print_string") 
        }
    } else if resultado.Tipo == "int" {
        // Manejar enteros
        if resultado.EsLiteral {
            registroResultado := v.expresionesProcessor.NuevoRegistroTmp()
            v.armGen.Mov(registroResultado, resultado.Valor.(int))
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("mov x0, %s", registroResultado))
            v.armGen.LlamarFuncion("print_int") 
        } else {
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("mov x0, %s", resultado.Registro))
            v.armGen.LlamarFuncion("print_int")
        }
    } else if resultado.Tipo == "float" {
        if resultado.EsLiteral {
            regTmp := v.expresionesProcessor.NuevoRegistroFloatTmp()
            v.generarFloatInmediato(regTmp, resultado.Valor.(float64))
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("fmov d0, %s", regTmp))
            v.armGen.LlamarFuncion("print_float")
        } else {
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("fmov d0, %s", resultado.Registro))
            v.armGen.LlamarFuncion("print_float") 
        }
    } else if resultado.Tipo == "bool" {
        // Manejar booleanos
        if resultado.EsLiteral {
            valor := 0
            if resultado.Valor.(bool) {
                valor = 1
            }
            registroResultado := v.expresionesProcessor.NuevoRegistroTmp()
            v.armGen.Mov(registroResultado, valor)
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("mov x0, %s", registroResultado))
            v.armGen.LlamarFuncion("print_bool") 
        } else {
            v.armGen.Instructions = append(v.armGen.Instructions,
                fmt.Sprintf("mov x0, %s", resultado.Registro))
            v.armGen.LlamarFuncion("print_bool") 
        }
    }

    // Agregar saltos de línea
    if esPrintln {
        v.armGen.LlamarFuncion("print_newline") 
        v.armGen.LlamarFuncion("print_newline") 
    } else {
        v.armGen.LlamarFuncion("print_newline") 
    }

    v.armGen.Instructions = append(v.armGen.Instructions, "")
}

func (v *VisitorARM64) generarFloatInmediato(registro string, valor float64) {
	if valor == 0.0 {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("fmov %s, wzr", registro))
		return
	}

	etiqueta := v.expresionesProcessor.AgregarConstanteFloat(valor)
	regTmp := v.expresionesProcessor.NuevoRegistroTmp()
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
		fmt.Sprintf("ldr %s, [%s]", registro, regTmp))
}
