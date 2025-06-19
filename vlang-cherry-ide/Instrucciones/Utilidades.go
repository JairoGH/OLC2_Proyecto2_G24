package instrucciones

import (
    "main/tiposDeDato"
    "strings"
)

// CadenaAVector: convierte una cadena en un vector de caracteres individuales
func CadenaAVector(s *tiposDeDato.ValorCadena) *TipoVector {
    items := make([]tiposDeDato.ValorInterno, 0, len(s.ValorInterno))
    for _, rune := range s.ValorInterno {
        items = append(items, &tiposDeDato.ValorCadena{
            ValorInterno: string(rune),
        })
    }

    // 🔧 CAMBIO: usar formato []tipo
    tag := "[]" + tiposDeDato.TIPO_CADENA
    return NewTipoVector(items, tag, tiposDeDato.TIPO_CADENA)
}

// 🔧 FUNCIÓN ACTUALIZADA: EsTipoVector
// Verifica si un tipo es vector de un solo nivel: []tipo
func EsTipoVector(tipo string) bool {
    // ✅ Vector formato V-Lang: []tipo (ej: []int, []string)
    
    // Debe empezar con [] exactamente
    if !strings.HasPrefix(tipo, "[]") {
        return false
    }
    
    // No debe ser matriz ([][]tipo, [][][]tipo, etc.)
    if strings.HasPrefix(tipo, "[][]") {
        return false
    }
    
    // Debe tener contenido después de []
    tipoElemento := strings.TrimPrefix(tipo, "[]")
    if len(tipoElemento) == 0 {
        return false
    }
    
    // El tipo de elemento no debe contener más corchetes
    if strings.Contains(tipoElemento, "[") || strings.Contains(tipoElemento, "]") {
        return false
    }
    
    return true
}

// 🔧 FUNCIÓN ACTUALIZADA: EsTipoMatriz
// Verifica si un tipo es matriz de múltiples niveles: [][]tipo, [][][]tipo, etc.
func EsTipoMatriz(tipo string) bool {
    // ✅ Matriz formato V-Lang: [][]tipo, [][][]tipo, etc.
    
    // Debe empezar con al menos [][]
    if !strings.HasPrefix(tipo, "[][]") {
        return false
    }
    
    // Contar cuántos [] hay al inicio
    temp := tipo
    nivelAnidamiento := 0
    
    for strings.HasPrefix(temp, "[]") {
        nivelAnidamiento++
        temp = strings.TrimPrefix(temp, "[]")
    }
    
    // Es matriz si tiene al menos 2 niveles y el resto es un tipo válido
    if nivelAnidamiento < 2 {
        return false
    }
    
    // El tipo base no debe contener corchetes
    if len(temp) == 0 || strings.Contains(temp, "[") || strings.Contains(temp, "]") {
        return false
    }
    
    return true
}

// 🔧 FUNCIÓN ACTUALIZADA: EliminarCorchetes
// Remueve un nivel de corchetes del formato []tipo
func EliminarCorchetes(s string) string {
    // Para formato []tipo (quitar un nivel)
    if strings.HasPrefix(s, "[]") {
        return strings.TrimPrefix(s, "[]")
    }
    
    // Para formato legacy [tipo] (compatibilidad)
    if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
        return strings.Trim(s, "[]")
    }
    
    return s
}

// 🆕 FUNCIÓN NUEVA: ObtenerTipoElemento
// Obtiene el tipo de elemento base de un vector o matriz
func ObtenerTipoElemento(tipo string) string {
    temp := tipo
    
    // Remover todos los [] del inicio
    for strings.HasPrefix(temp, "[]") {
        temp = strings.TrimPrefix(temp, "[]")
    }
    
    return temp
}

// 🆕 FUNCIÓN NUEVA: ObtenerNivelAnidamiento
// Obtiene el nivel de anidamiento de un vector/matriz
func ObtenerNivelAnidamiento(tipo string) int {
    temp := tipo
    nivel := 0
    
    // Contar cuántos [] hay al inicio
    for strings.HasPrefix(temp, "[]") {
        nivel++
        temp = strings.TrimPrefix(temp, "[]")
    }
    
    return nivel
}

// 🆕 FUNCIÓN NUEVA: EsTipoIterable
// Verifica si un tipo se puede iterar (vector, matriz o cadena)
func EsTipoIterable(tipo string) bool {
    return EsTipoVector(tipo) || EsTipoMatriz(tipo) || tipo == tiposDeDato.TIPO_CADENA
}

// 🆕 FUNCIÓN NUEVA: ConstruirTipoVector
// Construye un tipo vector a partir del tipo de elemento
func ConstruirTipoVector(tipoElemento string) string {
    return "[]" + tipoElemento
}

// 🆕 FUNCIÓN NUEVA: ConstruirTipoMatriz
// Construye un tipo matriz a partir del tipo de elemento y nivel
func ConstruirTipoMatriz(tipoElemento string, nivel int) string {
    resultado := tipoElemento
    for i := 0; i < nivel; i++ {
        resultado = "[]" + resultado
    }
    return resultado
}

// 🆕 FUNCIÓN NUEVA: ValidarTipoVector
// Valida si un string representa un tipo vector válido
func ValidarTipoVector(tipo string) bool {
    if !EsTipoVector(tipo) {
        return false
    }
    
    tipoElemento := ObtenerTipoElemento(tipo)
    
    // Verificar que el tipo de elemento es válido
    tiposValidos := []string{
        tiposDeDato.TIPO_ENTERO,
        tiposDeDato.TIPO_DECIMAL,
        tiposDeDato.TIPO_CADENA,
        tiposDeDato.TIPO_BOOLEAN,
    }
    
    for _, tipoValido := range tiposValidos {
        if tipoElemento == tipoValido {
            return true
        }
    }
    
    return false
}

// 🆕 FUNCIÓN NUEVA: ValidarTipoMatriz
// Valida si un string representa un tipo matriz válido
func ValidarTipoMatriz(tipo string) bool {
    if !EsTipoMatriz(tipo) {
        return false
    }
    
    tipoElemento := ObtenerTipoElemento(tipo)
    nivel := ObtenerNivelAnidamiento(tipo)
    
    // Verificar límites razonables
    if nivel > 10 { // Máximo 10 niveles de anidamiento
        return false
    }
    
    // Verificar que el tipo de elemento es válido
    tiposValidos := []string{
        tiposDeDato.TIPO_ENTERO,
        tiposDeDato.TIPO_DECIMAL,
        tiposDeDato.TIPO_CADENA,
        tiposDeDato.TIPO_BOOLEAN,
    }
    
    for _, tipoValido := range tiposValidos {
        if tipoElemento == tipoValido {
            return true
        }
    }
    
    return false
}