package errores

import (
	instrucciones "main/Instrucciones"

	"github.com/antlr4-go/antlr/v4"
)

// SyntaxError: maneja errores sintácticos durante el análisis y los registra
// en la tabla de errores
type SyntaxError struct {
	*antlr.DefaultErrorListener
	TablaError *instrucciones.TablaError
}

// NewSyntaxError: constructor para crear un nuevo manejador de errores sintácticos
func NewSyntaxError(TablaError *instrucciones.TablaError) *SyntaxError {
	return &SyntaxError{
		TablaError: TablaError,
	}
}

// SyntaxError: captura errores sintácticos de ANTLR y los agrega a la tabla
func (e *SyntaxError) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, descripcion string, _ antlr.RecognitionException) {
	e.TablaError.AgregarError(
		linea,
		columna,
		descripcion,
		instrucciones.ErrorSintactico,
	)
}

// LexicalErrorListener: maneja errores léxicos durante el análisis lexicográfico
type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	TablaError *instrucciones.TablaError
}

// NewErrorLexico: constructor para crear un nuevo manejador de errores léxicos
func NewErrorLexico() *LexicalErrorListener {
	return &LexicalErrorListener{
		TablaError: instrucciones.NewTablaError(),
	}
}

// SyntaxError: captura errores léxicos de ANTLR y los agrega a la tabla
func (e *LexicalErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, descripcion string, _ antlr.RecognitionException) {
	e.TablaError.AgregarError(
		linea,
		columna,
		descripcion,
		instrucciones.ErrorLexico,
	)
}
