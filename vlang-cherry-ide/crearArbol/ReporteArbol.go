package crearArbol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// Estructura para almacenar la respuesta del servicio de ANTLR
type CSTresponse struct {
	SVGTree string `json:"svgtree"`
}

// Lee el contenido de un archivo y lo devuelve como cadena
func LeerArchivo(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contenido, _ := io.ReadAll(file)
	return string(contenido)
}

// Genera un reporte del árbol sintáctico a partir del código fuente
func ReporteArbol(input string) string {
	// Obtener ruta al directorio actual
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	// Obtener directorio raíz del proyecto
	rootPath := filepath.Join(currentDir, "..")

	// Construir rutas a los archivos de gramática
	parserPath := filepath.Join(rootPath, "parser", "VGrammar.g4")
	lexerPath := filepath.Join(rootPath, "parser", "VLexer.g4")

	// Leer gramáticas
	parser, err := json.Marshal(LeerArchivo(parserPath))
	if err != nil {
		fmt.Println("Error al leer la gramática del parser:", err)
		return ""
	}
	parserContent := string(parser)

	lexer, err := json.Marshal(LeerArchivo(lexerPath))
	if err != nil {
		fmt.Println("Error al leer la gramática del lexer:", err)
		return ""
	}
	lexerContent := string(lexer)

	jinput, err := json.Marshal(input)
	finput := string(jinput)

	// Crear payload para la petición
	payload := []byte(
		fmt.Sprintf(
			`{
				"grammar": %s,
				"input": %s,
				"lexgrammar": %s,
				"start": "%s"
			}`,
			parserContent,
			finput,
			lexerContent,
			"program",
		),
	)

	// Enviar solicitud al servicio de ANTLR
	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return ""
	}

	// Procesar la respuesta JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error al procesar el JSON:", err)
		return ""
	}

	result := data["result"].(map[string]interface{})
	return result["svgtree"].(string)
}
