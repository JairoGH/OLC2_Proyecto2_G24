package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Menus struct {
	MainMenu *fyne.MainMenu
}

func NewMenus() *Menus {
	return &Menus{}
}

func (m *Menus) SetupMenus(mainWindow *MainWindow) {
	// Menú Archivo
	fileMenu := fyne.NewMenu("Archivo",
		fyne.NewMenuItem("Nuevo", func() {
			mainWindow.NewFile()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Abrir", func() {
			mainWindow.OpenFile()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Guardar", func() {
			mainWindow.SaveFile()
		}),
		fyne.NewMenuItem("Guardar como...", func() {
			mainWindow.Editor.SaveFileAs(mainWindow.Window)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Salir", func() {
			mainWindow.Window.Close()
		}),
	)

	// Menú Herramientas
	toolsMenu := fyne.NewMenu("Herramientas",
		fyne.NewMenuItem("Ejecutar", func() {
			mainWindow.ExecuteCode()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Limpiar Consola", func() {
			mainWindow.Console.ClearOutput()
		}),
	)

	// Menú Ayuda
	helpMenu := fyne.NewMenu("Ayuda",
		fyne.NewMenuItem("Acerca de", func() {
			dialog.ShowInformation("V-Lang Cherry IDE",
				"IDE para el lenguaje V-Lang Cherry\nVersionn 1.0",
				mainWindow.Window)
		}),
	)

	m.MainMenu = fyne.NewMainMenu(fileMenu, toolsMenu, helpMenu)
}
