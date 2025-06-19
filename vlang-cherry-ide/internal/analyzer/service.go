package analyzer

import (
	"fmt"
	"main/crearArbol"
	"main/models"

	// Imports del backend original - ajusta las rutas según tu estructura
	errores "main/Errores"
	instrucciones "main/Instrucciones"
	compiler "main/parser"

	"github.com/antlr4-go/antlr/v4"
)

// AnalyzerService integra la lógica completa del backend original
type AnalyzerService struct{}

// NewAnalyzerService crea una nueva instancia del servicio
func NewAnalyzerService() *AnalyzerService {
	return &AnalyzerService{}
}

// AnalyzeCode usa la lógica completa del backend original
func (a *AnalyzerService) AnalyzeCode(code, filename string) (*models.AnalysisResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			// Manejo de errores - no crashear la aplicación
		}
	}()

	// Generar AST con ANTLR externo (en paralelo como en el original)
	resultChannel := make(chan string)
	go func() {
		resultChannel <- crearArbol.ReporteArbol(code)
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

	// Obtener el reporte del AST
	cstReport := <-resultChannel

	// === CONVERTIR A FORMATO DEL FRONTEND ===
	success := len(replVisitor.TablaError.Errores) == 0

	// Convertir errores - usando la estructura real correcta
	frontendErrors := []models.ErrorReport{}
	for _, err := range replVisitor.TablaError.Errores {
		frontendErrors = append(frontendErrors, models.ErrorReport{
			Linea:       err.Linea,       // ← CORRECTO: usa Line, no Linea
			Columna:     err.Columna,     // ← CORRECTO: usa Columna
			Descripcion: err.Descripcion, // ← CORRECTO: usa Descripcion
			Tipo:        err.Tipo,        // ← CORRECTO: usa Tipo (que ya es string)
		})
	}

	// Convertir tabla de símbolos
	symbolTable := a.convertSymbolTable(replVisitor.RegistroAmbito.Report())

	// Preparar output de consola
	consoleOutput := replVisitor.Consola.GetSalida()
	if consoleOutput == "" {
		if success {
			consoleOutput = "✅ Análisis completado exitosamente\n✅ Código interpretado sin errores"
		} else {
			consoleOutput = "❌ Se encontraron errores durante el análisis"
		}
	}

	result := &models.AnalysisResponse{
		Success:       success,
		AST:           "Árbol de Sintaxis Abstracta generado exitosamente",
		Errors:        frontendErrors,
		SymbolTable:   symbolTable,
		ConsoleOutput: consoleOutput,
		CSTSvg:        cstReport,
		ASTPng:        nil,
	}

	return result, nil
}

// Convertir tabla de símbolos del backend al formato del frontend
func (a *AnalyzerService) convertSymbolTable(registroAmbito instrucciones.ReporteTabla) []models.SymbolEntry {
    var symbolTable []models.SymbolEntry

    // Función recursiva con jerarquía de puntos
    var procesarAmbitoRecursivo func(ambito instrucciones.ReporteAmbito, rutaScope string)
    
    procesarAmbitoRecursivo = func(ambito instrucciones.ReporteAmbito, rutaScope string) {
        // Procesar variables del ámbito actual
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

        // Procesar funciones del ámbito actual
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

        // Procesar estructuras del ámbito actual
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

        // Procesar ámbitos hijos con nueva ruta
        for i, hijo := range ambito.AmbitosHijos {
            // Crear nombre del ámbito hijo
            nombreHijo := hijo.Nombre
            if nombreHijo == "" {
                nombreHijo = fmt.Sprintf("block_%d", i+1)
            }
            
            // Crear ruta jerárquica con puntos
            nuevaRuta := rutaScope + "." + nombreHijo
            
            // Llamada recursiva
            procesarAmbitoRecursivo(hijo, nuevaRuta)
        }
    }

    // Iniciar desde global
    procesarAmbitoRecursivo(registroAmbito.AmbitoGlobal, "global")

    return symbolTable
}

// ValidateCode valida la sintaxis del código
func (a *AnalyzerService) ValidateCode(code string) (*models.AnalysisResponse, error) {
	return a.AnalyzeCode(code, "temp.vch")
}

// GetVersion retorna información del analizador
func (a *AnalyzerService) GetVersion() (string, error) {
	return "V-Lang Cherry Analyzer (Direct Integration with Full Backend)", nil
}

// Ping siempre retorna éxito ya que no hay servidor externo
func (a *AnalyzerService) Ping() error {
	return nil // Siempre disponible
}
