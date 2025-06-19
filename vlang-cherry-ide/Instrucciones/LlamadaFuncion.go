package instrucciones

import "main/tiposDeDato"

const (
	Detener   = "break"
	Continuar = "continue"
	Retornar  = "return"
	MAX_RECURSION_DEPTH = 15  
)

// LlamadaFunciones: representa un elemento en la pila de llamadas para controlar
// el flujo de ejecución (break, continue, return)
type LlamadaFunciones struct {
	RetornarValor tiposDeDato.ValorInterno
	Tipo          []string
	Accion        string
}

// EsDeTipo: verifica si la llamada es de un tipo específico (break, continue, return)
func (lf *LlamadaFunciones) EsDeTipo(t string) bool {
	for _, i := range lf.Tipo {
		if i == t {
			return true
		}
	}
	return false
}

// EsAccion: verifica si la llamada tiene una acción específica
func (lf *LlamadaFunciones) EsAccion(a string) bool {
	return lf.Accion == a
}

// ResetAccion: limpia la acción de la llamada
func (lf *LlamadaFunciones) ResetAccion() {
	lf.Accion = ""
}

// PilaLlamada: estructura de pila para manejar llamadas de funciones y control de flujo
type PilaLlamada struct {
	Items []*LlamadaFunciones
	recursionDepth int
}

// Push: agrega un elemento al tope de la pila
func (pl *PilaLlamada) Push(item *LlamadaFunciones) {
	pl.Items = append(pl.Items, item)
}

// Pop: remueve y retorna el elemento del tope de la pila
func (pl *PilaLlamada) Pop() *LlamadaFunciones {
	item := pl.Items[len(pl.Items)-1]
	pl.Items = pl.Items[:len(pl.Items)-1]
	return item
}

// Peek: retorna el elemento del tope sin removerlo
func (pl *PilaLlamada) Peek() *LlamadaFunciones {
	return pl.Items[len(pl.Items)-1]
}

// In: verifica si un elemento específico está en la pila
func (pl *PilaLlamada) In(item *LlamadaFunciones) bool {
	for _, i := range pl.Items {
		if i == item {
			return true
		}
	}
	return false
}

// Limpiar: elimina elementos de la pila hasta encontrar el elemento buscado
func (pl *PilaLlamada) Limpiar(item *LlamadaFunciones) {
	if !pl.In(item) {
		return
	}

	for {
		peek := pl.Pop()
		if peek == item {
			break
		}
	}
}

// PuedeContinuar: verifica si se puede ejecutar continue en el contexto actual
// Continue solo puede estar en un bucle, no puede atravesar funciones
func (pl *PilaLlamada) PuedeContinuar() (bool, *LlamadaFunciones) {
	start := len(pl.Items) - 1

	for i := start; i >= 0; i-- {
		if pl.Items[i].EsDeTipo(Continuar) {
			return true, pl.Items[i]
		}

		if pl.Items[i].EsDeTipo(Retornar) {
			return false, nil
		}
	}

	return false, nil
}

// PuedeDetener: verifica si se puede ejecutar break (debe ser el elemento superior)
func (pl *PilaLlamada) PuedeDetener() (bool, *LlamadaFunciones) {
	if len(pl.Items) == 0 {
		return false, nil
	}

	if pl.Items[len(pl.Items)-1].EsDeTipo(Detener) {
		return true, pl.Items[len(pl.Items)-1]
	}

	return false, nil
}

// PuedeRetornar: verifica si se puede ejecutar return (puede interferir con cualquier elemento)
func (pl *PilaLlamada) PuedeRetornar() (bool, *LlamadaFunciones) {
	for i := len(pl.Items) - 1; i >= 0; i-- {
		if pl.Items[i].EsDeTipo(Retornar) {
			return true, pl.Items[i]
		}
	}

	return false, nil
}

// Len: retorna el número de elementos en la pila
func (pl *PilaLlamada) Len() int {
	return len(pl.Items)
}

// NUEVOS métodos para control de recursión
func (pl *PilaLlamada) IncrementRecursion() bool {
    pl.recursionDepth++
    return pl.recursionDepth > MAX_RECURSION_DEPTH
}

func (pl *PilaLlamada) DecrementRecursion() {
    if pl.recursionDepth > 0 {
        pl.recursionDepth--
    }
}

func (pl *PilaLlamada) GetRecursionDepth() int {
    return pl.recursionDepth
}

func (pl *PilaLlamada) ResetRecursion() {
    pl.recursionDepth = 0
}

// NuevaLlamadaFuncion: constructor que crea una nueva pila de llamadas vacía
func NuevaLlamadaFuncion() *PilaLlamada {
    return &PilaLlamada{
        Items: []*LlamadaFunciones{}, 
        recursionDepth: 0, 
    }
}
