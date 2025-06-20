package instrucciones

import (
	"fmt"
	"main/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorARM64 struct {
	parser.BaseVGrammarVisitor
	armGen           *ARMGenerator
	tmpCounter       int
	contadorVar      int
	contadorEtiqueta int
	contadorMensaje  int
	// Para almacenar los mensajes de datos
	MensajesDatos []string
	TablaError    *TablaError
}

type ResultadoExpresion struct {
	Registro  string
	Tipo      string
	EsLiteral bool
	Valor     interface{}
}

func NewVisitorARM64() *VisitorARM64 {
	return &VisitorARM64{
		armGen:          NewARMGenerator(),
		MensajesDatos:   []string{},
		contadorMensaje: 1,
	}
}

func (v *VisitorARM64) nuevoRegistroTmp() string {
	if 9+v.tmpCounter > 30 {
		panic("Se agotaron los registros temporales disponibles (x9-x30)")
	}
	reg := fmt.Sprintf("X%d", 9+v.tmpCounter)
	v.tmpCounter++
	return reg
}

func (v *VisitorARM64) GetCodigo() string {
	codigo := ""

	// Sección de datos
	if len(v.MensajesDatos) > 0 {
		codigo += ".section .data\n"
		for i, mensaje := range v.MensajesDatos {
			codigo += fmt.Sprintf("msg%d: .ascii \"%s\"\n", i+1, mensaje)
		}
		codigo += "\n"
	}

	// Sección de texto
	codigo += ".section .text\n"
	codigo += ".global _start\n\n"
	codigo += "_start:\n\n"

	// Instrucciones del programa
	for _, instr := range v.armGen.Instructions {
		if instr == "" {
			continue
		}
		codigo += "    " + instr + "\n"
	}

	// Salida del programa
	codigo += "\n    // Salida del programa\n"
	codigo += "    MOV X0, #0\n"
	codigo += "    MOV X8, #93\n"
	codigo += "    SVC #0\n"

	return codigo
}

// *** SIGUIENDO EXACTAMENTE LA MISMA ESTRUCTURA DE Visitor.go ***

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
	// Procesar todas las sentencias dentro de main
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

// **IGUAL QUE Visitor.go**: Implementar VisitStmt con la misma estructura
func (v *VisitorARM64) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones (decl_stmt) con sus 7 alternativas
	if ds := ctx.Stmt_declaracion(); ds != nil {
		// Para ARM64, por ahora ignoramos las declaraciones
		return nil
	}

	// 2) Asignaciones
	if ctx.Stmt_asignar() != nil {
		// Por ahora ignorar asignaciones
		return nil
	}

	// 3) Resto de sentencias
	switch {
	case ctx.If_stmt() != nil:
		// Por ahora ignorar if
		return nil
	case ctx.Switch_stmt() != nil:
		// Por ahora ignorar switch
		return nil
	case ctx.While_stmt() != nil:
		// Por ahora ignorar while
		return nil
	case ctx.For_clasico_stmt() != nil:
		// Por ahora ignorar for clásico
		return nil
	case ctx.For_stmt() != nil:
		// Por ahora ignorar for
		return nil
	case ctx.Stmt_transferencia() != nil:
		// Por ahora ignorar transferencia
		return nil
	case ctx.Llamar_funcion() != nil:
		// Procesar llamadas a función
		return v.Visit(ctx.Llamar_funcion())
	case ctx.Declarar_funcion() != nil:
		// Por ahora ignorar declaración de funciones
		return nil
	case ctx.Declarar_struct() != nil:
		// Por ahora ignorar structs
		return nil
	case ctx.Fun_slice() != nil:
		// Por ahora ignorar fun_slice
		return nil
	default:
		// **IGUAL QUE Visitor.go**: No hacer nada, no reportar error por ahora
		return nil
	}

	return nil
}

func (v *VisitorARM64) VisitLlamarFuncion(ctx *parser.LlamarFuncionContext) interface{} {
	nombreFuncion := ctx.PatronId().GetText()

	if nombreFuncion == "print" && ctx.Lista_argumentos() != nil {
		// Procesar argumentos de print
		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			return nil
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok {
			return nil
		}

		for _, arg := range argumentos {
			if resultado, ok := arg.(*ResultadoExpresion); ok {
				// Convertir el resultado a string
				valorStr := v.convertirAString(resultado)

				// Agregar mensaje a la sección de datos
				etiquetaMensaje := fmt.Sprintf("msg%d", v.contadorMensaje)
				v.MensajesDatos = append(v.MensajesDatos, valorStr)

				// Generar código para imprimir
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("// print() -> \"%s\"", valorStr),
					"MOV X0, #1", // stdout
					fmt.Sprintf("LDR X1, =%s", etiquetaMensaje), // dirección del mensaje
					fmt.Sprintf("MOV X2, #%d", len(valorStr)),   // longitud
					"MOV X8, #64", // syscall write
					"SVC #0",
					"", // línea vacía para separar
				)

				v.contadorMensaje++
			}
		}
	}
	return nil
}

func (v *VisitorARM64) convertirAString(resultado *ResultadoExpresion) string {
	if resultado.EsLiteral {
		switch resultado.Tipo {
		case "int":
			return fmt.Sprintf("%d", resultado.Valor.(int))
		case "float":
			return fmt.Sprintf("%.2f", resultado.Valor.(float64))
		case "string":
			return resultado.Valor.(string)
		case "bool":
			if resultado.Valor.(bool) {
				return "true"
			}
			return "false"
		}
	}
	return "unknown"
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

func (v *VisitorARM64) VisitBinarioExp(ctx *parser.BinarioExpContext) interface{} {
	leftResult := v.Visit(ctx.Expr(0))
	rightResult := v.Visit(ctx.Expr(1))

	if leftResult == nil || rightResult == nil {
		return &ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	left := leftResult.(*ResultadoExpresion)
	right := rightResult.(*ResultadoExpresion)

	// Si ambos son literales, calcular directamente
	if left.EsLiteral && right.EsLiteral && left.Tipo == "int" && right.Tipo == "int" {
		leftVal := left.Valor.(int)
		rightVal := right.Valor.(int)
		operador := ctx.GetOp().GetText()

		var resultado int
		switch operador {
		case "+":
			resultado = leftVal + rightVal
		case "-":
			resultado = leftVal - rightVal
		case "*":
			resultado = leftVal * rightVal
		case "/":
			if rightVal != 0 {
				resultado = leftVal / rightVal
			}
		case "%":
			if rightVal != 0 {
				resultado = leftVal % rightVal
			}
		}

		return &ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     resultado,
		}
	}

	// Si no son literales, generar código ARM64 
	registro := v.nuevoRegistroTmp()
	operador := ctx.GetOp().GetText()

	switch operador {
	case "+":
		v.armGen.Add(registro, left.Registro, right.Registro)
	case "-":
		v.armGen.Sub(registro, left.Registro, right.Registro)
	case "*":
		v.armGen.Mul(registro, left.Registro, right.Registro)
	case "/":
		v.armGen.Div(registro, left.Registro, right.Registro)
	case "%":
		tempDiv := v.nuevoRegistroTmp()
		tempMul := v.nuevoRegistroTmp()
		v.armGen.Div(tempDiv, left.Registro, right.Registro)
		v.armGen.Mul(tempMul, tempDiv, right.Registro)
		v.armGen.Sub(registro, left.Registro, tempMul)
	}

	return &ResultadoExpresion{
		Registro:  registro,
		Tipo:      "int",
		EsLiteral: false,
	}
}

// **IGUAL QUE Visitor.go**: Solo métodos específicos para cada literal
func (v *VisitorARM64) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	valor, _ := strconv.Atoi(ctx.INT_LITERAL().GetText())

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "int",
		EsLiteral: true,
		Valor:     valor,
	}
}

func (v *VisitorARM64) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	valor, _ := strconv.ParseFloat(ctx.FLOAT_LITERAL().GetText(), 64)

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "float",
		EsLiteral: true,
		Valor:     valor,
	}
}

func (v *VisitorARM64) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	valor := ctx.STRING_LITERAL().GetText()
	// Remover comillas
	valor = valor[1 : len(valor)-1]

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "string",
		EsLiteral: true,
		Valor:     valor,
	}
}

func (v *VisitorARM64) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	valor, _ := strconv.ParseBool(ctx.BOOL_LITERAL().GetText())

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "bool",
		EsLiteral: true,
		Valor:     valor,
	}
}

func (v *VisitorARM64) VisitNilLiteral(ctx *parser.NilLiteralContext) interface{} {
	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "nil",
		EsLiteral: true,
		Valor:     nil,
	}
}

// **IGUAL QUE Visitor.go**: VisitLiteralExp llama a Visit(ctx.Literal())
func (v *VisitorARM64) VisitLiteralExp(ctx *parser.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *VisitorARM64) VisitParentecisExp(ctx *parser.ParentecisExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *VisitorARM64) VisitID_Patron(ctx *parser.ID_PatronContext) interface{} {
	return ctx.GetText()
}
