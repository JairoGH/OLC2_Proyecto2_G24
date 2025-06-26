package analyzer

import (
	"fmt"
	"main/crearArbol"
	"main/models"
	"os"
	"path/filepath"

	errores "main/Errores"
	instrucciones "main/Instrucciones"
	compiler "main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type AnalyzerService struct{}

func NewAnalyzerService() *AnalyzerService {
	return &AnalyzerService{}
}

// 🆕 MÉTODO 1: Solo análisis léxico, sintáctico y semántico
func (a *AnalyzerService) AnalyzeOnly(code, filename string) (*models.AnalysisOnlyResponse, antlr.ParseTree, *instrucciones.PatronVIsitor, error) {
	defer func() {
		if r := recover(); r != nil {
			// Manejo de errores - no crashear la aplicación
		}
	}()

	// === ANÁLISIS LÉXICO ===
	lexicalErrorListener := errores.NewErrorLexico()
	lexer := compiler.NewVLexer(antlr.NewInputStream(code))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// === ANÁLISIS SINTÁCTICO ===
	parser := compiler.NewVGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errores.NewSyntaxError(lexicalErrorListener.TablaError)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errores.NewErrorPersonalizado())
	parser.AddErrorListener(syntaxErrorListener)

	tree := parser.Program()

	// === ANÁLISIS SEMÁNTICO ===
	dclVisitor := instrucciones.NewVisitanteDcl(syntaxErrorListener.TablaError)
	dclVisitor.Visit(tree)

	// === INTERPRETACIÓN ===
	replVisitor := instrucciones.NewVisitor(dclVisitor)
	replVisitor.Visit(tree)

	// === CONVERTIR ERRORES ===
	success := len(replVisitor.TablaError.Errores) == 0
	frontendErrors := []models.ErrorReport{}
	for _, err := range replVisitor.TablaError.Errores {
		frontendErrors = append(frontendErrors, models.ErrorReport{
			Linea:       err.Linea,
			Columna:     err.Columna,
			Descripcion: err.Descripcion,
			Tipo:        err.Tipo,
		})
	}

	// === CONVERTIR TABLA DE SÍMBOLOS ===
	symbolTable := a.convertSymbolTable(replVisitor.RegistroAmbito.Report())

	// === PREPARAR OUTPUT DE CONSOLA ===
	consoleOutput := replVisitor.Consola.GetSalida()
	if consoleOutput == "" {
		if success {
			consoleOutput = "✅ Análisis completado exitosamente\n✅ No se encontraron errores"
		} else {
			consoleOutput = "❌ Se encontraron errores durante el análisis"
		}
	}

	result := &models.AnalysisOnlyResponse{
		Success:       success,
		Errors:        frontendErrors,
		SymbolTable:   symbolTable,
		ConsoleOutput: consoleOutput,
		ParseTree:     tree,
		Visitor:       replVisitor,
	}

	return result, tree, replVisitor, nil
}

// 🆕 MÉTODO 2: Solo generación ARM64 (recibe el árbol ya parseado)
func (a *AnalyzerService) GenerateARM64Only(tree antlr.ParseTree, visitor *instrucciones.PatronVIsitor) (*models.ARM64Response, error) {
	defer func() {
		if r := recover(); r != nil {
			// Manejo de errores
		}
	}()

	// Verificar que no hay errores antes de generar ARM64
	if len(visitor.TablaError.Errores) > 0 {
		return &models.ARM64Response{
			Success:       false,
			ErrorMessage:  "No se puede generar ARM64 porque hay errores en el código",
			ConsoleOutput: "❌ No se generó código ARM64 porque hubo errores.",
		}, nil
	}

	// === GENERAR CÓDIGO ARM64 ===
	armVisitor := instrucciones.NewVisitorARM64()
	armVisitor.Visit(tree) // Usar el árbol ya parseado
	armCode := armVisitor.GetCodigo()

	// === CREAR DIRECTORIO Y GUARDAR ARCHIVO ===
	outputDir := "output"
	os.MkdirAll(outputDir, 0755)

	archivoAsm := filepath.Join(outputDir, "programa.s")
	err := os.WriteFile(archivoAsm, []byte(armCode), 0644)

	var consoleOutput string
	if err != nil {
		consoleOutput = fmt.Sprintf("⚠️ Error al guardar archivo ARM64: %v", err)
		return &models.ARM64Response{
			Success:       false,
			ErrorMessage:  err.Error(),
			ConsoleOutput: consoleOutput,
		}, err
	}

	consoleOutput = fmt.Sprintf("\n📁 Archivo: %s ", archivoAsm)

	return &models.ARM64Response{
		Success:       true,
		ConsoleOutput: consoleOutput,
		ARM64Code:     armCode,
		FilePath:      archivoAsm,
	}, nil
}

// 🆕 MÉTODO 3: Solo generación AST (independiente)
func (a *AnalyzerService) GenerateASTOnly(code string) (*models.ASTResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			// Manejo de errores
		}
	}()

	// === GENERAR AST EN PARALELO ===
	resultChannel := make(chan string)
	go func() {
		resultChannel <- crearArbol.ReporteArbol(code)
	}()

	// Obtener el reporte del AST
	cstReport := <-resultChannel

	if cstReport == "" {
		return &models.ASTResponse{
			Success:       false,
			ErrorMessage:  "No se pudo generar el AST",
			ConsoleOutput: "❌ Error al generar el árbol de sintaxis",
		}, fmt.Errorf("AST generation failed")
	}

	consoleOutput := "\n🌳 Árbol AST generado exitosamente\n✅ SVG generado correctamente"

	return &models.ASTResponse{
		Success:       true,
		ConsoleOutput: consoleOutput,
		SVGContent:    cstReport,
	}, nil
}

// Mantener métodos auxiliares sin cambios
func (a *AnalyzerService) convertSymbolTable(registroAmbito instrucciones.ReporteTabla) []models.SymbolEntry {
	var symbolTable []models.SymbolEntry

	var procesarAmbitoRecursivo func(ambito instrucciones.ReporteAmbito, rutaScope string)

	procesarAmbitoRecursivo = func(ambito instrucciones.ReporteAmbito, rutaScope string) {
		for _, variable := range ambito.Variables {
			line := variable.Linea
			column := variable.Columna
			if line == 0 && column == 0 {
				line = 1
				column = 1
			}

			symbolTable = append(symbolTable, models.SymbolEntry{
				ID:         variable.Nombre,
				SymbolType: "Variable",
				DataType:   variable.Tipo,
				Scope:      rutaScope,
				Line:       line,
				Column:     column,
			})
		}

		for _, function := range ambito.Funciones {
			line := function.Linea
			column := function.Columna
			if line == 0 && column == 0 {
				line = 1
				column = 1
			}

			symbolTable = append(symbolTable, models.SymbolEntry{
				ID:         function.Nombre,
				SymbolType: "Function",
				DataType:   function.Tipo,
				Scope:      rutaScope,
				Line:       line,
				Column:     column,
			})
		}

		for _, estructura := range ambito.Estructuras {
			line := estructura.Linea
			column := estructura.Columna
			if line == 0 && column == 0 {
				line = 1
				column = 1
			}

			symbolTable = append(symbolTable, models.SymbolEntry{
				ID:         estructura.Nombre,
				SymbolType: "Struct",
				DataType:   estructura.Tipo,
				Scope:      rutaScope,
				Line:       line,
				Column:     column,
			})
		}

		for i, hijo := range ambito.AmbitosHijos {
			nombreHijo := hijo.Nombre
			if nombreHijo == "" {
				nombreHijo = fmt.Sprintf("block_%d", i+1)
			}

			nuevaRuta := rutaScope + "." + nombreHijo
			procesarAmbitoRecursivo(hijo, nuevaRuta)
		}
	}

	procesarAmbitoRecursivo(registroAmbito.AmbitoGlobal, "global")
	return symbolTable
}

// ValidateCode valida la sintaxis del código
func (a *AnalyzerService) ValidateCode(code string) (*models.AnalysisResponse, error) {
	// Usar el método optimizado
	analysisResult, _, _, err := a.AnalyzeOnly(code, "validation.vch")
	if err != nil {
		return nil, err
	}

	// Convertir a formato legacy para compatibilidad
	result := &models.AnalysisResponse{
		Success:       analysisResult.Success,
		AST:           "Validación completada",
		Errors:        analysisResult.Errors,
		SymbolTable:   analysisResult.SymbolTable,
		ConsoleOutput: analysisResult.ConsoleOutput,
		CSTSvg:        "",
		ASTPng:        nil,
	}

	return result, nil
}

// GetVersion retorna información del analizador
func (a *AnalyzerService) GetVersion() (string, error) {
	return "V-Lang Cherry Analyzer (Direct Integration with Full Backend)", nil
}

// Ping siempre retorna éxito ya que no hay servidor externo
func (a *AnalyzerService) Ping() error {
	return nil // Siempre disponible
}
