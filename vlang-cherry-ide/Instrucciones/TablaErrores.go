package instrucciones

import "github.com/antlr4-go/antlr/v4"

const (
	ErrorLexico     = "Error Léxico"
	ErrorSintactico = "Error Sintáctico"
	ErrorSemantico  = "Error Semántico"
	ErrorEjecucion  = "Error de Ejecución"
)

// Error: representa un error encontrado durante el análisis o ejecución
// con información de ubicación y tipo
type Error struct {
	Linea       int    
	Columna     int    
	Descripcion string 
	Tipo        string 
}

// TablaError: almacena y maneja todos los errores encontrados durante
// el proceso de compilación/interpretación
type TablaError struct {
	Errores []Error 
}

// AgregarError: agrega un nuevo error a la tabla con información de ubicación
func (et *TablaError) AgregarError(Linea int, Columna int, descripcion string, TipoError string) {
	et.Errores = append(et.Errores, Error{Linea, Columna, descripcion, TipoError})
}

// NewLErrorLexico: crea y agrega un error léxico
func (et *TablaError) NewLErrorLexico(Linea int, Columna int, descripcion string) {
	et.AgregarError(Linea, Columna, descripcion, ErrorLexico)
}

// NewErrorSintactico: crea y agrega un error sintáctico
func (et *TablaError) NewErrorSintactico(Linea int, Columna int, descripcion string) {
	et.AgregarError(Linea, Columna, descripcion, ErrorSintactico)
}

// NewErrorSemantico: crea y agrega un error semántico usando información del token
func (et *TablaError) NewErrorSemantico(token antlr.Token, descripcion string) {
	et.AgregarError(token.GetLine(), token.GetColumn(), descripcion, ErrorSemantico)
}

// NewErrorEjecucion: crea y agrega un error de tiempo de ejecución
func (et *TablaError) NewErrorEjecucion(Linea int, Columna int, descripcion string) {
	et.AgregarError(Linea, Columna, descripcion, ErrorEjecucion)
}

// NewTablaError: constructor que crea una nueva tabla de errores vacía
func NewTablaError() *TablaError {
	return &TablaError{}
}
