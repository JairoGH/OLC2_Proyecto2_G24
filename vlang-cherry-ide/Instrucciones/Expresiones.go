package instrucciones

import (
	"main/tiposDeDato"
)

// Definición de tipos de funciones para evaluación y conversión
type funcionEvaluar func(tiposDeDato.ValorInterno, tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno)
type funcionConv func(tiposDeDato.ValorInterno) tiposDeDato.ValorInterno

type VerificacionBinaria struct {
	TipoIzq  string
	TipoDcha string
	ConvIzq  funcionConv
	ConvDcha funcionConv
	Evaluar  funcionEvaluar
}

type MetodoBinario struct {
	Nombre               string
	Validaciones         []VerificacionBinaria
	Viceversa            bool
	EvaluacionPorDefecto funcionEvaluar
}

// ValidarExp: valida y evalúa una expresión binaria verificando tipos compatibles
func (s *MetodoBinario) ValidarExp(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

	// Los valores nulos no permiten operaciones
	if left.Type() == tiposDeDato.TIPO_NULO || right.Type() == tiposDeDato.TIPO_NULO {
		return false, "No es posible realizar operaciones con valores nulos", tiposDeDato.NuloPorDefecto
	}

	for _, valid := range s.Validaciones {

		if valid.TipoIzq == left.Type() && valid.TipoDcha == right.Type() {

			if valid.ConvIzq != nil {
				left = valid.ConvIzq(left)
			}

			if valid.ConvDcha != nil {
				right = valid.ConvDcha(right)
			}

			if valid.Evaluar != nil {
				return valid.Evaluar(left, right)
			}

			return s.EvaluacionPorDefecto(left, right)
		}

		if s.Viceversa && valid.TipoIzq == right.Type() && valid.TipoDcha == left.Type() {

			if valid.ConvIzq != nil {
				right = valid.ConvIzq(right)
			}

			if valid.ConvDcha != nil {
				left = valid.ConvDcha(left)
			}

			if valid.Evaluar != nil {
				return valid.Evaluar(left, right)
			}

			return s.EvaluacionPorDefecto(left, right)
		}

	}

	msg := "No es posible realizar la operación '" + s.Nombre + "' con los tipos '" + left.Type() + "' y '" + right.Type() + "'"

	return false, msg, tiposDeDato.NuloPorDefecto
}

// Operadores aritméticos

// SumaExpresion: maneja la suma entre enteros, decimales y concatenación de cadenas
var SumaExpresion = MetodoBinario{
	Nombre:               "+",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: left.(*tiposDeDato.ValorEntero).ValorInterno + right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor + right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor + right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorCadena{
					ValorInterno: left.(*tiposDeDato.ValorCadena).ValorInterno + right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
		{
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorCadena{}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorCadena{
					ValorInterno: left.(*tiposDeDato.ValorCadena).ValorInterno + right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
	},
}

// RestaExpresion: maneja la resta entre enteros y decimales
var RestaExpresion = MetodoBinario{
	Nombre:               "-",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: left.(*tiposDeDato.ValorEntero).ValorInterno - right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor - right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor - right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// MultipliacionExpresion: maneja la multiplicación entre enteros y decimales
var MultipliacionExpresion = MetodoBinario{
	Nombre:               "*",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: left.(*tiposDeDato.ValorEntero).ValorInterno * right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor * right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor * right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// DivisionExpresion: maneja la división entre enteros y decimales con validación de división por cero
var DivisionExpresion = MetodoBinario{
	Nombre:               "/",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorEntero).ValorInterno == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: left.(*tiposDeDato.ValorEntero).ValorInterno / right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorDecimal).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor / right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorDecimal).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor / right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// ModuloExpresion: maneja el módulo entre enteros con validación de división por cero
var ModuloExpresion = MetodoBinario{
	Nombre:               "%",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorEntero).ValorInterno == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: left.(*tiposDeDato.ValorEntero).ValorInterno % right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
	},
}

// Operadores de comparación

// ExpresionTipoIgual: función generadora para operadores de igualdad y desigualdad
func ExpresionTipoIgual(Nombre string, eval funcionEvaluar) MetodoBinario {
	return MetodoBinario{
		Nombre:               Nombre,
		Viceversa:            true,
		EvaluacionPorDefecto: eval,
		Validaciones: []VerificacionBinaria{
			{
				TipoIzq:  tiposDeDato.TIPO_ENTERO,
				TipoDcha: tiposDeDato.TIPO_ENTERO,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_DECIMAL,
				TipoDcha: tiposDeDato.TIPO_DECIMAL,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_BOOLEAN,
				TipoDcha: tiposDeDato.TIPO_BOOLEAN,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_CADENA,
				TipoDcha: tiposDeDato.TIPO_CADENA,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			// Comparaciones cruzadas int == float64
			{
				TipoIzq:  tiposDeDato.TIPO_ENTERO,
				TipoDcha: tiposDeDato.TIPO_DECIMAL,
				ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
					return &tiposDeDato.ValorDecimal{
						InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
					}
				},
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_DECIMAL,
				TipoDcha: tiposDeDato.TIPO_ENTERO,
				ConvIzq:  nil,
				ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
					return &tiposDeDato.ValorDecimal{
						InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
					}
				},
				Evaluar: nil,
			},
		},
	}
}

var IgualExpresion = ExpresionTipoIgual("==", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.Value() == right.Value(),
	}
})

var NotExpresion = ExpresionTipoIgual("!=", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.Value() != right.Value(),
	}
})

// MenorExpresion: maneja comparaciones menor que entre enteros, decimales y cadenas
var MenorExpresion = MetodoBinario{
	Nombre:               "<",
	Viceversa:            false, 
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).ValorInterno < right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor < right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor < right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor < right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).ValorInterno < right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
	},
}

// MenorQueExpresion: maneja comparaciones menor o igual que
var MenorQueExpresion = MetodoBinario{
	Nombre:               "<=",
	Viceversa:            false, 
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).ValorInterno <= right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor <= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor <= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor <= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).ValorInterno <= right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
	},
}

// MayorExpresion: maneja comparaciones mayor que
var MayorExpresion = MetodoBinario{
	Nombre:               ">",
	Viceversa:            false, 
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).ValorInterno > right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor > right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor > right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor > right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).ValorInterno > right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
	},
}

// MayorQueExpresion: maneja comparaciones mayor o igual que
var MayorQueExpresion = MetodoBinario{
	Nombre:               ">=",
	Viceversa:            false, 
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).ValorInterno >= right.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor >= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor >= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).ValorInterno),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor >= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).ValorInterno >= right.(*tiposDeDato.ValorCadena).ValorInterno,
				}
			},
		},
	},
}

// Operadores lógicos

// LogicaBinariaGenerica: función generadora para operadores lógicos AND y OR
func LogicaBinariaGenerica(Nombre string, eval funcionEvaluar) MetodoBinario {
	return MetodoBinario{
		Nombre:               Nombre,
		Viceversa:            true,
		EvaluacionPorDefecto: eval,
		Validaciones: []VerificacionBinaria{
			{
				TipoIzq:  tiposDeDato.TIPO_BOOLEAN,
				TipoDcha: tiposDeDato.TIPO_BOOLEAN,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
		},
	}
}

var AndExpresion = LogicaBinariaGenerica("&&", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.(*tiposDeDato.ValorBool).InternalValor && right.(*tiposDeDato.ValorBool).InternalValor,
	}
})

var OrExpresion = LogicaBinariaGenerica("||", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.(*tiposDeDato.ValorBool).InternalValor || right.(*tiposDeDato.ValorBool).InternalValor,
	}
})

// Mapa con todas las expresiones binarias disponibles
var ExpresionBinaria = map[string]MetodoBinario{
	"+":  SumaExpresion,
	"-":  RestaExpresion,
	"*":  MultipliacionExpresion,
	"/":  DivisionExpresion,
	"%":  ModuloExpresion,
	"==": IgualExpresion,
	"!=": NotExpresion,
	"<":  MenorExpresion,
	"<=": MenorQueExpresion,
	">":  MayorExpresion,
	">=": MayorQueExpresion,
	"&&": AndExpresion,
	"||": OrExpresion,
}

// Expresiones unarias

type ValidarUnaria struct {
	Tipo       string
	Conversion funcionConv
	Eval       funcionEvaluar
}

type ExpresionUnaria struct {
	Nombre               string
	Validaciones         []ValidarUnaria
	EvaluacionPorDefecto funcionEvaluar
}

// ValidarExp: valida y evalúa una expresión unaria
func (s *ExpresionUnaria) ValidarExp(val tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

	if val.Type() == tiposDeDato.TIPO_NULO {
		return false, "No es posible realizar operaciones con valores nulos", tiposDeDato.NuloPorDefecto
	}

	for _, valid := range s.Validaciones {

		if valid.Tipo == val.Type() {

			if valid.Conversion != nil {
				val = valid.Conversion(val)
			}

			if valid.Eval != nil {
				return valid.Eval(val, nil)
			}

			return s.EvaluacionPorDefecto(val, nil)
		}

	}

	msg := "No es posible realizar la operación '" + s.Nombre + "' con el tipo '" + val.Type() + "'"

	return false, msg, tiposDeDato.NuloPorDefecto
}

// ExpresionNot: maneja la negación lógica (!valor)
var ExpresionNot = ExpresionUnaria{
	Nombre:               "!",
	EvaluacionPorDefecto: nil,
	Validaciones: []ValidarUnaria{
		{
			Tipo:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: !i1.(*tiposDeDato.ValorBool).InternalValor,
				}
			},
		},
	},
}

// ExpresionMenos: maneja la negación aritmética (-valor)
var ExpresionMenos = ExpresionUnaria{
	Nombre:               "-",
	EvaluacionPorDefecto: nil,
	Validaciones: []ValidarUnaria{
		{
			Tipo:       tiposDeDato.TIPO_ENTERO,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					ValorInterno: -i1.(*tiposDeDato.ValorEntero).ValorInterno,
				}
			},
		},
		{
			Tipo:       tiposDeDato.TIPO_DECIMAL,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: -i1.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

var ExpresionesUnarias = map[string]ExpresionUnaria{
	"!": ExpresionNot,
	"-": ExpresionMenos,
}

// Evaluación de cortocircuito

// RetornoAnd: maneja la evaluación de cortocircuito para AND (&&)
var RetornoAnd = ExpresionUnaria{
	Nombre: "&&",
	Validaciones: []ValidarUnaria{
		{
			Tipo:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if !i1.(*tiposDeDato.ValorBool).InternalValor {
					return true, "", &tiposDeDato.ValorBool{
						InternalValor: false,
					}
				}

				return false, "", nil
			},
		},
	},
}

// RetornoOr: maneja la evaluación de cortocircuito para OR (||)
var RetornoOr = ExpresionUnaria{
	Nombre: "||",
	Validaciones: []ValidarUnaria{
		{
			Tipo:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if i1.(*tiposDeDato.ValorBool).InternalValor {
					return true, "", &tiposDeDato.ValorBool{
						InternalValor: true,
					}
				}

				return false, "", nil
			},
		},
	},
}

var RetornoAnticipado = map[string]ExpresionUnaria{
	"&&": RetornoAnd,
	"||": RetornoOr,
}
