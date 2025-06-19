package tiposDeDato

type ValorDecimal struct {
	InternalValor float64
}

func (d ValorDecimal) Value() interface{} {
	return d.InternalValor
}

func (d ValorDecimal) Type() string {
	return TIPO_DECIMAL 
}

func (d ValorDecimal) Copy() ValorInterno {
	return &ValorDecimal{d.InternalValor}
}
