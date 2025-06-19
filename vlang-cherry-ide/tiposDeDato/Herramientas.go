package tiposDeDato

func PRIMITIVO(t string) bool {
	switch t {
	case TIPO_BOOLEAN, TIPO_ENTERO, TIPO_DECIMAL, TIPO_CADENA, TIPO_NULO:
		return true
	default:
		return false
	}
}

func Casteo(TipoCasteo string, Valor ValorInterno) (ValorInterno, bool) {

	if TipoCasteo == Valor.Type() {
		return Valor, true
	}

	// conversión implícita

	// 1. Los enteros pueden convertirse a decimales
	if TipoCasteo == TIPO_DECIMAL && Valor.Type() == TIPO_ENTERO {
		return &ValorDecimal{
			InternalValor: float64(Valor.(*ValorEntero).ValorInterno),
		}, true
	}

	return nil, false
}
