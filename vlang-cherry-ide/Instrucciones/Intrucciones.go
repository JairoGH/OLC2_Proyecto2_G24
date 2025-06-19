package instrucciones

// InstruccionesContexto: contexto de ejecución que contiene todos los componentes
// necesarios para interpretar y ejecutar instrucciones del lenguaje
type InstruccionesContexto struct {
	// La consola es la salida del REPL
	Consola *Consola
	// El registro de ámbito maneja el ámbito actual del REPL
	RegistroAmbito *RegistroAmbito
	// La pila de llamadas almacena elementos que pueden ser interrumpidos, continuados o retornados
	PilaLlamada *PilaLlamada
	// La tabla de errores almacena todos los errores encontrados
	TablaError *TablaError
}
