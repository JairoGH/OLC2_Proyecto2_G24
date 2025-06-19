package tiposDeDato

const (
	TIPO_ENTERO         = "int"
	TIPO_DECIMAL        = "float64"
	TIPO_CADENA         = "string"
	TIPO_BOOLEAN        = "bool"
	TIPO_NULO           = "nil"
	TIPO_FUNCIONNATIVA  = "builtin_function"
	TIPO_FUNCION        = "function"
	TIPO_VECTOR         = "vector"
	TIPO_OBJECTO        = "object"
	TIPO_ANY            = "any"
	TIPO_PUNTERO        = "pointer"
	TIPO_MATRIZ         = "matrix"
	TIPO_METODO         = "metodo"
	TIPO_PROPIO         = "self"
	TIPO_NOINICIALIZADO = "uninitialized"
)

// Valor Interno de los Objetos
type ValorInterno interface {
	Value() interface{}
	Type() string
	Copy() ValorInterno
}
