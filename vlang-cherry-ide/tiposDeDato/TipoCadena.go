package tiposDeDato

type ValorCadena struct {
	ValorInterno string
}

func (s ValorCadena) Value() interface{} {
	return s.ValorInterno
}

func (s ValorCadena) Type() string {
	return TIPO_CADENA
}

func (s ValorCadena) Copy() ValorInterno {
	return &ValorCadena{ValorInterno: s.ValorInterno}
}
