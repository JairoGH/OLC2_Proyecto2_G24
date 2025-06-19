package tiposDeDato

type ValorNulo struct {
}

func (s ValorNulo) Value() interface{} {
	return nil
}

func (s ValorNulo) Type() string {
	return TIPO_NULO
}

func (s ValorNulo) Copy() ValorInterno {
	return NuloPorDefecto
}

var NuloPorDefecto = &ValorNulo{}

type UnInitializedValor struct {
}

func (s UnInitializedValor) Value() interface{} {
	return nil
}

func (s UnInitializedValor) Type() string {
	return TIPO_NOINICIALIZADO
}

func (s UnInitializedValor) Copy() ValorInterno {
	return ValorNoIniPorDefecto
}

var ValorNoIniPorDefecto = &UnInitializedValor{}
