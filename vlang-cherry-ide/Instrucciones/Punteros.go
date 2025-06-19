package instrucciones

import "main/tiposDeDato"

// TipoPuntero: representa un puntero que apunta a una variable específica,
// permitiendo el paso por referencia y modificación de la variable original
type TipoPuntero struct {
	VariableAsociada *Variable 
}

func (p TipoPuntero) Value() interface{} {
	return p
}

func (p TipoPuntero) Type() string {
	return tiposDeDato.TIPO_PUNTERO
}

// Copy: retorna el mismo puntero (los punteros no se copian profundamente)
func (p TipoPuntero) Copy() tiposDeDato.ValorInterno {
	return p
}
