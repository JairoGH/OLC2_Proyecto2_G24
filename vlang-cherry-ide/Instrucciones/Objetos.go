package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

// TipoObjeto: representa una instancia de un struct con su ámbito interno
// y funcionalidad asociada
type TipoObjeto struct {
	AmbitoInterno  *AmbitoBase
	ObjetoAuxiliar interface{}
	TipoConcreto   string
	Visitante      *PatronVIsitor
	Token          antlr.Token
}

func (o TipoObjeto) Value() interface{} {
	return o
}

func (o TipoObjeto) Type() string {
	if o.TipoConcreto != "" {
		return o.TipoConcreto
	}

	return tiposDeDato.TIPO_OBJECTO
}

// Copy: crea una copia del objeto con todos sus campos
func (o *TipoObjeto) Copy() tiposDeDato.ValorInterno {
	args := make([]*Argumento, 0)

	for _, prop := range o.AmbitoInterno.Variables {
		args = append(args, &Argumento{
			Nombre: prop.Nombre,
			Valor:  prop.Valor,
		})
	}

	return NewTipoObjeto(o.Visitante, o.TipoConcreto, o.Token, args, true)
}

// NuevoTipoObjeto: constructor que crea una nueva instancia de struct validando
// campos y argumentos del constructor
func NewTipoObjeto(visitante *PatronVIsitor, targetStruct string, tokenObjetivo antlr.Token, args []*Argumento, allowReinitialize bool) tiposDeDato.ValorInterno {

	// Verificar si la estructura existe
	plantillaStruct, msg := visitante.RegistroAmbito.AmbitoGlobal.GetEstructura(targetStruct)

	if plantillaStruct == nil {
		visitante.TablaError.NewErrorSemantico(tokenObjetivo, msg)
		return tiposDeDato.NuloPorDefecto
	}

	ambitoInterno := NuevoAmbitoGlobal()

	ambitoAnterior := visitante.RegistroAmbito.AmbitoActual
	visitante.RegistroAmbito.AmbitoActual = ambitoInterno

	defer func() {
		// Restaurar ámbito
		visitante.RegistroAmbito.AmbitoActual = ambitoAnterior
	}()

	// Agregar campos al ámbito interno con verificaciones
	for _, field := range plantillaStruct.Atributos {
		if field != nil { // Verificar que field no sea nil
			result := visitante.Visit(field)
			if result == nil {
				// Si hay error procesando el campo, continuar con el siguiente
				continue
			}
		}
	}

	// Crear mapa de argumentos
	mapaArgumentos := make(map[string]*Argumento)

	for _, arg := range args {
		// repeat arg
		if _, ok := mapaArgumentos[arg.Nombre]; ok {
			visitante.TablaError.NewErrorSemantico(arg.Token, "El argumento "+arg.Nombre+" ya fue definido")
			return tiposDeDato.NuloPorDefecto
		}

		mapaArgumentos[arg.Nombre] = arg
	}

	// Validar argumentos del constructor
	eraConstante := false
	argumentosUsados := make(map[string]bool)

	for _, prop := range ambitoInterno.Variables {
		arg, found := mapaArgumentos[prop.Nombre]

		if !found {
			if prop.Valor == tiposDeDato.ValorNoIniPorDefecto {
				visitante.TablaError.NewErrorSemantico(tokenObjetivo, "El campo "+prop.Nombre+" no fue inicializado en el constructor")
				return tiposDeDato.NuloPorDefecto
			}

			continue
		}

		// El argumento existe
		if prop.esConstante {
			if (prop.Valor != tiposDeDato.ValorNoIniPorDefecto) && !allowReinitialize {
				visitante.TablaError.NewErrorSemantico(tokenObjetivo, "El campo "+prop.Nombre+" es inmutable y ya fue inicializado")
				return tiposDeDato.NuloPorDefecto
			}

			eraConstante = true
			prop.esConstante = false
		}

		var lanzarError bool = false
		var msg string = ""
		var valorAsignar tiposDeDato.ValorInterno = arg.Valor

		// Soporte para punteros
		if arg.esReferencia {
			if arg.VariableRef == nil {
				msg = "No es posible pasar por referencia un valor que no este asociado a una variable"
				lanzarError = true
			}

			// Crear el puntero
			valorAsignar = &TipoPuntero{
				VariableAsociada: arg.VariableRef,
			}
		}

		if !lanzarError {
			lanzarError, msg = prop.AsignarVariable(valorAsignar, false)
		}

		if eraConstante {
			prop.esConstante = true
			eraConstante = false
		}

		if !lanzarError {
			visitante.TablaError.NewErrorSemantico(tokenObjetivo, msg)
			return tiposDeDato.NuloPorDefecto
		}

		argumentosUsados[prop.Nombre] = true
	}

	// Validar argumentos no utilizados
	for _, arg := range args {
		if _, ok := argumentosUsados[arg.Nombre]; !ok {
			visitante.TablaError.NewErrorSemantico(arg.Token, "El argumento "+arg.Nombre+" no es utilizado en el constructor")
		}
	}

	// Marcar propiedades como mutables
	for _, prop := range ambitoInterno.Variables {
		prop.esPropiedad = true
	}

	// Objeto self
	selfObject := &TipoObjeto{
		AmbitoInterno: ambitoInterno,
		TipoConcreto:  tiposDeDato.TIPO_PROPIO,
	}

	ambitoInternoInstancia := NuevoAmbitoGlobal()

	ambitoInternoInstancia.AgregarVariable("self", tiposDeDato.TIPO_PROPIO, selfObject, true, false, nil)

	// Hacer que las funciones usen el ámbito de la instancia
	for _, function := range ambitoInterno.Funciones {
		f, ok := function.(*Funcion)

		if !ok {
			continue
		}

		f.AmbitoDefault = ambitoInternoInstancia
	}

	// Crear instancia
	return &TipoObjeto{
		AmbitoInterno: ambitoInterno,
		TipoConcreto:  targetStruct,
		Visitante:     visitante,
		Token:         tokenObjetivo,
	}
}

// ArgumentoValidoEstructura: verifica que todos los argumentos tengan nombre
// (requerido para constructores de structs)
func ArgumentoValidoEstructura(argumentos []*Argumento) bool {
	for _, a := range argumentos {

		if a.Nombre == "" {
			return false
		}
	}
	return true
}
