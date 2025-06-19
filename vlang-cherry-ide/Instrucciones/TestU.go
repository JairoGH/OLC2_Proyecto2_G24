package instrucciones

import "testing"

// ProbarTipoVector: prueba la función EsTipoVector con diferentes casos de entrada
func ProbarTipoVector(t *testing.T) {
    // 🔧 CASOS ACTUALIZADOS AL FORMATO []tipo
    casos := map[string]bool{
        // ✅ VECTORES VÁLIDOS (formato []tipo)
        "[]int":    true,
        "[]string": true,
        "[]float64":  true,
        "[]bool":   true,
        
        // ❌ NO SON VECTORES
        "int":      false,  // Tipo primitivo
        "string":   false,  // Tipo primitivo
        "[]":       false,  // Sin tipo de elemento
        "":         false,  // Cadena vacía
        
        // ❌ SON MATRICES (no vectores de un nivel)
        "[][]int":     false,  // Matriz 2D
        "[][][]int":   false,  // Matriz 3D
        "[][][][]int": false,  // Matriz 4D
        
        // ❌ FORMATO INCORRECTO
        "[int]":    false,  // Formato antiguo/legacy
        "[string]": false,  // Formato antiguo/legacy
        "[[int]]":  false,  // Formato antiguo/legacy
        "[":        false,  // Incompleto
        "]":        false,  // Incompleto
        "[]int]":   false,  // Mal formado
        "[[]int":   false,  // Mal formado
    }

    for tipo, esperado := range casos {
        resultado := EsTipoVector(tipo)
        if resultado != esperado {
            t.Errorf("EsTipoVector(%q) = %t, esperado %t", tipo, resultado, esperado)
        }
    }
}

// ProbarTipoMatriz: prueba la función EsTipoMatriz con diferentes casos de entrada
func ProbarTipoMatriz(t *testing.T) {
    // 🔧 CASOS ACTUALIZADOS AL FORMATO [][]tipo
    casos := map[string]bool{
        // ✅ MATRICES VÁLIDAS (formato [][]tipo, [][][]tipo, etc.)
        "[][]int":       true,   // Matriz 2D
        "[][][]string":  true,   // Matriz 3D
        "[][][][]float64": true,   // Matriz 4D
        "[][]bool":      true,   // Matriz 2D booleana
        
        // ❌ NO SON MATRICES
        "[]int":    false,  // Vector (un solo nivel)
        "[]string": false,  // Vector (un solo nivel)
        "int":      false,  // Tipo primitivo
        "string":   false,  // Tipo primitivo
        "[]":       false,  // Sin tipo de elemento
        "":         false,  // Cadena vacía
        
        // ❌ FORMATO INCORRECTO
        "[[int]]":     false,  // Formato antiguo/legacy
        "[[[int]]]":   false,  // Formato antiguo/legacy
        "[int]":       false,  // Formato antiguo/legacy
        "[":           false,  // Incompleto
        "]":           false,  // Incompleto
        "[[]":         false,  // Incompleto
        "[]]":         false,  // Mal formado
        "[][]int]":    false,  // Mal formado
        "[[][]int":    false,  // Mal formado
        "[][] int":    false,  // Con espacio
        "[] []int":    false,  // Con espacio
    }

    for tipo, esperado := range casos {
        resultado := EsTipoMatriz(tipo)
        if resultado != esperado {
            t.Errorf("EsTipoMatriz(%q) = %t, esperado %t", tipo, resultado, esperado)
        }
    }
}

// 🆕 NUEVA FUNCIÓN: ProbarObtenerTipoElemento
func ProbarObtenerTipoElemento(t *testing.T) {
    casos := map[string]string{
        "[]int":       "int",
        "[]string":    "string",
        "[][]int":     "int",
        "[][][]bool":  "bool",
        "[][][][]float64": "float64",
        "int":         "int",     // Sin corchetes
        "":            "",        // Cadena vacía
    }

    for entrada, esperado := range casos {
        resultado := ObtenerTipoElemento(entrada)
        if resultado != esperado {
            t.Errorf("ObtenerTipoElemento(%q) = %q, esperado %q", entrada, resultado, esperado)
        }
    }
}

// 🆕 NUEVA FUNCIÓN: ProbarObtenerNivelAnidamiento
func ProbarObtenerNivelAnidamiento(t *testing.T) {
    casos := map[string]int{
        "[]int":         1,  // Vector
        "[][]int":       2,  // Matriz 2D
        "[][][]int":     3,  // Matriz 3D
        "[][][][]int":   4,  // Matriz 4D
        "int":           0,  // Sin corchetes
        "":              0,  // Cadena vacía
    }

    for entrada, esperado := range casos {
        resultado := ObtenerNivelAnidamiento(entrada)
        if resultado != esperado {
            t.Errorf("ObtenerNivelAnidamiento(%q) = %d, esperado %d", entrada, resultado, esperado)
        }
    }
}

// 🆕 NUEVA FUNCIÓN: ProbarConstruirTipoVector
func ProbarConstruirTipoVector(t *testing.T) {
    casos := map[string]string{
        "int":    "[]int",
        "string": "[]string",
        "bool":   "[]bool",
        "float64":  "[]float64",
        "":       "[]",  // Caso edge
    }

    for entrada, esperado := range casos {
        resultado := ConstruirTipoVector(entrada)
        if resultado != esperado {
            t.Errorf("ConstruirTipoVector(%q) = %q, esperado %q", entrada, resultado, esperado)
        }
    }
}

// 🆕 NUEVA FUNCIÓN: ProbarConstruirTipoMatriz
func ProbarConstruirTipoMatriz(t *testing.T) {
    casos := []struct {
        tipoElemento string
        nivel        int
        esperado     string
    }{
        {"int", 1, "[]int"},
        {"int", 2, "[][]int"},
        {"int", 3, "[][][]int"},
        {"string", 2, "[][]string"},
        {"bool", 4, "[][][][]bool"},
        {"float64", 0, "float64"},  // Nivel 0
    }

    for _, caso := range casos {
        resultado := ConstruirTipoMatriz(caso.tipoElemento, caso.nivel)
        if resultado != caso.esperado {
            t.Errorf("ConstruirTipoMatriz(%q, %d) = %q, esperado %q", 
                caso.tipoElemento, caso.nivel, resultado, caso.esperado)
        }
    }
}

// 🆕 NUEVA FUNCIÓN: ProbarEsTipoIterable
func ProbarEsTipoIterable(t *testing.T) {
    casos := map[string]bool{
        // ✅ ITERABLES
        "[]int":     true,  // Vector
        "[][]int":   true,  // Matriz
        "string":    true,  // Cadena (asumiendo que TIPO_CADENA = "string")
        
        // ❌ NO ITERABLES
        "int":       false, // Primitivo
        "bool":      false, // Primitivo
        "float64":     false, // Primitivo
        "":          false, // Vacío
    }

    for tipo, esperado := range casos {
        resultado := EsTipoIterable(tipo)
        if resultado != esperado {
            t.Errorf("EsTipoIterable(%q) = %t, esperado %t", tipo, resultado, esperado)
        }
    }
}