package instrucciones

import (
	"fmt"
)

// Consola: maneja la salida de texto del programa, almacenando y mostrando
// los resultados de las operaciones de impresión
type Consola struct {
	salida string 
}

// Print: agrega texto al buffer de salida con salto de línea
func (c *Consola) Print(s string) {
	c.salida += s + "\n"
}

// Show: muestra todo el contenido del buffer en la consola del sistema
func (c *Consola) Show() {
	fmt.Println(c.salida)
}

// Clear: limpia completamente el buffer de salida
func (c *Consola) Clear() {
	c.salida = ""
}

// NewConsola: crea una nueva instancia de consola con buffer vacío
func NewConsola() *Consola {
	return &Consola{}
}

// GetSalida: retorna el contenido actual del buffer de salida
func (c *Consola) GetSalida() string {
	return c.salida
}
