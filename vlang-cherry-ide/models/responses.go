/*
package models

// Respuesta que espera tu frontend de escritorio
type AnalysisResponse struct {
	Success       bool          `json:"success"`
	AST           string        `json:"ast"`
	Errors        []ErrorReport `json:"errors"`
	SymbolTable   []SymbolEntry `json:"symbolTable"`
	ConsoleOutput string        `json:"consoleOutput"`
	CSTSvg        string        `json:"cstSvg"` // Nuevo: SVG del backend
}

// Respuesta real del backend (lo que devuelve /compile)
type BackendResponse struct {
	Errors         []BackendError `json:"errors"`
	Output         string         `json:"output"`
	CSTSvg         string         `json:"cstSvg"`
	RegistroAmbito AmbitoGlobal   `json:"RegistroAmbito"` // CAMBIADO: ahora es un objeto
}

// Estructura del ámbito global que viene del backend
type AmbitoGlobal struct {
	AmbitoGlobal AmbitoData `json:"AmbitoGlobal"`
}

// Datos del ámbito
type AmbitoData struct {
	Name        string        `json:"Name"`
	Vars        []AmbitoVar   `json:"Vars"`
	Funcs       []AmbitoFunc  `json:"Funcs"`
	Structs     []interface{} `json:"Structs"`
	ChildScopes []interface{} `json:"ChildScopes"`
}

// Variable en el ámbito
type AmbitoVar struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	Line   int    `json:"Line"`
	Column int    `json:"Column"`
}

// Función en el ámbito
type AmbitoFunc struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	Line   int    `json:"Line"`
	Column int    `json:"Column"`
}

// Estructura de error del backend
type BackendError struct {
	Linea       int    `json:"linea"`
	Columna     int    `json:"columna"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

// Estructura de errores para el frontend
type ErrorReport struct {
	Linea       int    `json:"linea"`
	Columna     int    `json:"columna"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

// Entrada de la tabla de símbolos para el frontend
type SymbolEntry struct {
	ID         string `json:"id"`
	SymbolType string `json:"symbolType"`
	DataType   string `json:"dataType"`
	Scope      string `json:"scope"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`
}

// Request que NO se usa (el backend usa FormData, no JSON)
type AnalysisRequest struct {
	Code     string `json:"code"`
	Filename string `json:"filename"`
}
*/

package models

// Respuesta que espera tu frontend de escritorio
type AnalysisResponse struct {
	Success       bool          `json:"success"`
	AST           string        `json:"ast"`
	Errors        []ErrorReport `json:"errors"`
	SymbolTable   []SymbolEntry `json:"symbolTable"`
	ConsoleOutput string        `json:"consoleOutput"`
	CSTSvg        string        `json:"cstSvg"` // SVG original del backend
	ASTPng        []byte        `json:"-"`      // PNG convertido (no se serializa en JSON)
}

// Respuesta real del backend (lo que devuelve /compile)
type BackendResponse struct {
	Errors         []BackendError `json:"errors"`
	Output         string         `json:"output"`
	CSTSvg         string         `json:"cstSvg"`
	RegistroAmbito AmbitoGlobal   `json:"RegistroAmbito"` // CAMBIADO: ahora es un objeto
}

// Estructura del ámbito global que viene del backend
type AmbitoGlobal struct {
	AmbitoGlobal AmbitoData `json:"AmbitoGlobal"`
}

// Datos del ámbito
type AmbitoData struct {
	Name        string        `json:"Name"`
	Vars        []AmbitoVar   `json:"Vars"`
	Funcs       []AmbitoFunc  `json:"Funcs"`
	Structs     []interface{} `json:"Structs"`
	ChildScopes []interface{} `json:"ChildScopes"`
}

// Variable en el ámbito
type AmbitoVar struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	Line   int    `json:"Line"`
	Column int    `json:"Column"`
}

// Función en el ámbito
type AmbitoFunc struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	Line   int    `json:"Line"`
	Column int    `json:"Column"`
}

// Estructura de error del backend
type BackendError struct {
	Linea       int    `json:"linea"`
	Columna     int    `json:"columna"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

// Estructura de errores para el frontend
type ErrorReport struct {
	Linea       int    `json:"linea"`
	Columna     int    `json:"columna"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

// Entrada de la tabla de símbolos para el frontend
type SymbolEntry struct {
	ID         string `json:"id"`
	SymbolType string `json:"symbolType"`
	DataType   string `json:"dataType"`
	Scope      string `json:"scope"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`
}

// Request que NO se usa (el backend usa FormData, no JSON)
type AnalysisRequest struct {
	Code     string `json:"code"`
	Filename string `json:"filename"`
}
