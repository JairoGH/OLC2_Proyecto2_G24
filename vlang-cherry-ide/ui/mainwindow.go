package ui

import (
	"main/api"
	"main/models"
	"strings"

	instrucciones "main/Instrucciones" // 🆕 AGREGAR ESTE IMPORT

	"github.com/antlr4-go/antlr/v4" // 🆕 AGREGAR ESTE IMPORT

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

// 🔧 ACTUALIZAR: Estado para controlar las fases (tipos correctos)
type AnalysisState struct {
	IsAnalyzed   bool
	HasErrors    bool
	Code         string
	Filename     string
	ParseTree    antlr.ParseTree              // 🔧 Tipo correcto
	Visitor      *instrucciones.PatronVIsitor // 🔧 Tipo correcto
	LastAnalysis *models.AnalysisOnlyResponse // 🔧 Tipo correcto
}

type MainWindow struct {
	App           fyne.App
	Window        fyne.Window
	Editor        *Editor
	Console       *Console
	Reports       *Reports
	MenuBar       *Menus
	APIClient     *api.Client
	AnalysisState *AnalysisState
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("V-Lang Cherry IDE")
	window.SetMaster()
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()

	mainWindow := &MainWindow{
		App:     app,
		Window:  window,
		Editor:  NewEditor(),
		Console: NewConsole(),
		Reports: NewReports(),

		// 🆕 INICIALIZAR estado
		AnalysisState: &AnalysisState{
			IsAnalyzed: false,
			HasErrors:  false,
		},
	}

	// Crear cliente API
	mainWindow.APIClient = api.NewClient()

	// Crear menús pasando la referencia a mainWindow
	mainWindow.MenuBar = NewMenus(mainWindow)

	// Configurar layout
	content := mainWindow.createLayout()
	mainWindow.Window.SetContent(content)

	// Configurar menús
	mainWindow.Window.SetMainMenu(mainWindow.MenuBar.MainMenu)

	return mainWindow
}

// 🔧 MÉTODO 1: Ejecutar solo análisis (IMPLEMENTACIÓN REAL)
func (mw *MainWindow) ExecuteAnalysis() {
	instrucciones.ResetAllRecursionFlags()
	code := mw.Editor.GetCode()
	if strings.TrimSpace(code) == "" {
		dialog.ShowInformation("Advertencia", "No hay código para analizar", mw.Window)
		return
	}

	mw.Console.ClearOutput()
	mw.Reports.ClearAll()
	mw.Console.AddOutput("🔍 Iniciando análisis léxico, sintáctico y semántico...\n")

	// 🚀 IMPLEMENTACIÓN REAL - Usar el servicio separado
	result, tree, visitor, err := mw.APIClient.AnalyzeOnly(code, "temp.vch")
	if err != nil {
		mw.Console.AddOutput("❌ Error durante el análisis: " + err.Error() + "\n")
		mw.AnalysisState.IsAnalyzed = false
		return
	}

	// 🎯 GUARDAR ESTADO PARA LOS SIGUIENTES MÉTODOS
	mw.AnalysisState = &AnalysisState{
		IsAnalyzed:   true,
		HasErrors:    !result.Success,
		Code:         code,
		Filename:     "temp.vch",
		ParseTree:    tree,
		Visitor:      visitor,
		LastAnalysis: result,
	}

	// 🎨 MOSTRAR RESULTADOS EN LA UI
	mw.Console.AddOutput(result.ConsoleOutput + "\n")
	mw.Reports.UpdateErrors(result.Errors)
	mw.Reports.UpdateSymbolTable(result.SymbolTable)

	// 🎯 MOSTRAR ESTADO FINAL
	if result.Success {
		mw.Console.AddOutput("\n✅ Análisis completado - ¡Listo para ARM64 y AST!\n")
		mw.Console.AddOutput("💡 Usa '⚙️ Generar ARM64' para compilar\n")
		mw.Console.AddOutput("💡 Usa '🌳 Generar AST' para ver el árbol\n")
	} else {
		mw.Console.AddOutput("\n⚠️ Análisis completado con errores\n")
		mw.Console.AddOutput("❌ ARM64 no disponible hasta corregir errores\n")
		mw.Console.AddOutput("✅ AST sí está disponible para depuración\n")
	}
}

// 🔧 MÉTODO 2: Generar solo ARM64 (IMPLEMENTACIÓN REAL)
func (mw *MainWindow) GenerateARM64() {
    if !mw.AnalysisState.IsAnalyzed {
        dialog.ShowInformation("Análisis requerido",
            "Primero ejecuta '🔍 Ejecutar Análisis'", mw.Window)
        return
    }

    if mw.AnalysisState.HasErrors {
        dialog.ShowInformation("Errores encontrados",
            "No se puede generar ARM64 con errores presentes.\n\nRevisa la pestaña 'Errores' y corrige el código.", mw.Window)
        return
    }

    mw.Console.AddOutput("⚙️ Generando ARM64...\n")

    // 🚀 IMPLEMENTACIÓN REAL - Usar árbol y visitor ya procesados
    armResult, err := mw.APIClient.GenerateARM64Only(mw.AnalysisState.ParseTree, mw.AnalysisState.Visitor)
    if err != nil {
        mw.Console.AddOutput("❌ Error generando ARM64: " + err.Error() + "\n")
        return
    }

    // 🎨 MOSTRAR RESULTADO EN CONSOLA
    mw.Console.AddOutput(armResult.ConsoleOutput + "\n")

    if armResult.Success {
        // 🆕 MOSTRAR CÓDIGO EN LA PESTAÑA ARM64
        mw.Reports.UpdateARM64Content(armResult.ARM64Code, armResult.FilePath)
        mw.Console.AddOutput("✅ ARM64 generado - Ve a la pestaña ARM64\n")
    }
}

// 🔧 MÉTODO 3: Generar solo AST (IMPLEMENTACIÓN REAL)
// 🔧 MÉTODO 3: Generar solo AST (IMPLEMENTACIÓN REAL)
func (mw *MainWindow) GenerateAST() {
    if !mw.AnalysisState.IsAnalyzed {
        dialog.ShowInformation("Análisis requerido",
            "Primero ejecuta '🔍 Ejecutar Análisis'", mw.Window)
        return
    }

    mw.Console.AddOutput("🌳 Generando AST...\n")

    // 🚀 IMPLEMENTACIÓN REAL - Generar AST independiente
    astResult, err := mw.APIClient.GenerateASTOnly(mw.AnalysisState.Code)
    if err != nil {
        mw.Console.AddOutput("❌ Error generando AST: " + err.Error() + "\n")
        return
    }

    // 🎨 MOSTRAR RESULTADO
    mw.Console.AddOutput(astResult.ConsoleOutput + "\n")

    if astResult.Success {
        // 🔧 CONECTAR CON LA PESTAÑA AST - ESTO FALTABA
        mw.Reports.UpdateASTWithSVG("", astResult.SVGContent)
        mw.Console.AddOutput("✅ AST generado - Ve a la pestaña 'AST'\n")
    }
}

func (mw *MainWindow) createLayout() *container.Split {
	// Panel superior: Editor y Reportes
	topPanel := container.NewHSplit(
		mw.Editor.Container,  // Izquierda: Editor
		mw.Reports.Container, // Derecha: Reportes
	)
	topPanel.SetOffset(0.5) // 50% para editor, 50% para reportes

	// Layout principal: Split vertical con consola en la parte inferior
	mainSplit := container.NewVSplit(
		topPanel,             // Arriba: Editor + Reportes
		mw.Console.Container, // Abajo: Consola (ancho completo)
	)
	mainSplit.SetOffset(0.6) // 60% para panel superior, 40% para consola

	return mainSplit
}

func (mw *MainWindow) ShowAndRun() {
	mw.Window.ShowAndRun()
}

// Métodos para manejar acciones del menú
func (mw *MainWindow) NewFile() {
	mw.Editor.NewFile()
	// 🆕 RESETEAR ESTADO AL CREAR NUEVO ARCHIVO
	mw.AnalysisState = &AnalysisState{
		IsAnalyzed: false,
		HasErrors:  false,
	}
}

func (mw *MainWindow) OpenFile() {
	mw.Editor.OpenFile(mw.Window)
	// 🆕 RESETEAR ESTADO AL ABRIR ARCHIVO
	mw.AnalysisState = &AnalysisState{
		IsAnalyzed: false,
		HasErrors:  false,
	}
}

func (mw *MainWindow) SaveFile() {
	mw.Editor.SaveFile(mw.Window)
}

