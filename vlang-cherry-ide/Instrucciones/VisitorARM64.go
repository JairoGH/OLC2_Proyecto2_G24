package instrucciones

import (
	"fmt"
	assembly "main/Assembly"
	"main/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorARM64 struct {
	parser.BaseVGrammarVisitor
	armGen           *assembly.ARMGenerator
	tmpCounter       int
	tmpFloatCounter  int // ← Contador separado para registros float
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
		armGen:          assembly.NewARMGenerator(),
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
	codigoBase := assembly.GenerateCode(v)
	if len(v.MensajesDatos) == 0 {
		return codigoBase
	}

	lines := strings.Split(codigoBase, "\n")
	var out []string
	inserted := false

	for _, line := range lines {
		out = append(out, line)
		if !inserted && strings.Contains(line, ".section .data") {
			out = append(out,
				"    .align 3    // alinea dobles a 8 bytes",
			)
			for _, c := range v.MensajesDatos {
				out = append(out, "    "+c)
			}
			inserted = true
		}
	}

	// Fallback: si no vio .data, lo mete antes de .text
	if !inserted {
		for i, line := range out {
			if strings.Contains(line, ".section .text") {
				prefix := append([]string{}, out[:i]...)
				prefix = append(prefix,
					"    .align 3    // alinea dobles a 8 bytes",
				)
				for _, c := range v.MensajesDatos {
					prefix = append(prefix, "    "+c)
				}
				out = append(prefix, out[i:]...)
				break
			}
		}
	}

	return strings.Join(out, "\n")
}

func (v *VisitorARM64) nuevoRegistroTmp() string {
	if 9+v.tmpCounter > 30 {
		panic("Se agotaron los registros temporales disponibles (x9-x30)")
	}
	reg := fmt.Sprintf("x%d", 9+v.tmpCounter)
	v.tmpCounter++
	return reg
}

// ✅ NUEVO: Registros float (d0-d31)
func (v *VisitorARM64) nuevoRegistroFloatTmp() string {
	if v.tmpFloatCounter > 31 {
		panic("Se agotaron los registros float d0–d31")
	}
	reg := fmt.Sprintf("d%d", v.tmpFloatCounter)
	v.tmpFloatCounter++
	return reg
}

// ✅ MEJORADO: Determinar tipo resultante de operación
func (v *VisitorARM64) determinarTipoResultante(tipo1, tipo2 string) string {
	// Si cualquiera es float, el resultado es float
	if tipo1 == "float" || tipo2 == "float" {
		return "float"
	}
	// Si ambos son int, el resultado es int
	return "int"
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

// En VisitorARM64.go
func (v *VisitorARM64) VisitFuncionMain(ctx *parser.FuncionMainContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		// Resetear contadores para evitar agotar registros
		v.tmpCounter = 0
		v.tmpFloatCounter = 0
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
	operador := ctx.GetOp().GetText()

	// Determinar tipo resultante
	tipoResultante := v.determinarTipoResultante(left.Tipo, right.Tipo)

	// ✅ MANEJAR OPERACIONES POR TIPO
	if tipoResultante == "float" {
		return v.procesarOperacionFloat(left, right, operador)
	} else {
		return v.procesarOperacionInt(left, right, operador)
	}
}

// ✅ NUEVO: Procesar operaciones enteras
func (v *VisitorARM64) procesarOperacionInt(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	registroIzq := v.nuevoRegistroTmp()
	registroDer := v.nuevoRegistroTmp()
	registroResultado := v.nuevoRegistroTmp()

	// Cargar operando izquierdo
	if left.EsLiteral && left.Tipo == "int" {
		v.armGen.Mov(registroIzq, left.Valor.(int))
	} else if left.EsLiteral && left.Tipo == "float" {
		// Convertir float literal a int
		v.armGen.Mov(registroIzq, int(left.Valor.(float64)))
	} else {
		v.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cargar operando derecho
	if right.EsLiteral && right.Tipo == "int" {
		v.armGen.Mov(registroDer, right.Valor.(int))
	} else if right.EsLiteral && right.Tipo == "float" {
		// Convertir float literal a int
		v.armGen.Mov(registroDer, int(right.Valor.(float64)))
	} else {
		v.armGen.MovReg(registroDer, right.Registro)
	}

	// Generar operación
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
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("udiv %s, %s, %s", registroResultado, registroIzq, registroDer),
			fmt.Sprintf("msub %s, %s, %s, %s", registroResultado, registroResultado, registroDer, registroIzq))
	}

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "int",
		EsLiteral: false,
	}
}

// ✅ NUEVO: Procesar operaciones float
func (v *VisitorARM64) procesarOperacionFloat(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	registroIzq := v.nuevoRegistroFloatTmp()
	registroDer := v.nuevoRegistroFloatTmp()
	registroResultado := v.nuevoRegistroFloatTmp()

	// Cargar operando izquierdo
	if left.EsLiteral {
		if left.Tipo == "float" {
			v.cargarConstanteFloat(registroIzq, left.Valor.(float64))
		} else if left.Tipo == "int" {
			// Convertir int a float
			regTmp := v.nuevoRegistroTmp()
			v.armGen.Mov(regTmp, left.Valor.(int))
			v.armGen.ScvtfIntToFloat(registroIzq, regTmp)
		}
	} else {
		if left.Tipo == "int" {
			v.armGen.ScvtfIntToFloat(registroIzq, left.Registro)
		} else {
			v.armGen.FMov(registroIzq, left.Registro)
		}
	}

	// Cargar operando derecho
	if right.EsLiteral {
		if right.Tipo == "float" {
			v.cargarConstanteFloat(registroDer, right.Valor.(float64))
		} else if right.Tipo == "int" {
			regTmp := v.nuevoRegistroTmp()
			v.armGen.Mov(regTmp, right.Valor.(int))
			v.armGen.ScvtfIntToFloat(registroDer, regTmp)
		}
	} else {
		if right.Tipo == "int" {
			v.armGen.ScvtfIntToFloat(registroDer, right.Registro)
		} else {
			v.armGen.FMov(registroDer, right.Registro)
		}
	}

	// Generar operación float
	switch operador {
	case "+":
		v.armGen.FAdd(registroResultado, registroIzq, registroDer)
	case "-":
		v.armGen.FSub(registroResultado, registroIzq, registroDer)
	case "*":
		v.armGen.FMul(registroResultado, registroIzq, registroDer)
	case "/":
		v.armGen.FDiv(registroResultado, registroIzq, registroDer)
	case "%":
		// como antes, sólo un comentario
		v.armGen.Comment("Operador % no soportado directamente para floats")
	}

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "float",
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
	if resultado.Tipo == "string" {
		// Manejar strings
		if resultado.EsLiteral {
			etiqueta := v.agregarMensajeString(resultado.Valor.(string))
			regTmp := v.nuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adrp %s, %s@PAGE", regTmp, etiqueta),
				fmt.Sprintf("add x0, %s, %s@PAGEOFF", regTmp, etiqueta),
				"bl print_string")
		}
	} else if resultado.Tipo == "int" {
		// Manejar enteros
		if resultado.EsLiteral {
			registroResultado := v.nuevoRegistroTmp()
			v.armGen.Mov(registroResultado, resultado.Valor.(int))
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov x0, %s", registroResultado),
				"bl print_int")
		} else {
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov x0, %s", resultado.Registro),
				"bl print_int")
		}
	} else if resultado.Tipo == "float" {
		if resultado.EsLiteral {
			regTmp := v.nuevoRegistroFloatTmp()
			v.generarFloatInmediato(regTmp, resultado.Valor.(float64))
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("fmov d0, %s", regTmp),
				"bl print_float",
			)
		} else {
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("fmov d0, %s", resultado.Registro),
				"bl print_float",
			)
		}

	} else if resultado.Tipo == "bool" {
		// Manejar booleanos
		if resultado.EsLiteral {
			valor := 0
			if resultado.Valor.(bool) {
				valor = 1
			}
			registroResultado := v.nuevoRegistroTmp()
			v.armGen.Mov(registroResultado, valor)
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov x0, %s", registroResultado),
				"bl print_bool")
		} else {
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov x0, %s", resultado.Registro),
				"bl print_bool")
		}
	}

	// Agregar saltos de línea
	if esPrintln {
		v.armGen.Instructions = append(v.armGen.Instructions,
			"bl print_newline",
			"bl print_newline")
	} else {
		v.armGen.Instructions = append(v.armGen.Instructions,
			"bl print_newline")
	}
}

func (v *VisitorARM64) agregarMensajeString(mensaje string) string {
	etiqueta := fmt.Sprintf("msg_%d", v.contadorMensaje)
	v.contadorMensaje++
	v.MensajesDatos = append(v.MensajesDatos,
		fmt.Sprintf("%s: .string \"%s\"", etiqueta, mensaje))
	return etiqueta
}

func (v *VisitorARM64) generarFloatInmediato(registro string, valor float64) {
	// 1) Cero → fmov inmediato válido
	if valor == 0.0 {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("fmov %s, wzr", registro),
		)
		return
	}

	// 2) Literal pequeño → fmov inmediato (opcional)
	if esFloatInmediato(valor) {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("fmov %s, #%.2f", registro, valor),
		)
		return
	}

	// 3) Cualquiera → cargar desde memoria (.double) usando relocations GNU as
	etiqueta := v.agregarConstanteFloat(valor)
	regTmp := v.nuevoRegistroTmp()
	v.armGen.Instructions = append(v.armGen.Instructions,
		// ADRP obtiene los 21 bits altos de la dirección de etiqueta
		fmt.Sprintf("adrp %s, :pg_hi21:%s", regTmp, etiqueta),
		// LDR suma el desplazamiento bajo de 12 bits y carga 64 bits
		fmt.Sprintf("ldr  %s, [%s, :lo12:%s]", registro, regTmp, etiqueta),
	)
}

// 2. Función auxiliar para verificar si un float se puede cargar como inmediato
func esFloatInmediato(_ float64) bool {
	return false
}

// 3. Agregar constantes float a la sección de datos
func (v *VisitorARM64) agregarConstanteFloat(valor float64) string {
	etiqueta := fmt.Sprintf("float_const_%d", v.contadorMensaje)
	v.contadorMensaje++
	v.MensajesDatos = append(v.MensajesDatos,
		fmt.Sprintf("%s: .double %.6f", etiqueta, valor),
	)
	return etiqueta
}

// 4. Cargar constante float desde memoria
func (v *VisitorARM64) cargarConstanteFloat(registro string, valor float64) {
	etiqueta := v.agregarConstanteFloat(valor)
	regTmp := v.nuevoRegistroTmp()

	// Usar adr para cargar la dirección, luego ldr para cargar el valor
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
		fmt.Sprintf("ldr %s, [%s]", registro, regTmp))
}
