/*
package api

import (

	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"
	"vlang-cherry-ide/models"

)

	type Client struct {
		baseURL    string
		httpClient *http.Client
	}

	func NewClient() *Client {
		return &Client{
			baseURL: "http://127.0.0.1:3000", // Backend real
			httpClient: &http.Client{
				Timeout: 30 * time.Second,
			},
		}
	}

// AnalyzeCode envía código al backend usando FormData como lo hace React

	func (c *Client) AnalyzeCode(code, filename string) (*models.AnalysisResponse, error) {
		// Crear FormData como lo hace React
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		// Agregar el campo "code" (no "input" como antes pensamos)
		err := writer.WriteField("code", code)
		if err != nil {
			return nil, fmt.Errorf("error creando form data: %v", err)
		}

		// Cerrar el writer
		err = writer.Close()
		if err != nil {
			return nil, fmt.Errorf("error cerrando form writer: %v", err)
		}

		// Crear request POST
		req, err := http.NewRequest("POST", c.baseURL+"/compile", &buf)
		if err != nil {
			return nil, fmt.Errorf("error creando request: %v", err)
		}

		// Establecer Content-Type correcto para multipart/form-data
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Ejecutar request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error al conectar con el servidor: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("servidor retornó error: %d", resp.StatusCode)
		}

		// Parsear respuesta del backend real
		var backendResponse models.BackendResponse
		if err := json.NewDecoder(resp.Body).Decode(&backendResponse); err != nil {
			return nil, fmt.Errorf("error al parsear respuesta: %v", err)
		}

		// Convertir respuesta del backend al formato del frontend
		result := convertBackendResponse(backendResponse)

		return &result, nil
	}

// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend
// Convertir respuesta del backend al formato esperado por el frontend

	func convertBackendResponse(backendResp models.BackendResponse) models.AnalysisResponse {
		// Convertir errores
		frontendErrors := make([]models.ErrorReport, len(backendResp.Errors))
		for i, err := range backendResp.Errors {
			frontendErrors[i] = models.ErrorReport{
				Linea:       err.Linea,
				Columna:     err.Columna,
				Descripcion: err.Descripcion,
				Tipo:        err.Tipo,
			}
		}

		// Convertir tabla de símbolos/ámbito
		symbolTable := []models.SymbolEntry{}

		// Extraer variables del ámbito global
		for _, variable := range backendResp.RegistroAmbito.AmbitoGlobal.Vars {
			symbolTable = append(symbolTable, models.SymbolEntry{
				ID:         variable.Name,
				SymbolType: "Variable",
				DataType:   variable.Type,
				Scope:      "global",
				Line:       variable.Line,
				Column:     variable.Column,
			})
		}

		// Extraer funciones del ámbito global
		for _, function := range backendResp.RegistroAmbito.AmbitoGlobal.Funcs {
			symbolTable = append(symbolTable, models.SymbolEntry{
				ID:         function.Name,
				SymbolType: "Function",
				DataType:   function.Type,
				Scope:      "global",
				Line:       function.Line,
				Column:     function.Column,
			})
		}

		// El CST/AST viene como SVG - mostrarlo nativamente en Fyne
		astString := "AST disponible como SVG nativo en el IDE"
		if backendResp.CSTSvg != "" {
			astString = fmt.Sprintf("AST generado exitosamente\nMostrándose directamente en el IDE (%d caracteres)", len(backendResp.CSTSvg))
		}

		return models.AnalysisResponse{
			Success:       len(backendResp.Errors) == 0,
			AST:           astString,
			Errors:        frontendErrors,
			SymbolTable:   symbolTable,
			ConsoleOutput: backendResp.Output,
			CSTSvg:        backendResp.CSTSvg, // SVG para mostrar nativamente
			ASTPng:        nil,                // No necesitamos PNG
		}
	}

// ValidateCode solo valida sintaxis

	func (c *Client) ValidateCode(code string) (*models.AnalysisResponse, error) {
		return c.AnalyzeCode(code, "temp.vch")
	}

// GetVersion retorna información del compilador

	func (c *Client) GetVersion() (string, error) {
		return "V-Lang Cherry Compiler (Backend Real)", nil
	}

// Ping verifica que el backend esté funcionando
// Ping verifica que el backend esté funcionando

	func (c *Client) Ping() error {
		// En lugar de GET /, hacemos una prueba pequeña con POST /compile
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		// Enviar código simple para probar
		err := writer.WriteField("code", "// ping test")
		if err != nil {
			return fmt.Errorf("error creando ping request: %v", err)
		}

		err = writer.Close()
		if err != nil {
			return fmt.Errorf("error cerrando ping writer: %v", err)
		}

		// Crear request de prueba
		req, err := http.NewRequest("POST", c.baseURL+"/compile", &buf)
		if err != nil {
			return fmt.Errorf("error creando ping request: %v", err)
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Ejecutar con timeout corto para ping
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("servidor no disponible: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			return fmt.Errorf("servidor retornó error: %d", resp.StatusCode)
		}

		return nil
	}

// SetBaseURL permite cambiar la URL del backend

	func (c *Client) SetBaseURL(url string) {
		c.baseURL = url
	}

// GetBaseURL retorna la URL actual del backend

	func (c *Client) GetBaseURL() string {
		return c.baseURL
	}
*/
package api

import (
	"main/internal/analyzer"
	"main/models"
)

type Client struct {
	analyzerService *analyzer.AnalyzerService
}

func NewClient() *Client {
	return &Client{
		analyzerService: analyzer.NewAnalyzerService(),
	}
}

// AnalyzeCode ahora usa el servicio local completo - NO más HTTP
func (c *Client) AnalyzeCode(code, filename string) (*models.AnalysisResponse, error) {
	// Llamada directa al servicio local que integra tu backend completo
	return c.analyzerService.AnalyzeCode(code, filename)
}

// ValidateCode solo valida sintaxis
func (c *Client) ValidateCode(code string) (*models.AnalysisResponse, error) {
	return c.analyzerService.ValidateCode(code)
}

// GetVersion retorna información del compilador
func (c *Client) GetVersion() (string, error) {
	return c.analyzerService.GetVersion()
}

// Ping verifica que el analizador esté funcionando
func (c *Client) Ping() error {
	return c.analyzerService.Ping()
}

// Métodos de compatibilidad (ya no necesarios pero los mantengo por compatibilidad)
func (c *Client) SetBaseURL(url string) {
	// Ya no hace nada - sin servidor HTTP
}

func (c *Client) GetBaseURL() string {
	return "local://direct-integration"
}
