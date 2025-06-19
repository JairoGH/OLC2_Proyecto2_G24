package instrucciones

import (
	"main/parser"
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

// Struct: representa la definición de una estructura (struct) en el lenguaje.
// Contiene el nombre del struct, sus atributos/campos y el token para información de posición.
type Struct struct {
	Nombre    string
	Atributos []parser.IPropiedad_structContext
	Token     antlr.Token
}

// AtributoStruct: representa un campo individual dentro de un struct.
// Define el nombre, tipo y posición de cada propiedad del struct.
type AtributoStruct struct {
	Nombre string
	Tipo   string
	Token  antlr.Token
}

// ValorInicialStruct: representa un valor de inicialización para un campo de struct.
// Se usa cuando se crea una instancia del struct con valores específicos.
type ValorInicialStruct struct {
	Nombre string
	Valor  tiposDeDato.ValorInterno
	Token  antlr.Token
}
