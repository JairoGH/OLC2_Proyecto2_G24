package errores

import (
	"github.com/antlr4-go/antlr/v4"
)

// ErrorPersonalizado: estrategia personalizada para manejo de errores con
// mensajes en español
type ErrorPersonalizado struct {
	*antlr.DefaultErrorStrategy
}

// NewErrorPersonalizado: constructor que crea una nueva estrategia de errores
func NewErrorPersonalizado() *ErrorPersonalizado {
	return &ErrorPersonalizado{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
	}
}

// ReportarEntradaNoCoincidente: reporta errores de tokens no esperados
func (es *ErrorPersonalizado) ReportarEntradaNoCoincidente(recognizer antlr.Parser, e *antlr.InputMisMatchException) {
	t1 := recognizer.GetTokenStream().LT(-1)
	msg := "Se recibió " + t1.GetText() + ", se esperaba " + es.GetExpectedTokens(recognizer).String()
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}
