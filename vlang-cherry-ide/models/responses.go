package models

// Respuesta que espera tu frontend de escritorio
type AnalysisResponse struct {
	Success       bool          `json:"success"`
	AST           string        `json:"ast"`
	Errors        []ErrorReport `json:"errors"`
	SymbolTable   []SymbolEntry `json:"symbolTable"`
	ConsoleOutput string        `json:"consoleOutput"`
	CSTSvg        string        `json:"cstSvg"` // SVG original del backend
	ASTPng        []byte        `json:"-"`      // PNG convertido (no se serializa en JSON)
}

// Estructura de errores para el frontend
type ErrorReport struct {
	Linea       int    `json:"linea"`
	Columna     int    `json:"columna"`
	Descripcion string `json:"descripcion"`
	Tipo        string `json:"tipo"`
}

// Entrada de la tabla de símbolos para el frontend
type SymbolEntry struct {
	ID         string `json:"id"`
	SymbolType string `json:"symbolType"`
	DataType   string `json:"dataType"`
	Scope      string `json:"scope"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`
}

// ARM64Response para la generación de código ARM64 separada
type ARM64Response struct {
	Success       bool   `json:"success"`
	ConsoleOutput string `json:"console_output"`
	ARM64Code     string `json:"arm64_code"`
	FilePath      string `json:"file_path"`
	ErrorMessage  string `json:"error_message,omitempty"`
}

// ASTResponse para la generación de AST separada
type ASTResponse struct {
	Success       bool   `json:"success"`
	ConsoleOutput string `json:"console_output"`
	SVGContent    string `json:"svg_content"`
	PNGData       []byte `json:"png_data,omitempty"`
	ErrorMessage  string `json:"error_message,omitempty"`
}

// AnalysisOnlyResponse para solo análisis (extendiendo AnalysisResponse)
type AnalysisOnlyResponse struct {
	Success       bool          `json:"success"`
	Errors        []ErrorReport `json:"errors"`
	SymbolTable   []SymbolEntry `json:"symbol_table"`
	ConsoleOutput string        `json:"console_output"`
	ErrorMessage  string        `json:"error_message,omitempty"`

	// 🆕 Campos para el estado interno (no se envían al frontend como JSON)
	ParseTree interface{} `json:"-"` // El árbol ANTLR para reutilizar
	Visitor   interface{} `json:"-"` // El visitor con estado
}
