package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

// FuncionNativaObjeto: representa una función nativa asociada a un objeto específico
// con lógica de ejecución personalizada
type FuncionNativaObjeto struct {
	*Funcion
	Objeto     *TipoObjeto
	CustomExec func(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token)
}

func (b FuncionNativaObjeto) Type() string {
	return tiposDeDato.TIPO_FUNCION
}

func (b FuncionNativaObjeto) Value() interface{} {
	return b
}

func (b FuncionNativaObjeto) Copy() tiposDeDato.ValorInterno {
	return b
}

// Exec: ejecuta la función nativa del objeto validando argumentos y delegando
// la ejecución a la función personalizada
func (f *FuncionNativaObjeto) Exec(visitor *PatronVIsitor, args []*Argumento, token antlr.Token) {

	context := visitor.GetInstruccionesContexto()

	argsOk, argsMap := f.ValidarArgumentos(context, args, token)

	if !argsOk {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	f.CustomExec(f, visitor, argsMap, token)
}

// Funciones nativas para vectores

// agregarParametros: parámetros para la función append de vectores
var agregarParametros = []*Parametros{
	{
		NombreExterno:      "_",
		NombreInterno:      "_",
		Tipo:               tiposDeDato.TIPO_ANY,
		PasarPorReferencia: false,
		Token:              nil,
	},
}

// EjecucionPersonalizada: agrega un elemento al final del vector validando tipos
func EjecucionPersonalizada(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	vector := builtinRef.Objeto.ObjetoAuxiliar.(*TipoVector)
	arg := args["_"]

	if vector.TipoElemento != arg.Valor.Type() {
		visitor.TablaError.NewErrorSemantico(arg.Token, "No se puede agregar un valor de tipo "+arg.Valor.Type()+" a un vector de tipo "+vector.TipoElemento)
		return
	}

	vector.ValorInterno = append(vector.ValorInterno, arg.Valor)
	vector.actualizarPropiedades()
}

// eliminarParametros: parámetros para la función remove de vectores
var eliminarParametros = []*Parametros{
	{
		NombreExterno:      "at",
		NombreInterno:      "at",
		Tipo:               tiposDeDato.TIPO_ENTERO,
		PasarPorReferencia: false,
		Token:              nil,
	},
}

// RemoverEjecucion: elimina un elemento del vector en la posición especificada
func RemoverEjecucion(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	vector := builtinRef.Objeto.ObjetoAuxiliar.(*TipoVector)
	arg := args["at"]

	if arg.Valor.Type() != tiposDeDato.TIPO_ENTERO {
		visitor.TablaError.NewErrorSemantico(arg.Token, "El Argumento 'at' debe ser de tipo Int")
		return
	}

	// Validar límites del vector
	if arg.Valor.Value().(int) >= vector.Size() || arg.Valor.Value().(int) < 0 {
		visitor.TablaError.NewErrorSemantico(arg.Token, "El indice esta fuera de rango")
		return
	}

	vector.ValorInterno = append(vector.ValorInterno[:arg.Valor.Value().(int)], vector.ValorInterno[arg.Valor.Value().(int)+1:]...)
	vector.actualizarPropiedades()
}

// eliminarUltimosParametros: parámetros para la función removeLast (sin parámetros)
var eliminarUltimosParametros = []*Parametros{}

// EliminarUltimaEjecucion: elimina el último elemento del vector
func EliminarUltimaEjecucion(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	vector := builtinRef.Objeto.ObjetoAuxiliar.(*TipoVector)

	if vector.Size() == 0 {
		visitor.TablaError.NewErrorSemantico(token, "El vector esta vacio y no se puede remover el ultimo elemento")
		return
	}

	vector.ValorInterno = vector.ValorInterno[:vector.Size()-1]
	vector.actualizarPropiedades()
}

// AgregarVectorNativo: registra todas las funciones y propiedades nativas de un vector
func AgregarVectorNativo(vectorRef *TipoVector) {

	vectorScope := NuevoVectorGlobal()

	vectorInternalObject := &TipoObjeto{
		AmbitoInterno:  vectorScope,
		ObjetoAuxiliar: vectorRef,
	}

	// Registrar funciones nativas
	vectorScope.AgregarFuncion("append", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: agregarParametros,
		},
		Objeto:     vectorInternalObject,
		CustomExec: EjecucionPersonalizada,
	})

	vectorScope.AgregarFuncion("remove", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: eliminarParametros,
		},
		Objeto:     vectorInternalObject,
		CustomExec: RemoverEjecucion,
	})

	vectorScope.AgregarFuncion("removeLast", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: eliminarUltimosParametros,
		},
		Objeto:     vectorInternalObject,
		CustomExec: EliminarUltimaEjecucion,
	})

	// Registrar propiedades nativas
	vectorScope.AgregarVariable("isEmpty", tiposDeDato.TIPO_BOOLEAN, vectorRef.EstaVacio, true, false, nil)
	vectorScope.AgregarVariable("count", tiposDeDato.TIPO_ENTERO, vectorRef.ValorTamaño, true, false, nil)

	vectorRef.TipoObjeto = vectorInternalObject
}
