package tiposDeDato

type ValorBool struct {
	InternalValor bool
}

func (b ValorBool) Value() interface{} {
	return b.InternalValor
}

func (b ValorBool) Type() string {
	return TIPO_BOOLEAN
}

func (b ValorBool) Copy() ValorInterno {
	return &ValorBool{b.InternalValor}
}
