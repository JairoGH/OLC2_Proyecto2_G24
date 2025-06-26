package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Menus struct {
	MainMenu   *fyne.MainMenu
	mainWindow *MainWindow
}

func NewMenus(mainWindow *MainWindow) *Menus {
	menuBar := &Menus{
		mainWindow: mainWindow,
	}
	menuBar.createMenus()
	return menuBar
}

func (m *Menus) createMenus() {
	// Menú Archivo (mantener igual)
	fileMenu := fyne.NewMenu("Archivo",
		fyne.NewMenuItem("Nuevo", func() {
			m.mainWindow.NewFile()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Abrir", func() {
			m.mainWindow.OpenFile()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Guardar", func() {
			m.mainWindow.SaveFile()
		}),
		fyne.NewMenuItem("Guardar como...", func() {
			m.mainWindow.Editor.SaveFileAs(m.mainWindow.Window)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Salir", func() {
			m.mainWindow.Window.Close()
		}),
	)

	// 🔥 MENÚ HERRAMIENTAS LIMPIO 
	toolsMenu := fyne.NewMenu("Herramientas",
		fyne.NewMenuItem("🔍 Ejecutar Análisis", func() {
			m.mainWindow.ExecuteAnalysis()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("⚙️ Generar ARM64", func() {
			m.mainWindow.GenerateARM64()
		}),
		fyne.NewMenuItem("🌳 Generar AST", func() {
			m.mainWindow.GenerateAST()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("🧹 Limpiar Consola", func() {
			m.mainWindow.Console.ClearOutput()
		}),
	)

	// Menú Ayuda 
	helpMenu := fyne.NewMenu("Ayuda",
		fyne.NewMenuItem("Acerca de", func() {
			dialog.ShowInformation("V-Lang Cherry IDE",
				"IDE para el lenguaje V-Lang Cherry\nVersión 2.0 - Optimizada",
				m.mainWindow.Window)
		}),
	)

	m.MainMenu = fyne.NewMainMenu(
		fileMenu,
		toolsMenu,
		helpMenu,
	)
}
