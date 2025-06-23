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
}

func NewVisitorARM64() *VisitorARM64 {
	armGen := assembly.NewARMGenerator()
	expresionesProcessor := assembly.NewExpresionesProcessor(armGen)
	return &VisitorARM64{
		armGen:               armGen,
		expresionesProcessor: expresionesProcessor,
		variablesProcessor:   assembly.NewVariablesProcessor(armGen, expresionesProcessor),
		sliceProcessor:       assembly.NewSliceProcessor(armGen, expresionesProcessor),
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

func (v *VisitorARM64) VisitStmt_asignar(ctx *parser.Stmt_asignarContext) interface{} {
	if ctx.GetChildCount() > 0 {
		if child, ok := ctx.GetChild(0).(antlr.ParseTree); ok {
			// VERIFICAR SI ES DECLARACIÓN DE SLICE
			if declSliceCtx, ok := child.(*parser.DeclararSliceContext); ok {
				return v.VisitDeclararSlice(declSliceCtx)
			}

			// VERIFICAR SI ES ASIGNACIÓN A ELEMENTO DE SLICE
			if asignSliceItemCtx, ok := child.(*parser.AsignacionSliceItemContext); ok {
				return v.VisitAsignacionSliceItem(asignSliceItemCtx)
			}

			return v.Visit(child)
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
	v.armGen.Comment("Asignación aritmética no implementada aún")
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

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararTipo(ctx *parser.DeclararTipoContext) interface{} {
	nombre := ctx.ID().GetText()
	tipo := v.obtenerTipoString(ctx.Tipo())

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, nil)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclaraTipoValor(ctx *parser.DeclaraTipoValorContext) interface{} {
	nombre := ctx.ID().GetText()
	tipo := v.obtenerTipoString(ctx.Tipo())
	exprResult := v.Visit(ctx.Expr())
	valor := exprResult.(*assembly.ResultadoExpresion)

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararInferencia(ctx *parser.DeclararInferenciaContext) interface{} {
	nombre := ctx.ID().GetText()
	exprResult := v.Visit(ctx.Expr())
	valor := exprResult.(*assembly.ResultadoExpresion)
	tipo := valor.Tipo

	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	}
	return nil
}

func (v *VisitorARM64) VisitDeclararInferenciaMut(ctx *parser.DeclararInferenciaMutContext) interface{} {
	nombre := ctx.ID().GetText()
	exprResult := v.Visit(ctx.Expr())
	if exprResult == nil {
		v.armGen.Comment(fmt.Sprintf("Error: expresión inválida para %s", nombre))
		return nil
	}

	valor, ok := exprResult.(*assembly.ResultadoExpresion)
	if !ok {
		v.armGen.Comment(fmt.Sprintf("Error: tipo de expresión inválido para %s", nombre))
		return nil
	}

	tipo := valor.Tipo
	err := v.variablesProcessor.DeclararVariable(nombre, tipo, valor)
	if err != nil {
		v.armGen.Comment(fmt.Sprintf("Error: %s", err.Error()))
	} else {
		v.armGen.Comment(fmt.Sprintf("Declaración mut inferida: mut %s := ?", nombre))
	}

	return nil
}

func (v *VisitorARM64) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones  - USAR LA SINTAXIS CORRECTA
	if ds := ctx.Stmt_declaracion(); ds != nil {
		return v.Visit(ds)
	}

	// 2) Asignaciones  - USAR LA SINTAXIS CORRECTA
	if ctx.Stmt_asignar() != nil {
		return v.Visit(ctx.Stmt_asignar())
	}

	// 3) Resto de sentencias (igual que antes)
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
		} else {
			// Múltiples elementos para agregar
			var elementos []*assembly.ResultadoExpresion
			for i := 1; i < len(argumentos); i++ {
				elem, ok := argumentos[i].(*assembly.ResultadoExpresion)
				if !ok {
					v.armGen.Comment(fmt.Sprintf("Error: argumento %d de append no es válido", i+1))
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
				elementos = append(elementos, elem)
			}

			// Llamar a la función append múltiple
			resultado, err := v.sliceProcessor.AppendMultiple(nombreSlice, elementos)
			if err != nil {
				v.armGen.Comment(fmt.Sprintf("Error en append múltiple: %s", err.Error()))
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

// PROBLEMA 1: VisitIdExp no maneja slices correctamente
func (v *VisitorARM64) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	nombreVariable := ctx.PatronId().GetText()

	//  PRIMER PRIORITY: Verificar si es un slice
	if v.sliceProcessor.ExisteSlice(nombreVariable) {
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "slice_name",
			EsLiteral: true,
			Valor:     nombreVariable,
		}
	}

	//  SEGUNDO: Intentar cargar variable normal
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

	//  PRIMER PRIORITY: Verificar si es un slice
	if v.sliceProcessor.ExisteSlice(nombreVariable) {
		return &assembly.ResultadoExpresion{
			Registro:  "",
			Tipo:      "slice_name",
			EsLiteral: true,
			Valor:     nombreVariable,
		}
	}

	//  SEGUNDO: Intentar cargar variable desde el stack
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
