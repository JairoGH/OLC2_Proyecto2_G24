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
	variablesProcessor   *assembly.VariablesProcessor
	sliceProcessor       *assembly.SliceProcessor
	TablaError           *TablaError
	pilaControl          []*LoopControlItem
}

func NewVisitorARM64() *VisitorARM64 {
	armGen := assembly.NewARMGenerator()
	expresionesProcessor := assembly.NewExpresionesProcessor(armGen)
	return &VisitorARM64{
		armGen:               armGen,
		expresionesProcessor: expresionesProcessor,
		variablesProcessor:   assembly.NewVariablesProcessor(armGen, expresionesProcessor),
		sliceProcessor:       assembly.NewSliceProcessor(armGen, expresionesProcessor),
		TablaError:           NewTablaError(),
	}
}

// LoopControlItem simula el comportamiento de LlamadaFunciones del intérprete
type LoopControlItem struct {
	TipoLoop         string   // "for_clasico", "while", "switch", "for_range"
	Acciones         []string // ["break", "continue"] o solo ["break"] para switch
	EtiquetaSalida   string   // Para break
	EtiquetaContinue string   // Para continue
	AccionActual     string   // La acción que se está ejecutando
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

func (v *VisitorARM64) VisitStmt_asignar(ctx *parser.Stmt_asignarContext) interface{} {
	if ctx.GetChildCount() > 0 {
		if child, ok := ctx.GetChild(0).(antlr.ParseTree); ok {
			switch childTyped := child.(type) {
			case *parser.AsignacionDirectaContext:
				return v.Visit(childTyped)
			case *parser.AsignacionAritmeticaContext:
				return v.Visit(childTyped)
			case *parser.AsignacionSliceContext:
				return v.Visit(childTyped)
			case *parser.AsignacionSliceItemContext:
				return v.Visit(childTyped)
			case *parser.DeclararSliceContext:
				return v.VisitDeclararSlice(childTyped) 
			default:
				return v.Visit(child) // Fallback
			}
		}
	}
	return nil
}

func (v *VisitorARM64) VisitStmt_declaracion(ctx *parser.Stmt_declaracionContext) interface{} {
	if ctx.GetChildCount() > 0 {
		if child, ok := ctx.GetChild(0).(antlr.ParseTree); ok {
			return v.Visit(child)
		}
	}
	return nil
}

// VisitAsignacionDirecta
func (v *VisitorARM64) VisitAsignacionDirecta(ctx *parser.AsignacionDirectaContext) interface{} {
	nombreVar := ctx.PatronId().GetText()
	exprResult := v.Visit(ctx.Expr())

	if exprResult == nil {
		v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para %s", nombreVar))
		return nil
	}

	valor := exprResult.(*assembly.ResultadoExpresion)

	// Manejar asignación de slice (resultado de append)
	if valor.Tipo == "slice_name" {
		// Es resultado de append u otra operación de slice
		nombreSliceResultado := valor.Valor.(string)

		// Si es el mismo slice (autoasignación como append), no hacer nada extra
		if nombreVar == nombreSliceResultado {
			return nil
		} else {
			v.armGen.Comment(fmt.Sprintf("Asignación entre slices diferentes: %s = %s", nombreVar, nombreSliceResultado))
			return nil
		}
	}

	// Asignación normal de variables
	err := v.variablesProcessor.AsignarVariable(nombreVar, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitAsignacionAritmetica(ctx *parser.AsignacionAritmeticaContext) interface{} {
	// CORRECCIÓN: Obtener el nombre directamente del contexto
	nombreVariable := ctx.PatronId().GetText()

	v.armGen.Comment(fmt.Sprintf("=== ASIGNACIÓN ARITMÉTICA: %s %s ===", nombreVariable, ctx.GetOp().GetText()))

	// Buscar la variable
	variable, err := v.variablesProcessor.CargarVariable(nombreVariable, "")
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: Variable '%s' no encontrada", nombreVariable))
		return nil
	}

	valorIzq := variable
	valorDcha := v.Visit(ctx.Expr())

	if valorDcha == nil {
		v.armGen.Comment("Error: expresión del lado derecho es inválida")
		return nil
	}

	valorDerecho, ok := valorDcha.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment("Error: el valor derecho no es una expresión válida")
		return nil
	}

	// Obtener el operador base (quitar el "=")
	operadorCompleto := ctx.GetOp().GetText()   // "+=", "-="
	operadorBase := string(operadorCompleto[0]) // "+", "-"

	// Realizar la operación aritmética usando ExpresionesProcessor
	resultado := v.expresionesProcessor.ProcesarOperacionBinaria(valorIzq, valorDerecho, operadorBase)
	if resultado == nil {
		v.armGen.Comment(fmt.Sprintf("Error en operación %s", operadorBase))
		return nil
	}

	// Asignar el resultado de vuelta a la variable
	err = v.variablesProcessor.AsignarVariable(nombreVariable, resultado)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error asignando resultado a %s: %s", nombreVariable, err.Error()))
		return nil
	}

	v.armGen.Comment(fmt.Sprintf("=== FIN ASIGNACIÓN ARITMÉTICA: %s %s === \n", nombreVariable, operadorCompleto))
	return nil
}

func (v *VisitorARM64) VisitAsignacionSlice(ctx *parser.AsignacionSliceContext) interface{} {
	v.armGen.Comment("Asignación a slice no implementada aún")
	return nil
}

// Método auxiliar para obtener tipo como string
func (v *VisitorARM64) obtenerTipoString(tipoCtx parser.ITipoContext) string {
	if tipoCtx.RW_INT() != nil {
		return "int"
	} else if tipoCtx.RW_FLOAT64() != nil {
		return "float"
	} else if tipoCtx.RW_STRING() != nil {
		return "string"
	} else if tipoCtx.RW_BOOL() != nil {
		return "bool"
	} else if tipoCtx.ID() != nil {
		return tipoCtx.ID().GetText() // Para structs
	} else if tipoCtx.Tipos_slices() != nil {
		//MANEJAR TIPOS DE SLICE
		return v.Visit(tipoCtx.Tipos_slices()).(string)
	}
	return "unknown"
}

func (v *VisitorARM64) VisitDeclararSinMutValor(ctx *parser.DeclararSinMutValorContext) interface{} {
	nombre := ctx.ID().GetText()
	tipo := v.obtenerTipoString(ctx.Tipo())
	exprResult := v.Visit(ctx.Expr())
	valor := exprResult.(*assembly.ResultadoExpresion)

	v.armGen.Comment(fmt.Sprintf("=== DECLARAR VARIABLE : %s ===", nombre))

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararTipo(ctx *parser.DeclararTipoContext) interface{} {
	nombre := ctx.ID().GetText()
	tipo := v.obtenerTipoString(ctx.Tipo())

    v.armGen.Comment(fmt.Sprintf("=== DECLARAR VARIABLE MUT: %s ===", nombre))

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, nil)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclaraTipoValor(ctx *parser.DeclaraTipoValorContext) interface{} {
	nombre := ctx.ID().GetText()
	tipo := v.obtenerTipoString(ctx.Tipo())

	v.armGen.Comment(fmt.Sprintf("=== DECLARAR VARIABLE MUT: %s ===", nombre))

	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para %s", nombre))
		return nil
	}

	// CONVERSIÓN SEGURA
	valor, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment(fmt.Sprintf("Error: tipo de expresión inválido para %s", nombre))
		return nil
	}

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararInferencia(ctx *parser.DeclararInferenciaContext) interface{} {
	nombre := ctx.ID().GetText()
	v.armGen.Comment(fmt.Sprintf("=== DECLARAR VARIABLE: %s ===", nombre))

	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para %s", nombre))
		return nil
	}

	// CONVERSIÓN SEGURA
	valor, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment(fmt.Sprintf("Error: tipo de expresión inválido para %s", nombre))
		return nil
	}

	tipo := valor.Tipo
	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararInferenciaMut(ctx *parser.DeclararInferenciaMutContext) interface{} {
	nombre := ctx.ID().GetText()
	v.armGen.Comment(fmt.Sprintf("=== DECLARAR VARIABLE MUT: %s ===", nombre))

	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para %s", nombre))
		return nil
	}

	// CONVERSIÓN SEGURA
	valor, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment(fmt.Sprintf("Error: tipo de expresión inválido para %s", nombre))
		return nil
	}

	// VALIDAR QUE valor NO SEA NIL
	if valor == nil {
		v.armGen.Comment(fmt.Sprintf("Error: valor es nil para %s", nombre))
		return nil
	}

	tipo := valor.Tipo
	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	} 
	return nil
}

func (v *VisitorARM64) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones
	if ds := ctx.Stmt_declaracion(); ds != nil {
		return v.Visit(ds)
	}

	// 2) Asignaciones
	if ctx.Stmt_asignar() != nil {
		return v.Visit(ctx.Stmt_asignar())
	}

	// AGREGAR CASE FALTANTE PARA IF
	if ctx.If_stmt() != nil {
		return v.Visit(ctx.If_stmt())
	}

	// 3) Resto de sentencias
	switch {
	case ctx.Switch_stmt() != nil:
		return v.Visit(ctx.Switch_stmt())
	case ctx.While_stmt() != nil:
		return v.Visit(ctx.While_stmt())
	case ctx.For_clasico_stmt() != nil:
		return v.Visit(ctx.For_clasico_stmt())
	case ctx.For_stmt() != nil:
		return v.Visit(ctx.For_stmt())
	case ctx.Stmt_transferencia() != nil:
		return v.Visit(ctx.Stmt_transferencia())
	case ctx.Llamar_funcion() != nil:
		return v.Visit(ctx.Llamar_funcion())
	case ctx.Declarar_funcion() != nil:
		return v.Visit(ctx.Declarar_funcion())
	case ctx.Declarar_struct() != nil:
		return v.Visit(ctx.Declarar_struct())
	case ctx.Fun_slice() != nil:
		return v.Visit(ctx.Fun_slice())
	default:
		return nil
	}
}

// VisitLlamarFuncion
func (v *VisitorARM64) VisitLlamarFuncion(ctx *parser.LlamarFuncionContext) interface{} {
	nombreFuncion := ctx.PatronId().GetText()

	//  Manejar función append
	if nombreFuncion == "append" && ctx.Lista_argumentos() != nil {

		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			v.armGen.Comment("Error: argumentos inválidos para append")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "slice_name",
				EsLiteral: true,
				Valor:     "error",
			}
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok || len(argumentos) < 2 {
			v.armGen.Comment("Error: append requiere al menos 2 argumentos (slice, elemento)")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "slice_name",
				EsLiteral: true,
				Valor:     "error",
			}
		}

		// Primer argumento: slice
		primerArg, ok1 := argumentos[0].(*assembly.ResultadoExpresion)
		if !ok1 {
			v.armGen.Comment("Error: primer argumento de append no es válido")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "slice_name",
				EsLiteral: true,
				Valor:     "error",
			}
		}

		// Verificar que el primer argumento sea un slice
		if primerArg.Tipo != "slice_name" {
			v.armGen.Comment("Error: primer argumento de append debe ser un slice")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "slice_name",
				EsLiteral: true,
				Valor:     "error",
			}
		}

		nombreSlice := primerArg.Valor.(string)

		// Si solo hay un elemento adicional
		if len(argumentos) == 2 {
			segundoArg, ok2 := argumentos[1].(*assembly.ResultadoExpresion)
			if !ok2 {
				v.armGen.Comment("Error: segundo argumento de append no es válido")
				regError := v.expresionesProcessor.NuevoRegistroTmp()
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("mov %s, #0", regError))
				return &assembly.ResultadoExpresion{
					Registro:  regError,
					Tipo:      "slice_name",
					EsLiteral: true,
					Valor:     "error",
				}
			}

			// Llamar a la función append simple
			resultado, err := v.sliceProcessor.Append(nombreSlice, segundoArg)
			if err != nil {
				v.armGen.Comment(fmt.Sprintf("Error en append: %s", err.Error()))
				regError := v.expresionesProcessor.NuevoRegistroTmp()
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("mov %s, #0", regError))
				return &assembly.ResultadoExpresion{
					Registro:  regError,
					Tipo:      "slice_name",
					EsLiteral: true,
					Valor:     "error",
				}
			}

			return resultado
		} 
	}
	// Manejar función len
	if nombreFuncion == "len" && ctx.Lista_argumentos() != nil {

		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			v.armGen.Comment("Error: argumentos inválidos para len")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "int",
				EsLiteral: false,
			}
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok || len(argumentos) != 1 {
			v.armGen.Comment("Error: len requiere exactamente 1 argumento")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "int",
				EsLiteral: false,
			}
		}

		// El argumento debe ser un slice
		argumento, ok := argumentos[0].(*assembly.ResultadoExpresion)
		if !ok {
			v.armGen.Comment("Error: argumento de len no es válido")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "int",
				EsLiteral: false,
			}
		}

		// Verificar que el argumento sea un slice
		if argumento.Tipo != "slice_name" {
			v.armGen.Comment("Error: len() solo funciona con slices")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "int",
				EsLiteral: false,
			}
		}

		nombreSlice := argumento.Valor.(string)

		// Llamar a la función len del SliceProcessor
		resultado, err := v.sliceProcessor.Len(nombreSlice)
		if err != nil {
			v.armGen.Comment(fmt.Sprintf("Error en len: %s", err.Error()))
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", regError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "int",
				EsLiteral: false,
			}
		}

		return resultado
	}
	// Manejar función join
	if nombreFuncion == "join" && ctx.Lista_argumentos() != nil {

		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			v.armGen.Comment("Error: argumentos inválidos para join")
			etiquetaError := v.expresionesProcessor.AgregarMensajeString("")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regError, etiquetaError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "string",
				EsLiteral: false,
			}
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok || len(argumentos) != 2 {
			v.armGen.Comment("Error: join requiere exactamente 2 argumentos")
			etiquetaError := v.expresionesProcessor.AgregarMensajeString("")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regError, etiquetaError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "string",
				EsLiteral: false,
			}
		}

		// Primer argumento: slice de strings
		primerArg, ok1 := argumentos[0].(*assembly.ResultadoExpresion)
		segundoArg, ok2 := argumentos[1].(*assembly.ResultadoExpresion)

		if !ok1 || !ok2 {
			v.armGen.Comment("Error: argumentos de join no son válidos")
			etiquetaError := v.expresionesProcessor.AgregarMensajeString("")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regError, etiquetaError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "string",
				EsLiteral: false,
			}
		}

		// Verificar que el primer argumento sea un slice
		if primerArg.Tipo != "slice_name" {
			v.armGen.Comment("Error: primer argumento de join debe ser un slice")
			etiquetaError := v.expresionesProcessor.AgregarMensajeString("")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regError, etiquetaError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "string",
				EsLiteral: false,
			}
		}

		nombreSlice := primerArg.Valor.(string)
		separador := segundoArg

		// Llamar a la función join del SliceProcessor
		resultado, err := v.sliceProcessor.Join(nombreSlice, separador)
		if err != nil {
			v.armGen.Comment(fmt.Sprintf("Error en join: %s", err.Error()))
			etiquetaError := v.expresionesProcessor.AgregarMensajeString("")
			regError := v.expresionesProcessor.NuevoRegistroTmp()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regError, etiquetaError))
			return &assembly.ResultadoExpresion{
				Registro:  regError,
				Tipo:      "string",
				EsLiteral: false,
			}
		}

		return resultado
	}

	// MANEJAR indexOf
	if nombreFuncion == "indexOf" && ctx.Lista_argumentos() != nil {
		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			return &assembly.ResultadoExpresion{
				Registro:  "",
				Tipo:      "int",
				EsLiteral: true,
				Valor:     -1,
			}
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok || len(argumentos) != 2 {
			return &assembly.ResultadoExpresion{
				Registro:  "",
				Tipo:      "int",
				EsLiteral: true,
				Valor:     -1,
			}
		}

		primerArg, ok1 := argumentos[0].(*assembly.ResultadoExpresion)
		segundoArg, ok2 := argumentos[1].(*assembly.ResultadoExpresion)

		if !ok1 || !ok2 || primerArg.Tipo != "slice_name" {
			return &assembly.ResultadoExpresion{
				Registro:  "",
				Tipo:      "int",
				EsLiteral: true,
				Valor:     -1,
			}
		}

		nombreSlice := primerArg.Valor.(string)
		valorBuscado := segundoArg

		resultado, err := v.sliceProcessor.IndexOf(nombreSlice, valorBuscado)
		if err != nil {
			return &assembly.ResultadoExpresion{
				Registro:  "",
				Tipo:      "int",
				EsLiteral: true,
				Valor:     -1,
			}
		}

		return resultado
	}

	// MANEJAR print/println
	if (nombreFuncion == "print" || nombreFuncion == "println") && ctx.Lista_argumentos() != nil {
		argumentosResult := v.Visit(ctx.Lista_argumentos())
		if argumentosResult == nil {
			return nil
		}

		argumentos, ok := argumentosResult.([]interface{})
		if !ok {
			return nil
		}

		// Iterar argumentos SIN saltos de línea entre ellos
		for i, arg := range argumentos {
			if resultado, ok := arg.(*assembly.ResultadoExpresion); ok {
				// NO generar salto de línea para elementos intermedios
				v.generarCodigoImpresion(resultado)

				// SEPARAR ELEMENTOS CON ESPACIO (excepto el último)
				if i < len(argumentos)-1 {
					v.armGen.UsarFuncion("print_char")
					v.armGen.Instructions = append(v.armGen.Instructions,
						"mov x0, #32          // ' ' (espacio)",
						"bl print_char")
				}
			}
		}

		//print() vs println()
		if nombreFuncion == "println" {
			v.armGen.LlamarFuncion("print_newline")
			v.armGen.LlamarFuncion("print_newline")
		} else {
			// print() agrega solo UN salto de línea
			v.armGen.LlamarFuncion("print_newline")
		}

		return nil
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

func (v *VisitorARM64) VisitStmt_transferencia(ctx *parser.Stmt_transferenciaContext) interface{} {
	v.armGen.Comment("Transferencia de control no implementada aún")
	return nil
}

func (v *VisitorARM64) VisitFor_clasico_stmt(ctx *parser.For_clasico_stmtContext) interface{} {
	v.armGen.Comment("For clásico no implementado aún")
	return nil
}

func (v *VisitorARM64) VisitDeclarar_funcion(ctx *parser.Declarar_funcionContext) interface{} {
	v.armGen.Comment("Declaración de función no implementada aún")
	return nil
}

func (v *VisitorARM64) VisitDeclarar_struct(ctx *parser.Declarar_structContext) interface{} {
	v.armGen.Comment("Declaración de struct no implementada aún")
	return nil
}

func (v *VisitorARM64) VisitFun_slice(ctx *parser.Fun_sliceContext) interface{} {
	v.armGen.Comment("Función de slice no implementada aún")
	return nil
}

// ============= EXPRESIONES - DELEGANDO A ExpresionesProcessor =============

func (v *VisitorARM64) VisitBinarioExp(ctx *parser.BinarioExpContext) interface{} {
	operador := ctx.GetOp().GetText()
	v.armGen.Comment(fmt.Sprintf("===  EXPRESIÓN BINARIA %s ===", operador))

	// RESETEO OPCIONAL PARA OPERADORES LÓGICOS COMPLEJOS
	if operador == "&&" || operador == "||" {
		v.expresionesProcessor.ResetearRegistrosSiNecesario()
	}

	leftResult := v.Visit(ctx.Expr(0))

	if leftResult == nil {
		fmt.Printf("❌ ERROR: Operando izquierdo es nil\n")
		v.armGen.Comment("ERROR: operando izquierdo es nil")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	rightResult := v.Visit(ctx.Expr(1))

	if rightResult == nil {
		fmt.Printf("❌ ERROR: Operando derecho es nil\n")
		v.armGen.Comment("ERROR: operando derecho es nil")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	left, okLeft := leftResult.(*assembly.ResultadoExpresion)
	right, okRight := rightResult.(*assembly.ResultadoExpresion)

	if !okLeft {
		fmt.Printf("❌ ERROR: Operando izquierdo no es ResultadoExpresion, es: %T\n", leftResult)
		v.armGen.Comment("ERROR: operando izquierdo no es ResultadoExpresion")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	if !okRight {
		fmt.Printf("❌ ERROR: Operando derecho no es ResultadoExpresion, es: %T\n", rightResult)
		v.armGen.Comment("ERROR: operando derecho no es ResultadoExpresion")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	// DELEGAR a ExpresionesProcessor CON PROTECCIÓN
	var resultado *assembly.ResultadoExpresion

	// PROTECCIÓN CONTRA PANIC EN PROCESAMIENTO
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("❌ PANIC en ProcesarOperacionBinaria: %v\n", r)
				resultado = &assembly.ResultadoExpresion{
					Registro:  "",
					Tipo:      "int",
					EsLiteral: true,
					Valor:     0,
				}
			}
		}()
		resultado = v.expresionesProcessor.ProcesarOperacionBinaria(left, right, operador)
	}()

	if resultado == nil {
		fmt.Printf("❌ ERROR: ProcesarOperacionBinaria retornó nil\n")
		resultado = &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	v.armGen.Comment("=== FIN EXPRESIÓN BINARIA === \n")

	return resultado
}

func (v *VisitorARM64) VisitUnarioExp(ctx *parser.UnarioExpContext) interface{} {
	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment("Error: expresión unaria es nil")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	// CONVERSIÓN SEGURA
	expr, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment("Error: expresión unaria no es ResultadoExpresion")
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	operador := ctx.GetOp().GetText()

	// DELEGAR a ExpresionesProcessor
	resultado := v.expresionesProcessor.ProcesarOperacionUnaria(expr, operador)
	return resultado
}

// VisitIdExp  correctamente
func (v *VisitorARM64) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	nombreVariable := ctx.PatronId().GetText()

	// Protección contra panic
	var resultado *assembly.ResultadoExpresion
	var err error

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("❌ PANIC en CargarVariable(%s): %v\n", nombreVariable, r)
				err = fmt.Errorf("panic al cargar variable: %v", r)
			}
		}()
		resultado, err = v.variablesProcessor.CargarVariable(nombreVariable, "")
	}()

	if err != nil {
		fmt.Printf("❌ ERROR: %s\n", err.Error())
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	if resultado == nil {
		fmt.Printf("❌ ERROR: CargarVariable retornó nil para %s\n", nombreVariable)
		v.armGen.Comment(fmt.Sprintf("Error: CargarVariable retornó nil para %s", nombreVariable))
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	return resultado
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
	nombreVariable := ctx.GetText()

	// CASO 1: Si se usa en contexto de string (como PatronId en asignaciones)
	// En algunos contextos, necesitamos retornar solo el nombre

	// CASO 2: Verificar si es un slice
	if v.sliceProcessor.ExisteSlice(nombreVariable) {
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "slice_name",
			EsLiteral: true,
			Valor:     nombreVariable,
		}
	}

	// CASO 3: Intentar cargar variable desde el stack
	resultado, err := v.variablesProcessor.CargarVariable(nombreVariable, "")
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "int",
			EsLiteral: true,
			Valor:     0,
		}
	}

	return resultado
}

// ============= DECLARACIÓN DE SLICE - LÓGICA ESPECÍFICA DEL VISITOR =============
func (v *VisitorARM64) VisitDeclararSlice(ctx *parser.DeclararSliceContext) interface{} {
	// Obtener nombre del slice
	nombreSlice := ctx.ID().GetText()

	// Obtener tipo del slice
	var tipoSlice string
	if ctx.Tipo() != nil {
		tipoCompleto := v.obtenerTipoString(ctx.Tipo())

		if strings.HasPrefix(tipoCompleto, "[]") {
			tipoSlice = tipoCompleto[2:] // Eliminar [] para obtener el tipo base
		} else {
			tipoSlice = tipoCompleto
		}
	} else {
		tipoSlice = "int" // Valor por defecto
	}

	// Verificar si hay expresión (valor inicial)
	var elementos []*assembly.ResultadoExpresion
	var tipoElemento string

	if ctx.Expr() != nil {

		// Evaluar expresión (debería ser una lista de slice)
		exprResult := v.Visit(ctx.Expr())
		if exprResult == nil {
			v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para slice %s", nombreSlice))
			return nil
		}

		// Si es una lista de slice
		if listaElementos, ok := exprResult.([]interface{}); ok {

			var err error
			elementos, tipoElemento, err = v.sliceProcessor.ParsearListaSlice(listaElementos)
			if err != nil {
				v.armGen.Comment(fmt.Sprintf("Error parseando lista: %s", err.Error()))
				return nil
			}

			//  CORREGIR: Verificación de tipos más flexible
			if tipoElemento != "" && tipoElemento != tipoSlice {
				v.armGen.Comment(fmt.Sprintf("Warning: tipo de elementos (%s) no coincide exactamente con el tipo declarado (%s), pero continuando",
					tipoElemento, tipoSlice))
				//  NO RETORNAR ERROR, solo advertencia
				// return nil
			}
		} else {
			v.armGen.Comment("La expresión no es una lista de elementos")
		}
	} else {
		v.armGen.Comment("No hay expresión inicial - slice vacío")
	}

	// Si no se especificó tipo de elementos, usar el tipo declarado
	if tipoElemento == "" {
		tipoElemento = tipoSlice
	}

	// Declarar el slice
	err := v.sliceProcessor.DeclararSlice(nombreSlice, tipoElemento, elementos)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error declarando slice: %s", err.Error()))
		return nil
	}

	return nil
}

// IMPLEMENTAR VisitListaSlice
func (v *VisitorARM64) VisitListaSlice(ctx *parser.ListaSliceContext) interface{} {
	var elementos []interface{}

	// Procesar cada expresión en la lista
	for i, exprCtx := range ctx.AllExpr() {
		resultado := v.Visit(exprCtx)
		if resultado != nil {
			if _, ok := resultado.(*assembly.ResultadoExpresion); ok {
			}
			elementos = append(elementos, resultado)
		} else {
			v.armGen.Comment(fmt.Sprintf("Elemento %d resultó nil", i))
		}
	}

	return elementos
}

// IMPLEMENTAR VisitSliceExp
func (v *VisitorARM64) VisitSliceExp(ctx *parser.SliceExpContext) interface{} {
	resultado := v.Visit(ctx.Lista_slice())
	return resultado
}

// VisitTipos_slices
func (v *VisitorARM64) VisitTipos_slices(ctx *parser.Tipos_slicesContext) interface{} {
	return ctx.GetText() // Retorna "[]int", "[]string", etc.
}

// VisitVectorSimple
func (v *VisitorARM64) VisitVectorSimple(ctx *parser.VectorSimpleContext) interface{} {
	return ctx.GetText() // Retorna "[]int", "[]string", etc.
}

// VisitMatrizDoble (por completitud)
func (v *VisitorARM64) VisitMatrizDoble(ctx *parser.MatrizDobleContext) interface{} {
	return ctx.GetText() // Retorna "[][]int", "[][]string", etc.
}

func (v *VisitorARM64) VisitItemSliceExp(ctx *parser.ItemSliceExpContext) interface{} {

	// Delegar a Item_slice
	resultado := v.Visit(ctx.Item_slice())

	// El resultado ya es un ResultadoExpresion con el valor del elemento
	return resultado
}

// VisitItemSlice para acceso por índice
func (v *VisitorARM64) VisitItemSlice(ctx *parser.ItemSliceContext) interface{} {
	nombreSlice := ctx.PatronId().GetText()

	// Verificar que existe el slice
	if !v.sliceProcessor.ExisteSlice(nombreSlice) {
		v.armGen.Comment(fmt.Sprintf("Error: slice '%s' no encontrado", nombreSlice))
		regError := v.expresionesProcessor.NuevoRegistroTmp()
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("mov %s, #0", regError),
			fmt.Sprintf("sub %s, %s, #1", regError, regError)) // -1 para error
		return &assembly.ResultadoExpresion{
			Registro:  regError,
			Tipo:      "int",
			EsLiteral: false,
		}
	}

	// Solo manejar acceso simple: slice[indice]
	if len(ctx.AllExpr()) != 1 {
		v.armGen.Comment("Error: acceso multidimensional no soportado en slices")
		regError := v.expresionesProcessor.NuevoRegistroTmp()
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("mov %s, #0", regError),
			fmt.Sprintf("sub %s, %s, #1", regError, regError))
		return &assembly.ResultadoExpresion{
			Registro:  regError,
			Tipo:      "int",
			EsLiteral: false,
		}
	}

	// Evaluar el índice
	indiceResult := v.Visit(ctx.Expr(0))
	if indiceResult == nil {
		v.armGen.Comment("Error: índice inválido")
		regError := v.expresionesProcessor.NuevoRegistroTmp()
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("mov %s, #0", regError),
			fmt.Sprintf("sub %s, %s, #1", regError, regError))
		return &assembly.ResultadoExpresion{
			Registro:  regError,
			Tipo:      "int",
			EsLiteral: false,
		}
	}

	indiceExpr := indiceResult.(*assembly.ResultadoExpresion)

	// CORREGIR: Usar el tipo correcto de assembly.SliceAccessInfo
	sliceAccessInfo := &assembly.SliceAccessInfo{
		NombreSlice: nombreSlice,
		IndiceExpr:  indiceExpr,
	}

	// Llamar a la función de acceso por índice
	resultado, err := v.sliceProcessor.AccederElementoPorIndice(nombreSlice, indiceExpr)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error en acceso por índice: %s", err.Error()))
		regError := v.expresionesProcessor.NuevoRegistroTmp()
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("mov %s, #0", regError),
			fmt.Sprintf("sub %s, %s, #1", regError, regError))
		return &assembly.ResultadoExpresion{
			Registro:  regError,
			Tipo:      "int",
			EsLiteral: false,
		}
	}

	// AGREGAR INFORMACIÓN PARA ASIGNACIONES
	resultado.SliceInfo = sliceAccessInfo

	return resultado
}

// VisitAsignacionSliceItem para asignaciones por índice
func (v *VisitorARM64) VisitAsignacionSliceItem(ctx *parser.AsignacionSliceItemContext) interface{} {

	// Obtener la referencia al elemento del slice
	refItemSlice := v.Visit(ctx.Item_slice())
	if refItemSlice == nil {
		v.armGen.Comment("Error: referencia de slice inválida")
		return nil
	}

	resultado, ok := refItemSlice.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment("Error: tipo de referencia no válido")
		return nil
	}

	// VERIFICAR SI HAY INFORMACIÓN DE SLICE DISPONIBLE
	if resultado.SliceInfo == nil {
		v.armGen.Comment("Error: información de slice no disponible para asignación")
		return nil
	}

	// EVALUAR EL NUEVO VALOR A ASIGNAR
	nuevoValorResult := v.Visit(ctx.Expr())
	if nuevoValorResult == nil {
		v.armGen.Comment("Error: valor inválido para asignación")
		return nil
	}

	nuevoValor := nuevoValorResult.(*assembly.ResultadoExpresion)

	// USAR LA INFORMACIÓN DEL SLICE PARA LA ASIGNACIÓN
	nombreSlice := resultado.SliceInfo.NombreSlice
	indiceExpr := resultado.SliceInfo.IndiceExpr

	// LLAMAR AL SliceProcessor PARA REALIZAR LA ASIGNACIÓN
	err := v.sliceProcessor.AsignarElementoPorIndice(nombreSlice, indiceExpr, nuevoValor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error en asignación por índice: %s", err.Error()))
		return nil
	}

	v.armGen.Comment("Asignación a slice completada exitosamente")
	return nil
}

// ============= IMPRESIÓN - LÓGICA ESPECÍFICA DEL VISITOR =============
func (v *VisitorARM64) generarCodigoImpresion(resultado *assembly.ResultadoExpresion) {
	v.armGen.Comment(fmt.Sprintf("=== IMPRIMIR %s  ===", strings.ToUpper(resultado.Tipo)))

	if resultado.Tipo == "slice_name" {
		nombreSlice := resultado.Valor.(string)
		err := v.sliceProcessor.ImprimirSlice(nombreSlice)
		if err != nil {
			v.armGen.Comment(fmt.Sprintf("Error imprimiendo slice: %s", err.Error()))
		}
	} else if resultado.Tipo == "string" {
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
	} else {
		v.armGen.Comment(fmt.Sprintf("Tipo no soportado para impresión: %s", resultado.Tipo))
	}
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

// ============= ESTRUCTURAS DE CONTROL IF-ELSE =============

func (v *VisitorARM64) VisitIFstmt(ctx *parser.IFstmtContext) interface{} {
	v.armGen.Comment("=== INICIO ESTRUCTURA IF-ELSE ===")

	etiquetaFinal := v.expresionesProcessor.GenerarEtiquetaUnica("if_final")

	// Protección contra panic en el procesamiento de branches
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("❌ PANIC en procesamiento de IF branches: %v\n", r)
				v.armGen.Comment(fmt.Sprintf("ERROR: Panic en procesamiento IF: %v", r))
			}
		}()

		// Procesar todas las cadenas if/else-if
		for _, ifChain := range ctx.AllIf_chain() {
			v.procesarBranchIF(ifChain, etiquetaFinal)
		}

		// Procesar else si existe
		if ctx.Else_stmt() != nil {
			v.armGen.Comment("--- Ejecutando ELSE ---")
			v.Visit(ctx.Else_stmt())
		}
	}()

	// Etiqueta final
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaFinal))
	v.armGen.Comment("=== FIN ESTRUCTURA IF-ELSE ===")
	v.armGen.Instructions = append(v.armGen.Instructions, "")

	return nil
}

func (v *VisitorARM64) procesarBranchIF(ifChain interface{}, etiquetaFinal string) {

	// RESETEAR REGISTROS AL INICIO DE CADA BRANCH
	v.expresionesProcessor.ResetearRegistrosSiNecesario()

	ctx, ok := ifChain.(*parser.IFcadenaContext)
	if !ok {
		fmt.Printf("❌ ERROR: no se pudo convertir ifChain a IFcadenaContext\n")
		v.armGen.Comment("Error: no se pudo convertir ifChain a IFcadenaContext")
		return
	}

	// Protección contra panic en evaluación de expresión
	var resultado interface{}
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("❌ PANIC en evaluación de expresión IF: %v\n", r)
				resultado = &assembly.ResultadoExpresion{
					Registro:  "",
					Tipo:      "bool",
					EsLiteral: true,
					Valor:     false,
				}
			}
		}()
		resultado = v.Visit(ctx.Expr())
	}()

	// VALIDACIÓN SEGURA DEL RESULTADO
	cond := v.validarResultadoExpresionSeguro(resultado, "condición IF")
	if cond == nil {
		fmt.Printf("❌ ERROR: validarResultadoExpresionSeguro retornó nil\n")
		return
	}

	if cond.Tipo != "bool" {
		fmt.Printf("❌ ERROR: condición no es booleana, es: %s\n", cond.Tipo)
		v.armGen.Comment(fmt.Sprintf("Error: condición no booleana: %s", cond.Tipo))
		return
	}

	// Etiqueta para saltar si la condición es falsa
	etiquetaSaltarBranch := v.expresionesProcessor.GenerarEtiquetaUnica("skip_branch")

	if cond.EsLiteral {
		valor := v.obtenerValorBoolSeguro(cond)
		v.armGen.Comment(fmt.Sprintf("Condición literal: %t", valor))

		if valor {
			v.armGen.Comment("=== EJECUTANDO BLOQUE (literal true) ===")
			v.pushScope("if")
			for _, stmt := range ctx.AllStmt() {
				v.Visit(stmt)
			}
			v.popScope()
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("b %s", etiquetaFinal))
		} else {
		}
	} else {
		// Caso runtime: generar código condicional
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("cmp %s, #0", cond.Registro),
			fmt.Sprintf("beq %s", etiquetaSaltarBranch))

		v.armGen.Comment("=== EJECUTANDO BLOQUE (runtime true) ===")
		v.pushScope("if")
		for _, stmt := range ctx.AllStmt() {

			v.Visit(stmt)
		}
		v.popScope()

		// Saltar al final de la estructura completa
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaFinal))

		// Etiqueta para cuando la condición es falsa
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("%s:", etiquetaSaltarBranch))
	}

	v.armGen.Comment("=== FIN PROCESAMIENTO BRANCH ===")
}

func (v *VisitorARM64) VisitIFcadena(ctx *parser.IFcadenaContext) interface{} {
	// Esta función es llamada por procesarBranchIF, solo devolvemos el contexto
	return ctx
}

func (v *VisitorARM64) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {
	v.armGen.Comment("=== INICIO BLOQUE ELSE ===")

	v.pushScope("else")

	// Ejecutar todas las sentencias del bloque else
	for i, stmt := range ctx.AllStmt() {
		v.armGen.Comment(fmt.Sprintf("Ejecutando sentencia %d del bloque else", i+1))
		v.Visit(stmt)
	}

	v.popScope()

	v.armGen.Comment("=== FIN BLOQUE ELSE ===")

	return nil
}

func (v *VisitorARM64) pushScope(nombre string) {
	v.armGen.Comment(fmt.Sprintf("Push scope: %s", nombre))
}

func (v *VisitorARM64) popScope() {
	v.armGen.Comment("Pop scope")
}

// FUNCIÓN MEJORADA: Validación segura de ResultadoExpresion
func (v *VisitorARM64) validarResultadoExpresionSeguro(resultado interface{}, contexto string) *assembly.ResultadoExpresion {

	if resultado == nil {
		v.armGen.Comment(fmt.Sprintf("Error en %s: resultado nil", contexto))
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "bool",
			EsLiteral: true,
			Valor:     false,
		}
	}

	expr, ok := resultado.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment(fmt.Sprintf("Error en %s: no es ResultadoExpresion", contexto))
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "bool",
			EsLiteral: true,
			Valor:     false,
		}
	}

	return expr
}

func (v *VisitorARM64) obtenerValorBoolSeguro(expr *assembly.ResultadoExpresion) bool {

	if expr == nil || expr.Valor == nil {
		return false
	}

	if boolVal, ok := expr.Valor.(bool); ok {
		return boolVal
	}

	// Si no es bool, intentar convertir de otros tipos
	if intVal, ok := expr.Valor.(int); ok {
		return intVal != 0
	}

	if floatVal, ok := expr.Valor.(float64); ok {
		return floatVal != 0.0
	}

	if strVal, ok := expr.Valor.(string); ok {
		return strVal != ""
	}

	return false
}

// Agregar estas funciones a tu archivo VisitorARM64.go

// ============= IMPLEMENTACIÓN SWITCH-CASE CON DEBUG =============

func (v *VisitorARM64) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	v.armGen.Comment("=== INICIO SWITCH STATEMENT ===")

	// 1. Evaluar la expresión principal del switch
	if ctx.Expr() == nil {
		v.armGen.Comment("Error: No hay expresión en el switch")
		return nil
	}

	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment("Error: expresión del switch inválida")
		return nil
	}

	switchExpr, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment("Error: tipo de expresión inválido")
		return nil
	}
	v.armGen.Comment(fmt.Sprintf("Switch sobre expresión tipo: %s", switchExpr.Tipo))

	// 2. Generar etiquetas únicas para el switch
	contadorSwitch := v.getContadorSwitchSeguro()
	etiquetaFinSwitch := fmt.Sprintf(".Lswitch_end_%d", contadorSwitch)
	etiquetaDefault := fmt.Sprintf(".Lswitch_default_%d", contadorSwitch)

	// 2.1 AGREGAR CONTROL DE FLUJO (solo break, no continue en switch)
	switchItem := v.pushLoopControl("switch", []string{"break"}, etiquetaFinSwitch, "")
	defer func() {
		v.popLoopControl()
		v.limpiarItem(switchItem)
	}()

	// 3. Cargar valor del switch en registro temporal
	registroSwitch := v.getNuevoRegistroSeguro()

	if switchExpr.EsLiteral {
		switch switchExpr.Tipo {
		case "int":
			if valor, ok := switchExpr.Valor.(int); ok {
				v.armGen.Mov(registroSwitch, valor)
			} else {
				v.armGen.Mov(registroSwitch, 0)
			}
		case "bool":
			if valor, ok := switchExpr.Valor.(bool); ok {
				valorInt := 0
				if valor {
					valorInt = 1
				}
				v.armGen.Mov(registroSwitch, valorInt)
			} else {
				v.armGen.Mov(registroSwitch, 0)
			}
		case "string":
			if valor, ok := switchExpr.Valor.(string); ok {
				etiqueta := v.agregarMensajeStringSeguro(valor)
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("adr %s, %s", registroSwitch, etiqueta))
			} else {
				v.armGen.Mov(registroSwitch, 0)
			}
		default:
			v.armGen.Mov(registroSwitch, 0)
		}
	} else {
		v.armGen.MovReg(registroSwitch, switchExpr.Registro)
	}

	// 4. Procesar todos los cases
	cases := ctx.AllSwitch_case()
	if cases == nil {
		cases = []parser.ISwitch_caseContext{}
	}

	etiquetasCases := make([]string, len(cases))
	for i := range cases {
		etiquetasCases[i] = fmt.Sprintf(".Lcase_%d_%d", contadorSwitch, i)
	}

	// 5. Generar comparaciones para cada case
	for i, caseCtx := range cases {

		caseResult := v.GetCaseValue(caseCtx)
		if caseResult == nil {
			continue
		}

		caseExpr, ok := caseResult.(*assembly.ResultadoExpresion)
		if !ok {
			continue
		}

		if !v.sonTiposCompatibles(switchExpr.Tipo, caseExpr.Tipo) {
			continue
		}

		v.armGen.Comment(fmt.Sprintf("Comparación case %d", i))

		// Generar comparación según el tipo
		switch switchExpr.Tipo {
		case "int", "bool":
			registroCase := v.getNuevoRegistroSeguro()
			if caseExpr.EsLiteral {
				valor := 0
				if switchExpr.Tipo == "bool" && caseExpr.Valor.(bool) {
					valor = 1
				} else if switchExpr.Tipo == "int" {
					valor = caseExpr.Valor.(int)
				}
				v.armGen.Mov(registroCase, valor)
			} else {
				v.armGen.MovReg(registroCase, caseExpr.Registro)
			}

			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("cmp %s, %s", registroSwitch, registroCase),
				fmt.Sprintf("beq %s", etiquetasCases[i]))

		case "string":
			// FIX: Usar el método StrCmp existente del ARMGenerator
			registroCase := v.getNuevoRegistroSeguro()
			registroResultado := v.getNuevoRegistroSeguro()

			if caseExpr.EsLiteral {
				etiquetaStr := v.agregarMensajeStringSeguro(caseExpr.Valor.(string))
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("adr %s, %s", registroCase, etiquetaStr))
			} else {
				v.armGen.MovReg(registroCase, caseExpr.Registro)
			}

			// FIX: Usar StrCmp que ya maneja strcmp correctamente
			v.armGen.StrCmp(registroSwitch, registroCase, registroResultado)

			// FIX: strcmp retorna 0 cuando son iguales, así que comparamos con 0
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("cmp %s, #0", registroResultado),
				fmt.Sprintf("beq %s", etiquetasCases[i]))
		}
	}

	// 6. Si no hay coincidencias, saltar al default o al final
	if ctx.Default_case() != nil {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaDefault))
	} else {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaFinSwitch))
	}

	// 7. Generar código para cada case
	for i, caseCtx := range cases {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("%s:", etiquetasCases[i]))

		v.armGen.Comment(fmt.Sprintf("Ejecutando case %d", i))

		if caseCtx != nil {
			v.Visit(caseCtx)
		}

		// Break implícito: saltar al final
		// NOTA: El break explícito del usuario también saltará aquí
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaFinSwitch))
	}

	// 8. Generar código para default si existe
	if ctx.Default_case() != nil {
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("%s:", etiquetaDefault))

		v.armGen.Comment("Ejecutando case default")
		v.Visit(ctx.Default_case())
	}

	// 9. Etiqueta final del switch
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaFinSwitch))

	v.armGen.Comment("=== FIN SWITCH STATEMENT ===")
	return nil
}

// ============= FUNCIONES AUXILIARES  =============

func (v *VisitorARM64) GetCaseValue(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.SwitchCaseContext:
		if val.Expr() == nil {
			return nil
		}
		return v.Visit(val.Expr())
	default:
		v.armGen.Comment("Error interno: tipo de case no reconocido")
		return nil
	}
}

func (v *VisitorARM64) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {
	if ctx == nil {
		return nil
	}

	statements := ctx.AllStmt()
	for _, stmt := range statements {
		if stmt != nil {
			v.Visit(stmt)
		}
	}
	return nil
}

func (v *VisitorARM64) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	if ctx == nil {
		return nil
	}

	statements := ctx.AllStmt()
	for _, stmt := range statements {
		if stmt != nil {
			v.Visit(stmt)
		}
	}
	return nil
}

func (v *VisitorARM64) sonTiposCompatibles(tipo1, tipo2 string) bool {
	if tipo1 == tipo2 {
		return true
	}
	if (tipo1 == "int" && tipo2 == "bool") || (tipo1 == "bool" && tipo2 == "int") {
		return true
	}
	return false
}

// ============= IMPLEMENTACIÓN FOR TIPO WHILE =============
func (v *VisitorARM64) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	v.armGen.Comment("=== INICIO WHILE ARM64 ===")

	// Generar etiquetas únicas
	contadorLoop := v.expresionesProcessor.GetContadorEtiqueta()
	etiquetaInicio := fmt.Sprintf(".Lwhile_start_%d", contadorLoop)
	etiquetaFin := fmt.Sprintf(".Lwhile_end_%d", contadorLoop)
	etiquetaContinue := fmt.Sprintf(".Lwhile_continue_%d", contadorLoop)

	// Crear item de control
	whileItem := v.pushLoopControl("while", []string{"break", "continue"}, etiquetaFin, etiquetaContinue)

	// Manejo con defer
	defer func() {
		v.popLoopControl()
		v.limpiarItem(whileItem)
	}()

	// Evaluar condición inicial
	condicionResult := v.Visit(ctx.Expr())
	condicion := v.validarResultadoExpresionSeguro(condicionResult, "condición while")
	if condicion == nil || condicion.Tipo != "bool" {
		v.armGen.Comment("ERROR: Condición del while inválida")
		return nil
	}

	// Llamar función interna
	v.VisitInnerWhileARM64(ctx, condicion, etiquetaInicio, etiquetaFin, etiquetaContinue)

	v.armGen.Comment("=== FIN WHILE ARM64 ===")
	return nil
}

// FUNCIÓN INTERNA DEL WHILE
func (v *VisitorARM64) VisitInnerWhileARM64(ctx *parser.WhileStmtContext, condicionInicial *assembly.ResultadoExpresion, etiquetaInicio, etiquetaFin, etiquetaContinue string) {
	// ETIQUETA DE INICIO DEL BUCLE
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaInicio))

	// EVALUAR CONDICIÓN EN CADA ITERACIÓN
	v.armGen.Comment("Evaluando condición del while...")
	condicionActual := v.Visit(ctx.Expr())

	condicion := v.validarResultadoExpresionSeguro(condicionActual, "condición while en bucle")
	if condicion == nil {
		v.armGen.Comment("Error: condición del while inválida en bucle")
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaFin))
	} else if condicion.Tipo != "bool" {
		v.armGen.Comment("Error: la condición del ciclo debe ser un booleano")
		v.armGen.Instructions = append(v.armGen.Instructions,
			fmt.Sprintf("b %s", etiquetaFin))
	} else {
		// VERIFICAR CONDICIÓN Y SALIR SI ES FALSA
		if condicion.EsLiteral {
			valor := v.obtenerValorBoolSeguro(condicion)
			v.armGen.Comment(fmt.Sprintf("Condición literal: %t", valor))

			if !valor {
				// Si la condición es literalmente false, salir directamente
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("b %s", etiquetaFin))
			} else {
				// EJECUTAR CUERPO SI LA CONDICIÓN ES TRUE
				v.ejecutarCuerpoWhile(ctx, etiquetaInicio, etiquetaFin, etiquetaContinue)
			}
		} else {
			//  si es false (0), salir del bucle
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("cmp %s, #0", condicion.Registro),
				fmt.Sprintf("beq %s", etiquetaFin))

			// EJECUTAR CUERPO DEL BUCLE
			v.ejecutarCuerpoWhile(ctx, etiquetaInicio, etiquetaFin, etiquetaContinue)
		}
	}

	// SIEMPRE GENERAR ETIQUETA DE FIN DEL BUCLE
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaFin))
	v.armGen.Comment(">>> FIN DEL BUCLE WHILE <<<")
}

// FUNCIÓN SEPARADA PARA EJECUTAR EL CUERPO DEL WHILE
func (v *VisitorARM64) ejecutarCuerpoWhile(ctx *parser.WhileStmtContext, etiquetaInicio, etiquetaFin, etiquetaContinue string) {

	// Ejecutar todas las sentencias del cuerpo
	for i, stmt := range ctx.AllStmt() {
		v.armGen.Comment(fmt.Sprintf("Ejecutando sentencia %d del while", i+1))

		// PROTECCIÓN CONTRA ERRORES EN SENTENCIAS INDIVIDUALES
		func() {
			defer func() {
				if r := recover(); r != nil {
					v.armGen.Comment(fmt.Sprintf("Error en sentencia %d: %v", i+1, r))
				}
			}()

			v.expresionesProcessor.ResetearRegistrosSiNecesario()
			v.Visit(stmt)
		}()
	}

	// ETIQUETA CONTINUE
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaContinue))

	// SALTAR DE VUELTA AL INICIO PARA REEVALUAR LA CONDICIÓN
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", etiquetaInicio))
}

// ============= IMPLEMENTACIÓN FOR TIPO RANGE ============= // SEGUIR IMPLEMENTANDO.
func (v *VisitorARM64) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	v.armGen.Comment("=== INICIO FOR TIPO WHILE ===")

	// VALIDACIÓN INICIAL
	if ctx == nil {
		v.armGen.Comment("Error: contexto de for inválido")
		return nil
	}

	// VERIFICAR QUE TENEMOS AL MENOS UN ID
	if len(ctx.AllID()) == 0 {
		v.armGen.Comment("Error: el for debe tener al menos una variable")
		return nil
	}

	// DETECTAR SI HAY UNA O DOS VARIABLES
	var nombreIndice string = ""
	var nombreValor string = ""

	if len(ctx.AllID()) == 1 {
		// for val in collection
		if ctx.ID(0) == nil {
			v.armGen.Comment("Error: variable del for no válida")
			return nil
		}
		nombreValor = ctx.ID(0).GetText()
	} else if len(ctx.AllID()) == 2 {
		// for idx, val in collection
		if ctx.ID(0) == nil || ctx.ID(1) == nil {
			v.armGen.Comment("Error: variables del for no válidas")
			return nil
		}
		nombreIndice = ctx.ID(0).GetText()
		nombreValor = ctx.ID(1).GetText()
	} else {
		v.armGen.Comment("Error: el for debe tener 1 ó 2 variables")
		return nil
	}

	var iterableResult *assembly.ResultadoExpresion = nil

	// EVALUAR RANGE (opcional)
	if ctx.Range_() != nil {
		rangeResult := v.Visit(ctx.Range_())
		if rangeResult == nil {
			v.armGen.Comment("Error: error evaluando rango")
			return nil
		}

		// El rango debe retornar un slice_name
		if rangeExpr, ok := rangeResult.(*assembly.ResultadoExpresion); ok {
			if rangeExpr.Tipo == "slice_name" {
				iterableResult = rangeExpr
			} else {
				v.armGen.Comment("Error: el rango debe producir un slice")
				return nil
			}
		} else {
			v.armGen.Comment("Error: resultado de rango inválido")
			return nil
		}
	}

	// EVALUAR EXPRESIÓN (opcional)
	if ctx.Expr() != nil {
		exprResult := v.Visit(ctx.Expr())
		if exprResult == nil {
			v.armGen.Comment("Error: error evaluando expresión del for")
			return nil
		}

		if exprVal, ok := exprResult.(*assembly.ResultadoExpresion); ok {
			// VERIFICAR QUE SEA UN SLICE
			if exprVal.Tipo == "slice_name" {
				iterableResult = exprVal
			} else if exprVal.Tipo == "string" {
				// Convertir string a slice de caracteres si es necesario
				v.armGen.Comment("Iteración sobre string no implementada aún")
				return nil
			} else {
				v.armGen.Comment(fmt.Sprintf("Error: no se puede iterar sobre tipo %s", exprVal.Tipo))
				return nil
			}
		} else {
			v.armGen.Comment("Error: expresión del for no válida")
			return nil
		}
	}

	// VALIDAR QUE TENEMOS ALGO QUE ITERAR
	if iterableResult == nil {
		v.armGen.Comment("Error: no se pudo determinar qué iterar")
		return nil
	}

	nombreSlice := iterableResult.Valor.(string)

	// VERIFICAR QUE EL SLICE EXISTE
	if !v.sliceProcessor.ExisteSlice(nombreSlice) {
		v.armGen.Comment(fmt.Sprintf("Error: slice '%s' no encontrado", nombreSlice))
		return nil
	}

	// GENERAR CÓDIGO ARM64 PARA EL FOR
	v.generarForWhileARM64(ctx, nombreSlice, nombreIndice, nombreValor)

	v.armGen.Comment("=== FIN FOR TIPO WHILE ===")
	return nil
}

// FUNCIÓN PRINCIPAL PARA GENERAR CÓDIGO ARM64 DEL FOR
func (v *VisitorARM64) generarForWhileARM64(ctx *parser.ForStmtContext, nombreSlice, nombreIndice, nombreValor string) {
	v.armGen.Comment("=== GENERANDO FOR WHILE ARM64 ===")

	// GENERAR ETIQUETAS ÚNICAS
	contadorFor := v.expresionesProcessor.GetContadorEtiqueta()
	etiquetaInicioFor := fmt.Sprintf(".Lfor_start_%d", contadorFor)
	etiquetaFinFor := fmt.Sprintf(".Lfor_end_%d", contadorFor)
	etiquetaContinueFor := fmt.Sprintf(".Lfor_continue_%d", contadorFor)

	// OBTENER REGISTROS PARA CONTROL DEL LOOP
	regIndiceActual := v.expresionesProcessor.NuevoRegistroTmp()  // Para el índice actual
	regLongitudSlice := v.expresionesProcessor.NuevoRegistroTmp() // Para la longitud del slice

	v.armGen.Comment(">>> INICIALIZACIÓN DEL FOR <<<")

	// INICIALIZAR ÍNDICE EN 0
	v.armGen.Mov(regIndiceActual, 0)
	v.armGen.Comment(fmt.Sprintf("Inicializar índice del for en %s = 0", regIndiceActual))

	// OBTENER LONGITUD DEL SLICE
	longitudResult, err := v.sliceProcessor.Len(nombreSlice)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error obteniendo longitud del slice: %s", err.Error()))
		return
	}

	if longitudResult.EsLiteral {
		v.armGen.Mov(regLongitudSlice, longitudResult.Valor.(int))
	} else {
		v.armGen.MovReg(regLongitudSlice, longitudResult.Registro)
	}
	v.armGen.Comment(fmt.Sprintf("Longitud del slice en %s", regLongitudSlice))

	// CREAR ÁMBITO PARA LAS VARIABLES DEL FOR
	v.pushScope("for")

	// DECLARAR VARIABLE DEL ÍNDICE (si existe)
	if nombreIndice != "" {
		err := v.variablesProcessor.DeclararVariable(nombreIndice, "int", &assembly.ResultadoExpresion{
			Registro:  regIndiceActual,
			Tipo:      "int",
			EsLiteral: false,
		})
		if err != nil {
			v.armGen.Comment(fmt.Sprintf("Error declarando variable índice: %s", err.Error()))
		}
	}

	// DECLARAR VARIABLE DEL VALOR
	// Obtener tipo del slice para la variable valor
	tipoElemento, err := v.sliceProcessor.ObtenerTipoElemento(nombreSlice)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error obteniendo tipo del slice: %s", err.Error()))
		tipoElemento = "int" // Valor por defecto
	}

	err = v.variablesProcessor.DeclararVariable(nombreValor, tipoElemento, &assembly.ResultadoExpresion{
		Registro:  "",
		Tipo:      tipoElemento,
		EsLiteral: true,
		Valor:     0, // Valor inicial por defecto
	})
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error declarando variable valor: %s", err.Error()))
	}

	// ETIQUETA DE INICIO DEL LOOP
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaInicioFor))
	v.armGen.Comment(">>> INICIO DEL BUCLE <<<")

	// VERIFICAR CONDICIÓN: índice < longitud
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("cmp %s, %s", regIndiceActual, regLongitudSlice),
		fmt.Sprintf("bge %s", etiquetaFinFor)) // Si índice >= longitud, salir

	v.armGen.Comment("Condición del bucle verificada")

	// ACTUALIZAR VARIABLE DEL ÍNDICE (si existe)
	if nombreIndice != "" {
		err := v.variablesProcessor.AsignarVariable(nombreIndice, &assembly.ResultadoExpresion{
			Registro:  regIndiceActual,
			Tipo:      "int",
			EsLiteral: false,
		})
		if err != nil {
			v.armGen.Comment(fmt.Sprintf("Error actualizando variable índice: %s", err.Error()))
		}
	}

	// OBTENER ELEMENTO ACTUAL DEL SLICE
	elementoActual, err := v.sliceProcessor.AccederElementoPorIndice(nombreSlice, &assembly.ResultadoExpresion{
		Registro:  regIndiceActual,
		Tipo:      "int",
		EsLiteral: false,
	})
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error accediendo elemento del slice: %s", err.Error()))
		return
	}

	// ACTUALIZAR VARIABLE DEL VALOR
	err = v.variablesProcessor.AsignarVariable(nombreValor, elementoActual)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error actualizando variable valor: %s", err.Error()))
	}

	v.armGen.Comment("Variables del for actualizadas")

	// EJECUTAR CUERPO DEL BUCLE
	v.armGen.Comment(">>> EJECUTANDO CUERPO DEL BUCLE <<<")

	// Manejar break y continue con defer
	defer func() {
		if r := recover(); r != nil {
			// Manejar break y continue
			v.armGen.Comment("Manejo de break/continue en for")
		}
	}()

	for i, stmt := range ctx.AllStmt() {
		v.armGen.Comment(fmt.Sprintf("Ejecutando sentencia %d del for", i+1))
		v.expresionesProcessor.ResetearRegistrosSiNecesario() // Evitar agotar registros
		v.Visit(stmt)
	}

	// ETIQUETA CONTINUE (para continue statements)
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaContinueFor))
	v.armGen.Comment(">>> INCREMENTO E ITERACIÓN <<<")

	// INCREMENTAR ÍNDICE
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("add %s, %s, #1", regIndiceActual, regIndiceActual))
	v.armGen.Comment("Incrementar índice del for")

	// SALTAR DE VUELTA AL INICIO
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", etiquetaInicioFor))

	// ETIQUETA DE FIN DEL LOOP
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaFinFor))
	v.armGen.Comment(">>> FIN DEL BUCLE <<<")

	// LIMPIAR ÁMBITO
	v.popScope()

	v.armGen.Comment("=== FIN GENERACIÓN FOR WHILE ARM64 ===")
}

// IMPLEMENTAR VisitRangoNum PARA ARM64
func (v *VisitorARM64) VisitRangoNum(ctx *parser.RangoNumContext) interface{} {
	v.armGen.Comment("=== GENERANDO RANGO NUMÉRICO ===")

	// EVALUAR EXPRESIONES DEL RANGO
	leftResult := v.Visit(ctx.Expr(0))
	rightResult := v.Visit(ctx.Expr(1))

	if leftResult == nil || rightResult == nil {
		v.armGen.Comment("Error: expresiones del rango inválidas")
		return nil
	}

	leftExpr, okLeft := leftResult.(*assembly.ResultadoExpresion)
	rightExpr, okRight := rightResult.(*assembly.ResultadoExpresion)

	if !okLeft || !okRight {
		v.armGen.Comment("Error: tipos de expresiones del rango inválidos")
		return nil
	}

	if leftExpr.Tipo != "int" || rightExpr.Tipo != "int" {
		v.armGen.Comment("Error: los valores del rango deben ser enteros")
		return nil
	}

	// GENERAR SLICE TEMPORAL CON EL RANGO
	contadorRango := v.expresionesProcessor.GetContadorEtiqueta()
	nombreSliceRango := fmt.Sprintf("_range_slice_%d", contadorRango)

	// CREAR SLICE TEMPORAL VACÍO
	err := v.sliceProcessor.DeclararSlice(nombreSliceRango, "int", []*assembly.ResultadoExpresion{})
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error creando slice de rango: %s", err.Error()))
		return nil
	}

	// GENERAR CÓDIGO PARA LLENAR EL SLICE CON EL RANGO
	v.armGen.Comment(">>> LLENANDO SLICE DE RANGO <<<")

	regInicio := v.expresionesProcessor.NuevoRegistroTmp()
	regFin := v.expresionesProcessor.NuevoRegistroTmp()
	regActual := v.expresionesProcessor.NuevoRegistroTmp()

	// Cargar valores de inicio y fin
	if leftExpr.EsLiteral {
		v.armGen.Mov(regInicio, leftExpr.Valor.(int))
	} else {
		v.armGen.MovReg(regInicio, leftExpr.Registro)
	}

	if rightExpr.EsLiteral {
		v.armGen.Mov(regFin, rightExpr.Valor.(int))
	} else {
		v.armGen.MovReg(regFin, rightExpr.Registro)
	}

	// Verificar que inicio <= fin
	etiquetaRangoOk := fmt.Sprintf(".Lrange_ok_%d", contadorRango)
	etiquetaRangoEnd := fmt.Sprintf(".Lrange_end_%d", contadorRango)
	etiquetaRangoLoop := fmt.Sprintf(".Lrange_loop_%d", contadorRango)

	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("cmp %s, %s", regInicio, regFin),
		fmt.Sprintf("ble %s", etiquetaRangoOk))

	v.armGen.Comment("Error: valor izquierdo del rango debe ser <= valor derecho")
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", etiquetaRangoEnd))

	// Loop para llenar el slice
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaRangoOk),
		fmt.Sprintf("mov %s, %s", regActual, regInicio))

	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaRangoLoop),
		fmt.Sprintf("cmp %s, %s", regActual, regFin),
		fmt.Sprintf("bgt %s", etiquetaRangoEnd))

	// Agregar elemento actual al slice
	_, err = v.sliceProcessor.Append(nombreSliceRango, &assembly.ResultadoExpresion{
		Registro:  regActual,
		Tipo:      "int",
		EsLiteral: false,
	})
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error agregando elemento al rango: %s", err.Error()))
	}

	// Incrementar y continuar
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("add %s, %s, #1", regActual, regActual),
		fmt.Sprintf("b %s", etiquetaRangoLoop))

	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaRangoEnd))

	v.armGen.Comment("=== FIN GENERACIÓN RANGO ===")

	// Retornar referencia al slice creado
	return &assembly.ResultadoExpresion{
		Registro:  "",
		Tipo:      "slice_name",
		EsLiteral: true,
		Valor:     nombreSliceRango,
	}
}

// ================= IMPLEMENTACIÓN FOR CLÁSICO ARM64 CORREGIDA =================
func (v *VisitorARM64) VisitForClasicoStmt(ctx *parser.ForClasicoStmtContext) interface{} {
	v.armGen.Comment("=== INICIO FOR CLÁSICO ARM64 ===")

	// Generar etiquetas únicas
	id := v.expresionesProcessor.GetContadorEtiqueta()
	etiquetaStart := fmt.Sprintf(".Lforc_start_%d", id)
	etiquetaContinue := fmt.Sprintf(".Lforc_continue_%d", id)
	etiquetaEnd := fmt.Sprintf(".Lforc_end_%d", id)

	// Crear item de control
	forItem := v.pushLoopControl("for_clasico", []string{"break", "continue"}, etiquetaEnd, etiquetaContinue)

	// Manejo con defer (similar al intérprete)
	defer func() {
		v.popLoopControl()
		v.limpiarItem(forItem)
	}()

	// Llamar función interna
	v.generateInnerForClasicoARM64(ctx, etiquetaStart, etiquetaContinue, etiquetaEnd)

	v.armGen.Comment("=== FIN FOR CLÁSICO ARM64 ===")
	return nil
}

// generateInnerForClasicoARM64 inyecta las instrucciones ARM64 para el for clásico
func (v *VisitorARM64) generateInnerForClasicoARM64(
	ctx *parser.ForClasicoStmtContext,
	etiquetaStart, etiquetaContinue, etiquetaEnd string,
) {
	// INICIALIZACIÓN (opcional)
	// CORRECCIÓN: Usar los métodos correctos
	stmtAsignarList := ctx.AllStmt_asignar()
	if len(stmtAsignarList) > 0 && stmtAsignarList[0] != nil {
		v.Visit(stmtAsignarList[0])
	}

	// Etiqueta de inicio del bucle
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaStart))
	v.armGen.Comment("Evaluar condición del for clásico")

	// CONDICIÓN (opcional)
	// CORRECCIÓN: Usar el método correcto
	if ctx.Expr() != nil {
		condResult := v.Visit(ctx.Expr())
		cond := v.validarResultadoExpresionSeguro(condResult, "condición for clásico")
		if cond == nil || cond.Tipo != "bool" {
			// si inválida o no bool → salir
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("b %s", etiquetaEnd))
		} else if cond.EsLiteral {
			if !v.obtenerValorBoolSeguro(cond) {
				v.armGen.Instructions = append(v.armGen.Instructions,
					fmt.Sprintf("b %s", etiquetaEnd))
			}
			// Si es true literal, simplemente continúa (no genera código)
		} else {
			// variable: compare registro con 0
			v.armGen.Instructions = append(v.armGen.Instructions,
				fmt.Sprintf("cmp %s, #0", cond.Registro),
				fmt.Sprintf("beq %s", etiquetaEnd))
		}
	}

	// 3️⃣ CUERPO
	// CORRECCIÓN: Usar el método correcto
	stmtList := ctx.AllStmt()
	for i, stmt := range stmtList {
		v.armGen.Comment(fmt.Sprintf("Sentencia %d del for clásico", i+1))
		v.expresionesProcessor.ResetearRegistrosSiNecesario()
		v.Visit(stmt)
	}

	// Etiqueta de continue → aquí hacemos el incremento
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaContinue))

	// CORRECCIÓN: Usar los métodos correctos para el incremento
	if len(stmtAsignarList) > 1 && stmtAsignarList[1] != nil {
		v.Visit(stmtAsignarList[1])
	}

	// Salto de vuelta al inicio
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", etiquetaStart))

	// Etiqueta de fin
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s:", etiquetaEnd))
}

// ============= FUNCIONES AUXILIARES SEGURAS =============

func (v *VisitorARM64) getContadorSwitchSeguro() int {
	contadorSwitch++
	return contadorSwitch
}

var contadorSwitch int = 0

func (v *VisitorARM64) getNuevoRegistroSeguro() string {
	if v.expresionesProcessor != nil {
		return v.expresionesProcessor.NuevoRegistroTmp()
	}

	contadorRegistro++
	return fmt.Sprintf("x%d", 19+contadorRegistro%10)
}

var contadorRegistro int = 0

func (v *VisitorARM64) agregarMensajeStringSeguro(mensaje string) string {
	if v.expresionesProcessor != nil {
		return v.expresionesProcessor.AgregarMensajeString(mensaje)
	}

	contadorString++
	etiqueta := fmt.Sprintf("str_%d", contadorString)
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("%s: .asciz \"%s\"", etiqueta, mensaje))
	return etiqueta
}

var contadorString int = 0

// Métodos para LoopControlItem
func (item *LoopControlItem) PuedeDetener() bool {
	for _, accion := range item.Acciones {
		if accion == "break" {
			return true
		}
	}
	return false
}

func (item *LoopControlItem) PuedeContinuar() bool {
	for _, accion := range item.Acciones {
		if accion == "continue" {
			return true
		}
	}
	return false
}

func (item *LoopControlItem) EsAccion(accion string) bool {
	return item.AccionActual == accion
}

func (item *LoopControlItem) ResetAccion() {
	item.AccionActual = ""
}

// ============= PILA DE CONTROL EN VisitorARM64 =============

// Agregar al struct VisitorARM64:
// pilaControl []*LoopControlItem

// Métodos para manejar la pila
func (v *VisitorARM64) pushLoopControl(tipoLoop string, acciones []string, etiquetaSalida, etiquetaContinue string) *LoopControlItem {
	item := &LoopControlItem{
		TipoLoop:         tipoLoop,
		Acciones:         acciones,
		EtiquetaSalida:   etiquetaSalida,
		EtiquetaContinue: etiquetaContinue,
	}
	v.pilaControl = append(v.pilaControl, item)
	v.armGen.Comment(fmt.Sprintf("Push control: %s", tipoLoop))
	return item
}

func (v *VisitorARM64) popLoopControl() {
	if len(v.pilaControl) > 0 {
		ultimo := v.pilaControl[len(v.pilaControl)-1]
		v.pilaControl = v.pilaControl[:len(v.pilaControl)-1]
		v.armGen.Comment(fmt.Sprintf("Pop control: %s", ultimo.TipoLoop))
	}
}

func (v *VisitorARM64) puedeDetener() (bool, *LoopControlItem) {
	for i := len(v.pilaControl) - 1; i >= 0; i-- {
		if v.pilaControl[i].PuedeDetener() {
			return true, v.pilaControl[i]
		}
	}
	return false, nil
}

func (v *VisitorARM64) puedeContinuar() (bool, *LoopControlItem) {
	for i := len(v.pilaControl) - 1; i >= 0; i-- {
		if v.pilaControl[i].PuedeContinuar() {
			return true, v.pilaControl[i]
		}
	}
	return false, nil
}

func (v *VisitorARM64) limpiarItem(item *LoopControlItem) {
	for i := len(v.pilaControl) - 1; i >= 0; i-- {
		if v.pilaControl[i] == item {
			v.pilaControl = append(v.pilaControl[:i], v.pilaControl[i+1:]...)
			break
		}
	}
}

// ============= SENTENCIAS DE TRANSFERENCIA =============

func (v *VisitorARM64) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	v.armGen.Comment("=== BREAK STATEMENT ===")

	exists, item := v.puedeDetener()
	if !exists {
		v.armGen.Comment("ERROR: break debe estar dentro de un ciclo o switch")
		return nil
	}

	v.armGen.Comment(fmt.Sprintf("Break desde %s", item.TipoLoop))
	item.AccionActual = "break"

	// Generar salto a la etiqueta de salida
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", item.EtiquetaSalida))

	return nil
}

func (v *VisitorARM64) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	v.armGen.Comment("=== CONTINUE STATEMENT ===")

	exists, item := v.puedeContinuar()
	if !exists {
		v.armGen.Comment("ERROR: continue debe estar dentro de un ciclo")
		return nil
	}

	v.armGen.Comment(fmt.Sprintf("Continue en %s", item.TipoLoop))
	item.AccionActual = "continue"

	// Generar salto a la etiqueta de continue
	v.armGen.Instructions = append(v.armGen.Instructions,
		fmt.Sprintf("b %s", item.EtiquetaContinue))

	return nil
}

func (v *VisitorARM64) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	v.armGen.Comment("=== RETURN STATEMENT ===")

	// Evaluar expresión de retorno si existe
	if ctx.Expr() != nil {
		exprResult := v.Visit(ctx.Expr())
		if exprResult != nil {
			resultado := exprResult.(*assembly.ResultadoExpresion)

			// Mover resultado al registro de retorno apropiado
			switch resultado.Tipo {
			case "int", "bool":
				if resultado.EsLiteral {
					valor := 0
					if resultado.Tipo == "bool" && resultado.Valor.(bool) {
						valor = 1
					} else if resultado.Tipo == "int" {
						valor = resultado.Valor.(int)
					}
					v.armGen.Mov("x0", valor)
				} else {
					v.armGen.Instructions = append(v.armGen.Instructions,
						fmt.Sprintf("mov x0, %s", resultado.Registro))
				}
			case "float":
				if resultado.EsLiteral {
					regTmp := v.expresionesProcessor.NuevoRegistroFloatTmp()
					v.generarFloatInmediato(regTmp, resultado.Valor.(float64))
					v.armGen.Instructions = append(v.armGen.Instructions,
						fmt.Sprintf("fmov d0, %s", regTmp))
				} else {
					v.armGen.Instructions = append(v.armGen.Instructions,
						fmt.Sprintf("fmov d0, %s", resultado.Registro))
				}
			case "string":
				if resultado.EsLiteral {
					etiqueta := v.expresionesProcessor.AgregarMensajeString(resultado.Valor.(string))
					v.armGen.Instructions = append(v.armGen.Instructions,
						fmt.Sprintf("adr x0, %s", etiqueta))
				} else {
					v.armGen.Instructions = append(v.armGen.Instructions,
						fmt.Sprintf("mov x0, %s", resultado.Registro))
				}
			}
		}
	} else {
		// Return sin valor
		v.armGen.Mov("x0", 0)
	}

	// TODO: Restaurar stack frame
	v.armGen.Instructions = append(v.armGen.Instructions, "ret")
	return nil
}
