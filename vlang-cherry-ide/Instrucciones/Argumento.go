package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type Argumento struct {
	Nombre       string
	Valor        tiposDeDato.ValorInterno
	esReferencia bool
	Token        antlr.Token
	VariableRef  *Variable
}
