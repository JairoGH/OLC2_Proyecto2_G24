package tiposDeDato

type ValorEntero struct {
	ValorInterno int
}

func (i ValorEntero) Value() interface{} {
	return i.ValorInterno
}

func (i ValorEntero) Type() string {
	return TIPO_ENTERO
}

func (i ValorEntero) Copy() ValorInterno {
	return &ValorEntero{i.ValorInterno}
}
