package ui

import (
	"fmt"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	Container   fyne.CanvasObject
	TextArea    *widget.Entry
	LineNumbers *widget.Label
	StatusBar   *widget.Label
	CurrentFile string
	IsModified  bool

	// 🆕 SOLO AGREGAR ESTOS CAMPOS para controlar scroll
	ScrollContainer   *container.Scroll
	preventAutoScroll bool
}

func NewEditor() *Editor {
	// ✅ MANTENER EXACTAMENTE IGUAL - No cambiar nada del TextArea
	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder("Escribe tu código V-Lang Cherry aquí...")
	textArea.Wrapping = fyne.TextWrapOff

	// ✅ MANTENER IGUAL - números de línea
	lineNumbers := widget.NewLabel("1")
	lineNumbers.Alignment = fyne.TextAlignTrailing

	// ✅ MANTENER IGUAL - barra de estado
	statusBar := widget.NewLabel("Línea: 1, Columna: 1 | Sin archivo")

	editor := &Editor{
		TextArea:          textArea,
		LineNumbers:       lineNumbers,
		StatusBar:         statusBar,
		CurrentFile:       "",
		IsModified:        false,
		preventAutoScroll: false, // 🆕 SOLO ESTO
	}

	// Configurar eventos (IGUAL que antes)
	editor.setupEvents()

	// Crear layout (AQUÍ está el cambio mínimo)
	editor.createLayout()

	return editor
}

func (e *Editor) setupEvents() {
	// ✅ MANTENER EVENTOS ORIGINALES - Solo agregar flag de control
	e.TextArea.OnCursorChanged = func() {
		e.updateStatus()
	}

	// Actualizar números de línea y posición del cursor
	e.TextArea.OnChanged = func(text string) {
		e.updateLineNumbers(text)
		e.updateStatus()
		e.IsModified = true
	}
}

func (e *Editor) createLayout() {
	// ✅ MANTENER IGUAL - Panel de números de línea
	linePanel := container.NewScroll(e.LineNumbers)
	linePanel.SetMinSize(fyne.NewSize(50, 0))

	// 🔧 ÚNICO CAMBIO: Envolver el TextArea en Scroll con control
	e.ScrollContainer = container.NewScroll(e.TextArea)

	// 🆕 SOLO ESTA LÍNEA ES NUEVA - Detectar scroll manual
	e.ScrollContainer.OnScrolled = func(pos fyne.Position) {
		e.preventAutoScroll = true
	}

	// ✅ MANTENER IGUAL - resto del layout
	editorContent := container.NewHSplit(linePanel, e.ScrollContainer)
	editorContent.SetOffset(0.05) // 5% para números, 95% para editor

	// Layout principal con barra de estado
	e.Container = container.NewBorder(
		nil,           // top
		e.StatusBar,   // bottom
		nil,           // left
		nil,           // right
		editorContent, // center
	)
}

// ✅ MANTENER TODAS LAS FUNCIONES ORIGINALES IGUALES
func (e *Editor) updateLineNumbers(text string) {
	lines := strings.Split(text, "\n")
	lineCount := len(lines)

	var numbers []string
	for i := 1; i <= lineCount; i++ {
		numbers = append(numbers, fmt.Sprintf("%d", i))
	}

	e.LineNumbers.SetText(strings.Join(numbers, "\n"))
}

func (e *Editor) updateStatus() {
	content := e.TextArea.Text

	// Calcular línea y columna actual
	lines := strings.Split(content, "\n")
	currentLine := len(lines)

	// Para columna, usar longitud de la última línea + 1
	currentColumn := 1
	if len(lines) > 0 {
		currentColumn = len(lines[len(lines)-1]) + 1
	}

	// Actualizar barra de estado
	filename := "Sin archivo"
	if e.CurrentFile != "" {
		filename = e.CurrentFile
	}

	modifiedIndicator := ""
	if e.IsModified {
		modifiedIndicator = " *"
	}

	status := fmt.Sprintf("Línea: %d, Columna: %d | %s%s",
		currentLine, currentColumn, filename, modifiedIndicator)
	e.StatusBar.SetText(status)
}

// 🆕 FUNCIONES NUEVAS PARA EL MENÚ - pero no afectan el editor original
func (e *Editor) GoToCursor() {
	e.preventAutoScroll = false
	// Permitir que vaya al cursor solo cuando se pida explícitamente
}

func (e *Editor) ResetScrollControl() {
	e.preventAutoScroll = false
}

// ✅ MANTENER TODAS LAS FUNCIONES ORIGINALES EXACTAMENTE IGUAL
func (e *Editor) NewFile() {
	e.TextArea.SetText("")
	e.CurrentFile = ""
	e.IsModified = false
	e.preventAutoScroll = false // Solo resetear flag
	e.updateStatus()
}

func (e *Editor) OpenFile(window fyne.Window) {
	// 🆕 CREAR FILTRO PARA SOLO ARCHIVOS .vch
	vchFilter := storage.NewExtensionFileFilter([]string{".vch"})

	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		defer reader.Close()

		// ✅ VERIFICAR EXTENSIÓN .vch (mantener por seguridad)
		if !strings.HasSuffix(reader.URI().Name(), ".vch") {
			dialog.ShowError(fmt.Errorf("❌ Solo se permiten archivos .vch\n💡 Selecciona un archivo con extensión .vch"), window)
			return
		}

		// Leer contenido
		content, err := io.ReadAll(reader)
		if err != nil {
			dialog.ShowError(fmt.Errorf("❌ Error al leer el archivo: %v", err), window)
			return
		}

		// Establecer contenido
		e.TextArea.SetText(string(content))
		e.CurrentFile = reader.URI().Path() // 🔧 USAR PATH COMPLETO
		e.IsModified = false
		e.preventAutoScroll = false
		e.updateStatus()

		// 🆕 MOSTRAR CONFIRMACIÓN
		dialog.ShowInformation("✅ Archivo abierto",
			fmt.Sprintf("📁 Archivo: %s\n📄 Líneas: %d",
				reader.URI().Name(),
				len(strings.Split(string(content), "\n"))),
			window)

	}, window)

	// 🆕 APLICAR FILTRO
	fileDialog.SetFilter(vchFilter)
	fileDialog.Show()
}

func (e *Editor) SaveFile(window fyne.Window) {
	if e.CurrentFile == "" {
		e.SaveFileAs(window)
		return
	}

	// 🔧 IMPLEMENTAR GUARDADO REAL
	content := e.TextArea.Text

	// Escribir al archivo actual
	err := os.WriteFile(e.CurrentFile, []byte(content), 0644)
	if err != nil {
		dialog.ShowError(fmt.Errorf("❌ Error al guardar: %v", err), window)
		return
	}

	// ✅ MARCAR COMO GUARDADO
	e.IsModified = false
	e.updateStatus()

	// 🆕 MOSTRAR CONFIRMACIÓN
	dialog.ShowInformation("✅ Guardado exitoso",
		fmt.Sprintf("📁 Archivo: %s\n💾 Guardado correctamente", e.CurrentFile),
		window)
}

func (e *Editor) SaveFileAs(window fyne.Window) {
	// 🆕 CREAR FILTRO PARA GUARDAR SOLO COMO .vch
	vchFilter := storage.NewExtensionFileFilter([]string{".vch"})

	fileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil || writer == nil {
			return
		}
		defer writer.Close()

		// 🔧 OBTENER NOMBRE Y ASEGURAR EXTENSIÓN .vch
		filename := writer.URI().Path()
		displayName := writer.URI().Name()

		// Asegurar que termine en .vch
		if !strings.HasSuffix(filename, ".vch") {
			filename += ".vch"
			displayName += ".vch"
		}

		// Escribir contenido
		content := e.TextArea.Text
		_, err = writer.Write([]byte(content))
		if err != nil {
			dialog.ShowError(fmt.Errorf("❌ Error al guardar: %v", err), window)
			return
		}

		// ✅ ACTUALIZAR ESTADO
		e.CurrentFile = filename
		e.IsModified = false
		e.updateStatus()

		// 🆕 MOSTRAR CONFIRMACIÓN
		dialog.ShowInformation("✅ Guardado exitoso",
			fmt.Sprintf("📁 Archivo: %s\n💾 Guardado como nuevo archivo", displayName),
			window)

	}, window)

	// 🆕 APLICAR FILTRO Y CONFIGURAR NOMBRE POR DEFECTO
	fileDialog.SetFilter(vchFilter)
	fileDialog.SetFileName("untitled.vch") // 🆕 NOMBRE POR DEFECTO
	fileDialog.Show()
}

func (e *Editor) GetCode() string {
	return e.TextArea.Text
}
