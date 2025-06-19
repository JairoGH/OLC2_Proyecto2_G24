package instrucciones

import (
	"fmt"
	"main/parser"
	"main/tiposDeDato"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type PatronVIsitor struct {
	parser.BaseVGrammarVisitor
	RegistroAmbito     *RegistroAmbito
	PilaLlamada        *PilaLlamada
	Consola            *Consola
	TablaError         *TablaError
	NombresEstructuras []string
}

// valorPorDefectoParaTipo devuelve el ValorInterno por defecto para tipos primitivos
func valorPorDefectoParaTipo(typeName string) tiposDeDato.ValorInterno {
	switch typeName {
	case tiposDeDato.TIPO_ENTERO:
		return &tiposDeDato.ValorEntero{ValorInterno: 0}
	case tiposDeDato.TIPO_DECIMAL:
		return &tiposDeDato.ValorDecimal{InternalValor: 0.0}
	case tiposDeDato.TIPO_CADENA:
		return &tiposDeDato.ValorCadena{ValorInterno: ""}
	case tiposDeDato.TIPO_BOOLEAN:
		return &tiposDeDato.ValorBool{InternalValor: false}
	default:
		return tiposDeDato.NuloPorDefecto
	}
}

func NewVisitor(VisitanteDcl *VisitanteDcl) *PatronVIsitor {
	return &PatronVIsitor{
		RegistroAmbito:     VisitanteDcl.RegistroAmbito,
		TablaError:         VisitanteDcl.TablaError,
		NombresEstructuras: VisitanteDcl.NombresEstructuras,
		PilaLlamada:        NuevaLlamadaFuncion(),
		Consola:            NewConsola(),
	}
}

func (v *PatronVIsitor) GetInstruccionesContexto() *InstruccionesContexto {
	return &InstruccionesContexto{
		Consola:        v.Consola,
		RegistroAmbito: v.RegistroAmbito,
		PilaLlamada:    v.PilaLlamada,
		TablaError:     v.TablaError,
	}
}

func (v *PatronVIsitor) TipoValido(_type string) bool {
	return v.RegistroAmbito.AmbitoGlobal.ValidType(_type)
}

func (v *PatronVIsitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		token := val.GetSymbol()
		v.TablaError.NewErrorSintactico(token.GetLine(), token.GetColumn(), "Error de análisis: "+val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *PatronVIsitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	// Agregar el manejo de main_func
	if ctx.Main_func() != nil {
		v.Visit(ctx.Main_func())
	}
	return nil
}
func (v *PatronVIsitor) VisitFuncionMain(ctx *parser.FuncionMainContext) interface{} {

	// Push scope para main
	v.RegistroAmbito.PushAmbito("main")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.RegistroAmbito.PopAmbito()

	return nil
}

func (v *PatronVIsitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones (decl_stmt) con sus 7 alternativas
	if ds := ctx.Stmt_declaracion(); ds != nil {
		switch decl := ds.(type) {
		case *parser.DeclararInferenciaMutContext:
			v.VisitDeclararInferenciaMut(decl)
		case *parser.DeclararInferenciaContext:
			v.VisitDeclararInferencia(decl)
		case *parser.DeclaraTipoValorContext:
			v.VisitDeclaraTipoValor(decl)
		case *parser.DeclararTipoContext:
			v.VisitDeclararTipo(decl)
		case *parser.DeclararSinMutValorContext:
			v.VisitDeclararSinMutValor(decl)
		case *parser.DeclararSliceContext:
			v.VisitDeclararSlice(decl)
		default:
			v.Visit(ds) // fallback por si se agregan nuevas alternativas
		}
		return nil
	}

	// 2) Asignaciones
	if ctx.Stmt_asignar() != nil {
		v.Visit(ctx.Stmt_asignar())
		return nil
	}

	// 3) Resto de sentencias
	switch {
	case ctx.If_stmt() != nil:
		v.Visit(ctx.If_stmt())
	case ctx.Switch_stmt() != nil:
		v.Visit(ctx.Switch_stmt())
	case ctx.While_stmt() != nil:
		v.Visit(ctx.While_stmt())
	case ctx.For_clasico_stmt() != nil:
		v.Visit(ctx.For_clasico_stmt())
	case ctx.For_stmt() != nil:
		v.Visit(ctx.For_stmt())
	case ctx.Stmt_transferencia() != nil:
		v.Visit(ctx.Stmt_transferencia())
	case ctx.Llamar_funcion() != nil:
		v.Visit(ctx.Llamar_funcion())
	case ctx.Declarar_funcion() != nil:
		v.Visit(ctx.Declarar_funcion())
	case ctx.Declarar_struct() != nil:
		v.Visit(ctx.Declarar_struct())
	case ctx.Fun_slice() != nil:
		v.Visit(ctx.Fun_slice())
	default:
		token := ctx.GetStart()
		v.TablaError.NewErrorSemantico(token, "No se reconoció la declaración: "+ctx.GetText())
	}

	return nil
}

// 1) mut ID type = expr
func (v *PatronVIsitor) VisitDeclaraTipoValor(ctx *parser.DeclaraTipoValorContext) interface{} {
	nombreVariable := ctx.ID().GetText()
	tipoVariable := v.Visit(ctx.Tipo()).(string)
	valorVariable := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// 1) Verificar mismatch de tipo
	if valorVariable.Type() != tipoVariable {
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			fmt.Sprintf("No se puede asignar un valor de tipo %s a variable %s de tipo %s",
				valorVariable.Type(), nombreVariable, tipoVariable))
		return nil
	}

	// copy object/vector
	if obj, ok := valorVariable.(*TipoObjeto); ok {
		valorVariable = obj.Copy()
	}
	if EsTipoVector(valorVariable.Type()) {
		valorVariable = valorVariable.Copy()
	}

	// insertar mutable + inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(
		nombreVariable, tipoVariable, valorVariable,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

// vector x type = expr
func (v *PatronVIsitor) VisitDeclararSlice(ctx *parser.DeclararSliceContext) interface{} {
	nombreVariable := ctx.ID().GetText()

	// Verificar que ctx.Type_() no sea nil
	if ctx.Tipo() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo no especificado en declaración de vector")
		return nil
	}

	tipoResultado := v.Visit(ctx.Tipo())
	if tipoResultado == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar el tipo del vector")
		return nil
	}
	tipoVariable := tipoResultado.(string)

	// Verificar que ctx.Expr() no sea nil
	if ctx.Expr() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Valor no especificado en declaración de vector")
		return nil
	}

	valorResultado := v.Visit(ctx.Expr())
	if valorResultado == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar el valor del vector")
		return nil
	}
	valorVariable := valorResultado.(tiposDeDato.ValorInterno)

	// NUEVO: Manejar vectores vacíos especialmente
	if valorVariable.Type() == "[]" {
		// Es un vector vacío, asignarle el tipo correcto
		if vectorVal, ok := valorVariable.(*TipoVector); ok {
			vectorVal.TipoCompleto = tipoVariable
			vectorVal.TipoElemento = EliminarCorchetes(tipoVariable)
		}
	} else {
		// Verificar que el tipo del vector coincida
		if valorVariable.Type() != tipoVariable {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un vector de tipo %s a variable %s de tipo %s",
					valorVariable.Type(), nombreVariable, tipoVariable))
			return nil
		}
	}

	// copy object
	if objeto, ok := valorVariable.(*TipoObjeto); ok {
		valorVariable = objeto.Copy()
	}

	if EsTipoVector(valorVariable.Type()) {
		valorVariable = valorVariable.Copy()
	}

	// CAMBIO: true para mutable y TRUE para inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(nombreVariable, tipoVariable, valorVariable, false, true, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

// mut x := expr  (inferencia mutable)
func (v *PatronVIsitor) VisitDeclararInferenciaMut(ctx *parser.DeclararInferenciaMutContext) interface{} {
	nombre := ctx.ID().GetText()
	valor := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// copia objetos/vectores
	if obj, ok := valor.(*TipoObjeto); ok {
		valor = obj.Copy()
	}
	if EsTipoVector(valor.Type()) {
		valor = valor.Copy()
	}

	typ := valor.Type()
	// insertar mutable + inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(
		nombre, typ, valor,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

// x := expr  (inferencia sin mut)
func (v *PatronVIsitor) VisitDeclararInferencia(ctx *parser.DeclararInferenciaContext) interface{} {
	nombre := ctx.ID().GetText()
	valor := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	if obj, ok := valor.(*TipoObjeto); ok {
		valor = obj.Copy()
	}
	if EsTipoVector(valor.Type()) {
		valor = valor.Copy()
	}

	tipo := valor.Type()
	variable, msg := v.RegistroAmbito.AgregarVariable(
		nombre, tipo, valor,
		false, /* mutable = false */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

func (v *PatronVIsitor) VisitDeclararTipo(ctx *parser.DeclararTipoContext) interface{} {
	nombreVariable := ctx.ID().GetText()
	tipoVariable := v.Visit(ctx.Tipo()).(string)

	// valor por defecto según tipo
	defaultVal := valorPorDefectoParaTipo(tipoVariable)

	variable, msg := v.RegistroAmbito.AgregarVariable(
		nombreVariable, tipoVariable, defaultVal,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

func (v *PatronVIsitor) VisitListaSlice(ctx *parser.ListaSliceContext) interface{} {

	var elementosVector []tiposDeDato.ValorInterno

	if len(ctx.AllExpr()) == 0 {
		return NewTipoVector(elementosVector, "[]", "")
	}

	for _, item := range ctx.AllExpr() {
		elementosVector = append(elementosVector, v.Visit(item).(tiposDeDato.ValorInterno))
	}

	if len(elementosVector) == 0 {
		return NewTipoVector(elementosVector, "[]", "")
	}

	tipoElementos := elementosVector[0].Type()

	// NUEVO: Detectar si es una matriz (vector de vectores)
	esMatriz := strings.HasPrefix(tipoElementos, "[]")

	for _, item := range elementosVector {
		if item.Type() != tipoElementos {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Todos los items de la colección deben ser del mismo tipo")
			return tiposDeDato.NuloPorDefecto
		}
	}

	if esMatriz {
		// Es una matriz [][]tipo
		tipoBase := strings.TrimPrefix(tipoElementos, "[]") // extraer tipo base
		tipoMatriz := "[]" + tipoElementos                  // [][]int

		// Convertir vectores a matriz multidimensional
		matriz := make([][]tiposDeDato.ValorInterno, len(elementosVector))
		for i, item := range elementosVector {
			if vector, ok := item.(*TipoVector); ok {
				matriz[i] = vector.ValorInterno
			} else {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					"Error interno: elemento de matriz no es vector")
				return tiposDeDato.NuloPorDefecto
			}
		}

		return NewTipoMatrizMulti(matriz, tipoMatriz, tipoElementos, tipoBase)
	} else {
		// Es un vector simple []tipo
		_type := "[]" + tipoElementos
		return NewTipoVector(elementosVector, _type, tipoElementos)
	}
}

func (v *PatronVIsitor) VisitTipo(ctx *parser.TipoContext) interface{} {
	_type := ctx.GetText()

	// Verificar si es matriz multidimensional [][]tipo
	if strings.HasPrefix(_type, "[][]") && len(_type) > 4 {
		// Extraer el tipo base (después de [][])
		tipoBase := _type[4:] // eliminar [][]

		// Verificar si el tipo base es válido
		if v.TipoValido(tipoBase) {
			return _type // retornar [][]int, [][]string, etc.
		}

		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El tipo "+tipoBase+" no es válido para una matriz")
		return tiposDeDato.TIPO_NULO
	}

	// Verificar si es vector simple []tipo
	if strings.HasPrefix(_type, "[]") && len(_type) > 2 {
		internType := _type[2:] // eliminar []

		if v.TipoValido(internType) {
			return _type // retornar []int, []string, etc.
		}

		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El tipo "+internType+" no es válido para un vector")
		return tiposDeDato.TIPO_NULO
	}

	// Resto del código existente para tipos primitivos y formatos antiguos...
	if v.TipoValido(_type) {
		return _type
	}

	// Código existente para EsTipoVector y EsTipoMatriz...

	v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo "+ctx.GetText()+" no encontrado")
	return tiposDeDato.TIPO_NULO
}

func (v *PatronVIsitor) VisitTipos_slices(ctx *parser.Tipos_slicesContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitItemSlice(ctx *parser.ItemSliceContext) interface{} {

	nombreVariable := ctx.PatronId().GetText()
	variable := v.RegistroAmbito.GetVariable(nombreVariable)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+nombreVariable+" no encontrada")
		return nil
	}

	// Verificar que sea un vector o matriz (incluyendo formato [][]tipo)
	if !(EsTipoVector(variable.Tipo) || EsTipoMatriz(variable.Tipo) ||
		strings.HasPrefix(variable.Tipo, "[]")) {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+nombreVariable+" no es un vector o una matriz")
		return nil
	}

	indices := []int{}
	for _, expr := range ctx.AllExpr() {
		val := v.Visit(expr).(tiposDeDato.ValorInterno)
		if val.Type() != tiposDeDato.TIPO_ENTERO {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los indices de acceso deben ser enteros")
			return nil
		}
		indices = append(indices, val.Value().(int))
	}

	// ✅ NUEVO: Manejar TipoMatrizMulti primero
	if matrizMulti, ok := variable.Valor.(*TipoMatrizMulti); ok {
		if len(indices) == 2 {
			// Acceso completo: matriz[fila][columna]
			fila, columna := indices[0], indices[1]
			if !matrizMulti.ValidIndex(fila, columna) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: [%d][%d]. La matriz tiene %d filas",
						fila, columna, len(matrizMulti.ValorInterno)))

				// Retornar un valor que indique error pero que no rompa el programa
				return &tiposDeDato.ValorEntero{ValorInterno: -1}
			}
			return &MatrizMultiItemReference{
				Matriz:  matrizMulti,
				Fila:    fila,
				Columna: columna,
				Valor:   matrizMulti.Get(fila, columna),
			}
		} else if len(indices) == 1 {
			// ✅ ARREGLO: Acceso a fila: matriz[fila] -> retorna vector para que len() funcione
			fila := indices[0]
			if fila < 0 || fila >= len(matrizMulti.ValorInterno) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice de fila fuera de rango: %d", fila))
				return &tiposDeDato.ValorEntero{ValorInterno: -1}
			}

			// ✅ CLAVE: Crear un TipoVector para la fila y retornarlo como VectorItemReference
			filaVector := NewTipoVector(matrizMulti.ValorInterno[fila], matrizMulti.TipoElemento, matrizMulti.TipoBase)

			return &VectorItemReference{
				Vector: filaVector,
				Indice: -1,         // -1 indica que es una fila completa, no un elemento individual
				Valor:  filaVector, // ✅ El valor ES el vector completo
			}
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Número incorrecto de índices para matriz multidimensional")
			return nil
		}
	}

	// Resto del código existente para vectores simples y matrices regulares
	tipoStructs := tiposDeDato.TIPO_VECTOR
	indice := v.Visit(ctx.Expr(0)).(tiposDeDato.ValorInterno)

	if len(ctx.AllExpr()) != 1 {
		tipoStructs = tiposDeDato.TIPO_MATRIZ
	}

	if tipoStructs == tiposDeDato.TIPO_VECTOR {
		switch valorVector := variable.Valor.(type) {
		case *TipoVector:
			valorIndice := indice.(*tiposDeDato.ValorEntero).ValorInterno
			if !valorVector.IndiceValido(valorIndice) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: %d. El vector tiene %d elementos (rango válido: 0-%d)",
						valorIndice, len(valorVector.ValorInterno), len(valorVector.ValorInterno)-1))
				return &VectorItemReference{
					Vector: valorVector,
					Indice: valorIndice,
					Valor:  &tiposDeDato.ValorEntero{ValorInterno: -1},
				}
			}
			return &VectorItemReference{
				Vector: valorVector,
				Indice: valorIndice,
				Valor:  valorVector.Get(valorIndice),
			}
		default:
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+nombreVariable+" no es un vector")
			return nil
		}
	} else if tipoStructs == tiposDeDato.TIPO_MATRIZ {
		switch TipoMatriz := variable.Valor.(type) {
		case *TipoMatriz:
			if !TipoMatriz.IndicesValidos(indices) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: %v. La matriz tiene dimensiones válidas", indices))
				return nil
			}
			return &ElementoMatriz{
				Matriz: TipoMatriz,
				Indice: indices,
				Valor:  TipoMatriz.Get(indices),
			}
		default:
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+nombreVariable+" no es una matriz")
			return nil
		}
	}

	return nil
}

func (v *PatronVIsitor) VisitAsignacionDirecta(ctx *parser.AsignacionDirectaContext) interface{} {

	nombreVariable := v.Visit(ctx.PatronId()).(string)
	valorVariable := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	variable := v.RegistroAmbito.GetVariable(nombreVariable)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+nombreVariable+" no encontrada")
	} else {

		// copy object
		if obj, ok := valorVariable.(*TipoObjeto); ok {
			valorVariable = obj.Copy()
		}

		if EsTipoVector(valorVariable.Type()) {
			valorVariable = valorVariable.Copy()
		}

		puedeModificarse := true

		if v.RegistroAmbito.AmbitoActual.EsEstructura {
			puedeModificarse = v.RegistroAmbito.EsEntornoMutable()
		}

		ok, msg := variable.AsignarVariable(valorVariable, puedeModificarse)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		}
	}

	return nil

}

func (v *PatronVIsitor) VisitAsignacionAritmetica(ctx *parser.AsignacionAritmeticaContext) interface{} {
	nombreVariable := v.Visit(ctx.PatronId()).(string)

	variable := v.RegistroAmbito.GetVariable(nombreVariable)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable '"+nombreVariable+"' no encontrada")
		return nil
	}

	valorIzq := variable.Valor
	valorDcha := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	op := string(ctx.GetOp().GetText()[0])
	strat, ok := ExpresionBinaria[op]

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Operador binario no reconocido: "+op)
		return nil
	}

	ok, msg, valorVariable := strat.ValidarExp(valorIzq, valorDcha)
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	puedeModificarse := true
	if v.RegistroAmbito.AmbitoActual.EsEstructura {
		puedeModificarse = v.RegistroAmbito.EsEntornoMutable()
	}

	ok, msg = variable.AsignarVariable(valorVariable, puedeModificarse)
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

func (v *PatronVIsitor) VisitAsignacionSlice(ctx *parser.AsignacionSliceContext) interface{} {
	valorDcha := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	switch itemRef := v.Visit(ctx.Item_slice()).(type) {
	case *VectorItemReference:
		valorIzq := itemRef.Valor

		if valorDcha.Type() != itemRef.Vector.TipoElemento {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				"No se puede asignar un valor de tipo "+valorDcha.Type()+
					" a un vector de tipo "+itemRef.Vector.TipoElemento)
			return nil
		}

		op := string(ctx.GetOp().GetText()[0])

		if op == "=" {
			itemRef.Vector.ValorInterno[itemRef.Indice] = valorDcha
			return nil
		}

		strat, ok := ExpresionBinaria[op]
		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Operador binario no reconocido: "+op)
			return nil
		}

		ok, msg, varValue := strat.ValidarExp(valorIzq, valorDcha)
		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			return nil
		}

		itemRef.Vector.ValorInterno[itemRef.Indice] = varValue
		return nil

	case *ElementoMatriz:
		valorIzq := itemRef.Valor

		tipoEsperado := EliminarCorchetes(itemRef.Matriz.Type())
		if valorDcha.Type() != tipoEsperado {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				"No se puede asignar un valor de tipo "+valorDcha.Type()+
					" a una matriz de tipo "+tipoEsperado)
			return nil
		}

		op := string(ctx.GetOp().GetText()[0])

		if op == "=" {
			itemRef.Matriz.Set(itemRef.Indice, valorDcha)
			return nil
		}

		strat, ok := ExpresionBinaria[op]
		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Operador binario no reconocido: "+op)
			return nil
		}

		ok, msg, valorVariable := strat.ValidarExp(valorIzq, valorDcha)
		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			return nil
		}

		itemRef.Matriz.Set(itemRef.Indice, valorVariable)
		return nil
	}

	// Si no se puede hacer match con ninguno
	v.TablaError.NewErrorSemantico(ctx.GetStart(), "Referencia de slice o matriz inválida")
	return nil
}

func (v *PatronVIsitor) VisitAsignacionSliceItem(ctx *parser.AsignacionSliceItemContext) interface{} {

	// Obtener la referencia al elemento del vector/matriz
	refItemVector := v.Visit(ctx.Item_slice())
	if refItemVector == nil {
		return nil
	}

	// NUEVO: Verificar si es un error (ValorEntero con -1)
	if errorVal, ok := refItemVector.(*tiposDeDato.ValorEntero); ok && errorVal.ValorInterno == -1 {
		// Ya se reportó el error en VisitVectorItem, no hacer nada más
		return nil
	}

	// Evaluar el nuevo valor a asignar
	nuevoValor := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	switch refItem := refItemVector.(type) {
	case *VectorItemReference:
		// Caso: vector simple numeros[2] = valor

		// Verificar compatibilidad de tipos
		if nuevoValor.Type() != refItem.Vector.TipoElemento {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de tipo %s",
					nuevoValor.Type(), refItem.Vector.TipoElemento))
			return nil
		}

		// Asignar el nuevo valor
		refItem.Vector.ValorInterno[refItem.Indice] = nuevoValor
		return nil

	case *ElementoMatriz:
		// Caso: matriz antigua TipoMatriz[indices] = valor

		// Verificar compatibilidad de tipos para matriz
		tipoActual := EliminarCorchetes(refItem.Matriz.Type())
		if nuevoValor.Type() != tipoActual {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de matriz de tipo %s",
					nuevoValor.Type(), tipoActual))
			return nil
		}

		// Asignar el nuevo valor en la matriz
		refItem.Matriz.Set(refItem.Indice, nuevoValor)
		return nil

	case *MatrizMultiItemReference:
		// Caso: matriz multidimensional mtx2[0][1] = valor

		// Verificar compatibilidad de tipos
		if nuevoValor.Type() != refItem.Matriz.TipoBase {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de matriz de tipo %s",
					nuevoValor.Type(), refItem.Matriz.TipoBase))
			return nil
		}

		// Verificar que los índices estén en rango ANTES de intentar asignar
		if !refItem.Matriz.ValidIndex(refItem.Fila, refItem.Columna) {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("Índice fuera de rango: [%d][%d]. La matriz tiene %d filas",
					refItem.Fila, refItem.Columna, len(refItem.Matriz.ValorInterno)))
			return nil
		}

		// Asignar el nuevo valor
		refItem.Matriz.Set(refItem.Fila, refItem.Columna, nuevoValor)
		return nil

	default:
		// MEJORAR: No mostrar "Error interno" si ya se reportó el error de índice
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "No se puede realizar la asignación debido a un error previo")
		return nil
	}
}

func (v *PatronVIsitor) VisitVectorSimple(ctx *parser.VectorSimpleContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitMatrizDoble(ctx *parser.MatrizDobleContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitID_Patron(ctx *parser.ID_PatronContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {

	valorInt, _ := strconv.Atoi(ctx.GetText())

	return &tiposDeDato.ValorEntero{
		ValorInterno: valorInt,
	}

}

func (v *PatronVIsitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {

	valorFloat, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &tiposDeDato.ValorDecimal{
		InternalValor: valorFloat,
	}

}

func (v *PatronVIsitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	// remove quotes
	valorString := ctx.GetText()[1 : len(ctx.GetText())-1]

	// \" \\ \n \r \
	valorString = strings.ReplaceAll(valorString, "\\\"", "\"")
	valorString = strings.ReplaceAll(valorString, "\\\\", "\\")
	valorString = strings.ReplaceAll(valorString, "\\n", "\n")
	valorString = strings.ReplaceAll(valorString, "\\r", "\r")

	// CAMBIO: Las cadenas con comillas dobles siempre son string
	// Si quieres caracteres individuales, usa comillas simples (si las tienes implementadas)
	return &tiposDeDato.ValorCadena{
		ValorInterno: valorString,
	}
}

func (v *PatronVIsitor) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {

	valorBool, _ := strconv.ParseBool(ctx.GetText())

	return &tiposDeDato.ValorBool{
		InternalValor: valorBool,
	}

}

func (v *PatronVIsitor) VisitNilLiteral(ctx *parser.NilLiteralContext) interface{} {
	return tiposDeDato.NuloPorDefecto
}

func (v *PatronVIsitor) VisitLiteralExp(ctx *parser.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *PatronVIsitor) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	nombreVariable := ctx.PatronId().GetText()

	variable := v.RegistroAmbito.GetVariable(nombreVariable)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+nombreVariable+" no encontrada")
		return tiposDeDato.NuloPorDefecto
	}

	// ? pointer
	return variable.Valor
}

func (v *PatronVIsitor) VisitParentecisExp(ctx *parser.ParentecisExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *PatronVIsitor) VisitItemSliceExp(ctx *parser.ItemSliceExpContext) interface{} {

	resultado := v.Visit(ctx.Item_slice())

	switch refItem := resultado.(type) {
	case *VectorItemReference:
		// ✅ ARREGLO: Si es una referencia a fila completa (Indice == -1), retornar el vector
		if refItem.Indice == -1 {
			return refItem.Vector // Retorna el TipoVector para que len() funcione
		}
		// Si es un elemento individual, retornar el valor
		return refItem.Valor

	case *ElementoMatriz:
		return refItem.Valor

	case *MatrizMultiItemReference:
		return refItem.Valor

	case *tiposDeDato.ValorEntero:
		// Caso de error (índice fuera de rango)
		if refItem.ValorInterno == -1 {
			return &tiposDeDato.ValorEntero{ValorInterno: -1}
		}
		return refItem

	case *TipoVector:
		return refItem

	default:
		return tiposDeDato.NuloPorDefecto
	}
}

func (v *PatronVIsitor) VisitLlamarFuncionExp(ctx *parser.LlamarFuncionExpContext) interface{} {
	return v.Visit(ctx.Llamar_funcion())
}

func (v *PatronVIsitor) VisitSliceExp(ctx *parser.SliceExpContext) interface{} {
	return v.Visit(ctx.Lista_slice())
}

func (v *PatronVIsitor) VisitUnarioExp(ctx *parser.UnarioExpContext) interface{} {
	exp := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	op := ctx.GetOp().GetText()
	strat, ok := ExpresionesUnarias[op]

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), "Operador unario no reconocido: "+op)
		return tiposDeDato.NuloPorDefecto
	}

	ok, msg, resultado := strat.ValidarExp(exp)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	return resultado
}

func (v *PatronVIsitor) VisitBinarioExp(ctx *parser.BinarioExpContext) interface{} {
	op := ctx.GetOp().GetText()
	izq := v.Visit(ctx.GetLeft()).(tiposDeDato.ValorInterno)

	// Evaluación anticipada (por ejemplo, cortocircuito lógico)
	if chequeoInicial, ok := RetornoAnticipado[op]; ok {
		ok, _, resultado := chequeoInicial.ValidarExp(izq)
		if ok {
			return resultado
		}
	}

	dcha := v.Visit(ctx.GetRight()).(tiposDeDato.ValorInterno)

	strat, ok := ExpresionBinaria[op]
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), "Operador binario no reconocido: "+op)
		return tiposDeDato.NuloPorDefecto
	}

	ok, msg, resultado := strat.ValidarExp(izq, dcha)
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	return resultado
}

func (v *PatronVIsitor) VisitIFstmt(ctx *parser.IFstmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)

		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *PatronVIsitor) VisitIFcadena(ctx *parser.IFcadenaContext) interface{} {

	condition := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

	}

	if condition.(*tiposDeDato.ValorBool).InternalValor {

		// Push scope
		v.RegistroAmbito.PushAmbito("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.RegistroAmbito.PopAmbito()

		return true
	}

	return false
}

func (v *PatronVIsitor) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {

	// Push scope
	v.RegistroAmbito.PushAmbito("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.RegistroAmbito.PopAmbito()

	return nil
}

func (v *PatronVIsitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	// 🎯 EVALUAR LA EXPRESIÓN PRINCIPAL DEL SWITCH
	mainValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// 🎯 CREAR ÁMBITO PARA EL SWITCH
	v.RegistroAmbito.PushAmbito("switch")

	// 🎯 CREAR ELEMENTO DE CONTROL PARA BREAK (igual que while)
	switchItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo: []string{
			Detener, // Solo break, no continue en switch
		},
	}

	v.PilaLlamada.Push(switchItem)

	// 🎯 MANEJO DE BREAK Y CLEANUP (igual que while)
	defer func() {
		v.RegistroAmbito.PopAmbito()      // pop switch scope
		v.PilaLlamada.Limpiar(switchItem) // limpiar si está en la pila

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {
			// Si no es nuestro switchItem, propagar el panic
			if item != switchItem {
				panic(item)
			}
			// Si es nuestro switchItem, es un break normal
			return
		}
	}()

	caseExecuted := false

	// 🎯 EVALUAR TODOS LOS CASES
	for _, switchCase := range ctx.AllSwitch_case() {
		caseValue := v.GetCaseValue(switchCase)

		// 🔧 MEJORAR: Verificar compatibilidad de tipos
		if caseValue.Type() != mainValue.Type() {
			v.TablaError.NewErrorSemantico(switchCase.GetStart(),
				fmt.Sprintf("El valor del case (%s) no es compatible con el tipo del switch (%s)",
					caseValue.Type(), mainValue.Type()))
			continue
		}

		// 🎯 SI HAY COINCIDENCIA, EJECUTAR
		if caseValue.Value() == mainValue.Value() {
			v.Visit(switchCase)
			caseExecuted = true
			break // 🎯 BREAK IMPLÍCITO (no fall-through)
		}
	}

	// 🎯 SI NO SE EJECUTÓ NINGÚN CASE, EJECUTAR DEFAULT
	if ctx.Default_case() != nil && !caseExecuted {
		v.Visit(ctx.Default_case())
	}

	return nil
}

func (v *PatronVIsitor) GetCaseValue(tree antlr.ParseTree) tiposDeDato.ValorInterno {
	switch val := tree.(type) {
	case *parser.SwitchCaseContext:
		return v.Visit(val.Expr()).(tiposDeDato.ValorInterno)
	default:
		// Esto no debería pasar, pero por seguridad
		v.TablaError.NewErrorSemantico(tree.(antlr.ParserRuleContext).GetStart(),
			"Error interno: tipo de case no reconocido")
		return tiposDeDato.NuloPorDefecto
	}
}

func (v *PatronVIsitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {
	// Los cases comparten el mismo ámbito del switch
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *PatronVIsitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	// Default también comparte el ámbito del switch
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *PatronVIsitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	condition := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// Push scope
	whileScope := v.RegistroAmbito.PushAmbito("while")

	// Push whileItem to call stack [breakable, continuable]
	whileItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo: []string{
			Detener,
			Continuar,
		},
	}

	v.PilaLlamada.Push(whileItem)

	// 🔧 MANEJO SEGURO CON DEFER
	defer func() {
		v.RegistroAmbito.PopAmbito()     // pop while scope
		v.PilaLlamada.Limpiar(whileItem) // clean item if it's still in call stack

		// 🔧 CAPTURAR PANICS NO CONTROLADOS
		if r := recover(); r != nil {
			if item, ok := r.(*LlamadaFunciones); ok && item != whileItem {
				// Es un panic de otra función (return), propagarlo
				panic(item)
			}
			// Si es nuestro whileItem o otro tipo, no hacer nada (ya se manejó)
		}
	}()

	v.VisitInnerWhile(ctx, condition, whileScope, whileItem)

	return nil
}

func (v *PatronVIsitor) VisitInnerWhile(ctx *parser.WhileStmtContext, condition tiposDeDato.ValorInterno, whileScope *AmbitoBase, whileItem *LlamadaFunciones) {
	if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
		return
	}

	for condition.(*tiposDeDato.ValorBool).InternalValor {
		// ✅ CAMBIO: NO hacer Reset del ámbito completo
		// Solo resetear variables locales del while si es necesario
		// whileScope.Reset() // ❌ ELIMINAR ESTA LÍNEA

		shouldBreak := false
		shouldContinue := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					if item, ok := r.(*LlamadaFunciones); ok && item == whileItem {
						if item.EsAccion(Continuar) {
							item.ResetAccion()
							shouldContinue = true
							return
						} else if item.EsAccion(Detener) {
							item.ResetAccion()
							shouldBreak = true
							return
						}
					} else if item, ok := r.(*LlamadaFunciones); ok && item != whileItem {
						panic(item)
					} else {
						v.TablaError.NewErrorSemantico(ctx.GetStart(),
							fmt.Sprintf("Error interno en bucle while: %v", r))
						shouldBreak = true
						return
					}
				}
			}()

			for _, stmt := range ctx.AllStmt() {
				v.Visit(stmt)
			}
		}()

		if shouldBreak {
			break
		}

		if shouldContinue {
			condition = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
			if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
				break
			}
			continue
		}

		condition = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
		if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
			break
		}
	}
}

func (v *PatronVIsitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// 🔧 VALIDACIÓN INICIAL MEJORADA
	if ctx == nil {
		v.TablaError.NewErrorSemantico(nil, "Contexto de for inválido")
		return nil
	}

	// 🔧 VERIFICAR QUE TENEMOS AL MENOS UN ID
	if len(ctx.AllID()) == 0 {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El for debe tener al menos una variable")
		return nil
	}

	// 🆕 DETECTAR SI HAY UNA O DOS VARIABLES CON VALIDACIÓN
	var nombreIndice string = ""
	var nombreValor string = ""

	if len(ctx.AllID()) == 1 {
		// for val in collection (comportamiento anterior)
		if ctx.ID(0) == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable del for no válida")
			return nil
		}
		nombreValor = ctx.ID(0).GetText()
	} else if len(ctx.AllID()) == 2 {
		// for idx, val in collection (nuevo comportamiento)
		if ctx.ID(0) == nil || ctx.ID(1) == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variables del for no válidas")
			return nil
		}
		nombreIndice = ctx.ID(0).GetText()
		nombreValor = ctx.ID(1).GetText()
	} else {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El for debe tener 1 ó 2 variables")
		return nil
	}

	var iterableItem *TipoVector = TipoVectorVacio

	// Evaluar range (opcional)
	if ctx.Range_() != nil {
		rangeResult := v.Visit(ctx.Range_())
		if rangeResult == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error evaluando rango")
			return nil
		}
		rangeItem, ok := rangeResult.(*TipoVector)
		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "El valor del rango debe ser un vector")
			return nil
		}
		iterableItem = rangeItem
	}

	// Evaluar expresión (opcional)
	if ctx.Expr() != nil {
		exprResult := v.Visit(ctx.Expr())
		if exprResult == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error evaluando expresión del for")
			return nil
		}

		iterableValue := exprResult.(tiposDeDato.ValorInterno)
		// 🔧 MEJORAR DETECCIÓN DE VECTORES
		tipoIterable := iterableValue.Type()

		// 🎯 VERIFICACIÓN MÁS ROBUSTA DE VECTORES
		if vectorVal, esVector := iterableValue.(*TipoVector); esVector {
			// Es directamente un TipoVector
			iterableItem = vectorVal
		} else if strings.HasPrefix(tipoIterable, "[]") {
			// Es un tipo vector ([]int, []string, etc.) pero no se reconoció como TipoVector
			// Intentar cast a TipoVector
			if vectorCast, ok := iterableValue.(*TipoVector); ok {
				iterableItem = vectorCast
			} else {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("El valor es de tipo vector (%s) pero no se puede iterar", tipoIterable))
				return nil
			}
		} else if tipoIterable == tiposDeDato.TIPO_CADENA {
			// Es una cadena
			iterableItem = CadenaAVector(iterableValue.(*tiposDeDato.ValorCadena))
		} else {
			// 🔧 MENSAJE DE ERROR MÁS DESCRIPTIVO
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("El valor de tipo '%s' no se puede iterar. Solo se permiten vectores ([]tipo) o cadenas", tipoIterable))
			return nil
		}
	}

	// 🔧 VALIDAR QUE TENEMOS ALGO QUE ITERAR
	if iterableItem == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "No se pudo determinar qué iterar")
		return nil
	}

	if iterableItem.Size() == 0 {
		return nil // Vector vacío, no hacer nada pero no es error
	}

	// Crear ámbito externo
	outerForScope := v.RegistroAmbito.PushAmbito("forRange")

	// 🆕 CREAR VARIABLE DEL ÍNDICE (si existe) CON VALIDACIÓN
	var indiceVariable *Variable = nil
	if nombreIndice != "" {
		// 🔧 VALIDAR TOKEN ANTES DE USAR
		var indiceToken antlr.Token = nil
		if ctx.ID(0) != nil && ctx.ID(0).GetSymbol() != nil {
			indiceToken = ctx.ID(0).GetSymbol()
		} else {
			indiceToken = ctx.GetStart() // Fallback
		}

		var msg string
		indiceVariable, msg = outerForScope.AgregarVariable(
			nombreIndice,
			tiposDeDato.TIPO_ENTERO,
			&tiposDeDato.ValorEntero{ValorInterno: 0},
			false, false, indiceToken) // 🔧 CAMBIO: false en lugar de true
		if indiceVariable == nil {
			v.TablaError.NewErrorSemantico(indiceToken,
				"No se pudo crear la variable del índice '"+nombreIndice+"': "+msg)
			return nil
		}
	}

	// 🔧 CREAR VARIABLE DEL VALOR CON VALIDACIÓN MEJORADA
	var valorVariable *Variable
	var valorToken antlr.Token = nil

	if nombreIndice != "" {
		// Dos variables: usar ID(1) para el valor
		if ctx.ID(1) != nil && ctx.ID(1).GetSymbol() != nil {
			valorToken = ctx.ID(1).GetSymbol()
		} else {
			valorToken = ctx.GetStart() // Fallback
		}
	} else {
		// Una variable: usar ID(0) para el valor
		if ctx.ID(0) != nil && ctx.ID(0).GetSymbol() != nil {
			valorToken = ctx.ID(0).GetSymbol()
		} else {
			valorToken = ctx.GetStart() // Fallback
		}
	}

	valorVariable, msg := outerForScope.AgregarVariable(
		nombreValor,
		iterableItem.TipoElemento,
		iterableItem.Actual(),
		false, false, valorToken) // 🔧 CAMBIO: false en lugar de true
	if valorVariable == nil {
		v.TablaError.NewErrorSemantico(valorToken,
			"No se pudo crear la variable del valor '"+nombreValor+"': "+msg)
		return nil
	}

	// Preparar estructura de control (break, continue)
	forItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo:          []string{Detener, Continuar},
	}

	v.PilaLlamada.Push(forItem)

	// Crear ámbito interno del for
	innerForScope := v.RegistroAmbito.PushAmbito("forAnidado")

	// 🆕 LLAMAR FUNCIÓN MODIFICADA CON MANEJO DE ERRORES
	defer func() {
		// Limpiar estado
		if iterableItem != nil {
			iterableItem.Reset()
		}
		v.RegistroAmbito.PopAmbito() // inner_for
		v.RegistroAmbito.PopAmbito() // outer_for
		v.PilaLlamada.Limpiar(forItem)

		// 🔧 CAPTURAR PANICS INESPERADOS
		if r := recover(); r != nil {
			// Si es un panic de control de flujo, manejarlo normalmente
			if llamada, ok := r.(*LlamadaFunciones); ok {
				if llamada == forItem {
					return // Es nuestro panic de control
				}
				// Si es de otra función, re-lanzarlo
				panic(llamada)
			}
			// Si es otro tipo de panic, reportar error
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("Error interno en bucle for: %v", r))
		}
	}()

	v.VisitInnerForWithIndex(ctx, outerForScope, innerForScope, forItem,
		iterableItem, indiceVariable, valorVariable)

	return nil
}

// 🆕 NUEVA FUNCIÓN PARA MANEJAR ÍNDICE Y VALOR
func (v *PatronVIsitor) VisitInnerForWithIndex(ctx *parser.ForStmtContext,
    outerForScope *AmbitoBase, innerForScope *AmbitoBase, forItem *LlamadaFunciones,
    iterableItem *TipoVector, indiceVariable *Variable, valorVariable *Variable) {

    // handle break and continue statements from call stack
    defer func() {
        // ✅ CAMBIO: NO resetear el ámbito
        // innerForScope.Reset()  // ❌ ELIMINAR ESTA LÍNEA
        
        if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

            // Not a for item, propagate panic
            if item != forItem {
                panic(item)
            }

            // Continue
            if item.EsAccion(Continuar) {
                item.ResetAccion()  // reset action, can be used again
                iterableItem.Next() // next item
                v.VisitInnerForWithIndex(ctx, outerForScope, innerForScope, forItem,
                    iterableItem, indiceVariable, valorVariable) // continue
            }

            // Break
            if item.EsAccion(Detener) {
                return
            }
        }
    }()

    // BUCLE PRINCIPAL CON ÍNDICE
    currentIndex := 0
    for iterableItem.IndiceActual < iterableItem.Size() {

        // 🆕 ACTUALIZAR VARIABLE DEL ÍNDICE (si existe)
        if indiceVariable != nil {
            indiceVariable.Valor = &tiposDeDato.ValorEntero{ValorInterno: currentIndex}
        }

        // ACTUALIZAR VARIABLE DEL VALOR
        valorVariable.Valor = iterableItem.Actual()

        // EJECUTAR CUERPO DEL BUCLE
        for _, stmt := range ctx.AllStmt() {
            v.Visit(stmt)
        }

        // AVANZAR AL SIGUIENTE ELEMENTO
        iterableItem.Next()
        currentIndex++
        
        // ✅ CAMBIO: NO resetear el ámbito
        // innerForScope.Reset()  // ❌ ELIMINAR ESTA LÍNEA
    }
}

func (v *PatronVIsitor) VisitInnerFor(ctx *parser.ForStmtContext, outerForScope *AmbitoBase, innerForScope *AmbitoBase, forItem *LlamadaFunciones, iterableItem *TipoVector, iterableVariable *Variable) {

    // handle break and continue statements from call stack
    defer func() {

        // ✅ CAMBIO: NO resetear el ámbito
        // innerForScope.Reset()  // ❌ ELIMINAR ESTA LÍNEA
        
        if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

            // Not a for item, propagate panic
            if item != forItem {
                panic(item)
            }

            // Continue
            if item.EsAccion(Continuar) {
                item.ResetAccion()                                                                          // reset action, can be used again
                iterableItem.Next()                                                                         // next item
                v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable) // continue
            }

            // Break
            if item.EsAccion(Detener) {
                return
            }

        }
    }()

    // iterableItem.Size()
    for iterableItem.IndiceActual < iterableItem.Size() {

        // update variable value
        iterableVariable.Valor = iterableItem.Actual()

        for _, stmt := range ctx.AllStmt() {
            v.Visit(stmt)
        }

        iterableItem.Next()
        
        // ✅ CAMBIO: NO resetear el ámbito
        // innerForScope.Reset()  // ❌ ELIMINAR ESTA LÍNEA
    }
}

func (v *PatronVIsitor) VisitRangoNum(ctx *parser.RangoNumContext) interface{} {

	leftExpr := v.Visit(ctx.Expr(0)).(tiposDeDato.ValorInterno)
	rightExpr := v.Visit(ctx.Expr(1)).(tiposDeDato.ValorInterno)

	if leftExpr.Type() != tiposDeDato.TIPO_ENTERO || rightExpr.Type() != tiposDeDato.TIPO_ENTERO {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los valores de los rangos deben ser enteros")
		return tiposDeDato.NuloPorDefecto
	}

	left := leftExpr.(*tiposDeDato.ValorEntero).ValorInterno
	right := rightExpr.(*tiposDeDato.ValorEntero).ValorInterno

	if left > right {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El valor izquierdo del rango debe ser menor o igual al valor derecho")
	}

	var values []tiposDeDato.ValorInterno

	for i := left; i <= right; i++ {
		values = append(values, &tiposDeDato.ValorEntero{
			ValorInterno: i,
		})
	}

	return &TipoVector{
		ValorInterno: values,
		IndiceActual: 0,
		TipoElemento: tiposDeDato.TIPO_ENTERO,
	}
}

// For Clasico
func (v *PatronVIsitor) VisitForClasicoStmt(ctx *parser.ForClasicoStmtContext) interface{} {
	// Crear ámbito para el for clásico
	forClasicoScope := v.RegistroAmbito.PushAmbito("forClasico")

	// Crear elemento de control para break/continue
	forClasicoItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo:          []string{Detener, Continuar},
	}

	v.PilaLlamada.Push(forClasicoItem)

	// 🔧 MANEJO SEGURO CON DEFER
	defer func() {
		v.RegistroAmbito.PopAmbito()
		v.PilaLlamada.Limpiar(forClasicoItem)

		// 🔧 CAPTURAR PANICS NO CONTROLADOS
		if r := recover(); r != nil {
			if item, ok := r.(*LlamadaFunciones); ok && item != forClasicoItem {
				// Es un panic de otra función (return), propagarlo
				panic(item)
			}
			// Si es nuestro forClasicoItem o otro tipo, no hacer nada
		}
	}()

	// 🎯 EJECUTAR LA IMPLEMENTACIÓN INTERNA
	v.VisitInnerForClasico(ctx, forClasicoScope, forClasicoItem)

	return nil
}

// 🆕 IMPLEMENTACIÓN INTERNA DEL FOR CLÁSICO
func (v *PatronVIsitor) VisitInnerForClasico(ctx *parser.ForClasicoStmtContext, forClasicoScope *AmbitoBase, forClasicoItem *LlamadaFunciones) {
	// 🎯 1️⃣ INICIALIZACIÓN (opcional): i = 0
	if len(ctx.AllStmt_asignar()) > 0 && ctx.Stmt_asignar(0) != nil {
		v.Visit(ctx.Stmt_asignar(0))
	}

	for {
		var deberContinuar bool = true

		if ctx.Expr() != nil {
			condicionResult := v.Visit(ctx.Expr())
			if condicionResult == nil {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error evaluando condición del for clásico")
				return
			}

			condicion := condicionResult.(tiposDeDato.ValorInterno)

			if condicion.Type() != tiposDeDato.TIPO_BOOLEAN {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condición del for clásico debe ser un booleano")
				return
			}

			deberContinuar = condicion.(*tiposDeDato.ValorBool).InternalValor
		}

		if !deberContinuar {
			break
		}

		shouldBreak := false
		shouldContinue := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					if item, ok := r.(*LlamadaFunciones); ok && item == forClasicoItem {
						if item.EsAccion(Continuar) {
							item.ResetAccion()
							shouldContinue = true
							return
						} else if item.EsAccion(Detener) {
							item.ResetAccion()
							shouldBreak = true
							return
						}
					} else if item, ok := r.(*LlamadaFunciones); ok && item != forClasicoItem {
						panic(item)
					} else {
						v.TablaError.NewErrorSemantico(ctx.GetStart(),
							fmt.Sprintf("Error interno en for clásico: %v", r))
						shouldBreak = true
						return
					}
				}
			}()

			if ctx.AllStmt() != nil {
				for _, stmt := range ctx.AllStmt() {
					if stmt != nil {
						v.Visit(stmt)
					}
				}
			}
		}()

		if shouldBreak {
			break
		}

		// 🎯 3️⃣ INCREMENTO (siempre se ejecuta, incluso con continue)
		if len(ctx.AllStmt_asignar()) > 1 && ctx.Stmt_asignar(1) != nil {
			v.Visit(ctx.Stmt_asignar(1))
		}

		// ✅ CAMBIO: NO resetear variables del ámbito
		// Solo limpiar variables temporales si es absolutamente necesario
		// if forClasicoScope != nil {
		//     forClasicoScope.ResetVariables()  // ❌ ELIMINAR ESTA LÍNEA
		// }

		if shouldContinue {
			continue
		}
	}
}

func (v *PatronVIsitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeRetornar()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia return debe estar dentro de una funcion")
		return nil
	}

	item.RetornarValor = tiposDeDato.NuloPorDefecto
	item.Accion = Retornar

	if ctx.Expr() != nil {
		item.RetornarValor = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
	}

	panic(item)
}

func (v *PatronVIsitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeDetener()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia break debe estar dentro de un ciclo o un switch")
		return nil
	}

	item.Accion = Detener
	panic(item)
}

func (v *PatronVIsitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeContinuar()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia continue debe estar dentro de un ciclo")
		return nil
	}

	item.Accion = Continuar
	panic(item)
}

func (v *PatronVIsitor) VisitLlamarFuncion(ctx *parser.LlamarFuncionContext) interface{} {
	canditateName := v.Visit(ctx.PatronId()).(string)

	// 1. Funciones embebidas
	if nativeFunc, exists := FuncionesEmbebidas[canditateName]; exists {
		args := make([]*Argumento, 0)
		if ctx.Lista_argumentos() != nil {
			args = v.Visit(ctx.Lista_argumentos()).([]*Argumento)
		}

		ret, ok, msg := nativeFunc.Exec(v.GetInstruccionesContexto(), args)
		if !ok {
			if msg != "" {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			}
			return tiposDeDato.NuloPorDefecto
		}
		return ret
	}

	// 2. Funciones del usuario o constructores de structs
	funcObj, msg1 := v.RegistroAmbito.GetFuncion(canditateName)
	structObj, msg2 := v.RegistroAmbito.AmbitoGlobal.GetEstructura(canditateName)

	if funcObj == nil && structObj == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg1+msg2)
		return tiposDeDato.NuloPorDefecto
	}

	args := make([]*Argumento, 0)
	if ctx.Lista_argumentos() != nil {
		args = v.Visit(ctx.Lista_argumentos()).([]*Argumento)
	}

	// Prioridad: struct antes que función
	if structObj != nil {
		if ArgumentoValidoEstructura(args) {
			return NewTipoObjeto(v, canditateName, ctx.PatronId().GetStart(), args, false)
		}
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			"El struct '"+canditateName+"' no puede ser construido con los argumentos dados, y tampoco es una función válida.")
		return tiposDeDato.NuloPorDefecto
	}

	switch funcObj := funcObj.(type) {
	case *FuncionNativa:
		ret, ok, msg := funcObj.Exec(v.GetInstruccionesContexto(), args)
		if !ok {
			if msg != "" {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			}
			return tiposDeDato.NuloPorDefecto
		}
		return ret

	case *Funcion:
		// SIMPLIFICAR: Sin manejo especial de panics de recursión
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.RetornarValor

	case *FuncionNativaObjeto:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.RetornarValor

	default:
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo de función no soportado para '"+canditateName+"'")
		return tiposDeDato.NuloPorDefecto
	}
}

func (v *PatronVIsitor) VisitListaArgumentos(ctx *parser.ListaArgumentosContext) interface{} {

	args := make([]*Argumento, 0)

	for _, arg := range ctx.AllArgumento_fun() {
		args = append(args, v.Visit(arg).(*Argumento))
	}

	return args

}

func (v *PatronVIsitor) VisitFuncionArg(ctx *parser.FuncionArgContext) interface{} {
	argName := ""
	passByReference := false

	var argValue tiposDeDato.ValorInterno = tiposDeDato.NuloPorDefecto
	var argVariableRef *Variable = nil

	if ctx.PatronId() != nil {
		// NO asignar el nombre de la variable como nombre del argumento
		variableName := ctx.PatronId().GetText()
		argVariableRef = v.RegistroAmbito.GetVariable(variableName)

		if argVariableRef != nil {
			argValue = argVariableRef.Valor
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+variableName+" no encontrada")
		}
	} else {
		argValue = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
	}

	// Solo asignar argName si hay un nombre explícito con DOS_PUNTOS (argumentos con nombre)
	if ctx.ID() != nil {
		argName = ctx.ID().GetText()
	}

	return &Argumento{
		Nombre:       argName,
		Valor:        argValue,
		esReferencia: passByReference,
		Token:        ctx.GetStart(),
		VariableRef:  argVariableRef,
	}
}

func (v *PatronVIsitor) VisitFuncionDeclerada(ctx *parser.FuncionDecleradaContext) interface{} {

	if v.RegistroAmbito.AmbitoActual == v.RegistroAmbito.AmbitoGlobal {
		// aready declared by dcl_visitor
		return nil
	}

	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal && !v.RegistroAmbito.AmbitoActual.EsEstructura {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global o en un struct")
	}

	funcName := ctx.ID().GetText()

	params := make([]*Parametros, 0)

	if ctx.Lista_parametros() != nil {
		params = v.Visit(ctx.Lista_parametros()).([]*Parametros)
	}

	if len(params) > 0 {

		baseParamType := params[0].TipoParametro()

		for _, param := range params {
			if param.TipoParametro() != baseParamType {
				v.TablaError.NewErrorSemantico(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}
	}

	returnType := tiposDeDato.TIPO_NULO
	var returnTypeToken antlr.Token = nil

	if ctx.Tipo() != nil {
		returnType = v.Visit(ctx.Tipo()).(string)
		returnTypeToken = ctx.Tipo().GetStart()
	}

	body := ctx.AllStmt()

	function := &Funcion{ // pointer ?
		Nombre:           funcName,
		Parametros:       params,
		TipoRetorno:      returnType,
		Cuerpo:           body,
		AmbitoDeclaro:    v.RegistroAmbito.AmbitoActual,
		TokenTipoRetorno: returnTypeToken,
		Token:            ctx.GetStart(),
	}

	ok, msg := v.RegistroAmbito.AgregarFuncion(funcName, function)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	return function
}

func (v *PatronVIsitor) VisitListaParametros(ctx *parser.ListaParametrosContext) interface{} {

	params := make([]*Parametros, 0)

	for _, param := range ctx.AllParametro_fun() {
		params = append(params, v.Visit(param).(*Parametros))
	}

	return params
}

func (v *PatronVIsitor) VisitParametroFun(ctx *parser.ParametroFunContext) interface{} {

	externName := ""
	innerName := ""

	// at least ID(0) is defined
	// only 1 ID defined
	if ctx.ID(1) == nil {
		// innerName : type
		// _ : type
		innerName = ctx.ID(0).GetText()
	} else {
		// externName innerName : type
		externName = ctx.ID(0).GetText()
		innerName = ctx.ID(1).GetText()
	}

	passByReference := false

	paramType := v.Visit(ctx.Tipo()).(string)

	return &Parametros{
		NombreExterno:      externName,
		NombreInterno:      innerName,
		PasarPorReferencia: passByReference,
		Tipo:               paramType,
		Token:              ctx.GetStart(),
	}

}

// * Structs
func (v *PatronVIsitor) VisitDeclararStruct(ctx *parser.DeclararStructContext) interface{} {
	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los structs solo pueden ser declaradas en el scope global")
		return nil
	}

	structName := ctx.ID().GetText()

	// NUEVO: Crear scope temporal para procesar los campos
	structScope := NewAmbitoBase(structName+"_template", true)
	prevScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = structScope

	// Procesar cada campo del struct
	for _, fieldCtx := range ctx.AllPropiedad_struct() {
		v.Visit(fieldCtx)
	}

	// Restaurar scope
	v.RegistroAmbito.AmbitoActual = prevScope

	// Crear el struct con los campos procesados (mantener compatibilidad)
	newStruct := &Struct{
		Nombre:    structName,
		Atributos: ctx.AllPropiedad_struct(), // Mantener para compatibilidad con NewTipoObjeto
		Token:     ctx.GetStart(),
	}

	structAdded, msg := v.RegistroAmbito.AmbitoGlobal.AgregarEstructura(structName, newStruct)
	if !structAdded {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

func (v *PatronVIsitor) VisitPropiedadStruct(ctx *parser.PropiedadStructContext) interface{} {

	// Verificar que ctx y sus elementos no sean nil
	if ctx == nil {
		return nil
	}

	if ctx.Tipo() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo de campo no especificado")
		return nil
	}

	if ctx.ID() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Nombre de campo no especificado")
		return nil
	}

	// Nueva sintaxis: type ID = expr?
	fieldTypeResult := v.Visit(ctx.Tipo())
	if fieldTypeResult == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar tipo de campo")
		return nil
	}

	fieldType := fieldTypeResult.(string) // string, int, bool
	fieldName := ctx.ID().GetText()       // Nombre, Edad, EsEstudiante

	var defaultValue tiposDeDato.ValorInterno = tiposDeDato.ValorNoIniPorDefecto

	// Si hay valor por defecto: string Nombre = "Alice"
	if ctx.Expr() != nil {
		exprResult := v.Visit(ctx.Expr())
		if exprResult == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar valor por defecto")
			return nil
		}

		defaultValue = exprResult.(tiposDeDato.ValorInterno)

		// Verificar que el tipo coincida
		if defaultValue.Type() != fieldType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s al campo %s de tipo %s",
					defaultValue.Type(), fieldName, fieldType))
			return nil
		}
	}

	// Los campos de struct son siempre mutables por defecto
	variable, msg := v.RegistroAmbito.AgregarVariable(
		fieldName, fieldType, defaultValue,
		false, // mutable (siempre false para struct fields)
		false, // inicializada
		ctx.ID().GetSymbol())

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	return variable // Retornar la variable creada
}

func (v *PatronVIsitor) VisitCrearStruct(ctx *parser.CrearStructContext) interface{} {

	if ctx == nil {
		return tiposDeDato.NuloPorDefecto
	}

	if ctx.ID() == nil {
		return tiposDeDato.NuloPorDefecto
	}

	structName := ctx.ID().GetText()

	// Verificar que el struct existe
	structDef, msg := v.RegistroAmbito.AmbitoGlobal.GetEstructura(structName)
	if structDef == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	// Crear argumentos para el constructor existente
	args := make([]*Argumento, 0)

	if ctx.Lista_parametros_init() != nil {

		initFieldsResult := v.Visit(ctx.Lista_parametros_init())
		if initFieldsResult == nil {
			return tiposDeDato.NuloPorDefecto
		}

		initFields, ok := initFieldsResult.([]*ValorInicialStruct)
		if !ok {
			return tiposDeDato.NuloPorDefecto
		}

		for _, initField := range initFields {
			if initField == nil {
				continue
			}

			args = append(args, &Argumento{
				Nombre: initField.Nombre,
				Valor:  initField.Valor,
				Token:  initField.Token,
			})
		}
	}
	// Usar el constructor existente
	result := NewTipoObjeto(v, structName, ctx.GetStart(), args, false)

	return result
}

func (v *PatronVIsitor) VisitCrearStructExp(ctx *parser.CrearStructExpContext) interface{} {
	return v.Visit(ctx.Crear_struct())
}

func (v *PatronVIsitor) VisitListaParametrosInit(ctx *parser.ListaParametrosInitContext) interface{} {
	initFields := make([]*ValorInicialStruct, 0)

	for _, fieldCtx := range ctx.AllParametros_init_struct() {
		if field := v.Visit(fieldCtx); field != nil {
			if initField, ok := field.(*ValorInicialStruct); ok {
				initFields = append(initFields, initField)
			}
		}
	}

	return initFields
}

func (v *PatronVIsitor) VisitParametrosInitStruct(ctx *parser.ParametrosInitStructContext) interface{} {
	fieldName := ctx.ID().GetText()
	fieldValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	return &ValorInicialStruct{
		Nombre: fieldName,
		Valor:  fieldValue,
		Token:  ctx.GetStart(),
	}
}

func (v *PatronVIsitor) VisitVectorFuncExp(ctx *parser.VectorFuncExpContext) interface{} {
	return v.Visit(ctx.Fun_slice())
}

func (v *PatronVIsitor) VisitPropSliceExp(ctx *parser.PropSliceExpContext) interface{} {
	return v.Visit(ctx.Prop_slice())
}

func (v *PatronVIsitor) VisitPropSlice(ctx *parser.PropSliceContext) interface{} {

	var objectCandidate tiposDeDato.ValorInterno

	switch itemRef := v.Visit(ctx.Item_slice()).(type) {
	case *VectorItemReference:
		objectCandidate = itemRef.Valor
	case *ElementoMatriz:
		objectCandidate = itemRef.Valor
	}

	obj, ok := objectCandidate.(*TipoObjeto)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El item del vector no es un struct")
		return tiposDeDato.NuloPorDefecto
	}

	lastScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = obj.AmbitoInterno

	defer func() {
		v.RegistroAmbito.AmbitoActual = lastScope
	}()

	variable := v.RegistroAmbito.GetVariable(ctx.PatronId().GetText())

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Propiedad "+ctx.PatronId().GetText()+" no encontrada item del vector")
		return tiposDeDato.NuloPorDefecto
	}

	return variable.Valor
}

func (v *PatronVIsitor) VisitFuncionSlice(ctx *parser.FuncionSliceContext) interface{} {

	var objectCandidate tiposDeDato.ValorInterno

	switch itemRef := v.Visit(ctx.Item_slice()).(type) {
	case *VectorItemReference:
		objectCandidate = itemRef.Valor
	case *ElementoMatriz:
		objectCandidate = itemRef.Valor
	}

	obj, ok := objectCandidate.(*TipoObjeto)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El objeto no es un struct")
		return tiposDeDato.NuloPorDefecto
	}

	lastScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = obj.AmbitoInterno

	defer func() {
		v.RegistroAmbito.AmbitoActual = lastScope
	}()

	return v.Visit(ctx.Llamar_funcion())
}

func (s *RegistroAmbito) Print() {

	fmt.Println("Global Scope")
	fmt.Println("============")

	fmt.Println("Variables")
	for k, v := range s.AmbitoGlobal.Variables {
		fmt.Println(k, v.Valor.Value(), v.Tipo)
	}

	fmt.Println("Funciones")
	for k, v := range s.AmbitoGlobal.Funciones {
		fmt.Println(k, v)
	}

	fmt.Println("Child Scopes")
	fmt.Println("============")
	fmt.Println("")

	for _, child := range s.AmbitoGlobal.Hijos {

		fmt.Println(child.Nombre)
		fmt.Println("============")

		fmt.Println("Variables")
		for k, v := range child.Variables {
			fmt.Println(k, v.Valor.Value())
		}

		fmt.Println("Funciones")
		for k, v := range child.Funciones {
			fmt.Println(k, v)
		}

		fmt.Println("")
	}

}

// Declaración sin 'mut/let/var', con tipo y valor:  ID type = expr
func (v *PatronVIsitor) VisitDeclararSinMutValor(ctx *parser.DeclararSinMutValorContext) interface{} {
	nombreVariable := ctx.ID().GetText()
	tipoVariable := v.Visit(ctx.Tipo()).(string)
	valorVariable := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// Copiar objetos y vectores para evitar aliasing
	if objeto, ok := valorVariable.(*TipoObjeto); ok {
		valorVariable = objeto.Copy()
	}
	if EsTipoVector(valorVariable.Type()) {
		valorVariable = valorVariable.Copy()
	}

	// isConst = false, inferir inicialización completa => inferir=false
	variable, msg := v.RegistroAmbito.AgregarVariable(nombreVariable, tipoVariable, valorVariable, false, false, ctx.GetStart())
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}
