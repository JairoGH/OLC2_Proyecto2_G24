package instrucciones

import (
	"fmt"
	assembly "main/Assembly" // Importar el package assembly
	"main/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorARM64 struct {
	parser.BaseVGrammarVisitor
	armGen           *assembly.ARMGenerator // Usar el tipo del package assembly
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
		armGen:          assembly.NewARMGenerator(), // Usar el constructor del package assembly
		MensajesDatos:   []string{},
		contadorMensaje: 1,
	}
}

// Implementar la interfaz requerida por el generador de código
func (v *VisitorARM64) GetARMGenerator() *assembly.ARMGenerator {
	return v.armGen
}

// GetCodigo utiliza el generador de código del package assembly
func (v *VisitorARM64) GetCodigo() string {
	return assembly.GenerateCode(v)
}

func (v *VisitorARM64) nuevoRegistroTmp() string {
	if 9+v.tmpCounter > 30 {
		panic("Se agotaron los registros temporales disponibles (x9-x30)")
	}
	reg := fmt.Sprintf("X%d", 9+v.tmpCounter)
	v.tmpCounter++
	return reg
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
			if resultado, ok := arg.(*ResultadoExpresion); ok {
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

// ✅ MEJORADO: Manejo completo de expresiones binarias con precedencia
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

	// Generar código ARM64 para realizar la operación
	registroIzq := v.nuevoRegistroTmp()
	registroDer := v.nuevoRegistroTmp()
	registroResultado := v.nuevoRegistroTmp()

	// Cargar operando izquierdo
	if left.EsLiteral && left.Tipo == "int" {
		v.armGen.Mov(registroIzq, left.Valor.(int))
	} else {
		v.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cargar operando derecho
	if right.EsLiteral && right.Tipo == "int" {
		v.armGen.Mov(registroDer, right.Valor.(int))
	} else {
		v.armGen.MovReg(registroDer, right.Registro)
	}

	// Generar operación ARM64
	operador := ctx.GetOp().GetText()
	switch operador {
	case "+":
		v.armGen.Add(registroResultado, registroIzq, registroDer)
	case "-":
		v.armGen.Sub(registroResultado, registroIzq, registroDer)
	case "*":
		v.armGen.Mul(registroResultado, registroIzq, registroDer)
	case "/":
		v.armGen.Div(registroResultado, registroIzq, registroDer)
	case "%":
		// Módulo: result = left - (left/right) * right
		tempDiv := v.nuevoRegistroTmp()
		tempMul := v.nuevoRegistroTmp()
		v.armGen.Div(tempDiv, registroIzq, registroDer)
		v.armGen.Mul(tempMul, tempDiv, registroDer)
		v.armGen.Sub(registroResultado, registroIzq, tempMul)
	default:
		v.armGen.Comment(fmt.Sprintf("Operador no soportado: %s", operador))
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      "int",
			EsLiteral: false,
		}
	}

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "int",
		EsLiteral: false,
	}
}

// ✅ NUEVO: Soporte para expresiones unarias
func (v *VisitorARM64) VisitUnarioExp(ctx *parser.UnarioExpContext) interface{} {
	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		return &ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	expr := exprResult.(*ResultadoExpresion)
	operador := ctx.GetOp().GetText()

	switch operador {
	case "-":
		// Negación: -expr
		registroOrigen := v.nuevoRegistroTmp()
		registroResultado := v.nuevoRegistroTmp()

		if expr.EsLiteral && expr.Tipo == "int" {
			v.armGen.Mov(registroOrigen, expr.Valor.(int))
		} else {
			v.armGen.MovReg(registroOrigen, expr.Registro)
		}

		// Negar el valor: 0 - valor
		v.armGen.Mov(registroResultado, 0)
		v.armGen.Sub(registroResultado, registroResultado, registroOrigen)

		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      "int",
			EsLiteral: false,
		}

	case "!":
		// Negación lógica (para futuro uso con booleanos)
		v.armGen.Comment("Operador ! no implementado para ARM64")
		return expr

	default:
		v.armGen.Comment(fmt.Sprintf("Operador unario no soportado: %s", operador))
		return expr
	}
}

// ✅ NUEVO: Soporte para identificadores (variables)
func (v *VisitorARM64) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	// Por ahora, las variables no están implementadas en ARM64
	// Retornar un valor por defecto
	nombreVariable := ctx.PatronId().GetText()
	v.armGen.Comment(fmt.Sprintf("Variable no soportada: %s", nombreVariable))

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "int",
		EsLiteral: true,
		Valor:     0,
	}
}

// ✅ NUEVO: Soporte para llamadas a función en expresiones
func (v *VisitorARM64) VisitLlamarFuncionExp(ctx *parser.LlamarFuncionExpContext) interface{} {
	// Por ahora, solo manejar funciones básicas
	return v.Visit(ctx.Llamar_funcion())
}

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

func (v *VisitorARM64) VisitLiteralExp(ctx *parser.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *VisitorARM64) VisitParentecisExp(ctx *parser.ParentecisExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *VisitorARM64) VisitID_Patron(ctx *parser.ID_PatronContext) interface{} {
	return ctx.GetText()
}

// Generar código ARM64 para imprimir el resultado (usando funciones auxiliares)
func (v *VisitorARM64) generarCodigoImpresion(resultado *ResultadoExpresion, esPrintln bool) {
    var registroResultado string

    if resultado.EsLiteral {
        // Si es literal, cargar directamente el valor
        registroResultado = v.nuevoRegistroTmp()
        switch resultado.Tipo {
        case "int":
            valor := resultado.Valor.(int)
            v.armGen.Mov(registroResultado, valor)
        case "float":
            valor := int(resultado.Valor.(float64))
            v.armGen.Mov(registroResultado, valor)
        }
    } else {
        // Si no es literal, ya tenemos el registro con el resultado
        registroResultado = resultado.Registro
    }

    // Usar las funciones auxiliares para imprimir el número
    v.armGen.Instructions = append(v.armGen.Instructions,
        "// Imprimir resultado",
        fmt.Sprintf("mov x0, %s", strings.ToLower(registroResultado)),
        "bl print_int")

    if esPrintln {
        // println(): Agregar DOS saltos de línea (línea en blanco)
        v.armGen.Instructions = append(v.armGen.Instructions,
            "// println: Agregar dos saltos de línea",
            "bl print_newline",  // Primer salto de línea
            "bl print_newline")  // Segundo salto de línea (línea en blanco)
    } else {
        // print(): Agregar UN salto de línea
        v.armGen.Instructions = append(v.armGen.Instructions,
            "// print: Agregar un salto de línea",
            "bl print_newline")  // Solo un salto de línea
    }

    v.armGen.Instructions = append(v.armGen.Instructions, "")
}
