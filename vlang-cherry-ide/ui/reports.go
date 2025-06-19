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
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type Reports struct {
    Container   *container.AppTabs
    ErrorList   *widget.Table
    SymbolTable *widget.Table
    ASTContent  *fyne.Container

    // 🆕 PANELES PARA ERRORES Y SÍMBOLOS
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
    
    // 🆕 CREAR PANEL DE DETALLES PARA ERRORES (SIMPLE Y PEQUEÑO)
    reports.ErrorDetails = widget.NewCard(
        "Mensaje Completo",
        "",
        widget.NewLabel("Selecciona un error para ver el mensaje completo"),
    )
    reports.ErrorDetails.Hide() // 🎯 SIEMPRE OCULTO AL INICIO

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

    // 🆕 EVENTO DE SELECCIÓN PARA ERRORES
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

    // 🆕 CREAR LAYOUT PARA ERRORES
    reports.ErrorPanel = container.NewVSplit(
        reports.ErrorList,
        reports.ErrorDetails,
    )
    reports.ErrorPanel.SetOffset(1.0) // 🎯 INICIAR COMPLETAMENTE OCULTO

    // ========== CONFIGURACIÓN TABLA DE SÍMBOLOS ==========

    // 🆕 CREAR PANEL DE DETALLES PARA SÍMBOLOS
    reports.SymbolDetails = widget.NewCard(
        "Detalles del Símbolo",
        "",
        widget.NewLabel("Selecciona un símbolo para ver los detalles completos"),
    )
    reports.SymbolDetails.Hide() // 🎯 SIEMPRE OCULTO AL INICIO

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
                            text = text + " 🔍" // 🆕 INDICADOR PARA ID
                        }
                    case 1: // Tipo
                        text = reports.truncateText(symbol.SymbolType, 12)
                        label.Alignment = fyne.TextAlignCenter
                        if len(symbol.SymbolType) > 12 {
                            text = text + " ..." // 🆕 INDICADOR PARA TIPO
                        }
                    case 2: // Tipo de Dato
                        text = reports.truncateText(symbol.DataType, 20)
                        label.Alignment = fyne.TextAlignCenter
                        if len(symbol.DataType) > 20 {
                            text = text + " ..." // 🆕 INDICADOR PARA TIPO DE DATO
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
                    label.TextStyle = fyne.TextStyle{Bold: false}
                }
            }
        },
    )

    // 🆕 EVENTO DE SELECCIÓN PARA SÍMBOLOS
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

    // 🆕 CREAR LAYOUT PARA SÍMBOLOS
    reports.SymbolPanel = container.NewVSplit(
        reports.SymbolTable,
        reports.SymbolDetails,
    )
    reports.SymbolPanel.SetOffset(1.0) // 🎯 INICIAR COMPLETAMENTE OCULTO

    // ========== CONFIGURACIÓN AST ==========

    // Crear contenedor para AST
    reports.ASTContent = container.NewVBox(
        widget.NewLabel("No hay AST disponible"),
    )

    // 🆕 PESTAÑAS CON PANELES
    reports.Container = container.NewAppTabs(
        container.NewTabItem("Errores", reports.ErrorPanel),           // Panel con mensaje para errores
        container.NewTabItem("Tabla de Símbolos", reports.SymbolPanel), // Panel con detalles para símbolos
        container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
    )

    return reports
}

// 🆕 FUNCIÓN PARA MOSTRAR MENSAJE DE ERROR (MEJORADA)
func (r *Reports) showErrorMessage(error models.ErrorReport) {
    // 🎯 FORMATO MEJORADO COMO PEDISTE
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

// 🆕 FUNCIÓN PARA MOSTRAR DETALLES DE SÍMBOLO
func (r *Reports) showSymbolMessage(symbol models.SymbolEntry) {
    // 🎯 FORMATO SIMILAR AL DE ERRORES
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

// 🆕 FUNCIÓN PARA OCULTAR TODOS LOS PANELES AL CAMBIAR DE PESTAÑA
func (r *Reports) HideAllPanels() {
    // 🎯 OCULTAR PANEL DE ERRORES
    r.ErrorDetails.Hide()
    r.ErrorPanel.SetOffset(1.0)
    
    // 🎯 OCULTAR PANEL DE SÍMBOLOS
    r.SymbolDetails.Hide()
    r.SymbolPanel.SetOffset(1.0)
}

// UpdateErrors actualiza la lista de errores
func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
    r.errors = errors
    r.ErrorList.Refresh()
    
    // 🎯 ASEGURAR QUE ESTÉ OCULTO AL ACTUALIZAR
    r.ErrorDetails.Hide()
    r.ErrorPanel.SetOffset(1.0)
}

// UpdateSymbolTable actualiza la tabla de símbolos
func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
    r.symbolTable = symbols
    r.SymbolTable.Refresh()
    
    // 🎯 ASEGURAR QUE ESTÉ OCULTO AL ACTUALIZAR
    r.SymbolDetails.Hide()
    r.SymbolPanel.SetOffset(1.0)
}

// ClearAll limpia todos los reportes
func (r *Reports) ClearAll() {
    r.errors = []models.ErrorReport{}
    r.symbolTable = []models.SymbolEntry{}

    r.ErrorList.Refresh()
    r.SymbolTable.Refresh()

    // 🎯 OCULTAR TODOS LOS PANELES AL LIMPIAR
    r.HideAllPanels()

    // Limpiar AST
    label := widget.NewLabel("No hay AST disponible")
    r.ASTContent.Objects = []fyne.CanvasObject{label}
    r.ASTContent.Refresh()
}

// ============ RESTO DE FUNCIONES MANTENER IGUAL ============
// (UpdateASTWithSVG, convertAndShowAST, etc. - sin cambios)

// UpdateASTWithSVG - Conversión automática SVG a PNG
func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
    if svgContent != "" {
        svgPath, err := r.saveSVGToFile(svgContent)
        if err != nil {
            r.showError("Error guardando SVG: " + err.Error())
            return
        }
        r.convertAndShowAST(svgPath, svgContent)
    } else {
        r.showError("No hay AST disponible")
    }
}

func (r *Reports) convertAndShowAST(svgPath, svgContent string) {
    pngPath, pngData, err := r.convertSVGToPNG(svgPath)

    if err == nil && len(pngData) > 0 {
        r.showPNGSuccess(svgPath, pngPath, pngData, svgContent)
    } else {
        r.showConversionFailed(svgPath, svgContent, err.Error())
    }
}

func (r *Reports) convertSVGToPNG(svgPath string) (string, []byte, error) {
    dir := filepath.Dir(svgPath)
    base := strings.TrimSuffix(filepath.Base(svgPath), ".svg")
    pngPath := filepath.Join(dir, base+".png")

    converters := [][]string{
        {"inkscape", svgPath, "--export-type=png", "--export-dpi=300", "--export-width=1200", "--export-filename=" + pngPath},
        {"rsvg-convert", "-a", "-w", "1200", "-h", "900", "-f", "png", "-o", pngPath, svgPath},
        {"convert", "-density", "300", "-background", "white", "-resize", "1200x900", svgPath, pngPath},
        {"cairosvg", svgPath, "-o", pngPath, "--output-width", "1200"},
    }

    for _, cmd := range converters {
        execCmd := exec.Command(cmd[0], cmd[1:]...)
        if err := execCmd.Run(); err == nil {
            if pngData, err := os.ReadFile(pngPath); err == nil && len(pngData) > 1000 {
                return pngPath, pngData, nil
            }
        }
    }

    return "", nil, fmt.Errorf("no se encontró ningún convertidor")
}

func (r *Reports) showPNGSuccess(svgPath, pngPath string, pngData []byte, svgContent string) {
    pngResource := fyne.NewStaticResource("ast.png", pngData)
    astImage := canvas.NewImageFromResource(pngResource)
    astImage.FillMode = canvas.ImageFillContain
    astImage.SetMinSize(fyne.NewSize(600, 400))

    infoText := fmt.Sprintf(`🌳 AST Mostrado Como Imagen Real

✅ Conversión exitosa SVG → PNG
📁 SVG: %s
🖼️ PNG: %s  
📊 Tamaño PNG: %d KB

💡 Esta es la imagen EXACTA del SVG`,
        filepath.Base(svgPath),
        filepath.Base(pngPath),
        len(pngData)/1024)

    infoLabel := widget.NewLabel(infoText)
    infoLabel.Wrapping = fyne.TextWrapWord

    openSVGBtn := widget.NewButton("📄 Abrir SVG", func() {
        r.openFileInBrowser(svgPath)
    })

    openPNGBtn := widget.NewButton("🖼️ Abrir PNG", func() {
        r.openFileInBrowser(pngPath)
    })

    content := container.NewVBox(
        infoLabel,
        widget.NewSeparator(),
        container.NewHBox(openSVGBtn, openPNGBtn),
        widget.NewSeparator(),
        widget.NewLabel("🎯 AST Visualizado:"),
        container.NewScroll(astImage),
    )

    r.ASTContent.Objects = []fyne.CanvasObject{content}
    r.ASTContent.Refresh()
}

func (r *Reports) showConversionFailed(svgPath, svgContent string, errorMsg string) {
    installText := `🔧 INSTALAR CONVERTIDOR PARA VER AST REAL

Para ver el AST como imagen exacta en el IDE:

🐧 Linux:
   sudo apt install inkscape

🍎 macOS:
   brew install librsvg

🪟 Windows:
   Descargar Inkscape: https://inkscape.org`

    installLabel := widget.NewLabel(installText)
    installLabel.TextStyle = fyne.TextStyle{Monospace: true}

    currentInfo := fmt.Sprintf(`📊 Estado Actual:

❌ Convertidor no encontrado: %s
📁 SVG guardado: %s
📊 Tamaño: %d caracteres

🌐 Mientras tanto, usa el botón para ver en navegador:`,
        errorMsg, filepath.Base(svgPath), len(svgContent))

    currentLabel := widget.NewLabel(currentInfo)
    currentLabel.Wrapping = fyne.TextWrapWord

    openBtn := widget.NewButton("🌐 Abrir en Navegador", func() {
        r.openFileInBrowser(svgPath)
    })

    content := container.NewVBox(
        currentLabel,
        widget.NewSeparator(),
        openBtn,
        widget.NewSeparator(),
        widget.NewLabel("📋 Instrucciones:"),
        container.NewScroll(installLabel),
    )

    r.ASTContent.Objects = []fyne.CanvasObject{content}
    r.ASTContent.Refresh()
}

func (r *Reports) saveSVGToFile(svgContent string) (string, error) {
    outputDir := "generated_asts"
    os.MkdirAll(outputDir, 0755)

    timestamp := time.Now().Format("2006-01-02_15-04-05")
    filename := fmt.Sprintf("ast_%s.svg", timestamp)
    fullPath := filepath.Join(outputDir, filename)

    err := os.WriteFile(fullPath, []byte(svgContent), 0644)
    if err != nil {
        return "", err
    }

    absPath, err := filepath.Abs(fullPath)
    if err != nil {
        return fullPath, nil
    }

    return absPath, nil
}

func (r *Reports) showError(message string) {
    label := widget.NewLabel("❌ " + message)
    label.Alignment = fyne.TextAlignCenter
    r.ASTContent.Objects = []fyne.CanvasObject{label}
    r.ASTContent.Refresh()
}

func (r *Reports) UpdateASTWithPNG(astText string, pngData []byte) {
    if len(pngData) > 0 {
        resource := fyne.NewStaticResource("ast.png", pngData)
        astImage := canvas.NewImageFromResource(resource)
        astImage.FillMode = canvas.ImageFillContain
        astImage.SetMinSize(fyne.NewSize(400, 300))

        content := container.NewVBox(
            widget.NewLabel("🌳 AST generado como PNG"),
            widget.NewSeparator(),
            container.NewScroll(astImage),
        )

        r.ASTContent.Objects = []fyne.CanvasObject{content}
        r.ASTContent.Refresh()
    } else {
        label := widget.NewLabel("❌ No hay datos PNG disponibles")
        label.Alignment = fyne.TextAlignCenter
        r.ASTContent.Objects = []fyne.CanvasObject{label}
        r.ASTContent.Refresh()
    }
}

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