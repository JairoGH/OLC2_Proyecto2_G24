package ui

import (
	"fmt"
	"main/models"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Reports struct {
	Container    *container.AppTabs
	ErrorList    *widget.Table
	SymbolTable  *widget.Table
	ASTContent   *fyne.Container
	ARM64Content *fyne.Container

	// PANELES PARA ERRORES Y SÍMBOLOS
	ErrorDetails  *widget.Card
	ErrorPanel    *container.Split
	SymbolDetails *widget.Card
	SymbolPanel   *container.Split

	// Datos
	errors      []models.ErrorReport
	symbolTable []models.SymbolEntry
}

func NewReports() *Reports {
	reports := &Reports{
		errors:      []models.ErrorReport{},
		symbolTable: []models.SymbolEntry{},
	}

	// ========== CONFIGURACIÓN ERRORES ==========

	// CREAR PANEL DE DETALLES PARA ERRORES (SIMPLE Y PEQUEÑO)
	reports.ErrorDetails = widget.NewCard(
		"Mensaje Completo",
		"",
		widget.NewLabel("Selecciona un error para ver el mensaje completo"),
	)
	reports.ErrorDetails.Hide() //  SIEMPRE OCULTO AL INICIO

	// Crear tabla de errores
	reports.ErrorList = widget.NewTable(
		func() (int, int) {
			return len(reports.errors) + 1, 5
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("Template")
			return label
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			if id.Row == 0 {
				// Encabezados
				headers := []string{"No.", "Descripción", "Línea", "Columna", "Tipo"}
				if id.Col < len(headers) {
					label.SetText(headers[id.Col])
					label.TextStyle = fyne.TextStyle{Bold: true}
					label.Alignment = fyne.TextAlignCenter
				}
			} else {
				// Datos
				if id.Row-1 < len(reports.errors) {
					error := reports.errors[id.Row-1]
					var text string

					switch id.Col {
					case 0: // No.
						text = strconv.Itoa(id.Row)
						label.Alignment = fyne.TextAlignCenter
					case 1: // Descripción
						text = reports.truncateText(error.Descripcion, 44)
						label.Alignment = fyne.TextAlignLeading
						if len(error.Descripcion) > 44 {
							text = text + " 🔍"
						}
					case 2: // Línea
						text = strconv.Itoa(error.Linea)
						label.Alignment = fyne.TextAlignCenter
					case 3: // Columna
						text = strconv.Itoa(error.Columna)
						label.Alignment = fyne.TextAlignCenter
					case 4: // Tipo
						text = reports.truncateText(error.Tipo, 22)
						label.Alignment = fyne.TextAlignCenter
						if len(error.Tipo) > 22 {
							text = text + " ..."
						}
					}

					label.SetText(text)
					label.TextStyle = fyne.TextStyle{Bold: false}
				}
			}
		},
	)

	// EVENTO DE SELECCIÓN PARA ERRORES
	reports.ErrorList.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 && id.Row-1 < len(reports.errors) {
			reports.showErrorMessage(reports.errors[id.Row-1])
		}
	}

	// Configurar anchos de tabla de errores
	reports.ErrorList.SetColumnWidth(0, 35)  // No.
	reports.ErrorList.SetColumnWidth(1, 300) // Descripción
	reports.ErrorList.SetColumnWidth(2, 60)  // Línea
	reports.ErrorList.SetColumnWidth(3, 70)  // Columna
	reports.ErrorList.SetColumnWidth(4, 150) // Tipo

	// CREAR LAYOUT PARA ERRORES
	reports.ErrorPanel = container.NewVSplit(
		reports.ErrorList,
		reports.ErrorDetails,
	)
	reports.ErrorPanel.SetOffset(1.0) //  INICIAR COMPLETAMENTE OCULTO

	// ========== CONFIGURACIÓN TABLA DE SÍMBOLOS ==========

	// CREAR PANEL DE DETALLES PARA SÍMBOLOS
	reports.SymbolDetails = widget.NewCard(
		"Detalles del Símbolo",
		"",
		widget.NewLabel("Selecciona un símbolo para ver los detalles completos"),
	)
	reports.SymbolDetails.Hide() //  SIEMPRE OCULTO AL INICIO

	// Crear tabla de símbolos
	reports.SymbolTable = widget.NewTable(
		func() (int, int) {
			return len(reports.symbolTable) + 1, 6
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row == 0 {
				headers := []string{"ID", "Tipo", "Tipo de Dato", "Ámbito", "Línea", "Columna"}
				if id.Col < len(headers) {
					label.SetText(headers[id.Col])
					label.TextStyle = fyne.TextStyle{Bold: true}
					label.Alignment = fyne.TextAlignCenter
				}
			} else {
				if id.Row-1 < len(reports.symbolTable) {
					symbol := reports.symbolTable[id.Row-1]
					var text string

					switch id.Col {
					case 0: // ID
						text = reports.truncateText(symbol.ID, 20)
						label.Alignment = fyne.TextAlignLeading
						if len(symbol.ID) > 20 {
							text = text + " 🔍" // INDICADOR PARA ID
						}
					case 1: // Tipo
						text = reports.truncateText(symbol.SymbolType, 12)
						label.Alignment = fyne.TextAlignCenter
						if len(symbol.SymbolType) > 12 {
							text = text + " ..." // INDICADOR PARA TIPO
						}
					case 2: // Tipo de Dato
						text = reports.truncateText(symbol.DataType, 20)
						label.Alignment = fyne.TextAlignCenter
						if len(symbol.DataType) > 20 {
							text = text + " ..." // INDICADOR PARA TIPO DE DATO
						}
					case 3: // Ámbito
						text = reports.truncateText(symbol.Scope, 20)
						label.Alignment = fyne.TextAlignCenter
					case 4: // Línea
						text = strconv.Itoa(symbol.Line)
						label.Alignment = fyne.TextAlignCenter
					case 5: // Columna
						text = strconv.Itoa(symbol.Column)
						label.Alignment = fyne.TextAlignCenter
					}

					label.SetText(text)
					label.TextStyle = fyne.TextStyle{Bold: false} // 🔧 CORREGIDO
				}
			}
		},
	)

	// EVENTO DE SELECCIÓN PARA SÍMBOLOS
	reports.SymbolTable.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 && id.Row-1 < len(reports.symbolTable) {
			reports.showSymbolMessage(reports.symbolTable[id.Row-1])
		}
	}

	// Configurar anchos de tabla de símbolos
	reports.SymbolTable.SetColumnWidth(0, 140) // ID
	reports.SymbolTable.SetColumnWidth(1, 120) // Tipo
	reports.SymbolTable.SetColumnWidth(2, 95)  // Tipo de Dato
	reports.SymbolTable.SetColumnWidth(3, 130) // Ámbito
	reports.SymbolTable.SetColumnWidth(4, 60)  // Línea
	reports.SymbolTable.SetColumnWidth(5, 70)  // Columna

	// CREAR LAYOUT PARA SÍMBOLOS
	reports.SymbolPanel = container.NewVSplit(
		reports.SymbolTable,
		reports.SymbolDetails,
	)
	reports.SymbolPanel.SetOffset(1.0) //  INICIAR COMPLETAMENTE OCULTO

	// ========== CONFIGURACIÓN AST ==========

	// Crear contenedor para AST
	reports.ASTContent = container.NewVBox(
		widget.NewLabel("No hay AST disponible"),
	)

	// PESTAÑAS CON PANELES
	reports.Container = container.NewAppTabs(
		container.NewTabItem("Errores", reports.ErrorPanel),            // Panel con mensaje para errores
		container.NewTabItem("Tabla de Símbolos", reports.SymbolPanel), // Panel con detalles para símbolos
		container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
	)

	// ========== CONFIGURACIÓN ARM64 ==========
	reports.ARM64Content = container.NewVBox(
		widget.NewLabel("No hay código ARM64 generado"),
	)

	// PESTAÑAS CON ARM64 INCLUIDO
	reports.Container = container.NewAppTabs(
		container.NewTabItem("Errores", reports.ErrorPanel),
		container.NewTabItem("Tabla de Símbolos", reports.SymbolPanel),
		container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
		container.NewTabItem("ARM64", container.NewScroll(reports.ARM64Content)), // NUEVA PESTAÑA
	)

	return reports
}

// FUNCIÓN PARA MOSTRAR MENSAJE DE ERROR (MEJORADA)
func (r *Reports) showErrorMessage(error models.ErrorReport) {
	//  FORMATO MEJORADO COMO PEDISTE
	messageText := fmt.Sprintf("📝 Mensaje Completo - Línea %d          Tipo: %s\n\n%s",
		error.Linea, error.Tipo, error.Descripcion)

	messageLabel := widget.NewLabel(messageText)
	messageLabel.Wrapping = fyne.TextWrapWord
	messageLabel.TextStyle = fyne.TextStyle{Monospace: false}

	// Botón pequeño para cerrar
	closeBtn := widget.NewButton("Cerrar", func() {
		r.ErrorDetails.Hide()
		r.ErrorPanel.SetOffset(1.0) // Volver a ocultar completamente
	})

	// Contenido simple
	content := container.NewVBox(
		messageLabel,
		closeBtn,
	)

	// Configurar card
	r.ErrorDetails.SetTitle("🚨 Error Completo")
	r.ErrorDetails.SetSubTitle(fmt.Sprintf("Línea %d • Columna %d", error.Linea, error.Columna))
	r.ErrorDetails.SetContent(content)
	r.ErrorDetails.Show()
	r.ErrorPanel.SetOffset(0.85) // 85% tabla, 15% mensaje (PEQUEÑO)
}

// FUNCIÓN PARA MOSTRAR DETALLES DE SÍMBOLO
func (r *Reports) showSymbolMessage(symbol models.SymbolEntry) {
	//  FORMATO SIMILAR AL DE ERRORES
	detailsText := fmt.Sprintf(`🔍 Detalles del Símbolo - Línea %d

ID: %s
Tipo: %s  
Tipo de Dato: %s

Ubicación:
• Ámbito: %s
• Línea: %d
• Columna: %d

💡 Información adicional:
Este símbolo fue declarado en el ámbito '%s' y es de tipo '%s'`,
		symbol.Line,
		symbol.ID,
		symbol.SymbolType,
		symbol.DataType,
		symbol.Scope,
		symbol.Line,
		symbol.Column,
		symbol.Scope,
		symbol.SymbolType)

	detailsLabel := widget.NewLabel(detailsText)
	detailsLabel.Wrapping = fyne.TextWrapWord
	detailsLabel.TextStyle = fyne.TextStyle{Monospace: false}

	// Botón para cerrar
	closeBtn := widget.NewButton("Cerrar", func() {
		r.SymbolDetails.Hide()
		r.SymbolPanel.SetOffset(1.0) // Volver a ocultar completamente
	})

	// Contenido
	content := container.NewVBox(
		detailsLabel,
		closeBtn,
	)

	// Configurar card
	r.SymbolDetails.SetTitle("📋 Símbolo Completo")
	r.SymbolDetails.SetSubTitle(fmt.Sprintf("%s • %s", symbol.ID, symbol.SymbolType))
	r.SymbolDetails.SetContent(content)
	r.SymbolDetails.Show()
	r.SymbolPanel.SetOffset(0.85) // 85% tabla, 15% detalles (PEQUEÑO)
}

// 🔧 FUNCIÓN TRUNCATE (mantener igual)
func (r *Reports) truncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}

	if maxLength <= 3 {
		return ""
	}

	return text[:maxLength-3] + "..."
}

// FUNCIÓN PARA OCULTAR TODOS LOS PANELES AL CAMBIAR DE PESTAÑA
func (r *Reports) HideAllPanels() {
	//  OCULTAR PANEL DE ERRORES
	r.ErrorDetails.Hide()
	r.ErrorPanel.SetOffset(1.0)

	//  OCULTAR PANEL DE SÍMBOLOS
	r.SymbolDetails.Hide()
	r.SymbolPanel.SetOffset(1.0)
}

// UpdateErrors actualiza la lista de errores
func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
	r.errors = errors
	r.ErrorList.Refresh()

	//  ASEGURAR QUE ESTÉ OCULTO AL ACTUALIZAR
	r.ErrorDetails.Hide()
	r.ErrorPanel.SetOffset(1.0)
}

// UpdateSymbolTable actualiza la tabla de símbolos
func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
	r.symbolTable = symbols
	r.SymbolTable.Refresh()

	//  ASEGURAR QUE ESTÉ OCULTO AL ACTUALIZAR
	r.SymbolDetails.Hide()
	r.SymbolPanel.SetOffset(1.0)
}

// 🔧 MÉTODO SIMPLIFICADO - Solo SVG
func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
	if svgContent != "" {
		svgPath, err := r.saveSVGToFile(svgContent)
		if err != nil {
			r.showError("Error guardando SVG: " + err.Error())
			return
		}
		r.showSVGOnly(svgPath, svgContent)
	} else {
		r.showError("No hay AST disponible")
	}
}

// Mostrar solo SVG sin conversión
func (r *Reports) showSVGOnly(svgPath, svgContent string) {
	//  INFORMACIÓN CONCISA Y CLARA
	infoText := fmt.Sprintf(`🌳 Árbol AST Generado

📁 %s
📊 %d KB
✅ Listo para visualizar`,
		filepath.Base(svgPath),
		len(svgContent)/1024)

	infoLabel := widget.NewLabel(infoText)
	infoLabel.Wrapping = fyne.TextWrapWord
	infoLabel.TextStyle = fyne.TextStyle{Bold: false}

	// 🔥 BOTÓN PRINCIPAL - MÁS GRANDE Y DESTACADO
	openSVGBtn := widget.NewButton("🌐 Abrir SVG", func() {
		r.openFileInBrowser(svgPath)
	})
	openSVGBtn.Importance = widget.HighImportance

	//  INFORMACIÓN TÉCNICA SIMPLE Y LEGIBLE
	techInfo := r.getSimpleSVGInfo(svgContent)
	techLabel := widget.NewLabel(techInfo)
	techLabel.TextStyle = fyne.TextStyle{Monospace: true, Bold: false}
	techLabel.Wrapping = fyne.TextWrapWord

	//  LAYOUT LIMPIO Y ESPACIADO
	content := container.NewVBox(
		infoLabel,
		widget.NewSeparator(),
		openSVGBtn,
		widget.NewSeparator(),
		widget.NewLabel("📊 Información técnica:"),
		techLabel,
	)

	r.ASTContent.Objects = []fyne.CanvasObject{content}
	r.ASTContent.Refresh()
}

// Información técnica clara y concisa
func (r *Reports) getSimpleSVGInfo(svgContent string) string {
	//  VERIFICAR ELEMENTOS PRINCIPALES
	elements := map[string]string{
		"<g":      "Grupos",
		"<text":   "Textos",
		"<path":   "Líneas",
		"<rect":   "Cajas",
		"<circle": "Círculos",
	}

	found := []string{}
	for element, name := range elements {
		if strings.Contains(svgContent, element) {
			found = append(found, name)
		}
	}

	//  FORMATO LIMPIO Y LEGIBLE
	result := fmt.Sprintf(`Tamaño: %d caracteres
Elementos: %s
Estado: ✅ Válido para navegador`,
		len(svgContent),
		strings.Join(found, ", "))

	return result
}

// Guardar SVG (método simplificado)
func (r *Reports) saveSVGToFile(svgContent string) (string, error) {
	// Crear directorio
	outputDir := "generated_asts"
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return "", fmt.Errorf("error creando directorio: %v", err)
	}

	// Generar nombre único
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("ast_%s.svg", timestamp)
	fullPath := filepath.Join(outputDir, filename)

	// Escribir archivo SVG
	err = os.WriteFile(fullPath, []byte(svgContent), 0644)
	if err != nil {
		return "", fmt.Errorf("error escribiendo archivo SVG: %v", err)
	}

	// Obtener ruta absoluta
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return fullPath, nil
	}

	return absPath, nil
}

// 🔧 MANTENER: Método para abrir en navegador
func (r *Reports) openFileInBrowser(filePath string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	default:
		r.showError("Sistema operativo no soportado")
		return
	}

	if err := cmd.Start(); err != nil {
		r.showError("Error abriendo archivo: " + err.Error())
	}
}

// 🔧 MANTENER: Método para mostrar errores
func (r *Reports) showError(message string) {
	label := widget.NewLabel("❌ " + message)
	label.Alignment = fyne.TextAlignCenter
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// Limpiar todos los reportes
func (r *Reports) ClearAll() {
	r.ClearErrors()
	r.ClearSymbolTable()
	r.ClearAST()
	r.ClearARM64() // LIMPIAR ARM64 TAMBIÉN
}

// Limpiar errores
func (r *Reports) ClearErrors() {
	r.errors = []models.ErrorReport{}
	r.ErrorList.Refresh()

	// Ocultar detalles también
	r.ErrorDetails.Hide()
	r.ErrorPanel.SetOffset(1.0)
}

// Limpiar tabla de símbolos
func (r *Reports) ClearSymbolTable() {
	r.symbolTable = []models.SymbolEntry{}
	r.SymbolTable.Refresh()

	// Ocultar detalles también
	r.SymbolDetails.Hide()
	r.SymbolPanel.SetOffset(1.0)
}

// Limpiar AST
func (r *Reports) ClearAST() {
	// Limpiar el contenido del AST
	label := widget.NewLabel("No hay AST disponible")
	label.Alignment = fyne.TextAlignCenter

	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// LIMPIAR ARM64
func (r *Reports) ClearARM64() {
	label := widget.NewLabel("No hay código ARM64 generado")
	label.Alignment = fyne.TextAlignCenter

	r.ARM64Content.Objects = []fyne.CanvasObject{label}
	r.ARM64Content.Refresh()
}

// 🔧 VERSIÓN CORREGIDA: Abrir archivo en editor de texto
func (r *Reports) openFileInTextEditor(filePath string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		//  UBUNTU: Probar editores en orden de preferencia
		editors := []string{
			"gedit",    // Editor gráfico principal de Ubuntu
			"code",     // VS Code
			"mousepad", // Editor simple
			"nano",     // Terminal
			"vim",      // Terminal
		}

		var foundEditor string
		for _, editor := range editors {
			if _, err := exec.LookPath(editor); err == nil {
				foundEditor = editor
				break
			}
		}

		if foundEditor != "" {
			if foundEditor == "nano" || foundEditor == "vim" {
				// 🔧 TERMINAL: Usar gnome-terminal que está en Ubuntu
				cmd = exec.Command("gnome-terminal", "--", foundEditor, filePath)
			} else {
				// 🔧 GRÁFICO: Abrir directamente
				cmd = exec.Command(foundEditor, filePath)
			}
		} else {
			// 🔧 FALLBACK: xdg-open (siempre funciona)
			cmd = exec.Command("xdg-open", filePath)
		}

	case "darwin":
		cmd = exec.Command("open", "-a", "TextEdit", filePath)
	case "windows":
		cmd = exec.Command("notepad.exe", filePath)
	default:
		r.showError("Sistema operativo no soportado")
		return
	}

	// 🔧 MANEJO DE ERRORES MEJORADO
	if err := cmd.Start(); err != nil {
		// Si falla, mostrar mensaje y usar fallback
		dialog.ShowInformation("📝 Abriendo archivo",
			fmt.Sprintf("No se pudo abrir editor automáticamente.\nArchivo: %s\n\n💡 Puedes abrirlo manualmente con cualquier editor de texto.",
				filePath),
			fyne.CurrentApp().Driver().AllWindows()[0])

		// Fallback: intentar xdg-open
		fallbackCmd := exec.Command("xdg-open", filePath)
		fallbackCmd.Start()
	}
}

// 🔧 MEJORAR la función copyToClipboard
func (r *Reports) copyToClipboard(text string) {
	// Usar el portapapeles del sistema si está disponible
	clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
	clipboard.SetContent(text)

	dialog.ShowInformation("📋 Copiado",
		fmt.Sprintf("✅ Código ARM64 copiado al portapapeles\n📊 %d líneas copiadas",
			len(strings.Split(text, "\n"))),
		fyne.CurrentApp().Driver().AllWindows()[0])
}

func (r *Reports) UpdateARM64Content(armCode, filePath string) {
	//  TÍTULO PRINCIPAL
	titleLabel := widget.NewLabel("⚙️ Código ARM64 Generado")
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	titleLabel.Alignment = fyne.TextAlignCenter

	// 📊 INFORMACIÓN TÉCNICA SIMPLIFICADA
	fileInfo := r.getDetailedFileInfo(armCode, filePath)
	fileInfoLabel := widget.NewLabel(fileInfo)
	fileInfoLabel.TextStyle = fyne.TextStyle{Monospace: true}
	fileInfoLabel.Alignment = fyne.TextAlignCenter
	fileInfoLabel.Wrapping = fyne.TextWrapWord

	// 🔥 TRES BOTONES CON NOMBRES COMPLETOS
	openFileBtn := widget.NewButton("📂 Abrir en Editor", func() {
		r.openFileInTextEditor(filePath)
	})

	copyCodeBtn := widget.NewButton("📋 Copiar Código", func() {
		r.copyToClipboard(armCode)
	})

	// BOTÓN EJECUTAR CON NOMBRE COMPLETO
	executeBtn := widget.NewButton("🚀 Ejecutar ARM64", func() {
		r.executeARM64Script()
	})
	executeBtn.Importance = widget.HighImportance

	// 🎨 CONTENEDOR DE BOTONES CENTRADOS (3 BOTONES CON NOMBRES COMPLETOS)
	buttonContainer := container.NewHBox(
		widget.NewLabel("  "), // Spacer
		widget.NewLabel("  "), // Spacer
		openFileBtn,
		widget.NewLabel(" "), // Spacer pequeño
		copyCodeBtn,
		widget.NewLabel(" "), // Spacer pequeño
		executeBtn,
		widget.NewLabel("  "), // Spacer
	)

	//  LAYOUT CON BOTÓN DE EJECUCIÓN
	content := container.NewVBox(
		titleLabel,
		widget.NewLabel(""), // Spacer
		fileInfoLabel,
		widget.NewLabel(""), // Spacer
		buttonContainer,
		widget.NewLabel(""), // Spacer
		widget.NewLabel("✅ Archivo listo para compilación ARM64"),
		widget.NewLabel(""), // Spacer inferior
	)

	r.ARM64Content.Objects = []fyne.CanvasObject{content}
	r.ARM64Content.Refresh()
}

// MÉTODO PARA OBTENER INFORMACIÓN DETALLADA DEL ARCHIVO
func (r *Reports) executeARM64Script() {
	scriptPath := "./ejecutarARM.sh"

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		dialog.ShowError(fmt.Errorf("❌ Script no encontrado: %s", scriptPath),
			fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}

	os.Chmod(scriptPath, 0755)

	var cmd *exec.Cmd
	var foundTerminal string

	switch runtime.GOOS {
	case "linux":
		terminals := []string{
			"code", "xterm", "konsole", "xfce4-terminal", "gnome-terminal",
		}

		for _, terminal := range terminals {
			if _, err := exec.LookPath(terminal); err == nil {
				foundTerminal = terminal
				break
			}
		}

		if foundTerminal == "" {
			dialog.ShowError(fmt.Errorf("❌ No se encontró terminal disponible"),
				fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		absScriptPath, _ := filepath.Abs(scriptPath)
		currentDir := getCurrentDir()

		switch foundTerminal {
		case "code":
			var terminalCmd *exec.Cmd

			autoScript := fmt.Sprintf(`#!/bin/bash
clear
echo "🚀 Ejecutando ARM64 desde vlang-cherry-ide..."
echo "📋 Script: %s"
echo ""
bash "%s"
echo ""
echo "✅ Completado. Presiona Enter para cerrar..."
read
`, scriptPath, absScriptPath)

			tempPath := "/tmp/vscode_arm64_exec.sh"
			os.WriteFile(tempPath, []byte(autoScript), 0755)

			if _, err := exec.LookPath("xterm"); err == nil {
				terminalCmd = exec.Command("xterm",
					"-title", "🚀 ARM64 - vlang-cherry-ide",
					"-geometry", "120x35",
					"-fa", "Ubuntu Mono", "-fs", "14",
					"-e", "bash", "-c", fmt.Sprintf("bash '%s'", tempPath))
			} else if _, err := exec.LookPath("gnome-terminal"); err == nil {
				terminalCmd = exec.Command("bash", "-c",
					fmt.Sprintf(`unset SNAP && unset SNAP_DATA && /usr/bin/gnome-terminal --title="🚀 ARM64 - vlang-cherry-ide" --geometry=120x35 -- bash -c "bash '%s'"`, tempPath))
			}

			if terminalCmd != nil {
				cmd = terminalCmd
			} else {
				dialog.ShowError(fmt.Errorf("❌ No se encontró terminal disponible\n\n💡 Instala xterm: sudo apt install xterm"),
					fyne.CurrentApp().Driver().AllWindows()[0])
				return
			}

		case "xterm":
			cmd = exec.Command("xterm", "-title", "🚀 Ejecutando ARM64",
				"-geometry", "120x40", "-fa", "Ubuntu Mono", "-fs", "14",
				"-fg", "white", "-bg", "black", "-e", "bash", "-c",
				fmt.Sprintf(`cd '%s' && clear && echo "🚀 Ejecutando ARM64..." && bash '%s' && echo "✅ Completado. Presiona Enter..." && read`,
					currentDir, absScriptPath))

		case "konsole":
			cmd = exec.Command("konsole", "--title", "🚀 Ejecutando ARM64",
				"--workdir", currentDir, "-e", "bash", "-c",
				fmt.Sprintf("echo '🚀 Ejecutando ARM64...' && bash '%s' && read", absScriptPath))

		case "xfce4-terminal":
			cmd = exec.Command("xfce4-terminal", "--title=🚀 Ejecutando ARM64",
				"--working-directory", currentDir, "-e",
				fmt.Sprintf("bash -c \"echo '🚀 Ejecutando ARM64...' && bash '%s' && read\"", absScriptPath))

		case "gnome-terminal":
			cmd = exec.Command("bash", "-c",
				fmt.Sprintf(`unset SNAP && unset SNAP_DATA && /usr/bin/gnome-terminal --title="🚀 Ejecutando ARM64" --geometry=120x35 --working-directory="%s" -- bash -c "echo '🚀 Ejecutando ARM64...' && bash '%s' && read"`,
					currentDir, absScriptPath))
		}

	case "darwin":
		foundTerminal = "Terminal (macOS)"
		if _, err := exec.LookPath("code"); err == nil {
			foundTerminal = "VS Code (macOS)"
			cmd = exec.Command("code", getCurrentDir())
		}

	case "windows":
		foundTerminal = "CMD (Windows)"
		if _, err := exec.LookPath("code"); err == nil {
			foundTerminal = "VS Code (Windows)"
			cmd = exec.Command("code", getCurrentDir())
		}

	default:
		dialog.ShowError(fmt.Errorf("❌ SO no soportado: %s", runtime.GOOS),
			fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}

	if err := cmd.Start(); err != nil {
		dialog.ShowError(fmt.Errorf("❌ Error: %v", err),
			fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}

	var terminalName string
	switch foundTerminal {
	case "code":
		terminalName = "VS Code detectado"
	case "xterm":
		terminalName = "XTerm"
	case "gnome-terminal":
		terminalName = "Terminal nativo"
	default:
		terminalName = foundTerminal
	}

	dialog.ShowInformation("🚀 Ejecutando ARM64",
		fmt.Sprintf("✅ Script ejecutándose en %s", terminalName),
		fyne.CurrentApp().Driver().AllWindows()[0])
}

// FUNCIÓN AUXILIAR: Obtener directorio actual
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}

// FUNCIÓN: Información técnica detallada del archivo
func (r *Reports) getDetailedFileInfo(armCode, filePath string) string {
	lines := strings.Split(armCode, "\n")
	lineCount := len(lines)

	info := fmt.Sprintf(`📁 Archivo: %s
📊 Tamaño: %d líneas (%d KB)
 Arquitectura: ARM64 (AArch64)
✅ Estado: Listo para ensamblar

🔧 Compatible con: aarch64-linux-gnu-as
🚀 Ejecutable con: qemu-aarch64`,
		filepath.Base(filePath),
		lineCount,
		len(armCode)/1024)

	return info
}
