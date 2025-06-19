package instrucciones

import (
	"main/tiposDeDato"
)

// TipoMatriz: representa una matriz multidimensional basada en vectores anidados
type TipoMatriz struct {
	*TipoVector
}

// Copy: crea una copia profunda de la matriz
func (v *TipoMatriz) Copy() tiposDeDato.ValorInterno {

	internalCopy := make([]tiposDeDato.ValorInterno, len(v.ValorInterno))

	for i, item := range v.ValorInterno {
		internalCopy[i] = item.Copy()
	}

	return NewTipoMatriz(internalCopy, v.TipoCompleto, v.TipoElemento)
}

// IndicesValidos: verifica si los índices proporcionados son válidos para la matriz
func (v *TipoMatriz) IndicesValidos(indexes []int) bool {

	pivot := v.TipoVector

	for i, index := range indexes {
		if index < 0 || index >= pivot.Size() {
			return false
		}

		item := pivot.Get(index)

		switch s := item.(type) {
		case *TipoVector:
			pivot = s

			if i == len(indexes)-1 {
				return true
			}

		case *TipoMatriz:
			pivot = s.TipoVector

			if i == len(indexes)-1 {
				return true
			}

		default:
			if i != len(indexes)-1 {
				return false
			} else {
				return true
			}
		}
	}

	return false
}

// Get: obtiene un elemento de la matriz usando múltiples índices
func (v *TipoMatriz) Get(index []int) tiposDeDato.ValorInterno {

	if !v.IndicesValidos(index) {
		return nil
	}

	pivot := v.TipoVector

	for i := 0; i < len(index); i++ {
		item := pivot.Get(index[i])

		switch s := item.(type) {
		case *TipoVector:
			pivot = s

			if i == len(index)-1 {
				return pivot
			}

		case *TipoMatriz:
			pivot = s.TipoVector

			if i == len(index)-1 {
				return pivot
			}
		default:
			return item
		}
	}

	return nil
}

// Set: asigna un valor en la posición especificada por los índices
func (v *TipoMatriz) Set(index []int, value tiposDeDato.ValorInterno) bool {

	if !v.IndicesValidos(index) {
		return false
	}

	pivot := v.TipoVector

	for i := 0; i < len(index); i++ {
		item := pivot.Get(index[i])

		switch s := item.(type) {
		case *TipoVector:
			pivot = s
		case *TipoMatriz:
			pivot = s.TipoVector
		default:
			if i == len(index)-1 {
				pivot.ValorInterno[index[i]] = value
				return true
			}
		}
	}

	return false
}

// NewTipoMatriz: constructor que crea una nueva matriz
func NewTipoMatriz(elementosVector []tiposDeDato.ValorInterno, tipoCompleto, tipoElemento string) *TipoMatriz {
	vector := &TipoVector{
		ValorInterno: elementosVector,
		IndiceActual: 0,
		TipoElemento: tipoElemento,
		TipoCompleto: tipoCompleto,
		ValorTamaño:  &tiposDeDato.ValorEntero{ValorInterno: len(elementosVector)},
		EstaVacio:    &tiposDeDato.ValorBool{InternalValor: len(elementosVector) == 0},
	}

	EliminarIntegradosDeVector(elementosVector)

	return &TipoMatriz{
		TipoVector: vector,
	}
}

// EliminarIntegradosDeVector: limpia funciones integradas de elementos vector
func EliminarIntegradosDeVector(elementosVector []tiposDeDato.ValorInterno) {
	for i := 0; i < len(elementosVector); i++ {
		if item, ok := elementosVector[i].(*TipoVector); ok {
			item.TipoObjeto.AmbitoInterno.Reset()
		} else {
			break
		}
	}
}

// ElementoMatriz: representa una referencia a un elemento específico de la matriz
type ElementoMatriz struct {
	Matriz *TipoMatriz
	Indice []int
	Valor  tiposDeDato.ValorInterno
}

// TipoMatrizMulti: representa matrices multidimensionales con sintaxis [][]tipo
type TipoMatrizMulti struct {
	ValorInterno [][]tiposDeDato.ValorInterno
	TipoCompleto string
	TipoElemento string
	TipoBase     string
	Filas        int
	Columnas     []int
}

func (t *TipoMatrizMulti) Type() string {
	return t.TipoCompleto
}

func (t *TipoMatrizMulti) Value() interface{} {
	return t.ValorInterno
}

// Copy: crea una copia profunda de la matriz multidimensional
func (t *TipoMatrizMulti) Copy() tiposDeDato.ValorInterno {
	nuevaMatriz := make([][]tiposDeDato.ValorInterno, len(t.ValorInterno))
	for i, row := range t.ValorInterno {
		nuevaMatriz[i] = make([]tiposDeDato.ValorInterno, len(row))
		for j, item := range row {
			nuevaMatriz[i][j] = item.Copy()
		}
	}

	return &TipoMatrizMulti{
		ValorInterno: nuevaMatriz,
		TipoCompleto: t.TipoCompleto,
		TipoElemento: t.TipoElemento,
		TipoBase:     t.TipoBase,
		Filas:        t.Filas,
		Columnas:     append([]int{}, t.Columnas...),
	}
}

// ValidIndex: verifica si los índices de fila y columna son válidos
func (t *TipoMatrizMulti) ValidIndex(fila, columna int) bool {
	// Verificar que la fila esté en rango
	if fila < 0 || fila >= len(t.ValorInterno) {
		return false
	}

	// Verificar que la columna esté en rango para esa fila específica
	if columna < 0 || columna >= len(t.ValorInterno[fila]) {
		return false
	}

	return true
}

// Get: obtiene un elemento en la posición [fila, columna]
func (t *TipoMatrizMulti) Get(fila, columna int) tiposDeDato.ValorInterno {
	if !t.ValidIndex(fila, columna) {
		return &tiposDeDato.ValorEntero{ValorInterno: -1}
	}
	return t.ValorInterno[fila][columna]
}

// Set: asigna un valor en la posición [fila, columna]
func (t *TipoMatrizMulti) Set(fila, columna int, valor tiposDeDato.ValorInterno) {
	if t.ValidIndex(fila, columna) {
		t.ValorInterno[fila][columna] = valor
	}
}

// NewTipoMatrizMulti: constructor para crear una nueva matriz multidimensional
func NewTipoMatrizMulti(filas [][]tiposDeDato.ValorInterno, fullType, itemType, baseType string) *TipoMatrizMulti {
	columnas := make([]int, len(filas))
	for i, row := range filas {
		columnas[i] = len(row)
	}

	return &TipoMatrizMulti{
		ValorInterno: filas,
		TipoCompleto: fullType,
		TipoElemento: itemType,
		TipoBase:     baseType,
		Filas:        len(filas),
		Columnas:     columnas,
	}
}

// MatrizMultiItemReference: referencia a un elemento específico de matriz multidimensional
type MatrizMultiItemReference struct {
	Matriz  *TipoMatrizMulti
	Fila    int
	Columna int
	Valor   tiposDeDato.ValorInterno
}
