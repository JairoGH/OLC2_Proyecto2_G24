package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

// CreateSVGFile crea un archivo SVG temporal y devuelve su ruta
func CreateSVGFile(svgContent string) (string, error) {
	// Crear directorio de salida
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

// OpenSVGInBrowser abre el SVG en el navegador predeterminado
func OpenSVGInBrowser(svgPath string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{svgPath}
	case "darwin": // macOS
		cmd = "open"
		args = []string{svgPath}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", svgPath}
	default:
		return fmt.Errorf("sistema operativo no soportado: %s", runtime.GOOS)
	}

	return exec.Command(cmd, args...).Start()
}

// GetSVGInfo extrae información básica del SVG
func GetSVGInfo(svgContent string) map[string]string {
	info := make(map[string]string)

	// Tamaño del contenido
	info["size"] = fmt.Sprintf("%d caracteres", len(svgContent))

	// Buscar dimensiones
	if len(svgContent) > 100 {
		info["status"] = "✅ SVG válido"
	} else {
		info["status"] = "⚠️ SVG muy pequeño"
	}

	// Verificar si contiene elementos comunes
	if contains := []string{"<path", "<use", "<text", "<g"}; any(svgContent, contains) {
		info["content"] = "Contiene elementos gráficos"
	} else {
		info["content"] = "Sin elementos gráficos detectados"
	}

	return info
}

// any verifica si el string contiene alguno de los elementos
func any(content string, elements []string) bool {
	for _, element := range elements {
		if len(content) > 0 && len(element) > 0 {
			// Búsqueda simple
			for i := 0; i <= len(content)-len(element); i++ {
				if content[i:i+len(element)] == element {
					return true
				}
			}
		}
	}
	return false
}
