package instrucciones

import (
	"main/parser"
	"main/tiposDeDato"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// AmbitoFuncionCache: mapa global para cachear ámbitos de funciones
var AmbitoFuncionCache = make(map[string]*AmbitoBase)
var recursionErrorReported = false

// Función para resetear el flag de error de recursión
func ResetRecursionError() {
	recursionErrorReported = false
}

// Variable global para controlar si se debe detener la recursión
var shouldStopRecursion = false

// Función para resetear el flag de detención
func ResetStopRecursion() {
	shouldStopRecursion = false
}

// Estructura para manejar errores de recursión
type RecursionError struct {
	Message string
}

// ResetAllRecursionFlags: resetea todos los flags de recursión antes de cada ejecución
func ResetAllRecursionFlags() {
	recursionErrorReported = false
	shouldStopRecursion = false
	LimpiarCacheAmbitosFunciones()
}

// Funcion: representa una función definida por el usuario con sus parámetros,
// tipo de retorno y contexto de ejecución
type Funcion struct {
	Nombre           string
	Parametros       []*Parametros
	TipoRetorno      string
	TokenTipoRetorno antlr.Token
	Cuerpo           []parser.IStmtContext
	AmbitoDeclaro    *AmbitoBase
	RetornarValor    tiposDeDato.ValorInterno
	IsMutating       bool
	AmbitoDefault    *AmbitoBase
	Token            antlr.Token
}

func (f *Funcion) Value() interface{} {
	return f
}

func (f *Funcion) Type() string {
	return tiposDeDato.TIPO_FUNCION
}

func (f *Funcion) Copy() tiposDeDato.ValorInterno {
	return f
}

// Exec: ejecuta la función validando argumentos, manejando el ámbito
// y controlando el flujo de retorno
func (f *Funcion) Exec(visitor *PatronVIsitor, args []*Argumento, token antlr.Token) {
	context := visitor.GetInstruccionesContexto()

	//Verificar si se debe detener la recursión globalmente
	if shouldStopRecursion {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	// Verificar límite de recursión ANTES de ejecutar
	if context.PilaLlamada.IncrementRecursion() {
		if !recursionErrorReported {
			context.TablaError.NewErrorEjecucion(
				token.GetLine(),
				token.GetColumn(),
				fmt.Sprintf("La función '%s' excedio el limite maximo de llamadas recursivas",
					f.Nombre))
			recursionErrorReported = true
		}

		shouldStopRecursion = true
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		context.PilaLlamada.DecrementRecursion()
		return
	}

	// Cleanup al final
	defer func() {
		context.PilaLlamada.DecrementRecursion()

		if r := recover(); r != nil {
			if item, ok := r.(*LlamadaFunciones); item != nil && ok {
				f.ValidarRetorno(context, item.RetornarValor, token)
				return
			}
			panic(r)
		}
	}()

	// Validar argumentos
	argsOk, argsMap := f.ValidarArgumentos(context, args, token)
	if !argsOk {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	// Solo usar cache para funciones no recursivas (profundidad == 1)
	initialScope := context.RegistroAmbito.AmbitoActual
	var funcionScope *AmbitoBase

	depth := context.PilaLlamada.GetRecursionDepth()

	if depth == 1 {
		// Primera llamada - usar cache si existe
		cacheKey := f.Nombre + "_scope"
		if cachedScope, exists := AmbitoFuncionCache[cacheKey]; exists {
			funcionScope = cachedScope
			funcionScope.ResetVariables()
		} else {
			funcionScope = NewAmbitoBase("func_"+f.Nombre, false)
			funcionScope.Padre = f.AmbitoDeclaro
			AmbitoFuncionCache[cacheKey] = funcionScope
		}
	} else {
		// Llamada recursiva - crear ámbito nuevo siempre
		funcionScope = NewAmbitoBase("func_"+f.Nombre+"_rec_"+fmt.Sprintf("%d", depth), false)
		funcionScope.Padre = f.AmbitoDeclaro
	}

	context.RegistroAmbito.AmbitoActual = funcionScope

	wasMutating := funcionScope.EsMutante
	funcionScope.EsMutante = f.IsMutating

	// Agregar elemento de retorno a la pila de llamadas
	funcItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo:          []string{Retornar},
	}
	context.PilaLlamada.Push(funcItem)

	// Cleanup del ámbito al final
	defer func() {
		context.PilaLlamada.Limpiar(funcItem)
		funcionScope.EsMutante = wasMutating
		context.RegistroAmbito.AmbitoActual = initialScope

		if r := recover(); r != nil {
			if item, ok := r.(*LlamadaFunciones); item != nil && ok {
				if item != funcItem {
					context.TablaError.NewErrorEjecucion(token.GetLine(), token.GetColumn(), "Return Invalido")
				}
				f.ValidarRetorno(context, item.RetornarValor, token)
				return
			}
			panic(r)
		}
		f.ValidarRetorno(context, tiposDeDato.NuloPorDefecto, token)
	}()

	// Agregar argumentos al ámbito (siempre crear nuevas variables)
	for varName, arg := range argsMap {
		if arg.esReferencia {
			if arg.VariableRef == nil {
				context.TablaError.NewErrorEjecucion(arg.Token.GetLine(), arg.Token.GetColumn(),
					"No es posible pasar por referencia un valor que no este asociado a una variable")
				f.ValidarRetorno(context, tiposDeDato.NuloPorDefecto, token)
				return
			}

			pointer := &TipoPuntero{
				VariableAsociada: arg.VariableRef,
			}
			funcionScope.AgregarVariable(varName, tiposDeDato.TIPO_PUNTERO, pointer, false, false, arg.Token)
			continue
		}
		//Siempre usar AgregarVariable (no AgregarVariableOActualizar) para recursión
		funcionScope.AgregarVariable(varName, arg.Valor.Type(), arg.Valor.Copy(), false, false, arg.Token)
	}

	// Ejecutar cuerpo de la función
	for _, stmt := range f.Cuerpo {
		if shouldStopRecursion {
			f.RetornarValor = tiposDeDato.NuloPorDefecto
			return
		}
		visitor.Visit(stmt)
	}
}

// Función para limpiar cache cuando sea necesario
func LimpiarCacheAmbitosFunciones() {
	AmbitoFuncionCache = make(map[string]*AmbitoBase)
}

// ValidarArgumentos: valida que los argumentos pasados coincidan con los parámetros
// esperados en número, tipo y forma de paso (valor/referencia)
func (f *Funcion) ValidarArgumentos(context *InstruccionesContexto, args []*Argumento, token antlr.Token) (bool, map[string]*Argumento) {

	if len(args) != len(f.Parametros) {
		context.TablaError.NewErrorSemantico(token, "Numero de Argumentos invalido")
		return false, nil
	}

	argsMap := make(map[string]*Argumento)
	finalArgsMap := make(map[string]*Argumento)

	// Verificar si hay argumentos con nombres
	hasNamedArgs := false
	for _, arg := range args {
		if arg.Nombre != "" {
			argsMap[arg.Nombre] = arg
			hasNamedArgs = true
		}
	}

	errorFound := false

	for i, param := range f.Parametros {

		var argToValidate *Argumento = nil

		if !hasNamedArgs {
			argToValidate = args[i]
		} else if param.NombreExterno == "" {
			argToValidate = argsMap[param.NombreInterno]
		} else if param.NombreExterno == "_" {
			argToValidate = args[i]
		} else {
			argToValidate = argsMap[param.NombreExterno]
		}

		if argToValidate == nil {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no especificado", param.NombreInterno))
			errorFound = true
			continue
		}

		if argToValidate.Valor.Type() != param.Tipo && param.Tipo != tiposDeDato.TIPO_ANY {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de Argumento %s invalido", param.NombreInterno))
			errorFound = true
			continue
		}

		if argToValidate.esReferencia != param.PasarPorReferencia {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no es pasado por referencia", param.NombreInterno))
			errorFound = true
			continue
		}

		finalArgsMap[param.NombreInterno] = argToValidate
	}

	if errorFound {
		return false, nil
	}

	return true, finalArgsMap
}

// ValidarRetorno: verifica que el tipo del valor de retorno coincida
// con el tipo declarado en la función
func (f *Funcion) ValidarRetorno(context *InstruccionesContexto, val tiposDeDato.ValorInterno, token antlr.Token) {

	if val.Type() != f.TipoRetorno {
		if f.TokenTipoRetorno != nil {
			context.TablaError.NewErrorSemantico(f.TokenTipoRetorno, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.TipoRetorno, val.Type()))
		} else {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.TipoRetorno, val.Type()))
		}

		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	f.RetornarValor = val
}
