package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type Variable struct {
	Nombre       string
	Valor        tiposDeDato.ValorInterno
	Tipo         string
	esConstante  bool
	PuedeSerNulo bool
	Token        antlr.Token
	esPropiedad  bool
}

// TipoValidacion: verifica que el valor asignado a la variable sea compatible con su tipo declarado.
// Maneja conversiones implícitas y casos especiales como vectores vacíos y valores nulos.
func (v *Variable) TipoValidacion() (bool, string) {

	if v.Valor == tiposDeDato.ValorNoIniPorDefecto {
		return true, ""
	}

	if v.Valor == tiposDeDato.NuloPorDefecto && v.PuedeSerNulo {
		return true, ""
	}

	if v.Tipo != v.Valor.Type() {

		// Validación de tipos de vectores
		if EsTipoVector(v.Tipo) && EsTipoVector(v.Valor.Type()) {

			// Manejar vectores con tipo "[]" y vectores vacíos
			vectorValue := v.Valor.(*TipoVector)

			// Vector vacío genérico - asignar el tipo correcto
			if v.Valor.Type() == "[]" || len(vectorValue.ValorInterno) == 0 {
				vectorValue.TipoElemento = EliminarCorchetes(v.Tipo)
				vectorValue.TipoCompleto = v.Tipo
				return true, ""
			}

			// Conversión implícita de vectores para vectores no vacíos
			targetType := EliminarCorchetes(v.Tipo)
			newConvertedItems := make([]tiposDeDato.ValorInterno, 0)

			for _, item := range vectorValue.ValorInterno {
				convertedValue, ok := tiposDeDato.Casteo(targetType, item)

				if !ok {
					break
				}
				newConvertedItems = append(newConvertedItems, convertedValue)
			}

			if len(newConvertedItems) == len(vectorValue.ValorInterno) {
				// Actualizar el vector con los elementos convertidos
				vectorValue.ValorInterno = newConvertedItems
				vectorValue.TipoElemento = targetType
				vectorValue.TipoCompleto = v.Tipo
				return true, ""
			}

			msg := "No se puede asignar un vector de tipo " + v.Valor.Type() + " a una vector de tipo " + v.Tipo
			v.Valor = tiposDeDato.NuloPorDefecto
			return false, msg
		}

		// Caso especial: vector vacío genérico "[]" asignado a tipo específico
		if v.Valor.Type() == "[]" && EsTipoVector(v.Tipo) {
			if vectorValue, ok := v.Valor.(*TipoVector); ok {
				vectorValue.TipoElemento = EliminarCorchetes(v.Tipo)
				vectorValue.TipoCompleto = v.Tipo
				return true, ""
			}
		}

		// Intentar conversión implícita de tipos primitivos
		convertedValue, ok := tiposDeDato.Casteo(v.Tipo, v.Valor)

		if !ok {
			// Si la conversión falla, asignar nil y retornar error
			msg := "No se puede asignar un valor de tipo " + v.Valor.Type() + " a una variable de tipo " + v.Tipo
			v.Valor = tiposDeDato.NuloPorDefecto
			return false, msg
		}

		v.Valor = convertedValue
	}

	return true, ""
}

// AsignarVariable: asigna un nuevo valor a la variable verificando restricciones como
// constantes y propiedades en contextos no mutables. Valida la compatibilidad de tipos.
func (v *Variable) AsignarVariable(val tiposDeDato.ValorInterno, isMutatingContext bool) (bool, string) {

	if v.esConstante {
		return false, "No se puede asignar un valor a una constante"
	}

	if v.esPropiedad {
		if !isMutatingContext {
			return false, "No se puede asignar un valor a una propiedad desde un contexto de función no mutable"
		}
	}

	v.Valor = val

	return v.TipoValidacion()
}
