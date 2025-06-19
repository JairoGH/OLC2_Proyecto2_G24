package instrucciones

import (
	"main/parser"
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

// VisitanteDcl: visitor que maneja la primera pasada del análisis,
// registrando declaraciones de funciones y estructuras
type VisitanteDcl struct {
	parser.BaseVGrammarVisitor
	RegistroAmbito     *RegistroAmbito
	TablaError         *TablaError
	NombresEstructuras []string
}

// NewVisitanteDcl: constructor que crea un nuevo visitor de declaraciones
func NewVisitanteDcl(TablaError *TablaError) *VisitanteDcl {
	return &VisitanteDcl{
		RegistroAmbito:     NuevoRegistroAmbito(),
		TablaError:         TablaError,
		NombresEstructuras: []string{},
	}
}

// Visit: método base para visitar nodos del árbol sintáctico
func (v *VisitanteDcl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		token := val.GetSymbol()
		v.TablaError.NewErrorSintactico(token.GetLine(), token.GetColumn(), "Error de Análisis: "+val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

// VisitProgram: procesa el nodo raíz del programa visitando todas las declaraciones
func (v *VisitanteDcl) VisitProgram(ctx *parser.ProgramContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

// VisitStmt: procesa declaraciones filtrando solo funciones y estructuras
func (v *VisitanteDcl) VisitStmt(ctx *parser.StmtContext) interface{} {

	if ctx.Declarar_funcion() != nil {
		v.Visit(ctx.Declarar_funcion())
	} else if ctx.Declarar_struct() != nil {
		v.Visit(ctx.Declarar_struct())
	}

	return nil
}

// VisitFuncionDeclerada: procesa declaraciones de funciones validando parámetros y tipo de retorno
func (v *VisitanteDcl) VisitFuncionDeclerada(ctx *parser.FuncionDecleradaContext) interface{} {

	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el ámbito global")
	}

	funcName := ctx.ID().GetText()

	parametros := make([]*Parametros, 0)

	if ctx.Lista_parametros() != nil {
		parametros = v.Visit(ctx.Lista_parametros()).([]*Parametros)
	}

	// Validar que todos los parámetros sean del mismo tipo (posicional, nombrado, etc.)
	if len(parametros) > 0 {

		tipoBaseParametro := parametros[0].TipoParametro()

		for _, param := range parametros {
			if param.TipoParametro() != tipoBaseParametro {
				v.TablaError.NewErrorSemantico(param.Token, "Todos los parámetros de la función deben ser del mismo tipo")
				return nil
			}
		}
	}

	retornarTipo := tiposDeDato.TIPO_NULO
	var returnTypeToken antlr.Token = nil

	if ctx.Tipo() != nil {
		retornarTipo = ctx.Tipo().GetText()
		returnTypeToken = ctx.Tipo().GetStart()
	}

	body := ctx.AllStmt()

	function := &Funcion{
		Nombre:           funcName,
		Parametros:       parametros,
		TipoRetorno:      retornarTipo,
		Cuerpo:           body,
		AmbitoDeclaro:    v.RegistroAmbito.AmbitoActual,
		TokenTipoRetorno: returnTypeToken,
		Token:            ctx.GetStart(),
	}

	ok, msg := v.RegistroAmbito.AgregarFuncion(funcName, function)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

// VisitParamList: procesa la lista de parámetros de una función
func (v *VisitanteDcl) VisitListaParametros(ctx *parser.ListaParametrosContext) interface{} {

	parametros := make([]*Parametros, 0)

	for _, param := range ctx.AllParametro_fun() {
		parametros = append(parametros, v.Visit(param).(*Parametros))
	}

	return parametros
}

// VisitFuncParam: procesa un parámetro individual determinando nombres externo/interno
func (v *VisitanteDcl) VisitParametroFun(ctx *parser.ParametroFunContext) interface{} {

	NombreExterno := ""
	NombreInterno := ""

	// Si solo hay un ID definido
	if ctx.ID(1) == nil {
		// nombreInterno : tipo
		// _ : tipo
		NombreInterno = ctx.ID(0).GetText()
	} else {
		// nombreExterno nombreInterno : tipo
		NombreExterno = ctx.ID(0).GetText()
		NombreInterno = ctx.ID(1).GetText()
	}

	pasarPorReferencia := false

	TipoParametros := ctx.Tipo().GetText()

	return &Parametros{
		NombreExterno:      NombreExterno,
		NombreInterno:      NombreInterno,
		PasarPorReferencia: pasarPorReferencia,
		Tipo:               TipoParametros,
		Token:              ctx.GetStart(),
	}
}

// VisitDeclararStruct: registra el nombre de estructuras declaradas para validación posterior
func (v *VisitanteDcl) VisitDeclararStruct(ctx *parser.DeclararStructContext) interface{} {
	v.NombresEstructuras = append(v.NombresEstructuras, ctx.ID().GetText())
	return nil
}
