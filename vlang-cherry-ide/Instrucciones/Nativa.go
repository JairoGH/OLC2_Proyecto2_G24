package instrucciones

import (
	"main/tiposDeDato"
	"strconv"
	"strings"
)

type FuncionNativa struct {
	Nombre string
	Exec   func(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string)
}

func (n FuncionNativa) Type() string {
	return tiposDeDato.TIPO_FUNCIONNATIVA
}

func (n FuncionNativa) Value() interface{} {
	return n
}

func (n FuncionNativa) Copy() tiposDeDato.ValorInterno {
	return n
}

// Función auxiliar para formatear decimales eliminando ceros innecesarios
func formatearDecimal(valor float64, maxDecimales int) string {
    // Formatear con el máximo de decimales
    str := strconv.FormatFloat(valor, 'f', maxDecimales, 64)
    
    // Eliminar ceros innecesarios al final
    str = strings.TrimRight(str, "0")
    
    // Si termina en punto, agregar un cero (mantener al menos un decimal)
    if strings.HasSuffix(str, ".") {
        str += "0"
    }
    
    return str
}

// Función para imprimir - CORREGIDA
func Imprimir(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
    var output string

    for i, arg := range args {
        switch arg.Valor.Type() {
        case tiposDeDato.TIPO_BOOLEAN:
            output += strconv.FormatBool(arg.Valor.Value().(bool))
        case tiposDeDato.TIPO_ENTERO:
            output += strconv.Itoa(arg.Valor.Value().(int))
        case tiposDeDato.TIPO_DECIMAL:
            output += formatearDecimal(arg.Valor.Value().(float64), 5)
        case tiposDeDato.TIPO_CADENA:
            output += arg.Valor.Value().(string)
        case tiposDeDato.TIPO_NULO:
            output += "nil"
        default:
            // ✅ NUEVO: Soporte prioritario para TipoMatrizMulti
            if matrizMulti, ok := arg.Valor.(*TipoMatrizMulti); ok {
                output += "["
                for fila, row := range matrizMulti.ValorInterno {
                    if fila > 0 {
                        output += ", "
                    }
                    output += "["
                    for col, elemento := range row {
                        if col > 0 {
                            output += ", "
                        }
                        switch elemento.Type() {
                        case tiposDeDato.TIPO_BOOLEAN:
                            output += strconv.FormatBool(elemento.Value().(bool))
                        case tiposDeDato.TIPO_ENTERO:
                            output += strconv.Itoa(elemento.Value().(int))
                        case tiposDeDato.TIPO_DECIMAL:
                            output += formatearDecimal(elemento.Value().(float64), 5)
                        case tiposDeDato.TIPO_CADENA:
                            output += "\"" + elemento.Value().(string) + "\""
                        case tiposDeDato.TIPO_NULO:
                            output += "nil"
                        default:
                            output += "nil"
                        }
                    }
                    output += "]"
                }
                output += "]"
            } else if obj, ok := arg.Valor.(*TipoObjeto); ok {
                // Manejar structs
                output += formatStruct(obj)
            } else if EsTipoVector(arg.Valor.Type()) || strings.HasPrefix(arg.Valor.Type(), "[]") {
                // Manejar vectores y matrices tradicionales
                if vector, ok := arg.Valor.(*TipoVector); ok {
                    output += "["
                    for j, elemento := range vector.ValorInterno {
                        if j > 0 {
                            output += ", "
                        }
                        switch elemento.Type() {
                        case tiposDeDato.TIPO_BOOLEAN:
                            output += strconv.FormatBool(elemento.Value().(bool))
                        case tiposDeDato.TIPO_ENTERO:
                            output += strconv.Itoa(elemento.Value().(int))
                        case tiposDeDato.TIPO_DECIMAL:
                            output += formatearDecimal(elemento.Value().(float64), 5)
                        case tiposDeDato.TIPO_CADENA:
                            output += "\"" + elemento.Value().(string) + "\""
                        case tiposDeDato.TIPO_NULO:
                            output += "nil"
                        default:
                            if subVector, ok := elemento.(*TipoVector); ok {
                                // Manejar matrices multidimensionales tradicionales
                                output += "["
                                for k, subElemento := range subVector.ValorInterno {
                                    if k > 0 {
                                        output += ", "
                                    }
                                    switch subElemento.Type() {
                                    case tiposDeDato.TIPO_BOOLEAN:
                                        output += strconv.FormatBool(subElemento.Value().(bool))
                                    case tiposDeDato.TIPO_ENTERO:
                                        output += strconv.Itoa(subElemento.Value().(int))
                                    case tiposDeDato.TIPO_DECIMAL:
                                        output += formatearDecimal(subElemento.Value().(float64), 5)
                                    case tiposDeDato.TIPO_CADENA:
                                        output += "\"" + subElemento.Value().(string) + "\""
                                    default:
                                        output += "nil"
                                    }
                                }
                                output += "]"
                            } else if obj, ok := elemento.(*TipoObjeto); ok {
                                output += formatStruct(obj)
                            } else {
                                output += "nil"
                            }
                        }
                    }
                    output += "]"
                } else if matriz, ok := arg.Valor.(*TipoMatriz); ok {
                    // ✅ CORREGIDO: Calcular dimensiones dinámicamente
                    output += "["
                    filas := len(matriz.ValorInterno) // Número de filas
                    
                    for fila := 0; fila < filas; fila++ {
                        if fila > 0 {
                            output += ", "
                        }
                        output += "["
                        
                        // Obtener la fila actual como vector
                        filaVector := matriz.ValorInterno[fila].(*TipoVector)
                        columnas := len(filaVector.ValorInterno) // Número de columnas en esta fila
                        
                        for col := 0; col < columnas; col++ {
                            if col > 0 {
                                output += ", "
                            }
                            elemento := filaVector.ValorInterno[col]
                            switch elemento.Type() {
                            case tiposDeDato.TIPO_BOOLEAN:
                                output += strconv.FormatBool(elemento.Value().(bool))
                            case tiposDeDato.TIPO_ENTERO:
                                output += strconv.Itoa(elemento.Value().(int))
                            case tiposDeDato.TIPO_DECIMAL:
                                output += formatearDecimal(elemento.Value().(float64), 5)
                            case tiposDeDato.TIPO_CADENA:
                                output += "\"" + elemento.Value().(string) + "\""
                            default:
                                output += "nil"
                            }
                        }
                        output += "]"
                    }
                    output += "]"
                } else {
                    return tiposDeDato.NuloPorDefecto, false, "La función print no puede imprimir el tipo: " + arg.Valor.Type()
                }
            } else {
                return tiposDeDato.NuloPorDefecto, false, "La función print no puede imprimir el tipo: " + arg.Valor.Type()
            }
        }

        if i < len(args)-1 {
            output += " "
        }
    }

    context.Consola.Print(output)
    return tiposDeDato.NuloPorDefecto, true, ""
}

func ImprimirLn(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
    var output string

    for i, arg := range args {
        switch arg.Valor.Type() {
        case tiposDeDato.TIPO_BOOLEAN:
            output += strconv.FormatBool(arg.Valor.Value().(bool))
        case tiposDeDato.TIPO_ENTERO:
            output += strconv.Itoa(arg.Valor.Value().(int))
        case tiposDeDato.TIPO_DECIMAL:
            output += formatearDecimal(arg.Valor.Value().(float64), 5)
        case tiposDeDato.TIPO_CADENA:
            output += arg.Valor.Value().(string)
        case tiposDeDato.TIPO_NULO:
            output += "nil"
        default:
            // ✅ NUEVO: Soporte prioritario para TipoMatrizMulti
            if matrizMulti, ok := arg.Valor.(*TipoMatrizMulti); ok {
                output += "["
                for fila, row := range matrizMulti.ValorInterno {
                    if fila > 0 {
                        output += ", "
                    }
                    output += "["
                    for col, elemento := range row {
                        if col > 0 {
                            output += ", "
                        }
                        switch elemento.Type() {
                        case tiposDeDato.TIPO_BOOLEAN:
                            output += strconv.FormatBool(elemento.Value().(bool))
                        case tiposDeDato.TIPO_ENTERO:
                            output += strconv.Itoa(elemento.Value().(int))
                        case tiposDeDato.TIPO_DECIMAL:
                            output += formatearDecimal(elemento.Value().(float64), 5)
                        case tiposDeDato.TIPO_CADENA:
                            output += "\"" + elemento.Value().(string) + "\""
                        case tiposDeDato.TIPO_NULO:
                            output += "nil"
                        default:
                            output += "nil"
                        }
                    }
                    output += "]"
                }
                output += "]"
            } else if obj, ok := arg.Valor.(*TipoObjeto); ok {
                // Manejar structs
                output += formatStruct(obj)
            } else if EsTipoVector(arg.Valor.Type()) || strings.HasPrefix(arg.Valor.Type(), "[]") {
                // Manejar vectores y matrices tradicionales
                if vector, ok := arg.Valor.(*TipoVector); ok {
                    output += "["
                    for j, elemento := range vector.ValorInterno {
                        if j > 0 {
                            output += ", "
                        }
                        switch elemento.Type() {
                        case tiposDeDato.TIPO_BOOLEAN:
                            output += strconv.FormatBool(elemento.Value().(bool))
                        case tiposDeDato.TIPO_ENTERO:
                            output += strconv.Itoa(elemento.Value().(int))
                        case tiposDeDato.TIPO_DECIMAL:
                            output += formatearDecimal(elemento.Value().(float64), 5)
                        case tiposDeDato.TIPO_CADENA:
                            output += "\"" + elemento.Value().(string) + "\""
                        case tiposDeDato.TIPO_NULO:
                            output += "nil"
                        default:
                            if subVector, ok := elemento.(*TipoVector); ok {
                                // Manejar matrices multidimensionales tradicionales
                                output += "["
                                for k, subElemento := range subVector.ValorInterno {
                                    if k > 0 {
                                        output += ", "
                                    }
                                    switch subElemento.Type() {
                                    case tiposDeDato.TIPO_BOOLEAN:
                                        output += strconv.FormatBool(subElemento.Value().(bool))
                                    case tiposDeDato.TIPO_ENTERO:
                                        output += strconv.Itoa(subElemento.Value().(int))
                                    case tiposDeDato.TIPO_DECIMAL:
                                        output += formatearDecimal(subElemento.Value().(float64), 5)
                                    case tiposDeDato.TIPO_CADENA:
                                        output += "\"" + subElemento.Value().(string) + "\""
                                    default:
                                        output += "nil"
                                    }
                                }
                                output += "]"
                            } else if obj, ok := elemento.(*TipoObjeto); ok {
                                output += formatStruct(obj)
                            } else {
                                output += "nil"
                            }
                        }
                    }
                    output += "]"
                } else if matriz, ok := arg.Valor.(*TipoMatriz); ok {
                    // ✅ CORREGIDO: Calcular dimensiones dinámicamente
                    output += "["
                    filas := len(matriz.ValorInterno) // Número de filas
                    
                    for fila := 0; fila < filas; fila++ {
                        if fila > 0 {
                            output += ", "
                        }
                        output += "["
                        
                        // Obtener la fila actual como vector
                        filaVector := matriz.ValorInterno[fila].(*TipoVector)
                        columnas := len(filaVector.ValorInterno) // Número de columnas en esta fila
                        
                        for col := 0; col < columnas; col++ {
                            if col > 0 {
                                output += ", "
                            }
                            elemento := filaVector.ValorInterno[col]
                            switch elemento.Type() {
                            case tiposDeDato.TIPO_BOOLEAN:
                                output += strconv.FormatBool(elemento.Value().(bool))
                            case tiposDeDato.TIPO_ENTERO:
                                output += strconv.Itoa(elemento.Value().(int))
                            case tiposDeDato.TIPO_DECIMAL:
                                output += formatearDecimal(elemento.Value().(float64), 5)
                            case tiposDeDato.TIPO_CADENA:
                                output += "\"" + elemento.Value().(string) + "\""
                            default:
                                output += "nil"
                            }
                        }
                        output += "]"
                    }
                    output += "]"
                } else {
                    return tiposDeDato.NuloPorDefecto, false, "La función println no puede imprimir el tipo: " + arg.Valor.Type()
                }
            } else {
                return tiposDeDato.NuloPorDefecto, false, "La función println no puede imprimir el tipo: " + arg.Valor.Type()
            }
        }

        if i < len(args)-1 {
            output += " "
        }
    }

    output += "\n"
    context.Consola.Print(output)
    return tiposDeDato.NuloPorDefecto, true, ""
}

// Función auxiliar para formatear structs
func formatStruct(obj *TipoObjeto) string {
	if obj == nil || obj.AmbitoInterno == nil {
		return "nil"
	}

	result := obj.TipoConcreto + "{"
	first := true

	// Iterar sobre las variables del scope interno del struct
	for name, variable := range obj.AmbitoInterno.Variables {
		if !first {
			result += ", "
		}
		first = false

		result += name + ": "

		// Formatear el valor según su tipo
		switch variable.Valor.Type() {
		case tiposDeDato.TIPO_CADENA:
			result += "\"" + variable.Valor.Value().(string) + "\""
		case tiposDeDato.TIPO_ENTERO:
			result += strconv.Itoa(variable.Valor.Value().(int))
		case tiposDeDato.TIPO_DECIMAL:
			result += strconv.FormatFloat(variable.Valor.Value().(float64), 'f', 2, 64)
		case tiposDeDato.TIPO_BOOLEAN:
			result += strconv.FormatBool(variable.Valor.Value().(bool))
		case tiposDeDato.TIPO_NULO:
			result += "nil"
		default:
			// Si es otro struct anidado
			if nestedObj, ok := variable.Valor.(*TipoObjeto); ok {
				result += formatStruct(nestedObj)
			} else {
				result += "nil"
			}
		}
	}

	result += "}"
	return result
}

func Len(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

    if len(args) != 1 {
        return tiposDeDato.NuloPorDefecto, false, "La función len solo acepta un argumento"
    }

    argValue := args[0].Valor

    switch v := argValue.(type) {
    case *TipoVector:
        // ✅ Manejar vectores normales Y filas de matrices
        return &tiposDeDato.ValorEntero{
            ValorInterno: v.Size(),
        }, true, ""
        
    case *tiposDeDato.ValorCadena:
        return &tiposDeDato.ValorEntero{
            ValorInterno: len(v.ValorInterno),
        }, true, ""
        
    case *TipoMatriz:
        // Para matrices tradicionales, contar filas
        return &tiposDeDato.ValorEntero{
            ValorInterno: v.Size(),
        }, true, ""
        
    case *TipoMatrizMulti:
        return &tiposDeDato.ValorEntero{
            ValorInterno: v.Filas,
        }, true, ""
        
    default:
        return tiposDeDato.NuloPorDefecto, false, "La función len solo acepta vectores, matrices o cadenas"
    }
}

// Función IndexOf - soporta matrices multidimensionales
func IndexOf(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false, "La función indexOf requiere exactamente 2 argumentos: indexOf(vector/matriz, valor)"
	}

	vectorArg := args[0].Valor
	searchValue := args[1].Valor

	// Verificar si es una matriz multidimensional
	if matrizMulti, ok := vectorArg.(*TipoMatrizMulti); ok {
		for filaIdx, fila := range matrizMulti.ValorInterno {
			for colIdx, item := range fila {
				if item.Value() == searchValue.Value() && item.Type() == searchValue.Type() {
					items := []tiposDeDato.ValorInterno{
						&tiposDeDato.ValorEntero{ValorInterno: filaIdx},
						&tiposDeDato.ValorEntero{ValorInterno: colIdx},
					}

					resultVector := &TipoVector{
						ValorInterno: items,
						TipoCompleto: "[]int",
						TipoElemento: "int",
					}
					return resultVector, true, ""
				}
			}
		}
		// Si no se encuentra, retornar [-1, -1]
		items := []tiposDeDato.ValorInterno{
			&tiposDeDato.ValorEntero{ValorInterno: -1},
			&tiposDeDato.ValorEntero{ValorInterno: -1},
		}

		resultVector := &TipoVector{
			ValorInterno: items,
			TipoCompleto: "[]int",
			TipoElemento: "int",
		}
		return resultVector, true, ""
	}

	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de indexOf debe ser un vector o matriz"
	}

	for i, item := range vector.ValorInterno {
		if item.Value() == searchValue.Value() && item.Type() == searchValue.Type() {
			return &tiposDeDato.ValorEntero{
				ValorInterno: i,
			}, true, ""
		}
	}

	// Si no se encuentra, retornar -1
	return &tiposDeDato.ValorEntero{
		ValorInterno: -1,
	}, true, ""
}

func Join(ctx *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false,
			"La función join requiere exactamente 2 argumentos: join(vector, separador)"
	}

	vectorArg := args[0].Valor
	separatorArg := args[1].Valor

	if vectorArg == nil {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de join es nil"
	}
	if separatorArg == nil {
		return tiposDeDato.NuloPorDefecto, false, "El segundo argumento de join es nil"
	}

	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false,
			"El primer argumento de join debe ser un vector (tipo: " + vectorArg.Type() + ")"
	}

	// Solo vectores de []string
	if vector.TipoCompleto != "[]string" {
		return tiposDeDato.NuloPorDefecto, false,
			"La función join solo acepta vectores de tipo []string (tipo recibido: " + vector.TipoCompleto + ")"
	}

	sepCadena, ok := separatorArg.(*tiposDeDato.ValorCadena)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false,
			"El segundo argumento de join debe ser una cadena (tipo recibido: " + separatorArg.Type() + ")"
	}
	separatorStr := sepCadena.ValorInterno

	var sb strings.Builder
	for i, item := range vector.ValorInterno {
		strItem := item.(*tiposDeDato.ValorCadena).ValorInterno
		sb.WriteString(strItem)
		if i < len(vector.ValorInterno)-1 {
			sb.WriteString(separatorStr)
		}
	}

	return &tiposDeDato.ValorCadena{ValorInterno: sb.String()}, true, ""
}

// Función Append - soporta matrices multidimensionales
func Append(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false, "La función append requiere exactamente 2 argumentos: append(vector/matriz, elemento)"
	}

	vectorArg := args[0].Valor
	elementArg := args[1].Valor

	// Verificar si es una matriz multidimensional
	if matrizMulti, ok := vectorArg.(*TipoMatrizMulti); ok {
		if elementVector, ok := elementArg.(*TipoVector); ok {
			if elementVector.TipoCompleto != matrizMulti.TipoElemento {
				return tiposDeDato.NuloPorDefecto, false, "No se puede agregar un vector de tipo " + elementVector.TipoCompleto + " a una matriz de tipo " + matrizMulti.TipoCompleto
			}

			// Crear una copia de la matriz original
			newMatrix := make([][]tiposDeDato.ValorInterno, len(matrizMulti.ValorInterno))
			for i, row := range matrizMulti.ValorInterno {
				newMatrix[i] = make([]tiposDeDato.ValorInterno, len(row))
				copy(newMatrix[i], row)
			}

			newMatrix = append(newMatrix, elementVector.ValorInterno)

			return NewTipoMatrizMulti(newMatrix, matrizMulti.TipoCompleto, matrizMulti.TipoElemento, matrizMulti.TipoBase), true, ""
		} else {
			return tiposDeDato.NuloPorDefecto, false, "Solo se pueden agregar vectores (filas) a una matriz multidimensional"
		}
	}

	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de append debe ser un vector o matriz"
	}

	expectedElementType := vector.TipoElemento
	actualElementType := elementArg.Type()

	if expectedElementType != actualElementType {
		return tiposDeDato.NuloPorDefecto, false, "No se puede agregar un elemento de tipo " + actualElementType + " a un vector de tipo " + vector.TipoCompleto
	}

	// Crear una copia del vector original
	newItems := make([]tiposDeDato.ValorInterno, len(vector.ValorInterno))
	copy(newItems, vector.ValorInterno)

	newItems = append(newItems, elementArg)

	newVector := NewTipoVector(newItems, vector.TipoCompleto, vector.TipoElemento)

	return newVector, true, ""
}

func Atoi(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' solo acepta un argumento"
	}

	arg := args[0].Valor

	if arg.Type() != tiposDeDato.TIPO_CADENA {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' solo acepta argumentos de tipo string"
	}

	str := arg.Value().(string)

	// Verificar si contiene punto (decimal)
	if _, err := strconv.ParseFloat(str, 64); err == nil && containsDot(str) {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' no permite valores decimales"
	}

	valor, err := strconv.Atoi(str)
	if err != nil {
		return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el string a entero"
	}

	return &tiposDeDato.ValorEntero{
		ValorInterno: valor,
	}, true, ""
}

// Función auxiliar para detectar decimales
func containsDot(input string) bool {
	for _, ch := range input {
		if ch == '.' {
			return true
		}
	}
	return false
}

func ParseFloat(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función 'parseFloat' solo acepta un argumento"
	}

	arg := args[0].Valor

	if arg.Type() != tiposDeDato.TIPO_CADENA {
		return tiposDeDato.NuloPorDefecto, false, "La función 'parseFloat' solo acepta argumentos de tipo string"
	}

	str := arg.Value().(string)

	valor, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el string a decimal"
	}

	return &tiposDeDato.ValorDecimal{
		InternalValor: valor,
	}, true, ""
}

func TypeOf(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función typeOf espera exactamente un argumento"
	}

	if args[0] == nil {
		return tiposDeDato.NuloPorDefecto, false, "El argumento de typeOf no puede ser nil"
	}

	if args[0].Valor == nil {
		return tiposDeDato.NuloPorDefecto, false, "El valor del argumento de typeOf no puede ser nil"
	}

	valor := args[0].Valor
	var tipo string

	switch v := valor.(type) {
	case *TipoVector:
		if v.TipoElemento != "" {
			// Mapear tipos específicos
			baseType := v.TipoElemento
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO || baseType == "int" {
				baseType = "int"
			} else if baseType == tiposDeDato.TIPO_CADENA || baseType == "string" {
				baseType = "string"
			} else if baseType == tiposDeDato.TIPO_BOOLEAN || baseType == "bool" {
				baseType = "bool"
			}
			tipo = "[]" + baseType
		} else if v.TipoCompleto != "" {
			tipo = v.TipoCompleto
		} else if len(v.ValorInterno) > 0 {
			// Inferir del primer elemento
			primerElemento := v.ValorInterno[0]
			if primerElemento != nil {
				switch primerElemento.Type() {
				case tiposDeDato.TIPO_ENTERO:
					tipo = "[]int"
				case tiposDeDato.TIPO_DECIMAL:
					tipo = "[]f64"
				case tiposDeDato.TIPO_CADENA:
					tipo = "[]string"
				case tiposDeDato.TIPO_BOOLEAN:
					tipo = "[]bool"
				default:
					tipo = "[]any"
				}
			} else {
				tipo = "[]any"
			}
		} else {
			tipo = "[]any"
		}

	case *TipoMatriz:
		if v.TipoElemento != "" {
			baseType := v.TipoElemento
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO {
				baseType = "int"
			}
			tipo = "[][]" + baseType
		} else {
			tipo = "[][]any"
		}

	case *TipoMatrizMulti:
		if v.TipoBase != "" {
			baseType := v.TipoBase
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO {
				baseType = "int"
			}
			tipo = "[][]" + baseType
		} else {
			tipo = "[][]any"
		}

	default:
		tipoOriginal := valor.Type()

		switch tipoOriginal {
		case tiposDeDato.TIPO_ENTERO:
			tipo = "int"
		case tiposDeDato.TIPO_DECIMAL:
			tipo = "f64"
		case tiposDeDato.TIPO_CADENA:
			tipo = "string"
		case tiposDeDato.TIPO_BOOLEAN:
			tipo = "bool"
		case tiposDeDato.TIPO_NULO:
			tipo = "nil"
		default:
			if tipoOriginal != "" {
				tipo = tipoOriginal
			} else {
				tipo = "unknown"
			}
		}
	}

	// Verificar que el resultado no sea vacío
	if tipo == "" {
		tipo = "unknown"
	}

	return &tiposDeDato.ValorCadena{
		ValorInterno: tipo,
	}, true, ""
}

var FuncionesEmbebidas = map[string]*FuncionNativa{
	"print": {
		Nombre: "print",
		Exec:   Imprimir,
	},
	"println": {
		Nombre: "println",
		Exec:   ImprimirLn,
	},
	"atoi": {
		Nombre: "atoi",
		Exec:   Atoi,
	},
	"parseFloat": {
		Nombre: "parseFloat",
		Exec:   ParseFloat,
	},
	"typeOf": {
		Nombre: "typeOf",
		Exec:   TypeOf,
	},
	"len": {
		Nombre: "len",
		Exec:   Len,
	},
	"indexOf": {
		Nombre: "indexOf",
		Exec:   IndexOf,
	},
	"join": {
		Nombre: "join",
		Exec:   Join,
	},
	"append": {
		Nombre: "append",
		Exec:   Append,
	},
}
