package assembly

import (
	"fmt"
	"strings"
)

// SliceInfo representa la información de un slice en ARM64
type SliceInfo struct {
	Nombre       string
	TipoElemento string
	Tamaño       int
	OffsetStack  int
	EsLiteral    bool
	Elementos    []*ResultadoExpresion
}

// SliceProcessor maneja todas las operaciones de slices
type SliceProcessor struct {
	armGen               *ARMGenerator
	expresionesProcessor *ExpresionesProcessor
	slicesDeclarados     map[string]*SliceInfo
	offsetStack          int
	stackAdjusted        bool
}

// NewSliceProcessor crea un nuevo procesador de slices
func NewSliceProcessor(armGen *ARMGenerator, expresionesProcessor *ExpresionesProcessor) *SliceProcessor {
	return &SliceProcessor{
		armGen:               armGen,
		expresionesProcessor: expresionesProcessor,
		slicesDeclarados:     make(map[string]*SliceInfo),
		offsetStack:          0,     //  CAMBIAR: Empezar desde 0
		stackAdjusted:        false, //  NUEVO
	}
}

// ============= MÉTODO PARA OBTENER TIPO DE ELEMENTO =============
func (sp *SliceProcessor) ObtenerTipoElemento(nombre string) (string, error) {
	sliceInfo, err := sp.CargarSlice(nombre)
	if err != nil {
		return "", err
	}
	return sliceInfo.TipoElemento, nil
}

// ============= DECLARACIÓN DE SLICES =============
func (sp *SliceProcessor) DeclararSlice(nombre, tipoElemento string, elementos []*ResultadoExpresion) error {
	sp.armGen.Comment(fmt.Sprintf("=== DECLARACIÓN SLICE %s: []%s ===", nombre, tipoElemento))

	// Verificar si ya existe
	if _, existe := sp.slicesDeclarados[nombre]; existe {
		return fmt.Errorf("slice '%s' ya está declarado", nombre)
	}

	// Determinar tamaño
	tamaño := len(elementos)

	// AJUSTAR STACK SI ES EL PRIMER SLICE Y GUARDAR REGISTRO BASE
	if !sp.stackAdjusted && tamaño > 0 {
		// Calcular espacio total necesario para todos los elementos posibles
		espacioTotal := 256 // 256 bytes debería ser suficiente para la mayoría de casos
		sp.armGen.Comment(fmt.Sprintf("Ajustar stack para slices: %d bytes", espacioTotal))
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("sub sp, sp, #%d", espacioTotal),
			// 🔥 SOLUCIÓN SIMPLE: Usar offset fijo desde frame pointer
			"mov x28, x29", // x28 = frame pointer
			fmt.Sprintf("sub x28, x28, #%d", espacioTotal)) // x28 = dirección fija del slice
		sp.stackAdjusted = true
	}

	// Resto del código igual...
	offsetActual := sp.offsetStack
	sliceInfo := &SliceInfo{
		Nombre:       nombre,
		TipoElemento: tipoElemento,
		Tamaño:       tamaño,
		OffsetStack:  offsetActual,
		EsLiteral:    false,
		Elementos:    elementos,
	}

	if tamaño == 0 {
		sp.armGen.Comment(fmt.Sprintf("Slice vacío: %s", nombre))
		sp.DeclararSliceVacio(sliceInfo)
	} else {
		sp.DeclararSliceConElementos(sliceInfo)
		sp.offsetStack += (tamaño * 8)
	}

	sp.slicesDeclarados[nombre] = sliceInfo
	return nil
}

func (sp *SliceProcessor) DeclararSliceVacio(sliceInfo *SliceInfo) {
	sp.armGen.Comment("Inicializar slice vacío")
	// Para slices vacíos, no necesitamos hacer nada especial en ARM64
}

func (sp *SliceProcessor) DeclararSliceConElementos(sliceInfo *SliceInfo) {
	// Usar offsets positivos desde la base del stack ajustado
	for i, elemento := range sliceInfo.Elementos {
		// USAR FÓRMULA CORRECTA: offset base + (i * 8)
		offsetElemento := sliceInfo.OffsetStack + (i * 8)

		if elemento.EsLiteral {
			sp.armGen.Comment(fmt.Sprintf("Elemento %d (literal): %v", i, elemento.Valor))
			sp.AlmacenarElementoLiteral(elemento, offsetElemento)
		} else {
			sp.armGen.Comment(fmt.Sprintf("Elemento %d (registro): %s", i, elemento.Registro))
			sp.AlmacenarElementoRegistro(elemento, offsetElemento)
		}
	}
}

func (sp *SliceProcessor) AlmacenarElementoLiteral(elemento *ResultadoExpresion, offset int) {
	switch elemento.Tipo {
	case "int":
		valor := elemento.Valor.(int)
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov x0, #%d", valor),
			fmt.Sprintf("str x0, [x28, #%d]", offset)) // CAMBIO: usar x28 como base

	case "float":
		// Para floats, usar constante en sección de datos
		regFloat := sp.expresionesProcessor.NuevoRegistroFloatTmp()
		sp.expresionesProcessor.CargarConstanteFloat(regFloat, elemento.Valor.(float64))
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("str %s, [x28, #%d]", regFloat, offset)) // CAMBIO: usar x28 como base

	case "string":
		// Para strings, almacenar dirección
		etiqueta := sp.expresionesProcessor.AgregarMensajeString(elemento.Valor.(string))
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("adr x0, %s", etiqueta),
			fmt.Sprintf("str x0, [x28, #%d]", offset)) // CAMBIO: usar x28 como base

	case "bool":
		valor := 0
		if elemento.Valor.(bool) {
			valor = 1
		}
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov x0, #%d", valor),
			fmt.Sprintf("str x0, [x28, #%d]", offset)) // CAMBIO: usar x28 como base
	}
}

func (sp *SliceProcessor) AlmacenarElementoRegistro(elemento *ResultadoExpresion, offset int) {
	switch elemento.Tipo {
	case "int", "bool":
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("str %s, [x28, #%d]", elemento.Registro, offset)) // CAMBIO: usar x28 como base

	case "float":
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("str %s, [x28, #%d]", elemento.Registro, offset)) // CAMBIO: usar x28 como base

	case "string":
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("str %s, [x28, #%d]", elemento.Registro, offset)) // CAMBIO: usar x28 como base
	}
}

// ============= ACCESO A SLICES =============

func (sp *SliceProcessor) CargarSlice(nombre string) (*SliceInfo, error) {
	sliceInfo, existe := sp.slicesDeclarados[nombre]
	if !existe {
		return nil, fmt.Errorf("slice '%s' no está declarado", nombre)
	}
	return sliceInfo, nil
}

func (sp *SliceProcessor) AccederElemento(nombre string, indice int) (*ResultadoExpresion, error) {
	sliceInfo, err := sp.CargarSlice(nombre)
	if err != nil {
		return nil, err
	}

	if indice < 0 || indice >= sliceInfo.Tamaño {
		return nil, fmt.Errorf("índice %d fuera de rango para slice '%s' (tamaño: %d)",
			indice, nombre, sliceInfo.Tamaño)
	}

	sp.armGen.Comment(fmt.Sprintf("Acceder elemento %d del slice %s", indice, nombre))

	// USAR LA MISMA FÓRMULA QUE PARA ALMACENAR
	offsetElemento := sliceInfo.OffsetStack + (indice * 8)
	registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()

	// Cargar elemento desde la base fija x28
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("ldr %s, [x28, #%d]    // cargar elemento %d desde base fija",
			registroResultado, offsetElemento, indice))

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      sliceInfo.TipoElemento,
		EsLiteral: false,
	}, nil
}

// ============= IMPRESIÓN DE SLICES =============
func (sp *SliceProcessor) ImprimirSlice(nombre string) error {
	sliceInfo, err := sp.CargarSlice(nombre)
	if err != nil {
		return err
	}

	if sliceInfo.Tamaño == 0 {
		sp.ImprimirSliceVacio()
		return nil
	}

	sp.ImprimirSliceConElementos(sliceInfo)
	return nil
}

func (sp *SliceProcessor) ImprimirSliceVacio() {
	// Imprimir "[]"
	sp.armGen.UsarFuncion("print_char")
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		"mov x0, #91          // '['",
		"bl print_char",
		"mov x0, #93          // ']'",
		"bl print_char")
}

func (sp *SliceProcessor) ImprimirSliceConElementos(sliceInfo *SliceInfo) {
	// Imprimir "["
	sp.armGen.UsarFuncion("print_char")
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		"mov x0, #91          // '['",
		"bl print_char")

	// Imprimir cada elemento
	for i := 0; i < sliceInfo.Tamaño; i++ {
		if i > 0 {
			// Imprimir ", " entre elementos
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				"mov x0, #44          // ','",
				"bl print_char",
				"mov x0, #32          // ' '",
				"bl print_char")
		}

		// Cargar y imprimir elemento
		sp.ImprimirElementoEnIndice(sliceInfo, i)
	}

	// Imprimir "]"
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		"mov x0, #93          // ']'",
		"bl print_char")
}

func (sp *SliceProcessor) ImprimirElementoEnIndice(sliceInfo *SliceInfo, indice int) {
	// USAR LA MISMA FÓRMULA QUE PARA ALMACENAR
	offsetElemento := sliceInfo.OffsetStack + (indice * 8)

	switch sliceInfo.TipoElemento {
	case "int":
		sp.armGen.UsarFuncion("print_int")
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr x0, [x28, #%d]    // cargar int[%d] desde base fija", offsetElemento, indice),
			"bl print_int")

	case "float":
		sp.armGen.UsarFuncion("print_float")
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr d0, [x28, #%d]    // cargar float[%d] desde base fija", offsetElemento, indice),
			"bl print_float")

	case "string":
		sp.armGen.UsarFuncion("print_string")
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr x0, [x28, #%d]    // cargar string[%d] desde base fija", offsetElemento, indice),
			"bl print_string")

	case "bool":
		sp.armGen.UsarFuncion("print_bool")
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr x0, [x28, #%d]    // cargar bool[%d] desde base fija", offsetElemento, indice),
			"bl print_bool")
	}
}

// ============= UTILIDADES =============
func (sp *SliceProcessor) GetSlicesDeclarados() map[string]*SliceInfo {
	return sp.slicesDeclarados
}

func (sp *SliceProcessor) ExisteSlice(nombre string) bool {
    _, existe := sp.slicesDeclarados[nombre]
    sp.armGen.Comment(fmt.Sprintf("Verificando slice '%s': %t", nombre, existe))
    return existe
}

func (sp *SliceProcessor) TamañoSlice(nombre string) (int, error) {
	sliceInfo, err := sp.CargarSlice(nombre)
	if err != nil {
		return 0, err
	}
	return sliceInfo.Tamaño, nil
}

// ============= PARSE DE LISTA DE SLICES =============
func (sp *SliceProcessor) ParsearListaSlice(elementos []interface{}) ([]*ResultadoExpresion, string, error) {
	if len(elementos) == 0 {
		return []*ResultadoExpresion{}, "", nil
	}

	var resultados []*ResultadoExpresion
	var tipoElemento string

	for i, elem := range elementos {
		if resultado, ok := elem.(*ResultadoExpresion); ok {
			if i == 0 {
				// Primer elemento determina el tipo
				tipoElemento = resultado.Tipo
			} else {
				// Verificar que todos los elementos sean del mismo tipo
				if resultado.Tipo != tipoElemento {
					return nil, "", fmt.Errorf("todos los elementos del slice deben ser del mismo tipo. Esperado: %s, encontrado: %s", tipoElemento, resultado.Tipo)
				}
			}
			resultados = append(resultados, resultado)
		} else {
			return nil, "", fmt.Errorf("elemento %d no es una expresión válida", i)
		}
	}

	return resultados, tipoElemento, nil
}

// Método para limpiar el stack al final del programa
func (sp *SliceProcessor) LimpiarStack() {
	if sp.stackAdjusted {
		espacioTotal := 256
		sp.armGen.Comment("Limpiar stack de slices")
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("add sp, sp, #%d", espacioTotal))
	}
}

// ============= FUNCIONES EMBEBIDAS =============
func (sp *SliceProcessor) IndexOf(nombreSlice string, valorBuscado *ResultadoExpresion) (*ResultadoExpresion, error) {
	sp.armGen.Comment(fmt.Sprintf("=== FUNCIÓN indexOf(%s, %v) ===", nombreSlice, valorBuscado.Valor))

	sliceInfo, err := sp.CargarSlice(nombreSlice)
	if err != nil {
		return nil, err
	}

	// Verificar que el tipo del valor buscado coincida con el tipo del slice
	if valorBuscado.Tipo != sliceInfo.TipoElemento {
		return nil, fmt.Errorf("no se puede buscar un valor de tipo %s en slice de tipo %s",
			valorBuscado.Tipo, sliceInfo.TipoElemento)
	}

	sp.armGen.Comment(fmt.Sprintf("Buscando %v en slice %s (tamaño: %d)",
		valorBuscado.Valor, nombreSlice, sliceInfo.Tamaño))

	// Si el slice está vacío, retornar -1 inmediatamente
	if sliceInfo.Tamaño == 0 {
		sp.armGen.Comment("Slice vacío, retornando -1")
		registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()
		// USAR EL MISMO MÉTODO SIMPLE
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, #0", registroResultado),
			fmt.Sprintf("sub %s, %s, #1", registroResultado, registroResultado)) // 0 - 1 = -1
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      "int",
			EsLiteral: false,
		}, nil
	}

	// Generar código de búsqueda
	return sp.generarCodigoBusqueda(sliceInfo, valorBuscado)
}

// generarCodigoBusqueda
func (sp *SliceProcessor) generarCodigoBusqueda(sliceInfo *SliceInfo, valorBuscado *ResultadoExpresion) (*ResultadoExpresion, error) {
	// Etiquetas para control de flujo
	labelLoop := fmt.Sprintf("indexOf_loop_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelFound := fmt.Sprintf("indexOf_found_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelNotFound := fmt.Sprintf("indexOf_not_found_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelEnd := fmt.Sprintf("indexOf_end_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())

	// Registros necesarios
	regIndice := "x9"        // Índice actual
	regElemento := "x10"     // Elemento del slice
	regResultado := "x11"    // Resultado final
	regValorBuscado := "x12" // Valor a buscar

	// Cargar valor buscado en registro
	if valorBuscado.EsLiteral {
		sp.cargarLiteralEnRegistro(valorBuscado, regValorBuscado)
	} else {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, %s", regValorBuscado, valorBuscado.Registro))
	}

	// Inicializar índice en 0
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("mov %s, #0", regIndice))

	// Etiqueta del loop principal
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelLoop))

	// Verificar si hemos llegado al final del slice
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("cmp %s, #%d", regIndice, sliceInfo.Tamaño),
		fmt.Sprintf("bge %s", labelNotFound))

	// Cargar elemento directamente
	offsetBase := sliceInfo.OffsetStack
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("add %s, x28, #%d", regElemento, offsetBase),                     // Base del slice
		fmt.Sprintf("ldr %s, [%s, %s, lsl #3]", regElemento, regElemento, regIndice)) // elemento[índice]

	// Comparar con valor buscado
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("cmp %s, %s", regElemento, regValorBuscado),
		fmt.Sprintf("beq %s", labelFound))

	// Incrementar índice y continuar
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("add %s, %s, #1", regIndice, regIndice),
		fmt.Sprintf("b %s", labelLoop))

	// Elemento encontrado
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelFound),
		fmt.Sprintf("mov %s, %s", regResultado, regIndice),
		fmt.Sprintf("b %s", labelEnd))

	// Elemento no encontrado - Usar constante desde memoria
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelNotFound),
		fmt.Sprintf("mov %s, #0", regResultado),                   // Inicializar en 0
		fmt.Sprintf("sub %s, %s, #1", regResultado, regResultado)) // Restar 1 para obtener -1

	// Final
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelEnd))

	sp.armGen.Comment(fmt.Sprintf("indexOf completado, resultado en %s", regResultado))

	return &ResultadoExpresion{
		Registro:  regResultado,
		Tipo:      "int",
		EsLiteral: false,
	}, nil
}

func (sp *SliceProcessor) cargarLiteralEnRegistro(literal *ResultadoExpresion, registro string) {
	switch literal.Tipo {
	case "int":
		valor := literal.Valor.(int)
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, #%d", registro, valor))

	case "bool":
		valor := 0
		if literal.Valor.(bool) {
			valor = 1
		}
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, #%d", registro, valor))

	case "string":
		// Para strings, cargar dirección
		etiqueta := sp.expresionesProcessor.AgregarMensajeString(literal.Valor.(string))
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("adr %s, %s", registro, etiqueta))

	case "float":
		// Para floats, necesitamos un registro float temporal
		regFloat := sp.expresionesProcessor.NuevoRegistroFloatTmp()
		sp.expresionesProcessor.CargarConstanteFloat(regFloat, literal.Valor.(float64))
		// Mover a registro general para comparación
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("fmov %s, %s", registro, regFloat))
	}
}

func (sp *SliceProcessor) Join(nombreSlice string, separador *ResultadoExpresion) (*ResultadoExpresion, error) {
    sp.armGen.Comment(fmt.Sprintf("=== FUNCIÓN join(%s, \"%s\") ===", nombreSlice, separador.Valor))

    sliceInfo, err := sp.CargarSlice(nombreSlice)
    if err != nil {
        return nil, err
    }

    // Verificar que el slice sea de tipo string
    if sliceInfo.TipoElemento != "string" {
        return nil, fmt.Errorf("join() solo funciona con slices de tipo []string, recibido: []%s", sliceInfo.TipoElemento)
    }

    // Verificar que el separador sea string
    if separador.Tipo != "string" {
        return nil, fmt.Errorf("el separador de join() debe ser string, recibido: %s", separador.Tipo)
    }

    sp.armGen.Comment(fmt.Sprintf("Uniendo %d elementos con separador", sliceInfo.Tamaño))

    // Si el slice está vacío, retornar string vacío
    if sliceInfo.Tamaño == 0 {
        sp.armGen.Comment("Slice vacío, retornando string vacío")
        registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()
        etiquetaVacio := sp.expresionesProcessor.AgregarMensajeString("")
        sp.armGen.Instructions = append(sp.armGen.Instructions,
            fmt.Sprintf("adr %s, %s", registroResultado, etiquetaVacio))
        return &ResultadoExpresion{
            Registro:  registroResultado,
            Tipo:      "string",
            EsLiteral: false,
        }, nil
    }

    // 🔥 CORRECCIÓN: Si solo hay un elemento, retornarlo usando la base correcta
    if sliceInfo.Tamaño == 1 {
        sp.armGen.Comment("Solo un elemento, retornándolo directamente usando base x28")
        registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()
        offsetPrimerElemento := sliceInfo.OffsetStack
        
        // 🔥 CORRECCIÓN: Usar x28 como base en lugar de sp
        sp.armGen.Instructions = append(sp.armGen.Instructions,
            fmt.Sprintf("ldr %s, [x28, #%d]    // cargar único elemento del slice %s", 
                registroResultado, offsetPrimerElemento, nombreSlice))
        
        sp.armGen.Comment(fmt.Sprintf("Elemento único cargado desde x28 + %d", offsetPrimerElemento))
        
        return &ResultadoExpresion{
            Registro:  registroResultado,
            Tipo:      "string",
            EsLiteral: false,
        }, nil
    }

    // Generar código de concatenación para múltiples elementos
    return sp.generarCodigoJoin(sliceInfo, separador)
}

// CORREGIR: generarCodigoJoin - Error en intercambio de buffers
func (sp *SliceProcessor) generarCodigoJoin(sliceInfo *SliceInfo, separador *ResultadoExpresion) (*ResultadoExpresion, error) {
    sp.armGen.Comment("=== INICIO JOIN SIMPLIFICADO ===")

    joinId := sp.expresionesProcessor.GetContadorEtiqueta()
    labelLoop := fmt.Sprintf("join_loop_%s_%d", sliceInfo.Nombre, joinId)
    labelEnd := fmt.Sprintf("join_end_%s_%d", sliceInfo.Nombre, joinId)

    regIndice := "x9"
    regSeparador := "x11"
    regResultado := "x12"

    // PASO 1: Cargar separador
    if separador.EsLiteral {
        valorSeparador := separador.Valor.(string)
        if strings.HasPrefix(valorSeparador, "\"") && strings.HasSuffix(valorSeparador, "\"") {
            valorSeparador = valorSeparador[1 : len(valorSeparador)-1]
        }
        etiquetaSep := sp.expresionesProcessor.AgregarMensajeString(valorSeparador)
        sp.armGen.Instructions = append(sp.armGen.Instructions,
            fmt.Sprintf("adr %s, %s", regSeparador, etiquetaSep))
        sp.armGen.Comment(fmt.Sprintf("Separador: '%s'", valorSeparador))
    } else {
        sp.armGen.Instructions = append(sp.armGen.Instructions,
            fmt.Sprintf("mov %s, %s", regSeparador, separador.Registro))
    }

    // 🔥 SOLUCIÓN: CREAR BUFFER ÚNICO PARA CADA LLAMADA
    nombreBuffer := fmt.Sprintf("join_buffer_%d", joinId)
    
    // Declarar buffer único en la sección de datos
    bufferSize := 512
    sp.armGen.StringConstants[nombreBuffer] = strings.Repeat("\\0", bufferSize)
    
    // PASO 2: Obtener buffer resultado único
    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("adr %s, %s", regResultado, nombreBuffer))

    sp.armGen.Comment(fmt.Sprintf("Usando buffer único: %s", nombreBuffer))

    // PASO 3: Construcción manual (resto igual)
    regDestPtr := "x13"
    regSrcPtr := "x14"

    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("mov %s, %s", regDestPtr, regResultado))

    // PASO 4: Copiar primer elemento
    offsetBase := sliceInfo.OffsetStack
    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("add x16, x28, #%d", offsetBase),
        fmt.Sprintf("ldr %s, [x16, #0]", regSrcPtr))

    labelCopy1Loop := fmt.Sprintf("copy1_loop_%d", joinId)
    labelCopy1End := fmt.Sprintf("copy1_end_%d", joinId)

    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("%s:", labelCopy1Loop),
        fmt.Sprintf("ldrb w15, [%s], #1", regSrcPtr),
        "cmp w15, #0",
        fmt.Sprintf("beq %s", labelCopy1End),
        fmt.Sprintf("strb w15, [%s], #1", regDestPtr),
        fmt.Sprintf("b %s", labelCopy1Loop),
        fmt.Sprintf("%s:", labelCopy1End))

    // PASO 5: Loop para elementos restantes
    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("mov %s, #1", regIndice),
        fmt.Sprintf("%s:", labelLoop))

    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("cmp %s, #%d", regIndice, sliceInfo.Tamaño),
        fmt.Sprintf("bge %s", labelEnd))

    // PASO 6: Copiar separador
    labelCopySepLoop := fmt.Sprintf("copy_sep_loop_%d_%s", joinId, regIndice)
    labelCopySepEnd := fmt.Sprintf("copy_sep_end_%d_%s", joinId, regIndice)

    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("mov %s, %s", regSrcPtr, regSeparador),
        fmt.Sprintf("%s:", labelCopySepLoop),
        fmt.Sprintf("ldrb w15, [%s], #1", regSrcPtr),
        "cmp w15, #0",
        fmt.Sprintf("beq %s", labelCopySepEnd),
        fmt.Sprintf("strb w15, [%s], #1", regDestPtr),
        fmt.Sprintf("b %s", labelCopySepLoop),
        fmt.Sprintf("%s:", labelCopySepEnd))

    // PASO 7: Copiar elemento actual
    labelCopyElemLoop := fmt.Sprintf("copy_elem_loop_%d_%s", joinId, regIndice)
    labelCopyElemEnd := fmt.Sprintf("copy_elem_end_%d_%s", joinId, regIndice)

    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("add x16, x28, #%d", offsetBase),
        fmt.Sprintf("ldr %s, [x16, %s, lsl #3]", regSrcPtr, regIndice),
        fmt.Sprintf("%s:", labelCopyElemLoop),
        fmt.Sprintf("ldrb w15, [%s], #1", regSrcPtr),
        "cmp w15, #0",
        fmt.Sprintf("beq %s", labelCopyElemEnd),
        fmt.Sprintf("strb w15, [%s], #1", regDestPtr),
        fmt.Sprintf("b %s", labelCopyElemLoop),
        fmt.Sprintf("%s:", labelCopyElemEnd))

    // PASO 8: Incrementar índice
    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("add %s, %s, #1", regIndice, regIndice),
        fmt.Sprintf("b %s", labelLoop))

    // PASO 9: Terminar string
    sp.armGen.Instructions = append(sp.armGen.Instructions,
        fmt.Sprintf("%s:", labelEnd),
        "mov w15, #0",
        fmt.Sprintf("strb w15, [%s]", regDestPtr))

    sp.armGen.Comment("=== FIN JOIN SIMPLIFICADO ===")

    return &ResultadoExpresion{
        Registro:  regResultado,
        Tipo:      "string",
        EsLiteral: false,
    }, nil
}

// Función len para obtener tamaño de slice
func (sp *SliceProcessor) Len(nombreSlice string) (*ResultadoExpresion, error) {
	sp.armGen.Comment(fmt.Sprintf("=== FUNCIÓN len(%s) ===", nombreSlice))

	sliceInfo, err := sp.CargarSlice(nombreSlice)
	if err != nil {
		return nil, err
	}

	sp.armGen.Comment(fmt.Sprintf("Slice %s tiene %d elementos", nombreSlice, sliceInfo.Tamaño))

	// Crear resultado con el tamaño del slice
	 registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()

	// Cargar el tamaño del slice en el registro
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("mov %s, #%d", registroResultado, sliceInfo.Tamaño))

	return &ResultadoExpresion{
		Registro:  registroResultado,
		Tipo:      "int",
		EsLiteral: false,
	}, nil
}

// Función append para expandir slices
func (sp *SliceProcessor) Append(nombreSlice string, elemento *ResultadoExpresion) (*ResultadoExpresion, error) {
	sp.armGen.Comment(fmt.Sprintf("=== FUNCIÓN append(%s, %v) ===", nombreSlice, elemento.Valor))

	sliceInfo, err := sp.CargarSlice(nombreSlice)
	if err != nil {
		return nil, err
	}

	// Verificar que el elemento sea del mismo tipo que el slice
	if elemento.Tipo != sliceInfo.TipoElemento {
		return nil, fmt.Errorf("no se puede agregar elemento de tipo %s al slice de tipo []%s",
			elemento.Tipo, sliceInfo.TipoElemento)
	}

	sp.armGen.Comment(fmt.Sprintf("Expandiendo slice %s de %d a %d elementos",
		nombreSlice, sliceInfo.Tamaño, sliceInfo.Tamaño+1))

	// Crear nuevo slice con tamaño expandido
	nuevoTamaño := sliceInfo.Tamaño + 1
	nuevoOffset := sp.calcularNuevoOffset(nuevoTamaño)

	// Copiar elementos existentes al nuevo espacio
	sp.copiarElementosExistentes(sliceInfo, nuevoOffset)

	// Agregar nuevo elemento al final
	offsetNuevoElemento := nuevoOffset + (sliceInfo.Tamaño * 8)
	sp.agregarNuevoElemento(elemento, offsetNuevoElemento)

	// Actualizar información del slice
	sliceInfo.Tamaño = nuevoTamaño
	sliceInfo.OffsetStack = nuevoOffset
	sp.slicesDeclarados[nombreSlice] = sliceInfo

	// Actualizar offset global para futuros slices
	sp.offsetStack = nuevoOffset + (nuevoTamaño * 8)

	sp.armGen.Comment(fmt.Sprintf("append completado: %s ahora tiene %d elementos", nombreSlice, nuevoTamaño))

	// Retornar referencia al slice expandido
	return &ResultadoExpresion{
		Registro:  "",
		Tipo:      "slice_name",
		EsLiteral: true,
		Valor:     nombreSlice,
	}, nil
}

func (sp *SliceProcessor) calcularNuevoOffset(nuevoTamaño int) int {
	// Calcular espacio necesario y alinear a 8 bytes
	espacioNecesario := nuevoTamaño * 8
	if espacioNecesario%8 != 0 {
		espacioNecesario = ((espacioNecesario / 8) + 1) * 8
	}

	// Usar el offset actual del stack
	return sp.offsetStack
}

func (sp *SliceProcessor) copiarElementosExistentes(sliceInfo *SliceInfo, nuevoOffset int) {
	sp.armGen.Comment(fmt.Sprintf("Copiando %d elementos existentes", sliceInfo.Tamaño))

	if sliceInfo.Tamaño == 0 {
		return // No hay elementos que copiar
	}

	regElemento := "x18"

	for i := 0; i < sliceInfo.Tamaño; i++ {
		offsetOrigen := sliceInfo.OffsetStack + (i * 8)
		offsetDestino := nuevoOffset + (i * 8)

		// CAMBIO: Usar x28 como base en lugar de sp
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr %s, [x28, #%d]    // cargar elemento %d original", regElemento, offsetOrigen, i),
			fmt.Sprintf("str %s, [x28, #%d]    // guardar elemento %d en nueva posición", regElemento, offsetDestino, i))
	}
}

func (sp *SliceProcessor) agregarNuevoElemento(elemento *ResultadoExpresion, offset int) {
	sp.armGen.Comment(fmt.Sprintf("Agregando nuevo elemento en offset %d", offset))

	if elemento.EsLiteral {
		sp.AlmacenarElementoLiteral(elemento, offset)
	} else {
		sp.AlmacenarElementoRegistro(elemento, offset)
	}
}

// Función para acceso por índice
func (sp *SliceProcessor) AccederElementoPorIndice(nombreSlice string, indiceExpr *ResultadoExpresion) (*ResultadoExpresion, error) {
	sp.armGen.Comment(fmt.Sprintf("=== ACCESO POR ÍNDICE %s[%v] ===", nombreSlice, indiceExpr.Valor))

	sliceInfo, err := sp.CargarSlice(nombreSlice)
	if err != nil {
		return nil, err
	}

	// Verificar que el índice sea entero
	if indiceExpr.Tipo != "int" {
		return nil, fmt.Errorf("el índice debe ser de tipo int, recibido: %s", indiceExpr.Tipo)
	}

	// Si es un índice literal, verificar rango en tiempo de compilación
	if indiceExpr.EsLiteral {
		indice := indiceExpr.Valor.(int)
		if indice < 0 || indice >= sliceInfo.Tamaño {
			sp.armGen.Comment(fmt.Sprintf("Índice fuera de rango: %d (tamaño: %d)", indice, sliceInfo.Tamaño))
			// Retornar -1 para indicar error
			registroError := sp.expresionesProcessor.NuevoRegistroTmp()
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("mov %s, #0", registroError),
				fmt.Sprintf("sub %s, %s, #1", registroError, registroError)) // 0 - 1 = -1
			return &ResultadoExpresion{
				Registro:  registroError,
				Tipo:      sliceInfo.TipoElemento,
				EsLiteral: false,
			}, nil
		}

		// Acceso directo con índice literal
		return sp.generarAccesoDirecto(sliceInfo, indice)
	} else {
		// Acceso con índice en registro (verificación en tiempo de ejecución)
		return sp.generarAccesoConVerificacion(sliceInfo, indiceExpr)
	}
}

// generarAccesoDirecto genera código para acceso con índice literal
func (sp *SliceProcessor) generarAccesoDirecto(sliceInfo *SliceInfo, indice int) (*ResultadoExpresion, error) {
	offsetElemento := sliceInfo.OffsetStack + (indice * 8)

	switch sliceInfo.TipoElemento {
	case "int", "bool":
		registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr %s, [x28, #%d]    // cargar %s[%d] desde base fija",
				registroResultado, offsetElemento, sliceInfo.Nombre, indice))
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      sliceInfo.TipoElemento,
			EsLiteral: false,
		}, nil
	case "float":
		regFloat := sp.expresionesProcessor.NuevoRegistroFloatTmp()
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr %s, [x28, #%d]    // cargar float %s[%d] desde base fija",
				regFloat, offsetElemento, sliceInfo.Nombre, indice))
		return &ResultadoExpresion{
			Registro:  regFloat,
			Tipo:      sliceInfo.TipoElemento,
			EsLiteral: false,
		}, nil
	case "string":
		registroResultado := sp.expresionesProcessor.NuevoRegistroTmp()
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("ldr %s, [x28, #%d]    // cargar string %s[%d] desde base fija",
				registroResultado, offsetElemento, sliceInfo.Nombre, indice))
		return &ResultadoExpresion{
			Registro:  registroResultado,
			Tipo:      sliceInfo.TipoElemento,
			EsLiteral: false,
		}, nil
	default:
		return nil, fmt.Errorf("tipo de elemento no soportado: %s", sliceInfo.TipoElemento)
	}
}

// generarAccesoConVerificacion genera código para acceso con verificación de rango
func (sp *SliceProcessor) generarAccesoConVerificacion(sliceInfo *SliceInfo, indiceExpr *ResultadoExpresion) (*ResultadoExpresion, error) {
	sp.armGen.Comment("Acceso con verificación de rango en tiempo de ejecución")

	// Etiquetas para control de flujo
	labelRangoValido := fmt.Sprintf("range_valid_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelRangoInvalido := fmt.Sprintf("range_invalid_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelEnd := fmt.Sprintf("access_end_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())

	// Registros necesarios
	regIndice := "x9"     // Índice a verificar
	regElemento := "x10"  // Elemento cargado
	regResultado := "x11" // Resultado final

	// Cargar índice en registro
	if indiceExpr.EsLiteral {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, #%d", regIndice, indiceExpr.Valor.(int)))
	} else {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, %s", regIndice, indiceExpr.Registro))
	}

	// Verificar rango: 0 <= índice < tamaño
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		// Verificar si índice < 0
		fmt.Sprintf("cmp %s, #0", regIndice),
		fmt.Sprintf("blt %s", labelRangoInvalido),
		// Verificar si índice >= tamaño
		fmt.Sprintf("cmp %s, #%d", regIndice, sliceInfo.Tamaño),
		fmt.Sprintf("bge %s", labelRangoInvalido))

	// Rango válido: cargar elemento
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelRangoValido))

	offsetBase := sliceInfo.OffsetStack
	switch sliceInfo.TipoElemento {
	case "int", "bool":
		// CAMBIO: Usar x28 como base
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("add %s, x28, #%d", regElemento, offsetBase),
			fmt.Sprintf("ldr %s, [%s, %s, lsl #3]", regResultado, regElemento, regIndice))
	case "float":
		regFloat := sp.expresionesProcessor.NuevoRegistroFloatTmp()
		// CAMBIO: Usar x28 como base
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("add %s, x28, #%d", regElemento, offsetBase),
			fmt.Sprintf("ldr %s, [%s, %s, lsl #3]", regFloat, regElemento, regIndice))
		regResultado = regFloat
	case "string":
		// CAMBIO: Usar x28 como base
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("add %s, x28, #%d", regElemento, offsetBase),
			fmt.Sprintf("ldr %s, [%s, %s, lsl #3]", regResultado, regElemento, regIndice))
	}

	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("b %s", labelEnd))

	// Rango inválido: retornar -1
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelRangoInvalido),
		fmt.Sprintf("mov %s, #0", regResultado),
		fmt.Sprintf("sub %s, %s, #1", regResultado, regResultado)) // 0 - 1 = -1

	// Final
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelEnd))

	sp.armGen.Comment(fmt.Sprintf("Acceso completado, resultado en %s", regResultado))

	return &ResultadoExpresion{
		Registro:  regResultado,
		Tipo:      sliceInfo.TipoElemento,
		EsLiteral: false,
	}, nil
}

// Función para asignación por índice
func (sp *SliceProcessor) AsignarElementoPorIndice(nombreSlice string, indiceExpr *ResultadoExpresion, valorExpr *ResultadoExpresion) error {
	sp.armGen.Comment(fmt.Sprintf("=== ASIGNACIÓN POR ÍNDICE %s[%v] = %v ===", nombreSlice, indiceExpr.Valor, valorExpr.Valor))

	sliceInfo, err := sp.CargarSlice(nombreSlice)
	if err != nil {
		return err
	}

	// Verificar que el índice sea entero
	if indiceExpr.Tipo != "int" {
		return fmt.Errorf("el índice debe ser de tipo int, recibido: %s", indiceExpr.Tipo)
	}

	// Verificar que el valor sea compatible con el tipo del slice
	if valorExpr.Tipo != sliceInfo.TipoElemento {
		return fmt.Errorf("no se puede asignar valor de tipo %s al slice de tipo []%s",
			valorExpr.Tipo, sliceInfo.TipoElemento)
	}

	// Si es un índice literal, verificar rango en tiempo de compilación
	if indiceExpr.EsLiteral {
		indice := indiceExpr.Valor.(int)
		if indice < 0 || indice >= sliceInfo.Tamaño {
			sp.armGen.Comment(fmt.Sprintf("Error: Índice fuera de rango: %d (tamaño: %d)", indice, sliceInfo.Tamaño))
			return fmt.Errorf("índice fuera de rango: %d", indice)
		}

		// Asignación directa con índice literal
		return sp.generarAsignacionDirecta(sliceInfo, indice, valorExpr)
	} else {
		// Asignación con índice en registro (verificación en tiempo de ejecución)
		return sp.generarAsignacionConVerificacion(sliceInfo, indiceExpr, valorExpr)
	}
}

// generarAsignacionDirecta genera código para asignación con índice literal
func (sp *SliceProcessor) generarAsignacionDirecta(sliceInfo *SliceInfo, indice int, valorExpr *ResultadoExpresion) error {
	sp.armGen.Comment(fmt.Sprintf("Asignación directa al elemento %d", indice))

	offsetElemento := sliceInfo.OffsetStack + (indice * 8)

	if valorExpr.EsLiteral {
		// Valor literal: cargar directamente
		switch sliceInfo.TipoElemento {
		case "int":
			valor := valorExpr.Valor.(int)
			regTmp := sp.expresionesProcessor.NuevoRegistroTmp()
			// CAMBIO: Usar x28 como base
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("mov %s, #%d", regTmp, valor),
				fmt.Sprintf("str %s, [x28, #%d]    // %s[%d] = %d", regTmp, offsetElemento, sliceInfo.Nombre, indice, valor))
		case "bool":
			valor := 0
			if valorExpr.Valor.(bool) {
				valor = 1
			}
			regTmp := sp.expresionesProcessor.NuevoRegistroTmp()
			// CAMBIO: Usar x28 como base
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("mov %s, #%d", regTmp, valor),
				fmt.Sprintf("str %s, [x28, #%d]    // %s[%d] = %t", regTmp, offsetElemento, sliceInfo.Nombre, indice, valorExpr.Valor.(bool)))
		case "string":
			etiqueta := sp.expresionesProcessor.AgregarMensajeString(valorExpr.Valor.(string))
			regTmp := sp.expresionesProcessor.NuevoRegistroTmp()
			// CAMBIO: Usar x28 como base
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regTmp, etiqueta),
				fmt.Sprintf("str %s, [x28, #%d]    // %s[%d] = \"%s\"", regTmp, offsetElemento, sliceInfo.Nombre, indice, valorExpr.Valor.(string)))
		case "float":
			regFloat := sp.expresionesProcessor.NuevoRegistroFloatTmp()
			sp.expresionesProcessor.CargarConstanteFloat(regFloat, valorExpr.Valor.(float64))
			// CAMBIO: Usar x28 como base
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("str %s, [x28, #%d]    // %s[%d] = %f", regFloat, offsetElemento, sliceInfo.Nombre, indice, valorExpr.Valor.(float64)))
		}
	} else {
		// Valor en registro: copiar directamente
		// CAMBIO: Usar x28 como base
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("str %s, [x28, #%d]    // %s[%d] = registro", valorExpr.Registro, offsetElemento, sliceInfo.Nombre, indice))
	}

	return nil
}

// generarAsignacionConVerificacion genera código para asignación con verificación de rango
func (sp *SliceProcessor) generarAsignacionConVerificacion(sliceInfo *SliceInfo, indiceExpr *ResultadoExpresion, valorExpr *ResultadoExpresion) error {
	sp.armGen.Comment("Asignación con verificación de rango en tiempo de ejecución")

	// Etiquetas para control de flujo
	labelRangoValido := fmt.Sprintf("assign_valid_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelRangoInvalido := fmt.Sprintf("assign_invalid_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())
	labelEnd := fmt.Sprintf("assign_end_%s_%d", sliceInfo.Nombre, sp.expresionesProcessor.GetContadorEtiqueta())

	// Registros necesarios
	regIndice := "x9"        // Índice a verificar
	regElementoAddr := "x10" // Dirección del elemento
	regValor := "x11"        // Valor a asignar

	// Cargar índice en registro
	if indiceExpr.EsLiteral {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, #%d", regIndice, indiceExpr.Valor.(int)))
	} else {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, %s", regIndice, indiceExpr.Registro))
	}

	// Cargar valor en registro
	if valorExpr.EsLiteral {
		switch sliceInfo.TipoElemento {
		case "int":
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("mov %s, #%d", regValor, valorExpr.Valor.(int)))
		case "bool":
			valor := 0
			if valorExpr.Valor.(bool) {
				valor = 1
			}
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("mov %s, #%d", regValor, valor))
		case "string":
			etiqueta := sp.expresionesProcessor.AgregarMensajeString(valorExpr.Valor.(string))
			sp.armGen.Instructions = append(sp.armGen.Instructions,
				fmt.Sprintf("adr %s, %s", regValor, etiqueta))
		}
	} else {
		sp.armGen.Instructions = append(sp.armGen.Instructions,
			fmt.Sprintf("mov %s, %s", regValor, valorExpr.Registro))
	}

	// Verificar rango: 0 <= índice < tamaño
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		// Verificar si índice < 0
		fmt.Sprintf("cmp %s, #0", regIndice),
		fmt.Sprintf("blt %s", labelRangoInvalido),
		// Verificar si índice >= tamaño
		fmt.Sprintf("cmp %s, #%d", regIndice, sliceInfo.Tamaño),
		fmt.Sprintf("bge %s", labelRangoInvalido))

	// Rango válido: asignar elemento
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelRangoValido))

	offsetBase := sliceInfo.OffsetStack
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("add %s, x28, #%d", regElementoAddr, offsetBase),
		fmt.Sprintf("str %s, [%s, %s, lsl #3]    // %s[índice] = valor", regValor, regElementoAddr, regIndice, sliceInfo.Nombre),
		fmt.Sprintf("b %s", labelEnd))

	// Rango inválido: no hacer nada (silencioso)
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelRangoInvalido),
		fmt.Sprintf("// Índice fuera de rango, asignación ignorada"))

	// Final
	sp.armGen.Instructions = append(sp.armGen.Instructions,
		fmt.Sprintf("%s:", labelEnd))

	sp.armGen.Comment("Asignación por índice completada")

	return nil
}
