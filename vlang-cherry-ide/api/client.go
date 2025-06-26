package api

import (
	instrucciones "main/Instrucciones"
	"main/internal/analyzer"
	"main/models"

	"github.com/antlr4-go/antlr/v4"
)

type Client struct {
	analyzerService *analyzer.AnalyzerService
}

func NewClient() *Client {
	return &Client{
		analyzerService: analyzer.NewAnalyzerService(),
	}
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

// 🔧 MÉTODO 1: Solo análisis (CORREGIDO)
func (c *Client) AnalyzeOnly(code, filename string) (*models.AnalysisOnlyResponse, antlr.ParseTree, *instrucciones.PatronVIsitor, error) {
	return c.analyzerService.AnalyzeOnly(code, filename)
}

// 🔧 MÉTODO 2: Solo ARM64 (CORREGIDO)
func (c *Client) GenerateARM64Only(tree antlr.ParseTree, visitor *instrucciones.PatronVIsitor) (*models.ARM64Response, error) {
	return c.analyzerService.GenerateARM64Only(tree, visitor)
}

// 🆕 MÉTODO 3: Solo AST
func (c *Client) GenerateASTOnly(code string) (*models.ASTResponse, error) {
	return c.analyzerService.GenerateASTOnly(code)
}
