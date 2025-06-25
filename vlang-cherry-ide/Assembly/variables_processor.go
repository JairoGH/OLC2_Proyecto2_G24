package assembly

import (
	"fmt"
)

// Simbolo representa una variable en la tabla de símbolos
type Simbolo struct {
	Nombre       string
	Tipo         string
	Ambito       string
	Offset       int // Posición en el stack (relativo a sp)
	EsGlobal     bool
	Inicializada bool
	Linea        int
	Columna      int
}

// TablaSimbolos maneja todas las variables y sus ámbitos
type TablaSimbolos struct {
	Simbolos     map[string]*Simbolo // key: ambito:nombre
	AmbitoActual string
	PilaAmbitos  []string // Stack de ámbitos
	StackActual  int      // Cuántas variables hay en el stack actual
}

// VariablesProcessor maneja declaraciones y asignaciones de variables
type VariablesProcessor struct {
	armGen               *ARMGenerator
	expresionesProcessor *ExpresionesProcessor
	tablaSimbolos        *TablaSimbolos
	contadorAmbito       int
}

// NewVariablesProcessor crea un nuevo procesador de variables
func NewVariablesProcessor(armGen *ARMGenerator, expProc *ExpresionesProcessor) *VariablesProcessor {
	return &VariablesProcessor{
		armGen:               armGen,
		expresionesProcessor: expProc,
		tablaSimbolos: &TablaSimbolos{
			Simbolos:     make(map[string]*Simbolo),
			AmbitoActual: "global",
			PilaAmbitos:  []string{"global"},
			StackActual:  0,
		},
		contadorAmbito: 0,
	}
}

// ============= GESTIÓN DE ÁMBITOS =============

func (vp *VariablesProcessor) EntrarAmbito(nombre string) {
	vp.contadorAmbito++
	nuevoAmbito := fmt.Sprintf("%s_%d", nombre, vp.contadorAmbito)
	vp.tablaSimbolos.PilaAmbitos = append(vp.tablaSimbolos.PilaAmbitos, nuevoAmbito)
	vp.tablaSimbolos.AmbitoActual = nuevoAmbito
	vp.armGen.Comment(fmt.Sprintf("=== ENTRAR ÁMBITO: %s ===", nuevoAmbito))
}

func (vp *VariablesProcessor) SalirAmbito() {
	if len(vp.tablaSimbolos.PilaAmbitos) <= 1 {
		return // No salir del ámbito global
	}

	ambitoAnterior := vp.tablaSimbolos.AmbitoActual
	vp.tablaSimbolos.PilaAmbitos = vp.tablaSimbolos.PilaAmbitos[:len(vp.tablaSimbolos.PilaAmbitos)-1]
	vp.tablaSimbolos.AmbitoActual = vp.tablaSimbolos.PilaAmbitos[len(vp.tablaSimbolos.PilaAmbitos)-1]

	vp.armGen.Comment(fmt.Sprintf("=== SALIR ÁMBITO: %s -> %s ===", ambitoAnterior, vp.tablaSimbolos.AmbitoActual))
}

// ============= DECLARACIÓN DE VARIABLES =============

func (vp *VariablesProcessor) DeclararVariable(nombre, tipo string, valor *ResultadoExpresion) error {
	// 1. Verificar que no existe en el ámbito actual
	claveActual := vp.tablaSimbolos.AmbitoActual + ":" + nombre
	if _, existe := vp.tablaSimbolos.Simbolos[claveActual]; existe {
		return fmt.Errorf("variable '%s' ya declarada en ámbito '%s'", nombre, vp.tablaSimbolos.AmbitoActual)
	}

	// 2. Reservar espacio en stack PRIMERO
	vp.armGen.Instructions = append(vp.armGen.Instructions,
		fmt.Sprintf("sub sp, sp, #8  // Reservar espacio para %s", nombre))

	// 3. El offset de esta variable es 0 (está en [sp])
	offset := 0

	// 4. Crear símbolo
	simbolo := &Simbolo{
		Nombre:       nombre,
		Tipo:         tipo,
		Ambito:       vp.tablaSimbolos.AmbitoActual,
		Offset:       offset,
		EsGlobal:     vp.tablaSimbolos.AmbitoActual == "global",
		Inicializada: valor != nil,
	}

	// 5. Actualizar offsets de TODAS las variables anteriores
	// porque el sp se movió hacia abajo
	for _, sim := range vp.tablaSimbolos.Simbolos {
		sim.Offset += 8
	}

	// 6. Inicializar variable (ahora está en [sp])
	if valor != nil {
		vp.inicializarVariable(simbolo, valor)
	} else {
		vp.inicializarVariableDefault(simbolo)
	}

	// 7. Agregar a tabla de símbolos
	vp.tablaSimbolos.Simbolos[claveActual] = simbolo

	// 8. Incrementar contador de stack
	vp.tablaSimbolos.StackActual++

	vp.armGen.Comment(fmt.Sprintf("%s en [sp+%d] \n", nombre, offset))
	return nil
}

// ============= INICIALIZACIÓN DE VARIABLES =============

func (vp *VariablesProcessor) inicializarVariable(simbolo *Simbolo, valor *ResultadoExpresion) {
	switch simbolo.Tipo {
	case "int":
		vp.inicializarVariableInt(simbolo, valor)
	case "float64", "float":
		vp.inicializarVariableFloat(simbolo, valor)
	case "string":
		vp.inicializarVariableString(simbolo, valor)
	case "bool":
		vp.inicializarVariableBool(simbolo, valor)
	}
}

func (vp *VariablesProcessor) inicializarVariableInt(simbolo *Simbolo, valor *ResultadoExpresion) {
	// La variable recién declarada siempre está en [sp]
	direccion := "[sp]"

	if valor.EsLiteral {
		// Valor literal
		if valor.Tipo == "int" {
			regTmp := vp.nuevoRegistroTmp()
			vp.armGen.Mov(regTmp, valor.Valor.(int))
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else if valor.Tipo == "float" || valor.Tipo == "float64" {
			// Convertir float a int
			regTmp := vp.nuevoRegistroTmp()
			vp.armGen.Mov(regTmp, int(valor.Valor.(float64)))
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		}
	} else {
		// Valor en registro
		if valor.Tipo == "int" {
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", valor.Registro, direccion))
		} else if valor.Tipo == "float" || valor.Tipo == "float64" {
			// Convertir float a int
			regTmp := vp.nuevoRegistroTmp()
			vp.armGen.FcvtzsFloatToInt(regTmp, valor.Registro)
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		}
	}
}

func (vp *VariablesProcessor) inicializarVariableFloat(simbolo *Simbolo, valor *ResultadoExpresion) {
	// La variable recién declarada siempre está en [sp]
	direccion := "[sp]"

	if valor.EsLiteral {
		if valor.Tipo == "float" || valor.Tipo == "float64" {
			regTmp := vp.nuevoRegistroFloatTmp()
			vp.cargarConstanteFloat(regTmp, valor.Valor.(float64))
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else if valor.Tipo == "int" {
			// Convertir int a float
			regInt := vp.nuevoRegistroTmp()
			regFloat := vp.nuevoRegistroFloatTmp()
			vp.armGen.Mov(regInt, valor.Valor.(int))
			vp.armGen.ScvtfIntToFloat(regFloat, regInt)
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regFloat, direccion))
		}
	} else {
		if valor.Tipo == "float" || valor.Tipo == "float64" {
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", valor.Registro, direccion))
		} else if valor.Tipo == "int" {
			regFloat := vp.nuevoRegistroFloatTmp()
			vp.armGen.ScvtfIntToFloat(regFloat, valor.Registro)
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regFloat, direccion))
		}
	}
}

func (vp *VariablesProcessor) inicializarVariableString(simbolo *Simbolo, valor *ResultadoExpresion) {
	// La variable recién declarada siempre está en [sp]
	direccion := "[sp]"

	regTmp := vp.nuevoRegistroTmp()

	if valor.EsLiteral {
		etiqueta := vp.agregarMensajeString(valor.Valor.(string))
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", regTmp, etiqueta))
	} else {
		vp.armGen.MovReg(regTmp, valor.Registro)
	}

	vp.armGen.Instructions = append(vp.armGen.Instructions,
		fmt.Sprintf("str %s, %s", regTmp, direccion))
}

func (vp *VariablesProcessor) inicializarVariableBool(simbolo *Simbolo, valor *ResultadoExpresion) {
	// La variable recién declarada siempre está en [sp]
	direccion := "[sp]"
	regTmp := vp.nuevoRegistroTmp()

	if valor.EsLiteral {
		valorInt := 0
		if valor.Valor.(bool) {
			valorInt = 1
		}
		vp.armGen.Mov(regTmp, valorInt)
	} else {
		vp.armGen.MovReg(regTmp, valor.Registro)
	}

	vp.armGen.Instructions = append(vp.armGen.Instructions,
		fmt.Sprintf("str %s, %s", regTmp, direccion))
}

// ============= VALORES POR DEFECTO =============

func (vp *VariablesProcessor) inicializarVariableDefault(simbolo *Simbolo) {
	// La variable recién declarada siempre está en [sp]
	direccion := "[sp]"
	regTmp := vp.nuevoRegistroTmp()

	switch simbolo.Tipo {
	case "int":
		vp.armGen.Mov(regTmp, 0)
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("str %s, %s", regTmp, direccion))
	case "float", "float64":
		regFloat := vp.nuevoRegistroFloatTmp()
		vp.cargarConstanteFloat(regFloat, 0.0)
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("str %s, %s", regFloat, direccion))
	case "string":
		etiqueta := vp.agregarMensajeString("")
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
			fmt.Sprintf("str %s, %s", regTmp, direccion))
	case "bool":
		vp.armGen.Mov(regTmp, 0) // false
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("str %s, %s", regTmp, direccion))
	}
}

// ============= ACCESO A VARIABLES =============

func (vp *VariablesProcessor) ObtenerVariable(nombre string) (*Simbolo, error) {
	// Buscar desde el ámbito más interno hacia afuera
	for i := len(vp.tablaSimbolos.PilaAmbitos) - 1; i >= 0; i-- {
		ambito := vp.tablaSimbolos.PilaAmbitos[i]
		clave := ambito + ":" + nombre
		if simbolo, existe := vp.tablaSimbolos.Simbolos[clave]; existe {
			return simbolo, nil
		}
	}
	return nil, fmt.Errorf("variable '%s' no declarada", nombre)
}

func (vp *VariablesProcessor) CargarVariable(nombre string, tipoDestino string) (*ResultadoExpresion, error) {
	simbolo, err := vp.ObtenerVariable(nombre)
	if err != nil {
		return nil, err
	}

	vp.armGen.Comment(fmt.Sprintf("=== CARGAR VARIABLE: %s (offset: %d) ===", nombre, simbolo.Offset))

	// Calcular dirección basada en el offset actual
	direccion := fmt.Sprintf("[sp, #%d]", simbolo.Offset)

	switch simbolo.Tipo {
	case "int":
		reg := vp.nuevoRegistroTmp()
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("ldr %s, %s", reg, direccion))
		return &ResultadoExpresion{
			Registro:  reg,
			Tipo:      "int",
			EsLiteral: false,
		}, nil

	case "float", "float64":
		reg := vp.nuevoRegistroFloatTmp()
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("ldr %s, %s", reg, direccion))
		return &ResultadoExpresion{
			Registro:  reg,
			Tipo:      "float",
			EsLiteral: false,
		}, nil

	case "string":
		reg := vp.nuevoRegistroTmp()
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("ldr %s, %s", reg, direccion))
		return &ResultadoExpresion{
			Registro:  reg,
			Tipo:      "string",
			EsLiteral: false,
		}, nil

	case "bool":
		reg := vp.nuevoRegistroTmp()
		vp.armGen.Instructions = append(vp.armGen.Instructions,
			fmt.Sprintf("ldr %s, %s", reg, direccion))
		return &ResultadoExpresion{
			Registro:  reg,
			Tipo:      "bool",
			EsLiteral: false,
		}, nil
	}

	return nil, fmt.Errorf("tipo no soportado: %s", simbolo.Tipo)
}

// ============= ASIGNACIÓN DE VARIABLES =============

func (vp *VariablesProcessor) AsignarVariable(nombre string, valor *ResultadoExpresion) error {
	simbolo, err := vp.ObtenerVariable(nombre)
	if err != nil {
		return err
	}

	vp.armGen.Comment(fmt.Sprintf("=== ASIGNAR VARIABLE: %s = %s (offset: %d) ===", nombre, valor.Tipo, simbolo.Offset))

	// Verificar compatibilidad de tipos
	if !vp.sonTiposCompatibles(simbolo.Tipo, valor.Tipo) {
		return fmt.Errorf("no se puede asignar %s a variable %s de tipo %s", valor.Tipo, nombre, simbolo.Tipo)
	}

	// Asignación a variable existente usando el offset actual
	direccion := fmt.Sprintf("[sp, #%d]", simbolo.Offset)

	switch simbolo.Tipo {
	case "int":
		if valor.EsLiteral {
			regTmp := vp.nuevoRegistroTmp()
			if valor.Tipo == "int" {
				vp.armGen.Mov(regTmp, valor.Valor.(int))
			} else if valor.Tipo == "float" || valor.Tipo == "float64" {
				vp.armGen.Mov(regTmp, int(valor.Valor.(float64)))
			}
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else {
			if valor.Tipo == "int" {
				vp.armGen.Instructions = append(vp.armGen.Instructions,
					fmt.Sprintf("str %s, %s", valor.Registro, direccion))
			} else if valor.Tipo == "float" || valor.Tipo == "float64" {
				regTmp := vp.nuevoRegistroTmp()
				vp.armGen.FcvtzsFloatToInt(regTmp, valor.Registro)
				vp.armGen.Instructions = append(vp.armGen.Instructions,
					fmt.Sprintf("str %s, %s", regTmp, direccion))
			}
		}
	case "float", "float64":
		if valor.EsLiteral {
			regTmp := vp.nuevoRegistroFloatTmp()
			if valor.Tipo == "float" || valor.Tipo == "float64" {
				vp.cargarConstanteFloat(regTmp, valor.Valor.(float64))
			} else if valor.Tipo == "int" {
				regInt := vp.nuevoRegistroTmp()
				vp.armGen.Mov(regInt, valor.Valor.(int))
				vp.armGen.ScvtfIntToFloat(regTmp, regInt)
			}
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else {
			if valor.Tipo == "float" || valor.Tipo == "float64" {
				vp.armGen.Instructions = append(vp.armGen.Instructions,
					fmt.Sprintf("str %s, %s", valor.Registro, direccion))
			} else if valor.Tipo == "int" {
				regFloat := vp.nuevoRegistroFloatTmp()
				vp.armGen.ScvtfIntToFloat(regFloat, valor.Registro)
				vp.armGen.Instructions = append(vp.armGen.Instructions,
					fmt.Sprintf("str %s, %s", regFloat, direccion))
			}
		}
	case "string":
		if valor.EsLiteral {
			regTmp := vp.nuevoRegistroTmp()
			etiqueta := vp.agregarMensajeString(valor.Valor.(string))
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else {
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", valor.Registro, direccion))
		}
	case "bool":
		if valor.EsLiteral {
			regTmp := vp.nuevoRegistroTmp()
			valorInt := 0
			if valor.Valor.(bool) {
				valorInt = 1
			}
			vp.armGen.Mov(regTmp, valorInt)
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", regTmp, direccion))
		} else {
			vp.armGen.Instructions = append(vp.armGen.Instructions,
				fmt.Sprintf("str %s, %s", valor.Registro, direccion))
		}
	}

	simbolo.Inicializada = true
	return nil
}

// ============= UTILIDADES =============

func (vp *VariablesProcessor) sonTiposCompatibles(tipoVariable, tipoValor string) bool {
	if tipoVariable == tipoValor {
		return true
	}
	// Conversión implícita: int -> float
	if (tipoVariable == "float" || tipoVariable == "float64") && tipoValor == "int" {
		return true
	}
	// float y float64 son compatibles
	if (tipoVariable == "float" || tipoVariable == "float64") &&
		(tipoValor == "float" || tipoValor == "float64") {
		return true
	}
	return false
}

func (vp *VariablesProcessor) InferirTipo(valor *ResultadoExpresion) string {
	return valor.Tipo
}

// Métodos auxiliares para registros y constantes
func (vp *VariablesProcessor) nuevoRegistroTmp() string {
	return vp.expresionesProcessor.NuevoRegistroTmp()
}

func (vp *VariablesProcessor) nuevoRegistroFloatTmp() string {
	return vp.expresionesProcessor.NuevoRegistroFloatTmp()
}

func (vp *VariablesProcessor) cargarConstanteFloat(registro string, valor float64) {
	vp.expresionesProcessor.CargarConstanteFloat(registro, valor)
}

func (vp *VariablesProcessor) agregarMensajeString(mensaje string) string {
	return vp.expresionesProcessor.AgregarMensajeString(mensaje)
}

// ============= GETTERS PARA REPORTES =============

func (vp *VariablesProcessor) GetTablaSimbolos() map[string]*Simbolo {
	return vp.tablaSimbolos.Simbolos
}

func (vp *VariablesProcessor) GetAmbitoActual() string {
	return vp.tablaSimbolos.AmbitoActual
}
