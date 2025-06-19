package instrucciones

import "main/tiposDeDato"

// TipoVector: representa un vector dinámico con funcionalidad de iteración
// y métodos nativos integrados
type TipoVector struct {
	*TipoObjeto

	ValorInterno []tiposDeDato.ValorInterno
	IndiceActual int
	TipoElemento string
	TipoCompleto string
	ValorTamaño  *tiposDeDato.ValorEntero
	EstaVacio    *tiposDeDato.ValorBool
}

func (v TipoVector) Value() interface{} {
	return v
}

func (v TipoVector) Type() string {
	return v.TipoCompleto
}

// Size: retorna el número de elementos en el vector
func (v TipoVector) Size() int {
	return len(v.ValorInterno)
}

// IndiceValido: verifica si un índice está dentro del rango válido
func (v TipoVector) IndiceValido(index int) bool {

	if index < 0 || index >= len(v.ValorInterno) {
		return false
	}

	return true

}

// Get: obtiene el elemento en la posición especificada
func (v TipoVector) Get(index int) tiposDeDato.ValorInterno {
	return v.ValorInterno[index]
}

// Next: avanza al siguiente elemento en la iteración
func (v *TipoVector) Next() bool {
	if v.IndiceActual < len(v.ValorInterno) {
		v.IndiceActual++
		return true
	}
	return false
}

// Actual: retorna el elemento actual en la iteración
func (v *TipoVector) Actual() tiposDeDato.ValorInterno {
	return v.ValorInterno[v.IndiceActual]
}

// Reset: reinicia el índice de iteración
func (v *TipoVector) Reset() {
	v.IndiceActual = 0
}

// Copy: crea una copia profunda del vector
func (v *TipoVector) Copy() tiposDeDato.ValorInterno {

	copiainterna := make([]tiposDeDato.ValorInterno, len(v.ValorInterno))

	for i, item := range v.ValorInterno {
		copiainterna[i] = item.Copy()
	}

	return NewTipoVector(copiainterna, v.TipoCompleto, v.TipoElemento)

}

// actualizarPropiedades: actualiza las propiedades de tamaño y estado vacío
func (v *TipoVector) actualizarPropiedades() {

	v.ValorTamaño.ValorInterno = len(v.ValorInterno)
	v.EstaVacio.InternalValor = len(v.ValorInterno) == 0

}

// NewTipoVector: constructor que crea un nuevo vector con métodos nativos
func NewTipoVector(elementosVector []tiposDeDato.ValorInterno, tipoCompleto, tipoElemento string) *TipoVector {
	vector := &TipoVector{
		TipoObjeto:   &TipoObjeto{},
		ValorInterno: elementosVector,
		IndiceActual: 0,
		TipoElemento: tipoElemento,
		TipoCompleto: tipoCompleto,
		ValorTamaño:  &tiposDeDato.ValorEntero{ValorInterno: len(elementosVector)},
		EstaVacio:    &tiposDeDato.ValorBool{InternalValor: len(elementosVector) == 0},
	}

	AgregarVectorNativo(vector)

	return vector
}

// TipoVectorVacio: instancia global de vector vacío
var TipoVectorVacio = &TipoVector{
	TipoObjeto:   &TipoObjeto{},
	ValorInterno: []tiposDeDato.ValorInterno{},
	IndiceActual: 0,
	TipoElemento: tiposDeDato.TIPO_ANY,
	TipoCompleto: "[" + tiposDeDato.TIPO_ANY + "]",
	ValorTamaño:  &tiposDeDato.ValorEntero{ValorInterno: 0},
	EstaVacio:    &tiposDeDato.ValorBool{InternalValor: true},
}

// VectorItemReference: referencia a un elemento específico del vector
type VectorItemReference struct {
	Vector *TipoVector
	Indice int
	Valor  tiposDeDato.ValorInterno
}
