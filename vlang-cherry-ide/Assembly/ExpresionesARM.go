package assembly

import (
	"fmt"
	"strconv"
)

// ResultadoExpresion representa el resultado de evaluar una expresión
// 1.1 - Agregar campos a ResultadoExpresion
type ResultadoExpresion struct {
	Registro       string
	Tipo           string
	EsLiteral      bool
	Valor          interface{}
	SliceInfo      *SliceAccessInfo `json:"-"`
	EsTypeOfResult bool             // 🔥 AGREGAR ESTE CAMPO
	TypeOfValue    string           // 🔥 AGREGAR ESTE CAMPO
}

type SliceAccessInfo struct {
	NombreSlice string
	IndiceExpr  *ResultadoExpresion
}

// ExpresionesProcessor maneja todas las operaciones de expresiones
type ExpresionesProcessor struct {
	armGen           *ARMGenerator
	tmpCounter       int
	tmpFloatCounter  int
	contadorEtiqueta int
	contadorMensaje  int
	MensajesDatos    []string
}

// NewExpresionesProcessor crea un nuevo procesador de expresiones
func NewExpresionesProcessor(armGen *ARMGenerator) *ExpresionesProcessor {
	return &ExpresionesProcessor{
		armGen:          armGen,
		MensajesDatos:   []string{},
		contadorMensaje: 1,
	}
}

// ============= GESTIÓN DE REGISTROS =============

func (ep *ExpresionesProcessor) NuevoRegistroTmp() string {
	// 🔥 EXCLUSIÓN CRÍTICA: No usar x28 (reservado para slices) ni x29 (frame pointer)

	// Usar registros x9-x18 de forma cíclica, evitando x28 y x29
	registroNum := 9 + (ep.tmpCounter % 10) // x9-x18 (10 registros disponibles)

	// 🔥 IMPORTANTE: Si por alguna razón llegamos a x28 o x29, volver al rango seguro
	if registroNum >= 19 {
		registroNum = 9 + (ep.tmpCounter % 10)
	}

	ep.tmpCounter++

	// PERMITIR RECICLAJE DE REGISTROS
	if ep.tmpCounter > 100 { // Evitar overflow del contador
		ep.armGen.Comment("ADVERTENCIA: Reciclando contadores de registros temporales")
		ep.tmpCounter = 0
	}

	return fmt.Sprintf("x%d", registroNum)
}

func (ep *ExpresionesProcessor) NuevoRegistroFloatTmp() string {
	if ep.tmpFloatCounter > 31 {
		ep.armGen.Comment("ADVERTENCIA: Reciclando registros float")
		ep.tmpFloatCounter = 0
	}
	reg := fmt.Sprintf("d%d", ep.tmpFloatCounter)
	ep.tmpFloatCounter++
	return reg
}

func (ep *ExpresionesProcessor) ResetearContadores() {
	ep.tmpCounter = 0
	ep.tmpFloatCounter = 0
}

func (ep *ExpresionesProcessor) ResetearRegistrosSiNecesario() {
	// Resetear si estamos cerca del límite
	if ep.tmpCounter > 50 { // Reducido para mayor seguridad
		ep.armGen.Comment("Reseteando contadores de registros por límite")
		ep.tmpCounter = 0
	}
	if ep.tmpFloatCounter > 20 {
		ep.tmpFloatCounter = 0
	}
}

// ============= DETERMINACIÓN DE TIPOS =============

func (ep *ExpresionesProcessor) DeterminarTipoResultante(tipo1, tipo2 string) string {
	if tipo1 == "float" || tipo2 == "float" {
		return "float"
	}
	return "int"
}

// ============= PROCESAMIENTO DE OPERACIONES BINARIAS =============

func (ep *ExpresionesProcessor) ProcesarOperacionBinaria(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	// CONCATENACIÓN DE STRINGS
	if operador == "+" && left.Tipo == "string" && right.Tipo == "string" {
		return ep.ProcesarConcatenacionString(left, right)
	}

	// OPERADORES LÓGICOS (&&, ||)
	if operador == "&&" || operador == "||" {
		return ep.ProcesarOperacionLogica(left, right, operador)
	}

	// OPERADORES RELACIONALES (==, !=, <, >, <=, >=)
	if operador == "==" || operador == "!=" || operador == "<" || operador == ">" || operador == "<=" || operador == ">=" {
		return ep.ProcesarOperacionRelacional(left, right, operador)
	}

	// OPERACIONES NUMÉRICAS
	tipoResultante := ep.DeterminarTipoResultante(left.Tipo, right.Tipo)
	if tipoResultante == "float" {
		return ep.ProcesarOperacionFloat(left, right, operador)
	} else {
		return ep.ProcesarOperacionInt(left, right, operador)
	}
}

// ============= CONCATENACIÓN DE STRINGS =============

func (ep *ExpresionesProcessor) ProcesarConcatenacionString(left, right *ResultadoExpresion) *ResultadoExpresion {
	ep.armGen.Comment("=== CONCATENACIÓN DE STRINGS ===")

	registroStr1 := ep.NuevoRegistroTmp()
	registroStr2 := ep.NuevoRegistroTmp()
	registroResultado := ep.NuevoRegistroTmp()

	// Cargar primer string
	if left.EsLiteral {
		etiqueta1 := ep.AgregarMensajeString(left.Valor.(string))
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", registroStr1, etiqueta1))
	} else {
		ep.armGen.MovReg(registroStr1, left.Registro)
	}

	// Cargar segundo string
	if right.EsLiteral {
		etiqueta2 := ep.AgregarMensajeString(right.Valor.(string))
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", registroStr2, etiqueta2))
	} else {
		ep.armGen.MovReg(registroStr2, right.Registro)
	}

	// Obtener buffer para resultado
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("adr %s, buffer_string", registroResultado))

	// Concatenar strings
	ep.armGen.ConcatStrings(registroResultado, registroStr1, registroStr2)

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "string",
		EsLiteral: false,
	}
}

// ============= OPERACIONES LÓGICAS =============

func (ep *ExpresionesProcessor) ProcesarOperacionLogica(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	ep.armGen.Comment(fmt.Sprintf("=== OPERACIÓN LÓGICA %s ===", operador))

	// Verificar que ambos operandos sean booleanos
	if left.Tipo != "bool" || right.Tipo != "bool" {
		ep.armGen.Comment(fmt.Sprintf("Error: Operadores lógicos requieren valores bool, recibido: %s %s %s", left.Tipo, operador, right.Tipo))
		registroError := ep.NuevoRegistroTmp()
		ep.armGen.Mov(registroError, 0)
		return &ResultadoExpresion{
			Registro:  registroError,
			Tipo:      "bool",
			EsLiteral: false,
		}
	}

	registroResultado := ep.NuevoRegistroTmp()

	if operador == "&&" {
		return ep.ProcesarOperacionAND(left, right, registroResultado)
	} else {
		return ep.ProcesarOperacionOR(left, right, registroResultado)
	}
}

func (ep *ExpresionesProcessor) ProcesarOperacionAND(left, right *ResultadoExpresion, registroResultado string) *ResultadoExpresion {
	// VALIDACIÓN PREVIA DE TIPOS
	if left == nil || right == nil {
		ep.armGen.Comment("Error: operando nil en operación AND")
		ep.armGen.Mov(registroResultado, 0)
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      "bool",
			EsLiteral: false,
		}
	}

	registroIzq := ep.NuevoRegistroTmp()
	registroDer := ep.NuevoRegistroTmp()

	// Cargar operando izquierdo con validación
	if left.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if left.Valor != nil {
			if boolVal, ok := left.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroIzq, valor)
	} else {
		ep.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cortocircuito para AND
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("cmp %s, #0", registroIzq),
		fmt.Sprintf("beq .Land_false_%d", ep.contadorEtiqueta))

	etiquetaFalse := ep.contadorEtiqueta
	etiquetaEnd := ep.contadorEtiqueta + 1
	ep.contadorEtiqueta += 2

	// Cargar operando derecho solo si el izquierdo es true
	if right.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if right.Valor != nil {
			if boolVal, ok := right.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroDer, valor)
	} else {
		ep.armGen.MovReg(registroDer, right.Registro)
	}

	// AND lógico
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("and %s, %s, %s", registroResultado, registroIzq, registroDer),
		fmt.Sprintf("b .Land_end_%d", etiquetaEnd))

	// Etiqueta para false (cortocircuito)
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf(".Land_false_%d:", etiquetaFalse),
		fmt.Sprintf("mov %s, #0", registroResultado))

	// Etiqueta final
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf(".Land_end_%d:", etiquetaEnd))

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "bool",
		EsLiteral: false,
	}
}

func (ep *ExpresionesProcessor) ProcesarOperacionOR(left, right *ResultadoExpresion, registroResultado string) *ResultadoExpresion {
	// VALIDACIÓN PREVIA DE TIPOS
	if left == nil || right == nil {
		ep.armGen.Comment("Error: operando nil en operación OR")
		ep.armGen.Mov(registroResultado, 0)
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      "bool",
			EsLiteral: false,
		}
	}

	registroIzq := ep.NuevoRegistroTmp()
	registroDer := ep.NuevoRegistroTmp()

	// Cargar operando izquierdo
	if left.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if left.Valor != nil {
			if boolVal, ok := left.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroIzq, valor)
	} else {
		ep.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cortocircuito para OR
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("cmp %s, #1", registroIzq),
		fmt.Sprintf("beq .Lor_true_%d", ep.contadorEtiqueta))

	etiquetaTrue := ep.contadorEtiqueta
	etiquetaEnd := ep.contadorEtiqueta + 1
	ep.contadorEtiqueta += 2

	// Cargar operando derecho solo si el izquierdo es false
	if right.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if right.Valor != nil {
			if boolVal, ok := right.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroDer, valor)
	} else {
		ep.armGen.MovReg(registroDer, right.Registro)
	}

	// OR lógico
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("orr %s, %s, %s", registroResultado, registroIzq, registroDer),
		fmt.Sprintf("b .Lor_end_%d", etiquetaEnd))

	// Etiqueta para true (cortocircuito)
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf(".Lor_true_%d:", etiquetaTrue),
		fmt.Sprintf("mov %s, #1", registroResultado))

	// Etiqueta final
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf(".Lor_end_%d:", etiquetaEnd))

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "bool",
		EsLiteral: false,
	}
}

// ============= OPERACIONES RELACIONALES =============

func (ep *ExpresionesProcessor) ProcesarOperacionRelacional(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	ep.armGen.Comment(fmt.Sprintf("=== OPERACIÓN RELACIONAL %s ===", operador))

	registroResultado := ep.NuevoRegistroTmp()

	// COMPARACIÓN DE STRINGS (solo == y !=)
	if left.Tipo == "string" && right.Tipo == "string" {
		if operador == "==" || operador == "!=" {
			return ep.ProcesarComparacionString(left, right, operador, registroResultado)
		} else {
			ep.armGen.Comment(fmt.Sprintf("Operador %s no soportado para strings", operador))
			ep.armGen.Mov(registroResultado, 0)
			return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
		}
	}

	// COMPARACIÓN DE NÚMEROS (int/float)
	if (left.Tipo == "int" || left.Tipo == "float") && (right.Tipo == "int" || right.Tipo == "float") {
		return ep.ProcesarComparacionNumerica(left, right, operador, registroResultado)
	}

	// COMPARACIÓN DE BOOLEANOS (solo == y !=)
	if left.Tipo == "bool" && right.Tipo == "bool" {
		if operador == "==" || operador == "!=" {
			return ep.ProcesarComparacionBooleana(left, right, operador, registroResultado)
		} else {
			ep.armGen.Comment(fmt.Sprintf("Operador %s no soportado para bool", operador))
			ep.armGen.Mov(registroResultado, 0)
			return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
		}
	}

	// Tipos incompatibles
	ep.armGen.Comment(fmt.Sprintf("Tipos incompatibles para %s: %s vs %s", operador, left.Tipo, right.Tipo))
	ep.armGen.Mov(registroResultado, 0)
	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarComparacionString(left, right *ResultadoExpresion, operador, registroResultado string) *ResultadoExpresion {
	registroStr1 := ep.NuevoRegistroTmp()
	registroStr2 := ep.NuevoRegistroTmp()

	// Cargar primer string
	if left.EsLiteral {
		etiqueta1 := ep.AgregarMensajeString(left.Valor.(string))
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", registroStr1, etiqueta1))
	} else {
		ep.armGen.MovReg(registroStr1, left.Registro)
	}

	// Cargar segundo string
	if right.EsLiteral {
		etiqueta2 := ep.AgregarMensajeString(right.Valor.(string))
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", registroStr2, etiqueta2))
	} else {
		ep.armGen.MovReg(registroStr2, right.Registro)
	}

	// Comparar strings
	ep.armGen.StrCmp(registroStr1, registroStr2, registroResultado)

	// Ajustar resultado según operador
	if operador == "==" {
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("cmp %s, #0", registroResultado),
			fmt.Sprintf("cset %s, eq", registroResultado))
	} else { // !=
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("cmp %s, #0", registroResultado),
			fmt.Sprintf("cset %s, ne", registroResultado))
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarComparacionNumerica(left, right *ResultadoExpresion, operador, registroResultado string) *ResultadoExpresion {
	esComparacionFloat := left.Tipo == "float" || right.Tipo == "float"

	if esComparacionFloat {
		return ep.ProcesarComparacionFloat(left, right, operador, registroResultado)
	} else {
		return ep.ProcesarComparacionInt(left, right, operador, registroResultado)
	}
}

func (ep *ExpresionesProcessor) ProcesarComparacionInt(left, right *ResultadoExpresion, operador, registroResultado string) *ResultadoExpresion {
	registroIzq := ep.NuevoRegistroTmp()
	registroDer := ep.NuevoRegistroTmp()

	// Cargar operando izquierdo
	if left.EsLiteral {
		ep.armGen.Mov(registroIzq, left.Valor.(int))
	} else {
		ep.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cargar operando derecho
	if right.EsLiteral {
		ep.armGen.Mov(registroDer, right.Valor.(int))
	} else {
		ep.armGen.MovReg(registroDer, right.Registro)
	}

	// Comparar y establecer resultado
	ep.armGen.Cmp(registroIzq, registroDer)

	switch operador {
	case "==":
		ep.armGen.CSet(registroResultado, "eq")
	case "!=":
		ep.armGen.CSet(registroResultado, "ne")
	case "<":
		ep.armGen.CSet(registroResultado, "lt")
	case ">":
		ep.armGen.CSet(registroResultado, "gt")
	case "<=":
		ep.armGen.CSet(registroResultado, "le")
	case ">=":
		ep.armGen.CSet(registroResultado, "ge")
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarComparacionFloat(left, right *ResultadoExpresion, operador, registroResultado string) *ResultadoExpresion {
	registroIzq := ep.NuevoRegistroFloatTmp()
	registroDer := ep.NuevoRegistroFloatTmp()

	// Cargar operando izquierdo (convertir int->float si es necesario)
	if left.EsLiteral {
		if left.Tipo == "float" {
			ep.CargarConstanteFloat(registroIzq, left.Valor.(float64))
		} else { // int
			regTmp := ep.NuevoRegistroTmp()
			ep.armGen.Mov(regTmp, left.Valor.(int))
			ep.armGen.ScvtfIntToFloat(registroIzq, regTmp)
		}
	} else {
		if left.Tipo == "int" {
			ep.armGen.ScvtfIntToFloat(registroIzq, left.Registro)
		} else {
			ep.armGen.FMov(registroIzq, left.Registro)
		}
	}

	// Cargar operando derecho (convertir int->float si es necesario)
	if right.EsLiteral {
		if right.Tipo == "float" {
			ep.CargarConstanteFloat(registroDer, right.Valor.(float64))
		} else { // int
			regTmp := ep.NuevoRegistroTmp()
			ep.armGen.Mov(regTmp, right.Valor.(int))
			ep.armGen.ScvtfIntToFloat(registroDer, regTmp)
		}
	} else {
		if right.Tipo == "int" {
			ep.armGen.ScvtfIntToFloat(registroDer, right.Registro)
		} else {
			ep.armGen.FMov(registroDer, right.Registro)
		}
	}

	// Comparar floats y establecer resultado
	ep.armGen.FCmp(registroIzq, registroDer)

	switch operador {
	case "==":
		ep.armGen.CSet(registroResultado, "eq")
	case "!=":
		ep.armGen.CSet(registroResultado, "ne")
	case "<":
		ep.armGen.CSet(registroResultado, "lt")
	case ">":
		ep.armGen.CSet(registroResultado, "gt")
	case "<=":
		ep.armGen.CSet(registroResultado, "le")
	case ">=":
		ep.armGen.CSet(registroResultado, "ge")
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarComparacionBooleana(left, right *ResultadoExpresion, operador, registroResultado string) *ResultadoExpresion {
	registroIzq := ep.NuevoRegistroTmp()
	registroDer := ep.NuevoRegistroTmp()

	// Cargar operando izquierdo con validación
	if left.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if left.Valor != nil {
			if boolVal, ok := left.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroIzq, valor)
	} else {
		ep.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cargar operando derecho con validación
	if right.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if right.Valor != nil {
			if boolVal, ok := right.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroDer, valor)
	} else {
		ep.armGen.MovReg(registroDer, right.Registro)
	}

	// Comparar y establecer resultado
	ep.armGen.Cmp(registroIzq, registroDer)

	if operador == "==" {
		ep.armGen.CSet(registroResultado, "eq")
	} else { // !=
		ep.armGen.CSet(registroResultado, "ne")
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

// ============= FUNCIONES AUXILIARES ADICIONALES =============

// Función auxiliar para validar y obtener valor entero de forma segura
func (ep *ExpresionesProcessor) obtenerValorEnteroSeguro(expr *ResultadoExpresion, valorDefecto int) int {
	if expr == nil || expr.Valor == nil {
		return valorDefecto
	}

	if intVal, ok := expr.Valor.(int); ok {
		return intVal
	}

	if floatVal, ok := expr.Valor.(float64); ok {
		return int(floatVal)
	}

	return valorDefecto
}

// ============= OPERACIONES NUMÉRICAS =============

func (ep *ExpresionesProcessor) ProcesarOperacionInt(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	registroIzq := ep.NuevoRegistroTmp()
	registroDer := ep.NuevoRegistroTmp()
	registroResultado := ep.NuevoRegistroTmp()

	// Cargar operando izquierdo con validación segura
	if left.EsLiteral {
		valor := ep.obtenerValorEnteroSeguro(left, 0)
		ep.armGen.Mov(registroIzq, valor)
	} else {
		ep.armGen.MovReg(registroIzq, left.Registro)
	}

	// Cargar operando derecho con validación segura
	if right.EsLiteral {
		valor := ep.obtenerValorEnteroSeguro(right, 0)
		ep.armGen.Mov(registroDer, valor)
	} else {
		ep.armGen.MovReg(registroDer, right.Registro)
	}

	// Generar operación
	switch operador {
	case "+":
		ep.armGen.Add(registroResultado, registroIzq, registroDer)
	case "-":
		ep.armGen.Sub(registroResultado, registroIzq, registroDer)
	case "*":
		ep.armGen.Mul(registroResultado, registroIzq, registroDer)
	case "/":
		ep.armGen.Div(registroResultado, registroIzq, registroDer)
	case "%":
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("udiv %s, %s, %s", registroResultado, registroIzq, registroDer),
			fmt.Sprintf("msub %s, %s, %s, %s", registroResultado, registroResultado, registroDer, registroIzq))
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "int", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarOperacionFloat(left, right *ResultadoExpresion, operador string) *ResultadoExpresion {
	registroIzq := ep.NuevoRegistroFloatTmp()
	registroDer := ep.NuevoRegistroFloatTmp()
	registroResultado := ep.NuevoRegistroFloatTmp()

	// Cargar operando izquierdo
	if left.EsLiteral {
		if left.Tipo == "float" {
			ep.CargarConstanteFloat(registroIzq, left.Valor.(float64))
		} else if left.Tipo == "int" {
			regTmp := ep.NuevoRegistroTmp()
			ep.armGen.Mov(regTmp, left.Valor.(int))
			ep.armGen.ScvtfIntToFloat(registroIzq, regTmp)
		}
	} else {
		if left.Tipo == "int" {
			ep.armGen.ScvtfIntToFloat(registroIzq, left.Registro)
		} else {
			ep.armGen.FMov(registroIzq, left.Registro)
		}
	}

	// Cargar operando derecho
	if right.EsLiteral {
		if right.Tipo == "float" {
			ep.CargarConstanteFloat(registroDer, right.Valor.(float64))
		} else if right.Tipo == "int" {
			regTmp := ep.NuevoRegistroTmp()
			ep.armGen.Mov(regTmp, right.Valor.(int))
			ep.armGen.ScvtfIntToFloat(registroDer, regTmp)
		}
	} else {
		if right.Tipo == "int" {
			ep.armGen.ScvtfIntToFloat(registroDer, right.Registro)
		} else {
			ep.armGen.FMov(registroDer, right.Registro)
		}
	}

	// Generar operación float
	switch operador {
	case "+":
		ep.armGen.FAdd(registroResultado, registroIzq, registroDer)
	case "-":
		ep.armGen.FSub(registroResultado, registroIzq, registroDer)
	case "*":
		ep.armGen.FMul(registroResultado, registroIzq, registroDer)
	case "/":
		ep.armGen.FDiv(registroResultado, registroIzq, registroDer)
	case "%":
		ep.armGen.Comment("Operador % no soportado directamente para floats")
	}

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "float", EsLiteral: false}
}

// ============= OPERACIONES UNARIAS =============

func (ep *ExpresionesProcessor) ProcesarOperacionUnaria(expr *ResultadoExpresion, operador string) *ResultadoExpresion {
	switch operador {
	case "-":
		if expr.Tipo == "int" {
			return ep.ProcesarNegacionAritmetica(expr)
		} else if expr.Tipo == "float" {
			return ep.ProcesarNegacionFloat(expr)
		} else {
			ep.armGen.Comment("Negación aritmética solo soporta int y float")
			return expr
		}
	case "!":
		if expr.Tipo == "bool" {
			return ep.ProcesarNegacionLogica(expr)
		} else {
			ep.armGen.Comment("Operador ! solo funciona con valores booleanos")
			return expr
		}
	default:
		ep.armGen.Comment(fmt.Sprintf("Operador unario no soportado: %s", operador))
		return expr
	}
}

func (ep *ExpresionesProcessor) ProcesarNegacionLogica(expr *ResultadoExpresion) *ResultadoExpresion {
	ep.armGen.Comment("=== NEGACIÓN LÓGICA ! ===")

	registroOrigen := ep.NuevoRegistroTmp()
	registroResultado := ep.NuevoRegistroTmp()

	// Cargar valor booleano con validación
	if expr.EsLiteral {
		valor := 0
		// VALIDACIÓN SEGURA DE TIPO
		if expr.Valor != nil {
			if boolVal, ok := expr.Valor.(bool); ok && boolVal {
				valor = 1
			}
		}
		ep.armGen.Mov(registroOrigen, valor)
	} else {
		ep.armGen.MovReg(registroOrigen, expr.Registro)
	}

	// Negación lógica: XOR con 1
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("eor %s, %s, #1", registroResultado, registroOrigen))

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "bool", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarNegacionAritmetica(expr *ResultadoExpresion) *ResultadoExpresion {
	ep.armGen.Comment("=== NEGACIÓN ARITMÉTICA - ===")

	registroOrigen := ep.NuevoRegistroTmp()
	registroResultado := ep.NuevoRegistroTmp()

	if expr.EsLiteral {
		ep.armGen.Mov(registroOrigen, expr.Valor.(int))
	} else {
		ep.armGen.MovReg(registroOrigen, expr.Registro)
	}

	// Negar el valor: 0 - valor
	ep.armGen.Mov(registroResultado, 0)
	ep.armGen.Sub(registroResultado, registroResultado, registroOrigen)

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "int", EsLiteral: false}
}

func (ep *ExpresionesProcessor) ProcesarNegacionFloat(expr *ResultadoExpresion) *ResultadoExpresion {
	ep.armGen.Comment("=== NEGACIÓN FLOAT - ===")

	registroOrigen := ep.NuevoRegistroFloatTmp()
	registroResultado := ep.NuevoRegistroFloatTmp()

	if expr.EsLiteral {
		ep.CargarConstanteFloat(registroOrigen, expr.Valor.(float64))
	} else {
		ep.armGen.FMov(registroOrigen, expr.Registro)
	}

	// Negar float: usar fneg
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("fneg %s, %s", registroResultado, registroOrigen))

	return &ResultadoExpresion{Registro: registroResultado, Tipo: "float", EsLiteral: false}
}

// ============= UTILIDADES =============

func (ep *ExpresionesProcessor) AgregarMensajeString(mensaje string) string {
	etiqueta := fmt.Sprintf("msg_%d", ep.contadorMensaje)
	ep.contadorMensaje++
	ep.MensajesDatos = append(ep.MensajesDatos,
		fmt.Sprintf("%s: .asciz \"%s\"", etiqueta, mensaje))
	return etiqueta
}

func (ep *ExpresionesProcessor) AgregarConstanteFloat(valor float64) string {
	etiqueta := fmt.Sprintf("float_const_%d", ep.contadorMensaje)
	ep.contadorMensaje++
	ep.MensajesDatos = append(ep.MensajesDatos,
		fmt.Sprintf("%s: .double %.6f", etiqueta, valor))
	return etiqueta
}

func (ep *ExpresionesProcessor) CargarConstanteFloat(registro string, valor float64) {
	if valor == 0.0 {
		ep.armGen.Instructions = append(ep.armGen.Instructions,
			fmt.Sprintf("fmov %s, wzr", registro))
		return
	}

	etiqueta := ep.AgregarConstanteFloat(valor)
	regTmp := ep.NuevoRegistroTmp()
	ep.armGen.Instructions = append(ep.armGen.Instructions,
		fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
		fmt.Sprintf("ldr %s, [%s]", registro, regTmp))
}

func (ep *ExpresionesProcessor) CrearLiteralInt(valor string) *ResultadoExpresion {
	val, _ := strconv.Atoi(valor)
	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "int",
		EsLiteral: true,
		Valor:     val,
	}
}

func (ep *ExpresionesProcessor) CrearLiteralFloat(valor string) *ResultadoExpresion {
	val, _ := strconv.ParseFloat(valor, 64)
	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "float",
		EsLiteral: true,
		Valor:     val,
	}
}

func (ep *ExpresionesProcessor) CrearLiteralString(valor string) *ResultadoExpresion {
	// Remover comillas
	val := valor[1 : len(valor)-1]
	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "string",
		EsLiteral: true,
		Valor:     val,
	}
}

func (ep *ExpresionesProcessor) CrearLiteralBool(valor string) *ResultadoExpresion {
	ep.armGen.Comment(fmt.Sprintf("=== LITERAL BOOL: %s ===", valor))

	// VALIDACIÓN SEGURA
	boolVal := false
	if valor == "true" {
		boolVal = true
	} else if valor == "false" {
		boolVal = false
	} else {
		ep.armGen.Comment(fmt.Sprintf("Advertencia: valor booleano no reconocido '%s', usando false", valor))
	}

	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "bool",
		EsLiteral: true,
		Valor:     boolVal, // SIEMPRE UN BOOL VÁLIDO
	}
}

func (ep *ExpresionesProcessor) GetMensajesDatos() []string {
	return ep.MensajesDatos
}
func (ep *ExpresionesProcessor) GetContadorEtiqueta() int {
	ep.contadorEtiqueta++
	return ep.contadorEtiqueta
}

func (ep *ExpresionesProcessor) GenerarEtiquetaUnica(prefijo string) string {
	etiqueta := fmt.Sprintf(".L%s_%d", prefijo, ep.contadorEtiqueta)
	ep.contadorEtiqueta++
	return etiqueta
}

func (ep *ExpresionesProcessor) IncrementarContadorEtiqueta() {
	ep.contadorEtiqueta++
}
