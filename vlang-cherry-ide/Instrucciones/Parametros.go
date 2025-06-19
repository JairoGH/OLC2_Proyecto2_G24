package instrucciones

import "github.com/antlr4-go/antlr/v4"

const (
	ParametroExterno = iota
	ParametroNombreUnico
	ParametroPosicional
)

// Parametros: representa un parámetro de función con su nombre externo e interno,
// tipo y forma de paso (por valor o referencia)
type Parametros struct {
	NombreExterno      string      
	NombreInterno      string      
	Tipo               string      
	PasarPorReferencia bool       
	Token              antlr.Token 
}

// TipoParametro: determina el tipo de parámetro basado en los nombres externos e internos
func (p *Parametros) TipoParametro() int {
	if p.NombreExterno != "" {
		if p.NombreExterno == "_" {
			return ParametroPosicional
		} else {
			return ParametroExterno
		}
	} else {
		return ParametroNombreUnico
	}
}
