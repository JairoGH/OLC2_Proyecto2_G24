/*

package ui

import (
	"fmt"
	"strings"
	"vlang-cherry-ide/api" // Importar el paquete api

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type MainWindow struct {
	App       fyne.App
	Window    fyne.Window
	Editor    *Editor
	Console   *Console
	Reports   *Reports
	MenuBar   *Menus
	APIClient *api.Client // AGREGAR ESTE CAMPO
}

func NewMainWindow(myApp fyne.App) *MainWindow {
	// Crear ventana principal
	window := myApp.NewWindow("V-Lang Cherry IDE")
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()

	// Crear componentes
	editor := NewEditor()
	console := NewConsole()
	reports := NewReports()
	menuBar := NewMenus()

	mainWin := &MainWindow{
		App:       myApp,
		Window:    window,
		Editor:    editor,
		Console:   console,
		Reports:   reports,
		MenuBar:   menuBar,
		APIClient: api.NewClient(), // INICIALIZAR EL CLIENTE
	}

	// Configurar menús
	menuBar.SetupMenus(mainWin)
	window.SetMainMenu(menuBar.MainMenu)

	// Crear layout principal
	content := mainWin.createLayout()
	window.SetContent(content)

	return mainWin
}

func (mw *MainWindow) createLayout() *container.Split {
	// Panel superior: Editor y Reportes
	topPanel := container.NewHSplit(
		mw.Editor.Container,  // Izquierda: Editor
		mw.Reports.Container, // Derecha: Reportes
	)
	topPanel.SetOffset(0.5) // 70% para editor, 30% para reportes

	// Layout principal: Split vertical con consola en la parte inferior
	mainSplit := container.NewVSplit(
		topPanel,             // Arriba: Editor + Reportes
		mw.Console.Container, // Abajo: Consola (ancho completo)
	)
	mainSplit.SetOffset(0.6) // 75% para panel superior, 25% para consola

	return mainSplit
}

func (mw *MainWindow) ShowAndRun() {
	mw.Window.ShowAndRun()
}

// Métodos para manejar acciones del menú
func (mw *MainWindow) NewFile() {
	mw.Editor.NewFile()
}

func (mw *MainWindow) OpenFile() {
	mw.Editor.OpenFile(mw.Window)
}

func (mw *MainWindow) SaveFile() {
	mw.Editor.SaveFile(mw.Window)
}

func (mw *MainWindow) ExecuteCode() {
	code := mw.Editor.GetCode()

	// Verificar que hay código
	if strings.TrimSpace(code) == "" {
		dialog.ShowInformation("Advertencia", "No hay código para analizar", mw.Window)
		return
	}

	// Limpiar consola y reportes antes de ejecutar
	mw.Console.ClearOutput()
	mw.Reports.ClearAll()

	// Verificar conexión con el servidor
	mw.Console.AddOutput("Verificando conexión con el servidor...\n")
	if err := mw.APIClient.Ping(); err != nil {
		mw.Console.AddOutput("❌ Error: No se puede conectar con el servidor del compilador\n")
		mw.Console.AddOutput("💡 Asegúrate de que el servidor esté ejecutándose:\n")
		mw.Console.AddOutput("   cd docs && go run main.go\n")
		dialog.ShowError(fmt.Errorf("servidor no disponible: %v", err), mw.Window)
		return
	}

	// Mostrar mensaje de carga
	mw.Console.AddOutput("✅ Conexión establecida\n")
	mw.Console.AddOutput("🔍 Analizando código...\n")

	// Llamar al backend para análisis
	filename := mw.Editor.CurrentFile
	if filename == "" {
		filename = "untitled.vch"
	}

	result, err := mw.APIClient.AnalyzeCode(code, filename)
	if err != nil {
		mw.Console.AddOutput("❌ Error al analizar código: " + err.Error() + "\n")
		dialog.ShowError(err, mw.Window)
		return
	}

	// Mostrar resultados en consola
	mw.Console.ClearOutput()
	mw.Console.AddOutput(result.ConsoleOutput)

	if result.Success {
		mw.Console.AddOutput("\n✅ Análisis completado exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Errores encontrados: %d\n", len(result.Errors)))
		mw.Console.AddOutput(fmt.Sprintf("🏷️  Símbolos en tabla: %d\n", len(result.SymbolTable)))
	} else {
		mw.Console.AddOutput("\n❌ Se encontraron errores en el código\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Total de errores: %d\n", len(result.Errors)))
	}

	// Actualizar reportes
	mw.Reports.UpdateErrors(result.Errors)
	mw.Reports.UpdateSymbolTable(result.SymbolTable)
	mw.Reports.UpdateASTWithSVG(result.AST, result.CSTSvg)

	// Si hay SVG, mostrar solo información, no el contenido completo
	if result.CSTSvg != "" {
		mw.Console.AddOutput("\n🌳 SVG del AST generado exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📄 Tamaño del SVG: %d caracteres\n", len(result.CSTSvg)))
		mw.Console.AddOutput("💡 Visualiza el AST en la pestaña 'AST'\n")
	}

	// Mostrar pestaña de errores si hay errores
	if len(result.Errors) > 0 {
		mw.Reports.Container.SelectTab(mw.Reports.Container.Items[0]) // Pestaña de errores
	}
}
*/
/*
package ui

import (
	"fmt"
	"strings"
	"vlang-cherry-ide/api" // Importar el paquete api

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type MainWindow struct {
	App       fyne.App
	Window    fyne.Window
	Editor    *Editor
	Console   *Console
	Reports   *Reports
	MenuBar   *Menus
	APIClient *api.Client // AGREGAR ESTE CAMPO
}

func NewMainWindow(myApp fyne.App) *MainWindow {
	// Crear ventana principal
	window := myApp.NewWindow("V-Lang Cherry IDE")
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()

	// Crear componentes
	editor := NewEditor()
	console := NewConsole()
	reports := NewReports()
	menuBar := NewMenus()

	mainWin := &MainWindow{
		App:       myApp,
		Window:    window,
		Editor:    editor,
		Console:   console,
		Reports:   reports,
		MenuBar:   menuBar,
		APIClient: api.NewClient(), // INICIALIZAR EL CLIENTE
	}

	// Configurar menús
	menuBar.SetupMenus(mainWin)
	window.SetMainMenu(menuBar.MainMenu)

	// Crear layout principal
	content := mainWin.createLayout()
	window.SetContent(content)

	return mainWin
}

func (mw *MainWindow) createLayout() *container.Split {
	// Panel superior: Editor y Reportes
	topPanel := container.NewHSplit(
		mw.Editor.Container,  // Izquierda: Editor
		mw.Reports.Container, // Derecha: Reportes
	)
	topPanel.SetOffset(0.5) // 70% para editor, 30% para reportes

	// Layout principal: Split vertical con consola en la parte inferior
	mainSplit := container.NewVSplit(
		topPanel,             // Arriba: Editor + Reportes
		mw.Console.Container, // Abajo: Consola (ancho completo)
	)
	mainSplit.SetOffset(0.6) // 75% para panel superior, 25% para consola

	return mainSplit
}

func (mw *MainWindow) ShowAndRun() {
	mw.Window.ShowAndRun()
}

// Métodos para manejar acciones del menú
func (mw *MainWindow) NewFile() {
	mw.Editor.NewFile()
}

func (mw *MainWindow) OpenFile() {
	mw.Editor.OpenFile(mw.Window)
}

func (mw *MainWindow) SaveFile() {
	mw.Editor.SaveFile(mw.Window)
}

func (mw *MainWindow) ExecuteCode() {
	code := mw.Editor.GetCode()

	// Verificar que hay código
	if strings.TrimSpace(code) == "" {
		dialog.ShowInformation("Advertencia", "No hay código para analizar", mw.Window)
		return
	}

	// Limpiar consola y reportes antes de ejecutar
	mw.Console.ClearOutput()
	mw.Reports.ClearAll()

	// Verificar conexión con el servidor
	mw.Console.AddOutput("Verificando conexión con el servidor...\n")
	if err := mw.APIClient.Ping(); err != nil {
		mw.Console.AddOutput("❌ Error: No se puede conectar con el servidor del compilador\n")
		mw.Console.AddOutput("💡 Asegúrate de que el servidor esté ejecutándose:\n")
		mw.Console.AddOutput("   cd docs && go run main.go\n")
		dialog.ShowError(fmt.Errorf("servidor no disponible: %v", err), mw.Window)
		return
	}

	// Mostrar mensaje de carga
	mw.Console.AddOutput("✅ Conexión establecida\n")
	mw.Console.AddOutput("🔍 Analizando código...\n")

	// Llamar al backend para análisis
	filename := mw.Editor.CurrentFile
	if filename == "" {
		filename = "untitled.vch"
	}

	result, err := mw.APIClient.AnalyzeCode(code, filename)
	if err != nil {
		mw.Console.AddOutput("❌ Error al analizar código: " + err.Error() + "\n")
		dialog.ShowError(err, mw.Window)
		return
	}

	// Mostrar resultados en consola
	mw.Console.ClearOutput()
	mw.Console.AddOutput(result.ConsoleOutput)

	if result.Success {
		mw.Console.AddOutput("\n✅ Análisis completado exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Errores encontrados: %d\n", len(result.Errors)))
		mw.Console.AddOutput(fmt.Sprintf("🏷️  Símbolos en tabla: %d\n", len(result.SymbolTable)))
	} else {
		mw.Console.AddOutput("\n❌ Se encontraron errores en el código\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Total de errores: %d\n", len(result.Errors)))
	}

	// Actualizar reportes
	mw.Reports.UpdateErrors(result.Errors)
	mw.Reports.UpdateSymbolTable(result.SymbolTable)
	mw.Reports.UpdateASTWithPNG(result.AST, result.ASTPng) // USAR PNG EN LUGAR DE SVG

	// Si hay PNG generado, mostrar información
	if len(result.ASTPng) > 0 {
		mw.Console.AddOutput("\n🌳 AST convertido a PNG exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📄 Tamaño del PNG: %d bytes\n", len(result.ASTPng)))
		mw.Console.AddOutput("💡 Visualiza el AST en la pestaña 'AST'\n")
	}

	// Mostrar pestaña de errores si hay errores
	if len(result.Errors) > 0 {
		mw.Reports.Container.SelectTab(mw.Reports.Container.Items[0]) // Pestaña de errores
	}
}
*/
/*
package ui

import (
	"fmt"
	"strings"
	"vlang-cherry-ide/api" // Importar el paquete api

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type MainWindow struct {
	App       fyne.App
	Window    fyne.Window
	Editor    *Editor
	Console   *Console
	Reports   *Reports
	MenuBar   *Menus
	APIClient *api.Client // AGREGAR ESTE CAMPO
}

func NewMainWindow(myApp fyne.App) *MainWindow {
	// Crear ventana principal
	window := myApp.NewWindow("V-Lang Cherry IDE")
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()

	// Crear componentes
	editor := NewEditor()
	console := NewConsole()
	reports := NewReports()
	menuBar := NewMenus()

	mainWin := &MainWindow{
		App:       myApp,
		Window:    window,
		Editor:    editor,
		Console:   console,
		Reports:   reports,
		MenuBar:   menuBar,
		APIClient: api.NewClient(), // INICIALIZAR EL CLIENTE
	}

	// Configurar menús
	menuBar.SetupMenus(mainWin)
	window.SetMainMenu(menuBar.MainMenu)

	// Crear layout principal
	content := mainWin.createLayout()
	window.SetContent(content)

	return mainWin
}

func (mw *MainWindow) createLayout() *container.Split {
	// Panel superior: Editor y Reportes
	topPanel := container.NewHSplit(
		mw.Editor.Container,  // Izquierda: Editor
		mw.Reports.Container, // Derecha: Reportes
	)
	topPanel.SetOffset(0.5) // 70% para editor, 30% para reportes

	// Layout principal: Split vertical con consola en la parte inferior
	mainSplit := container.NewVSplit(
		topPanel,             // Arriba: Editor + Reportes
		mw.Console.Container, // Abajo: Consola (ancho completo)
	)
	mainSplit.SetOffset(0.6) // 75% para panel superior, 25% para consola

	return mainSplit
}

func (mw *MainWindow) ShowAndRun() {
	mw.Window.ShowAndRun()
}

// Métodos para manejar acciones del menú
func (mw *MainWindow) NewFile() {
	mw.Editor.NewFile()
}

func (mw *MainWindow) OpenFile() {
	mw.Editor.OpenFile(mw.Window)
}

func (mw *MainWindow) SaveFile() {
	mw.Editor.SaveFile(mw.Window)
}

func (mw *MainWindow) ExecuteCode() {
	code := mw.Editor.GetCode()

	// Verificar que hay código
	if strings.TrimSpace(code) == "" {
		dialog.ShowInformation("Advertencia", "No hay código para analizar", mw.Window)
		return
	}

	// Limpiar consola y reportes antes de ejecutar
	mw.Console.ClearOutput()
	mw.Reports.ClearAll()

	// Verificar conexión con el servidor
	mw.Console.AddOutput("Verificando conexión con el servidor...\n")
	if err := mw.APIClient.Ping(); err != nil {
		mw.Console.AddOutput("❌ Error: No se puede conectar con el servidor del compilador\n")
		mw.Console.AddOutput("💡 Asegúrate de que el servidor esté ejecutándose:\n")
		mw.Console.AddOutput("   cd docs && go run main.go\n")
		dialog.ShowError(fmt.Errorf("servidor no disponible: %v", err), mw.Window)
		return
	}

	// Mostrar mensaje de carga
	mw.Console.AddOutput("✅ Conexión establecida\n")
	mw.Console.AddOutput("🔍 Analizando código...\n")

	// Llamar al backend para análisis
	filename := mw.Editor.CurrentFile
	if filename == "" {
		filename = "untitled.vch"
	}

	result, err := mw.APIClient.AnalyzeCode(code, filename)
	if err != nil {
		mw.Console.AddOutput("❌ Error al analizar código: " + err.Error() + "\n")
		dialog.ShowError(err, mw.Window)
		return
	}

	// Mostrar resultados en consola
	mw.Console.ClearOutput()
	mw.Console.AddOutput(result.ConsoleOutput)

	if result.Success {
		mw.Console.AddOutput("\n✅ Análisis completado exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Errores encontrados: %d\n", len(result.Errors)))
		mw.Console.AddOutput(fmt.Sprintf("🏷️  Símbolos en tabla: %d\n", len(result.SymbolTable)))
	} else {
		mw.Console.AddOutput("\n❌ Se encontraron errores en el código\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Total de errores: %d\n", len(result.Errors)))
	}

	// Actualizar reportes
	mw.Reports.UpdateErrors(result.Errors)
	mw.Reports.UpdateSymbolTable(result.SymbolTable)
	mw.Reports.UpdateASTWithPNG(result.AST, result.ASTPng) // USAR PNG EN LUGAR DE SVG

	// Si hay PNG generado, mostrar información y ruta del archivo
	if len(result.ASTPng) > 0 {
		mw.Console.AddOutput("\n🌳 AST convertido a PNG exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📄 Tamaño del PNG: %d bytes\n", len(result.ASTPng)))
		mw.Console.AddOutput("💡 Visualiza el AST en la pestaña 'AST'\n")

		// Mostrar información del archivo guardado si está disponible
		if strings.Contains(result.AST, "Archivo:") {
			mw.Console.AddOutput("💾 " + result.AST + "\n")
		}
	}

	// Mostrar pestaña de errores si hay errores
	if len(result.Errors) > 0 {
		mw.Reports.Container.SelectTab(mw.Reports.Container.Items[0]) // Pestaña de errores
	}
}*/
package ui

import (
	"fmt"
	instrucciones "main/Instrucciones"
	"main/api" // Importar el paquete api
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type MainWindow struct {
	App       fyne.App
	Window    fyne.Window
	Editor    *Editor
	Console   *Console
	Reports   *Reports
	MenuBar   *Menus
	APIClient *api.Client
}

func NewMainWindow(myApp fyne.App) *MainWindow {
	// Crear ventana principal
	window := myApp.NewWindow("V-Lang Cherry IDE")
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()

	// Crear componentes
	editor := NewEditor()
	console := NewConsole()
	reports := NewReports() // ✅ SIN PARÁMETROS
	menuBar := NewMenus()

	mainWin := &MainWindow{
		App:       myApp,
		Window:    window,
		Editor:    editor,
		Console:   console,
		Reports:   reports,
		MenuBar:   menuBar,
		APIClient: api.NewClient(),
	}

	// Configurar menús
	menuBar.SetupMenus(mainWin)
	window.SetMainMenu(menuBar.MainMenu)

	// Crear layout principal
	content := mainWin.createLayout()
	window.SetContent(content)

	return mainWin
}

// ...existing code...
// [El resto del código permanece igual]

func (mw *MainWindow) createLayout() *container.Split {
	// Panel superior: Editor y Reportes
	topPanel := container.NewHSplit(
		mw.Editor.Container,  // Izquierda: Editor
		mw.Reports.Container, // Derecha: Reportes
	)
	topPanel.SetOffset(0.5) // 70% para editor, 30% para reportes

	// Layout principal: Split vertical con consola en la parte inferior
	mainSplit := container.NewVSplit(
		topPanel,             // Arriba: Editor + Reportes
		mw.Console.Container, // Abajo: Consola (ancho completo)
	)
	mainSplit.SetOffset(0.6) // 75% para panel superior, 25% para consola

	return mainSplit
}

func (mw *MainWindow) ShowAndRun() {
	mw.Window.ShowAndRun()
}

// Métodos para manejar acciones del menú
func (mw *MainWindow) NewFile() {
	mw.Editor.NewFile()
}

func (mw *MainWindow) OpenFile() {
	mw.Editor.OpenFile(mw.Window)
}

func (mw *MainWindow) SaveFile() {
	mw.Editor.SaveFile(mw.Window)
}

func (mw *MainWindow) ExecuteCode() {
	instrucciones.ResetAllRecursionFlags()
	code := mw.Editor.GetCode()

	// Verificar que hay código
	if strings.TrimSpace(code) == "" {
		dialog.ShowInformation("Advertencia", "No hay código para analizar", mw.Window)
		return
	}

	// Limpiar consola y reportes antes de ejecutar
	mw.Console.ClearOutput()
	mw.Reports.ClearAll()

	// Verificar conexión con el servidor
	mw.Console.AddOutput("Verificando conexión con el servidor...\n")
	if err := mw.APIClient.Ping(); err != nil {
		mw.Console.AddOutput("❌ Error: No se puede conectar con el servidor del compilador\n")
		mw.Console.AddOutput("💡 Asegúrate de que el servidor esté ejecutándose:\n")
		mw.Console.AddOutput("   cd docs && go run main.go\n")
		dialog.ShowError(fmt.Errorf("servidor no disponible: %v", err), mw.Window)
		return
	}

	// Mostrar mensaje de carga
	mw.Console.AddOutput("✅ Conexión establecida\n")
	mw.Console.AddOutput("🔍 Analizando código...\n")

	// Llamar al backend para análisis
	filename := mw.Editor.CurrentFile
	if filename == "" {
		filename = "untitled.vch"
	}

	result, err := mw.APIClient.AnalyzeCode(code, filename)
	if err != nil {
		mw.Console.AddOutput("❌ Error al analizar código: " + err.Error() + "\n")
		dialog.ShowError(err, mw.Window)
		return
	}

	// Mostrar resultados en consola
	mw.Console.ClearOutput()
	mw.Console.AddOutput(result.ConsoleOutput)

	if result.Success {
		mw.Console.AddOutput("\n✅ Análisis completado exitosamente\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Errores encontrados: %d\n", len(result.Errors)))
		mw.Console.AddOutput(fmt.Sprintf("🏷️  Símbolos en tabla: %d\n", len(result.SymbolTable)))
	} else {
		mw.Console.AddOutput("\n❌ Se encontraron errores en el código\n")
		mw.Console.AddOutput(fmt.Sprintf("📊 Total de errores: %d\n", len(result.Errors)))
	}

	// Actualizar reportes con SVG nativo integrado
	mw.Reports.UpdateErrors(result.Errors)
	mw.Reports.UpdateSymbolTable(result.SymbolTable)
	mw.Reports.UpdateASTWithSVG(result.AST, result.CSTSvg) // SVG NATIVO EN EL IDE

	// Si hay SVG, mostrar información simple
	if result.CSTSvg != "" {
		mw.Console.AddOutput("\n🌳 AST renderizado en la pestaña AST\n")
		mw.Console.AddOutput("💡 Ve a la pestaña 'AST' para ver el árbol\n")
	}

	// Mostrar pestaña de errores si hay errores
	if len(result.Errors) > 0 {
		mw.Reports.Container.SelectTab(mw.Reports.Container.Items[0]) // Pestaña de errores
	}
}
