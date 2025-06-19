// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VGrammar struct {
	*antlr.BaseParser
}

var VGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vgrammarParserInit() {
	staticData := &VGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'mut'", "':='", "'+='", "'-='", "'='", "'int'", "",
		"'string'", "'bool'", "", "", "", "", "'nil'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'",
		"'||'", "'!'", "'main'", "'fn'", "'struct'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'",
		"'return'", "'in'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'.'",
		"','", "';'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMENTARIO", "COMENTARIOMULT", "RW_MUT", "OP_ASSIGN", "OP_INCREMENTO",
		"OP_DECREMENTO", "OP_DECLARACION", "RW_INT", "RW_FLOAT64", "RW_STRING",
		"RW_BOOL", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL",
		"NIL_LITERAL", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD",
		"OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ", "OP_MENORIGUAL",
		"OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "RW_MAIN", "RW_FN", "RW_STRUCT",
		"RW_IF", "RW_ELSE", "RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR",
		"RW_WHILE", "RW_BREAK", "RW_CONTINUE", "RW_RETURN", "RW_IN", "ID", "PAR_IZQ",
		"PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER",
		"PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS",
	}
	staticData.RuleNames = []string{
		"program", "main_func", "stmt", "stmt_declaracion", "stmt_asignar",
		"stmt_transferencia", "if_stmt", "if_chain", "else_stmt", "switch_stmt",
		"switch_case", "default_case", "while_stmt", "for_clasico_stmt", "for_stmt",
		"range", "declarar_funcion", "llamar_funcion", "lista_parametros", "parametro_fun",
		"lista_argumentos", "argumento_fun", "declarar_struct", "propiedad_struct",
		"crear_struct", "lista_parametros_init", "parametros_init_struct", "lista_slice",
		"item_slice", "prop_slice", "fun_slice", "tipo", "tipos_slices", "expr",
		"literal", "patronId",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 498, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 5, 0,
		74, 8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 0, 3, 0, 80, 8, 0, 1, 0, 3, 0, 83,
		8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1,
		94, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 134, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 140, 8, 3, 3, 3, 142, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 160, 8, 4,
		1, 5, 1, 5, 3, 5, 164, 8, 5, 1, 5, 1, 5, 3, 5, 168, 8, 5, 1, 6, 1, 6, 1,
		6, 5, 6, 173, 8, 6, 10, 6, 12, 6, 176, 9, 6, 1, 6, 3, 6, 179, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 185, 8, 7, 10, 7, 12, 7, 188, 9, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 5, 8, 195, 8, 8, 10, 8, 12, 8, 198, 9, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 206, 8, 9, 10, 9, 12, 9, 209, 9, 9, 1,
		9, 3, 9, 212, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 220,
		8, 10, 10, 10, 12, 10, 223, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 228, 8,
		11, 10, 11, 12, 11, 231, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 237,
		8, 12, 10, 12, 12, 12, 240, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 246,
		8, 13, 1, 13, 1, 13, 3, 13, 250, 8, 13, 1, 13, 1, 13, 3, 13, 254, 8, 13,
		1, 13, 1, 13, 5, 13, 258, 8, 13, 10, 13, 12, 13, 261, 9, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 269, 8, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 274, 8, 14, 1, 14, 1, 14, 5, 14, 278, 8, 14, 10, 14, 12, 14, 281,
		9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 295, 8, 16, 1, 16, 1, 16, 3, 16, 299, 8, 16, 1,
		16, 1, 16, 5, 16, 303, 8, 16, 10, 16, 12, 16, 306, 9, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 3, 17, 313, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 5, 18, 320, 8, 18, 10, 18, 12, 18, 323, 9, 18, 1, 18, 3, 18, 326, 8,
		18, 1, 19, 3, 19, 329, 8, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		5, 20, 337, 8, 20, 10, 20, 12, 20, 340, 9, 20, 1, 20, 3, 20, 343, 8, 20,
		1, 21, 1, 21, 3, 21, 347, 8, 21, 1, 21, 1, 21, 3, 21, 351, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 5, 22, 357, 8, 22, 10, 22, 12, 22, 360, 9, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 368, 8, 23, 1, 24, 1, 24,
		1, 24, 3, 24, 373, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 380,
		8, 25, 10, 25, 12, 25, 383, 9, 25, 1, 25, 3, 25, 386, 8, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 396, 8, 27, 10, 27,
		12, 27, 399, 9, 27, 3, 27, 401, 8, 27, 1, 27, 3, 27, 404, 8, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 4, 28, 413, 8, 28, 11, 28, 12,
		28, 414, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 431, 8, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 441, 8, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 3, 33, 458, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 5, 33, 478, 8, 33, 10, 33, 12, 33, 481, 9, 33, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 488, 8, 34, 1, 35, 1, 35, 1, 35, 5,
		35, 493, 8, 35, 10, 35, 12, 35, 496, 9, 35, 1, 35, 0, 1, 66, 36, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 7, 1, 0,
		6, 7, 1, 0, 6, 8, 2, 0, 19, 19, 31, 31, 1, 0, 20, 22, 1, 0, 18, 19, 1,
		0, 25, 28, 1, 0, 23, 24, 551, 0, 75, 1, 0, 0, 0, 2, 84, 1, 0, 0, 0, 4,
		109, 1, 0, 0, 0, 6, 141, 1, 0, 0, 0, 8, 159, 1, 0, 0, 0, 10, 167, 1, 0,
		0, 0, 12, 169, 1, 0, 0, 0, 14, 180, 1, 0, 0, 0, 16, 191, 1, 0, 0, 0, 18,
		201, 1, 0, 0, 0, 20, 215, 1, 0, 0, 0, 22, 224, 1, 0, 0, 0, 24, 232, 1,
		0, 0, 0, 26, 243, 1, 0, 0, 0, 28, 264, 1, 0, 0, 0, 30, 284, 1, 0, 0, 0,
		32, 290, 1, 0, 0, 0, 34, 309, 1, 0, 0, 0, 36, 316, 1, 0, 0, 0, 38, 328,
		1, 0, 0, 0, 40, 333, 1, 0, 0, 0, 42, 346, 1, 0, 0, 0, 44, 352, 1, 0, 0,
		0, 46, 363, 1, 0, 0, 0, 48, 369, 1, 0, 0, 0, 50, 376, 1, 0, 0, 0, 52, 387,
		1, 0, 0, 0, 54, 391, 1, 0, 0, 0, 56, 407, 1, 0, 0, 0, 58, 416, 1, 0, 0,
		0, 60, 420, 1, 0, 0, 0, 62, 430, 1, 0, 0, 0, 64, 440, 1, 0, 0, 0, 66, 457,
		1, 0, 0, 0, 68, 487, 1, 0, 0, 0, 70, 489, 1, 0, 0, 0, 72, 74, 3, 4, 2,
		0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 80, 3, 2, 1, 0,
		79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 83, 5,
		0, 0, 1, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 1, 1, 0, 0, 0, 84,
		85, 5, 33, 0, 0, 85, 86, 5, 32, 0, 0, 86, 87, 5, 47, 0, 0, 87, 88, 5, 48,
		0, 0, 88, 92, 5, 49, 0, 0, 89, 91, 3, 4, 2, 0, 90, 89, 1, 0, 0, 0, 91,
		94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0,
		0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 50, 0, 0, 96, 3, 1, 0, 0, 0, 97, 110,
		3, 6, 3, 0, 98, 110, 3, 8, 4, 0, 99, 110, 3, 10, 5, 0, 100, 110, 3, 12,
		6, 0, 101, 110, 3, 18, 9, 0, 102, 110, 3, 24, 12, 0, 103, 110, 3, 26, 13,
		0, 104, 110, 3, 28, 14, 0, 105, 110, 3, 34, 17, 0, 106, 110, 3, 60, 30,
		0, 107, 110, 3, 32, 16, 0, 108, 110, 3, 44, 22, 0, 109, 97, 1, 0, 0, 0,
		109, 98, 1, 0, 0, 0, 109, 99, 1, 0, 0, 0, 109, 100, 1, 0, 0, 0, 109, 101,
		1, 0, 0, 0, 109, 102, 1, 0, 0, 0, 109, 103, 1, 0, 0, 0, 109, 104, 1, 0,
		0, 0, 109, 105, 1, 0, 0, 0, 109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0,
		109, 108, 1, 0, 0, 0, 110, 5, 1, 0, 0, 0, 111, 112, 5, 4, 0, 0, 112, 113,
		5, 46, 0, 0, 113, 114, 5, 5, 0, 0, 114, 142, 3, 66, 33, 0, 115, 116, 5,
		46, 0, 0, 116, 117, 5, 5, 0, 0, 117, 142, 3, 66, 33, 0, 118, 119, 5, 4,
		0, 0, 119, 120, 5, 46, 0, 0, 120, 121, 3, 62, 31, 0, 121, 122, 5, 8, 0,
		0, 122, 123, 3, 66, 33, 0, 123, 142, 1, 0, 0, 0, 124, 125, 5, 4, 0, 0,
		125, 126, 5, 46, 0, 0, 126, 142, 3, 62, 31, 0, 127, 128, 5, 46, 0, 0, 128,
		129, 3, 62, 31, 0, 129, 130, 5, 8, 0, 0, 130, 131, 3, 66, 33, 0, 131, 142,
		1, 0, 0, 0, 132, 134, 5, 4, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 5, 46, 0, 0, 136, 137, 5, 8, 0, 0,
		137, 139, 3, 62, 31, 0, 138, 140, 3, 66, 33, 0, 139, 138, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 111, 1, 0, 0, 0, 141, 115,
		1, 0, 0, 0, 141, 118, 1, 0, 0, 0, 141, 124, 1, 0, 0, 0, 141, 127, 1, 0,
		0, 0, 141, 133, 1, 0, 0, 0, 142, 7, 1, 0, 0, 0, 143, 144, 3, 56, 28, 0,
		144, 145, 5, 8, 0, 0, 145, 146, 3, 66, 33, 0, 146, 160, 1, 0, 0, 0, 147,
		148, 3, 70, 35, 0, 148, 149, 5, 8, 0, 0, 149, 150, 3, 66, 33, 0, 150, 160,
		1, 0, 0, 0, 151, 152, 3, 70, 35, 0, 152, 153, 7, 0, 0, 0, 153, 154, 3,
		66, 33, 0, 154, 160, 1, 0, 0, 0, 155, 156, 3, 56, 28, 0, 156, 157, 7, 1,
		0, 0, 157, 158, 3, 66, 33, 0, 158, 160, 1, 0, 0, 0, 159, 143, 1, 0, 0,
		0, 159, 147, 1, 0, 0, 0, 159, 151, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 160,
		9, 1, 0, 0, 0, 161, 163, 5, 44, 0, 0, 162, 164, 3, 66, 33, 0, 163, 162,
		1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 168, 1, 0, 0, 0, 165, 168, 5, 42,
		0, 0, 166, 168, 5, 43, 0, 0, 167, 161, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0,
		167, 166, 1, 0, 0, 0, 168, 11, 1, 0, 0, 0, 169, 174, 3, 14, 7, 0, 170,
		171, 5, 36, 0, 0, 171, 173, 3, 14, 7, 0, 172, 170, 1, 0, 0, 0, 173, 176,
		1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 178, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 177, 179, 3, 16, 8, 0, 178, 177, 1, 0, 0, 0,
		178, 179, 1, 0, 0, 0, 179, 13, 1, 0, 0, 0, 180, 181, 5, 35, 0, 0, 181,
		182, 3, 66, 33, 0, 182, 186, 5, 49, 0, 0, 183, 185, 3, 4, 2, 0, 184, 183,
		1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 50, 0, 0,
		190, 15, 1, 0, 0, 0, 191, 192, 5, 36, 0, 0, 192, 196, 5, 49, 0, 0, 193,
		195, 3, 4, 2, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194,
		1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198, 196, 1, 0,
		0, 0, 199, 200, 5, 50, 0, 0, 200, 17, 1, 0, 0, 0, 201, 202, 5, 37, 0, 0,
		202, 203, 3, 66, 33, 0, 203, 207, 5, 49, 0, 0, 204, 206, 3, 20, 10, 0,
		205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 212,
		3, 22, 11, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1,
		0, 0, 0, 213, 214, 5, 50, 0, 0, 214, 19, 1, 0, 0, 0, 215, 216, 5, 38, 0,
		0, 216, 217, 3, 66, 33, 0, 217, 221, 5, 56, 0, 0, 218, 220, 3, 4, 2, 0,
		219, 218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 21, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 225, 5,
		39, 0, 0, 225, 229, 5, 56, 0, 0, 226, 228, 3, 4, 2, 0, 227, 226, 1, 0,
		0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0,
		230, 23, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233, 5, 40, 0, 0, 233,
		234, 3, 66, 33, 0, 234, 238, 5, 49, 0, 0, 235, 237, 3, 4, 2, 0, 236, 235,
		1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0,
		0, 0, 239, 241, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 242, 5, 50, 0, 0,
		242, 25, 1, 0, 0, 0, 243, 245, 5, 40, 0, 0, 244, 246, 3, 8, 4, 0, 245,
		244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 249,
		5, 55, 0, 0, 248, 250, 3, 66, 33, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1,
		0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 5, 55, 0, 0, 252, 254, 3, 8, 4,
		0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		259, 5, 49, 0, 0, 256, 258, 3, 4, 2, 0, 257, 256, 1, 0, 0, 0, 258, 261,
		1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0,
		0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 5, 50, 0, 0, 263, 27, 1, 0, 0, 0,
		264, 265, 5, 40, 0, 0, 265, 268, 5, 46, 0, 0, 266, 267, 5, 54, 0, 0, 267,
		269, 5, 46, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270,
		1, 0, 0, 0, 270, 273, 5, 45, 0, 0, 271, 274, 3, 66, 33, 0, 272, 274, 3,
		30, 15, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 279, 5, 49, 0, 0, 276, 278, 3, 4, 2, 0, 277, 276, 1, 0, 0, 0,
		278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283, 5, 50, 0, 0, 283, 29,
		1, 0, 0, 0, 284, 285, 3, 66, 33, 0, 285, 286, 5, 53, 0, 0, 286, 287, 5,
		53, 0, 0, 287, 288, 5, 53, 0, 0, 288, 289, 3, 66, 33, 0, 289, 31, 1, 0,
		0, 0, 290, 291, 5, 33, 0, 0, 291, 292, 5, 46, 0, 0, 292, 294, 5, 47, 0,
		0, 293, 295, 3, 36, 18, 0, 294, 293, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0,
		295, 296, 1, 0, 0, 0, 296, 298, 5, 48, 0, 0, 297, 299, 3, 62, 31, 0, 298,
		297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 304,
		5, 49, 0, 0, 301, 303, 3, 4, 2, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0,
		0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0,
		306, 304, 1, 0, 0, 0, 307, 308, 5, 50, 0, 0, 308, 33, 1, 0, 0, 0, 309,
		310, 3, 70, 35, 0, 310, 312, 5, 47, 0, 0, 311, 313, 3, 40, 20, 0, 312,
		311, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315,
		5, 48, 0, 0, 315, 35, 1, 0, 0, 0, 316, 321, 3, 38, 19, 0, 317, 318, 5,
		54, 0, 0, 318, 320, 3, 38, 19, 0, 319, 317, 1, 0, 0, 0, 320, 323, 1, 0,
		0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0,
		323, 321, 1, 0, 0, 0, 324, 326, 5, 54, 0, 0, 325, 324, 1, 0, 0, 0, 325,
		326, 1, 0, 0, 0, 326, 37, 1, 0, 0, 0, 327, 329, 5, 46, 0, 0, 328, 327,
		1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 331, 5, 46,
		0, 0, 331, 332, 3, 62, 31, 0, 332, 39, 1, 0, 0, 0, 333, 338, 3, 42, 21,
		0, 334, 335, 5, 54, 0, 0, 335, 337, 3, 42, 21, 0, 336, 334, 1, 0, 0, 0,
		337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339,
		342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341, 343, 5, 54, 0, 0, 342, 341,
		1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 41, 1, 0, 0, 0, 344, 345, 5, 46,
		0, 0, 345, 347, 5, 56, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0,
		347, 350, 1, 0, 0, 0, 348, 351, 3, 70, 35, 0, 349, 351, 3, 66, 33, 0, 350,
		348, 1, 0, 0, 0, 350, 349, 1, 0, 0, 0, 351, 43, 1, 0, 0, 0, 352, 353, 5,
		34, 0, 0, 353, 354, 5, 46, 0, 0, 354, 358, 5, 49, 0, 0, 355, 357, 3, 46,
		23, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0,
		358, 359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361,
		362, 5, 50, 0, 0, 362, 45, 1, 0, 0, 0, 363, 364, 3, 62, 31, 0, 364, 367,
		5, 46, 0, 0, 365, 366, 5, 8, 0, 0, 366, 368, 3, 66, 33, 0, 367, 365, 1,
		0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 47, 1, 0, 0, 0, 369, 370, 5, 46, 0,
		0, 370, 372, 5, 49, 0, 0, 371, 373, 3, 50, 25, 0, 372, 371, 1, 0, 0, 0,
		372, 373, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 5, 50, 0, 0, 375,
		49, 1, 0, 0, 0, 376, 381, 3, 52, 26, 0, 377, 378, 5, 54, 0, 0, 378, 380,
		3, 52, 26, 0, 379, 377, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1,
		0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0,
		0, 384, 386, 5, 54, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		51, 1, 0, 0, 0, 387, 388, 5, 46, 0, 0, 388, 389, 5, 56, 0, 0, 389, 390,
		3, 66, 33, 0, 390, 53, 1, 0, 0, 0, 391, 400, 5, 49, 0, 0, 392, 397, 3,
		66, 33, 0, 393, 394, 5, 54, 0, 0, 394, 396, 3, 66, 33, 0, 395, 393, 1,
		0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0,
		0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 392, 1, 0, 0, 0, 400,
		401, 1, 0, 0, 0, 401, 403, 1, 0, 0, 0, 402, 404, 5, 54, 0, 0, 403, 402,
		1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 406, 5, 50,
		0, 0, 406, 55, 1, 0, 0, 0, 407, 412, 3, 70, 35, 0, 408, 409, 5, 51, 0,
		0, 409, 410, 3, 66, 33, 0, 410, 411, 5, 52, 0, 0, 411, 413, 1, 0, 0, 0,
		412, 408, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414,
		415, 1, 0, 0, 0, 415, 57, 1, 0, 0, 0, 416, 417, 3, 56, 28, 0, 417, 418,
		5, 53, 0, 0, 418, 419, 3, 70, 35, 0, 419, 59, 1, 0, 0, 0, 420, 421, 3,
		56, 28, 0, 421, 422, 5, 53, 0, 0, 422, 423, 3, 34, 17, 0, 423, 61, 1, 0,
		0, 0, 424, 431, 5, 46, 0, 0, 425, 431, 5, 9, 0, 0, 426, 431, 5, 10, 0,
		0, 427, 431, 5, 11, 0, 0, 428, 431, 5, 12, 0, 0, 429, 431, 3, 64, 32, 0,
		430, 424, 1, 0, 0, 0, 430, 425, 1, 0, 0, 0, 430, 426, 1, 0, 0, 0, 430,
		427, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 429, 1, 0, 0, 0, 431, 63, 1,
		0, 0, 0, 432, 433, 5, 51, 0, 0, 433, 434, 5, 52, 0, 0, 434, 441, 3, 62,
		31, 0, 435, 436, 5, 51, 0, 0, 436, 437, 5, 52, 0, 0, 437, 438, 5, 51, 0,
		0, 438, 439, 5, 52, 0, 0, 439, 441, 3, 62, 31, 0, 440, 432, 1, 0, 0, 0,
		440, 435, 1, 0, 0, 0, 441, 65, 1, 0, 0, 0, 442, 443, 6, 33, -1, 0, 443,
		444, 5, 47, 0, 0, 444, 445, 3, 66, 33, 0, 445, 446, 5, 48, 0, 0, 446, 458,
		1, 0, 0, 0, 447, 458, 3, 34, 17, 0, 448, 458, 3, 70, 35, 0, 449, 458, 3,
		56, 28, 0, 450, 458, 3, 58, 29, 0, 451, 458, 3, 60, 30, 0, 452, 458, 3,
		68, 34, 0, 453, 458, 3, 54, 27, 0, 454, 458, 3, 48, 24, 0, 455, 456, 7,
		2, 0, 0, 456, 458, 3, 66, 33, 7, 457, 442, 1, 0, 0, 0, 457, 447, 1, 0,
		0, 0, 457, 448, 1, 0, 0, 0, 457, 449, 1, 0, 0, 0, 457, 450, 1, 0, 0, 0,
		457, 451, 1, 0, 0, 0, 457, 452, 1, 0, 0, 0, 457, 453, 1, 0, 0, 0, 457,
		454, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 479, 1, 0, 0, 0, 459, 460,
		10, 6, 0, 0, 460, 461, 7, 3, 0, 0, 461, 478, 3, 66, 33, 7, 462, 463, 10,
		5, 0, 0, 463, 464, 7, 4, 0, 0, 464, 478, 3, 66, 33, 6, 465, 466, 10, 4,
		0, 0, 466, 467, 7, 5, 0, 0, 467, 478, 3, 66, 33, 5, 468, 469, 10, 3, 0,
		0, 469, 470, 7, 6, 0, 0, 470, 478, 3, 66, 33, 4, 471, 472, 10, 2, 0, 0,
		472, 473, 5, 29, 0, 0, 473, 478, 3, 66, 33, 3, 474, 475, 10, 1, 0, 0, 475,
		476, 5, 30, 0, 0, 476, 478, 3, 66, 33, 2, 477, 459, 1, 0, 0, 0, 477, 462,
		1, 0, 0, 0, 477, 465, 1, 0, 0, 0, 477, 468, 1, 0, 0, 0, 477, 471, 1, 0,
		0, 0, 477, 474, 1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0,
		479, 480, 1, 0, 0, 0, 480, 67, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 488,
		5, 13, 0, 0, 483, 488, 5, 14, 0, 0, 484, 488, 5, 15, 0, 0, 485, 488, 5,
		16, 0, 0, 486, 488, 5, 17, 0, 0, 487, 482, 1, 0, 0, 0, 487, 483, 1, 0,
		0, 0, 487, 484, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 486, 1, 0, 0, 0,
		488, 69, 1, 0, 0, 0, 489, 494, 5, 46, 0, 0, 490, 491, 5, 53, 0, 0, 491,
		493, 5, 46, 0, 0, 492, 490, 1, 0, 0, 0, 493, 496, 1, 0, 0, 0, 494, 492,
		1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 71, 1, 0, 0, 0, 496, 494, 1, 0,
		0, 0, 54, 75, 79, 82, 92, 109, 133, 139, 141, 159, 163, 167, 174, 178,
		186, 196, 207, 211, 221, 229, 238, 245, 249, 253, 259, 268, 273, 279, 294,
		298, 304, 312, 321, 325, 328, 338, 342, 346, 350, 358, 367, 372, 381, 385,
		397, 400, 403, 414, 430, 440, 457, 477, 479, 487, 494,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VGrammarInit initializes any static state used to implement VGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VGrammarInit() {
	staticData := &VGrammarParserStaticData
	staticData.once.Do(vgrammarParserInit)
}

// NewVGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewVGrammar(input antlr.TokenStream) *VGrammar {
	VGrammarInit()
	this := new(VGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VGrammar.g4"

	return this
}

// VGrammar tokens.
const (
	VGrammarEOF            = antlr.TokenEOF
	VGrammarWS             = 1
	VGrammarCOMENTARIO     = 2
	VGrammarCOMENTARIOMULT = 3
	VGrammarRW_MUT         = 4
	VGrammarOP_ASSIGN      = 5
	VGrammarOP_INCREMENTO  = 6
	VGrammarOP_DECREMENTO  = 7
	VGrammarOP_DECLARACION = 8
	VGrammarRW_INT         = 9
	VGrammarRW_FLOAT64     = 10
	VGrammarRW_STRING      = 11
	VGrammarRW_BOOL        = 12
	VGrammarINT_LITERAL    = 13
	VGrammarFLOAT_LITERAL  = 14
	VGrammarSTRING_LITERAL = 15
	VGrammarBOOL_LITERAL   = 16
	VGrammarNIL_LITERAL    = 17
	VGrammarOP_SUMA        = 18
	VGrammarOP_RESTA       = 19
	VGrammarOP_MULT        = 20
	VGrammarOP_DIV         = 21
	VGrammarOP_MOD         = 22
	VGrammarOP_IGUAL       = 23
	VGrammarOP_DIFERENTE   = 24
	VGrammarOP_MENORQ      = 25
	VGrammarOP_MAYORQ      = 26
	VGrammarOP_MENORIGUAL  = 27
	VGrammarOP_MAYORIGUAL  = 28
	VGrammarOP_AND         = 29
	VGrammarOP_OR          = 30
	VGrammarOP_NOT         = 31
	VGrammarRW_MAIN        = 32
	VGrammarRW_FN          = 33
	VGrammarRW_STRUCT      = 34
	VGrammarRW_IF          = 35
	VGrammarRW_ELSE        = 36
	VGrammarRW_SWITCH      = 37
	VGrammarRW_CASE        = 38
	VGrammarRW_DEFAULT     = 39
	VGrammarRW_FOR         = 40
	VGrammarRW_WHILE       = 41
	VGrammarRW_BREAK       = 42
	VGrammarRW_CONTINUE    = 43
	VGrammarRW_RETURN      = 44
	VGrammarRW_IN          = 45
	VGrammarID             = 46
	VGrammarPAR_IZQ        = 47
	VGrammarPAR_DER        = 48
	VGrammarLLAVE_IZQ      = 49
	VGrammarLLAVE_DER      = 50
	VGrammarCORCHETE_IZQ   = 51
	VGrammarCORCHETE_DER   = 52
	VGrammarPUNTO          = 53
	VGrammarCOMA           = 54
	VGrammarPUNTO_Y_COMA   = 55
	VGrammarDOS_PUNTOS     = 56
)

// VGrammar rules.
const (
	VGrammarRULE_program                = 0
	VGrammarRULE_main_func              = 1
	VGrammarRULE_stmt                   = 2
	VGrammarRULE_stmt_declaracion       = 3
	VGrammarRULE_stmt_asignar           = 4
	VGrammarRULE_stmt_transferencia     = 5
	VGrammarRULE_if_stmt                = 6
	VGrammarRULE_if_chain               = 7
	VGrammarRULE_else_stmt              = 8
	VGrammarRULE_switch_stmt            = 9
	VGrammarRULE_switch_case            = 10
	VGrammarRULE_default_case           = 11
	VGrammarRULE_while_stmt             = 12
	VGrammarRULE_for_clasico_stmt       = 13
	VGrammarRULE_for_stmt               = 14
	VGrammarRULE_range                  = 15
	VGrammarRULE_declarar_funcion       = 16
	VGrammarRULE_llamar_funcion         = 17
	VGrammarRULE_lista_parametros       = 18
	VGrammarRULE_parametro_fun          = 19
	VGrammarRULE_lista_argumentos       = 20
	VGrammarRULE_argumento_fun          = 21
	VGrammarRULE_declarar_struct        = 22
	VGrammarRULE_propiedad_struct       = 23
	VGrammarRULE_crear_struct           = 24
	VGrammarRULE_lista_parametros_init  = 25
	VGrammarRULE_parametros_init_struct = 26
	VGrammarRULE_lista_slice            = 27
	VGrammarRULE_item_slice             = 28
	VGrammarRULE_prop_slice             = 29
	VGrammarRULE_fun_slice              = 30
	VGrammarRULE_tipo                   = 31
	VGrammarRULE_tipos_slices           = 32
	VGrammarRULE_expr                   = 33
	VGrammarRULE_literal                = 34
	VGrammarRULE_patronId               = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Main_func() IMain_funcContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) Main_func() IMain_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMain_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMain_funcContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(VGrammarEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VGrammarRULE_program)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(72)
				p.Stmt()
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_FN {
		{
			p.SetState(78)
			p.Main_func()
		}

	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(81)
			p.Match(VGrammarEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMain_funcContext is an interface to support dynamic dispatch.
type IMain_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMain_funcContext differentiates from other interfaces.
	IsMain_funcContext()
}

type Main_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_funcContext() *Main_funcContext {
	var p = new(Main_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
	return p
}

func InitEmptyMain_funcContext(p *Main_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
}

func (*Main_funcContext) IsMain_funcContext() {}

func NewMain_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_funcContext {
	var p = new(Main_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_main_func

	return p
}

func (s *Main_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_funcContext) CopyAll(ctx *Main_funcContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Main_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionMainContext struct {
	Main_funcContext
}

func NewFuncionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionMainContext {
	var p = new(FuncionMainContext)

	InitEmptyMain_funcContext(&p.Main_funcContext)
	p.parser = parser
	p.CopyAll(ctx.(*Main_funcContext))

	return p
}

func (s *FuncionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionMainContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionMainContext) RW_MAIN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MAIN, 0)
}

func (s *FuncionMainContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionMainContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionMainContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionMainContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionMainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionMainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionMain(s)
	}
}

func (s *FuncionMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionMain(s)
	}
}

func (s *FuncionMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Main_func() (localctx IMain_funcContext) {
	localctx = NewMain_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VGrammarRULE_main_func)
	var _la int

	localctx = NewFuncionMainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(VGrammarRW_MAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(89)
			p.Stmt()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stmt_declaracion() IStmt_declaracionContext
	Stmt_asignar() IStmt_asignarContext
	Stmt_transferencia() IStmt_transferenciaContext
	If_stmt() IIf_stmtContext
	Switch_stmt() ISwitch_stmtContext
	While_stmt() IWhile_stmtContext
	For_clasico_stmt() IFor_clasico_stmtContext
	For_stmt() IFor_stmtContext
	Llamar_funcion() ILlamar_funcionContext
	Fun_slice() IFun_sliceContext
	Declarar_funcion() IDeclarar_funcionContext
	Declarar_struct() IDeclarar_structContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Stmt_declaracion() IStmt_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_declaracionContext)
}

func (s *StmtContext) Stmt_asignar() IStmt_asignarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_asignarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_asignarContext)
}

func (s *StmtContext) Stmt_transferencia() IStmt_transferenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_transferenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_transferenciaContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_clasico_stmt() IFor_clasico_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_clasico_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_clasico_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *StmtContext) Fun_slice() IFun_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFun_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFun_sliceContext)
}

func (s *StmtContext) Declarar_funcion() IDeclarar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarar_funcionContext)
}

func (s *StmtContext) Declarar_struct() IDeclarar_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarar_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarar_structContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VGrammarRULE_stmt)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Stmt_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Stmt_asignar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Stmt_transferencia()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(101)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(103)
			p.For_clasico_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(104)
			p.For_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(105)
			p.Llamar_funcion()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(106)
			p.Fun_slice()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(107)
			p.Declarar_funcion()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(108)
			p.Declarar_struct()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_declaracionContext is an interface to support dynamic dispatch.
type IStmt_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_declaracionContext differentiates from other interfaces.
	IsStmt_declaracionContext()
}

type Stmt_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_declaracionContext() *Stmt_declaracionContext {
	var p = new(Stmt_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_declaracion
	return p
}

func InitEmptyStmt_declaracionContext(p *Stmt_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_declaracion
}

func (*Stmt_declaracionContext) IsStmt_declaracionContext() {}

func NewStmt_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_declaracionContext {
	var p = new(Stmt_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_declaracion

	return p
}

func (s *Stmt_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_declaracionContext) CopyAll(ctx *Stmt_declaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararSliceContext struct {
	Stmt_declaracionContext
}

func NewDeclararSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSliceContext {
	var p = new(DeclararSliceContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSliceContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSliceContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararSliceContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararSliceContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararSliceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSlice(s)
	}
}

func (s *DeclararSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSlice(s)
	}
}

func (s *DeclararSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaContext struct {
	Stmt_declaracionContext
}

func NewDeclararInferenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaContext {
	var p = new(DeclararInferenciaContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararInferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferencia(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclaraTipoValorContext struct {
	Stmt_declaracionContext
}

func NewDeclaraTipoValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaraTipoValorContext {
	var p = new(DeclaraTipoValorContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclaraTipoValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaraTipoValorContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclaraTipoValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclaraTipoValorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclaraTipoValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclaraTipoValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaraTipoValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclaraTipoValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararSinMutValorContext struct {
	Stmt_declaracionContext
}

func NewDeclararSinMutValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSinMutValorContext {
	var p = new(DeclararSinMutValorContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararSinMutValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSinMutValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSinMutValorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararSinMutValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararSinMutValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararSinMutValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSinMutValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaMutContext struct {
	Stmt_declaracionContext
}

func NewDeclararInferenciaMutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaMutContext {
	var p = new(DeclararInferenciaMutContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararInferenciaMutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaMutContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararInferenciaMutContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaMutContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaMutContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaMutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferenciaMut(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararTipoContext struct {
	Stmt_declaracionContext
}

func NewDeclararTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararTipoContext {
	var p = new(DeclararTipoContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararTipoContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararTipoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_declaracion() (localctx IStmt_declaracionContext) {
	localctx = NewStmt_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VGrammarRULE_stmt_declaracion)
	var _la int

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclararInferenciaMutContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.expr(0)
		}

	case 2:
		localctx = NewDeclararInferenciaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.expr(0)
		}

	case 3:
		localctx = NewDeclaraTipoValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Tipo()
		}
		{
			p.SetState(121)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.expr(0)
		}

	case 4:
		localctx = NewDeclararTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Tipo()
		}

	case 5:
		localctx = NewDeclararSinMutValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Tipo()
		}
		{
			p.SetState(129)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expr(0)
		}

	case 6:
		localctx = NewDeclararSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarRW_MUT {
			{
				p.SetState(132)
				p.Match(VGrammarRW_MUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(135)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Tipo()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(138)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_asignarContext is an interface to support dynamic dispatch.
type IStmt_asignarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_asignarContext differentiates from other interfaces.
	IsStmt_asignarContext()
}

type Stmt_asignarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_asignarContext() *Stmt_asignarContext {
	var p = new(Stmt_asignarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_asignar
	return p
}

func InitEmptyStmt_asignarContext(p *Stmt_asignarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_asignar
}

func (*Stmt_asignarContext) IsStmt_asignarContext() {}

func NewStmt_asignarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_asignarContext {
	var p = new(Stmt_asignarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_asignar

	return p
}

func (s *Stmt_asignarContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_asignarContext) CopyAll(ctx *Stmt_asignarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_asignarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_asignarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionSliceItemContext struct {
	Stmt_asignarContext
}

func NewAsignacionSliceItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionSliceItemContext {
	var p = new(AsignacionSliceItemContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionSliceItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionSliceItemContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *AsignacionSliceItemContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionSliceItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionSliceItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionSliceItem(s)
	}
}

func (s *AsignacionSliceItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionSliceItem(s)
	}
}

func (s *AsignacionSliceItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionSliceItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionDirectaContext struct {
	Stmt_asignarContext
}

func NewAsignacionDirectaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionDirectaContext {
	var p = new(AsignacionDirectaContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionDirectaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionDirectaContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *AsignacionDirectaContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionDirectaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionDirectaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionDirecta(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionSliceContext struct {
	Stmt_asignarContext
	op antlr.Token
}

func NewAsignacionSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionSliceContext {
	var p = new(AsignacionSliceContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionSliceContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionSliceContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *AsignacionSliceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionSliceContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionSliceContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionSliceContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionSlice(s)
	}
}

func (s *AsignacionSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionSlice(s)
	}
}

func (s *AsignacionSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionAritmeticaContext struct {
	Stmt_asignarContext
	op antlr.Token
}

func NewAsignacionAritmeticaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionAritmeticaContext {
	var p = new(AsignacionAritmeticaContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionAritmeticaContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionAritmeticaContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionAritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionAritmeticaContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *AsignacionAritmeticaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionAritmeticaContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionAritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_asignar() (localctx IStmt_asignarContext) {
	localctx = NewStmt_asignarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VGrammarRULE_stmt_asignar)
	var _la int

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionSliceItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Item_slice()
		}
		{
			p.SetState(144)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.expr(0)
		}

	case 2:
		localctx = NewAsignacionDirectaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.PatronId()
		}
		{
			p.SetState(148)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.expr(0)
		}

	case 3:
		localctx = NewAsignacionAritmeticaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.PatronId()
		}
		{
			p.SetState(152)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionAritmeticaContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_INCREMENTO || _la == VGrammarOP_DECREMENTO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionAritmeticaContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(153)
			p.expr(0)
		}

	case 4:
		localctx = NewAsignacionSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
			p.Item_slice()
		}
		{
			p.SetState(156)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionSliceContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionSliceContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(157)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_transferenciaContext is an interface to support dynamic dispatch.
type IStmt_transferenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_transferenciaContext differentiates from other interfaces.
	IsStmt_transferenciaContext()
}

type Stmt_transferenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_transferenciaContext() *Stmt_transferenciaContext {
	var p = new(Stmt_transferenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_transferencia
	return p
}

func InitEmptyStmt_transferenciaContext(p *Stmt_transferenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_transferencia
}

func (*Stmt_transferenciaContext) IsStmt_transferenciaContext() {}

func NewStmt_transferenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_transferenciaContext {
	var p = new(Stmt_transferenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_transferencia

	return p
}

func (s *Stmt_transferenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_transferenciaContext) CopyAll(ctx *Stmt_transferenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_transferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_transferenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Stmt_transferenciaContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) RW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CONTINUE, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Stmt_transferenciaContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) RW_BREAK() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BREAK, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Stmt_transferenciaContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RW_RETURN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_transferencia() (localctx IStmt_transferenciaContext) {
	localctx = NewStmt_transferenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VGrammarRULE_stmt_transferencia)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_RETURN:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(VGrammarRW_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(162)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VGrammarRW_BREAK:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(VGrammarRW_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_CONTINUE:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(VGrammarRW_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFstmtContext struct {
	If_stmtContext
}

func NewIFstmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFstmtContext {
	var p = new(IFstmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IFstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFstmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IFstmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_chainContext)
}

func (s *IFstmtContext) AllRW_ELSE() []antlr.TerminalNode {
	return s.GetTokens(VGrammarRW_ELSE)
}

func (s *IFstmtContext) RW_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, i)
}

func (s *IFstmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IFstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFstmt(s)
	}
}

func (s *IFstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFstmt(s)
	}
}

func (s *IFstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VGrammarRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIFstmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.If_chain()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.Match(VGrammarRW_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(171)
				p.If_chain()
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_ELSE {
		{
			p.SetState(177)
			p.Else_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFcadenaContext struct {
	If_chainContext
}

func NewIFcadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFcadenaContext {
	var p = new(IFcadenaContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IFcadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFcadenaContext) RW_IF() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IF, 0)
}

func (s *IFcadenaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IFcadenaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *IFcadenaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *IFcadenaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IFcadenaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IFcadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFcadena(s)
	}
}

func (s *IFcadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFcadena(s)
	}
}

func (s *IFcadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFcadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VGrammarRULE_if_chain)
	var _la int

	localctx = NewIFcadenaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(VGrammarRW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.expr(0)
	}
	{
		p.SetState(182)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(183)
			p.Stmt()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) RW_ELSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, 0)
}

func (s *ElseStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ElseStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VGrammarRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(VGrammarRW_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(193)
			p.Stmt()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) RW_SWITCH() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_SWITCH, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *SwitchStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VGrammarRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(VGrammarRW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.expr(0)
	}
	{
		p.SetState(203)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VGrammarRW_CASE {
		{
			p.SetState(204)
			p.Switch_case()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_DEFAULT {
		{
			p.SetState(210)
			p.Default_case()
		}

	}
	{
		p.SetState(213)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) RW_CASE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CASE, 0)
}

func (s *SwitchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VGrammarRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(VGrammarRW_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.expr(0)
	}
	{
		p.SetState(217)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(218)
			p.Stmt()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) RW_DEFAULT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_DEFAULT, 0)
}

func (s *DefaultCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VGrammarRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(VGrammarRW_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(226)
			p.Stmt()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *WhileStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VGrammarRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.expr(0)
	}
	{
		p.SetState(234)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(235)
			p.Stmt()
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(241)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_clasico_stmtContext is an interface to support dynamic dispatch.
type IFor_clasico_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_clasico_stmtContext differentiates from other interfaces.
	IsFor_clasico_stmtContext()
}

type For_clasico_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_clasico_stmtContext() *For_clasico_stmtContext {
	var p = new(For_clasico_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_clasico_stmt
	return p
}

func InitEmptyFor_clasico_stmtContext(p *For_clasico_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_clasico_stmt
}

func (*For_clasico_stmtContext) IsFor_clasico_stmtContext() {}

func NewFor_clasico_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_clasico_stmtContext {
	var p = new(For_clasico_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_for_clasico_stmt

	return p
}

func (s *For_clasico_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_clasico_stmtContext) CopyAll(ctx *For_clasico_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_clasico_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_clasico_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForClasicoStmtContext struct {
	For_clasico_stmtContext
}

func NewForClasicoStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForClasicoStmtContext {
	var p = new(ForClasicoStmtContext)

	InitEmptyFor_clasico_stmtContext(&p.For_clasico_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_clasico_stmtContext))

	return p
}

func (s *ForClasicoStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClasicoStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *ForClasicoStmtContext) AllPUNTO_Y_COMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO_Y_COMA)
}

func (s *ForClasicoStmtContext) PUNTO_Y_COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO_Y_COMA, i)
}

func (s *ForClasicoStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ForClasicoStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ForClasicoStmtContext) AllStmt_asignar() []IStmt_asignarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmt_asignarContext); ok {
			len++
		}
	}

	tst := make([]IStmt_asignarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmt_asignarContext); ok {
			tst[i] = t.(IStmt_asignarContext)
			i++
		}
	}

	return tst
}

func (s *ForClasicoStmtContext) Stmt_asignar(i int) IStmt_asignarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_asignarContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_asignarContext)
}

func (s *ForClasicoStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForClasicoStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClasicoStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForClasicoStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterForClasicoStmt(s)
	}
}

func (s *ForClasicoStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitForClasicoStmt(s)
	}
}

func (s *ForClasicoStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitForClasicoStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) For_clasico_stmt() (localctx IFor_clasico_stmtContext) {
	localctx = NewFor_clasico_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VGrammarRULE_for_clasico_stmt)
	var _la int

	localctx = NewForClasicoStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(244)
			p.Stmt_asignar()
		}

	}
	{
		p.SetState(247)
		p.Match(VGrammarPUNTO_Y_COMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&774058334216192) != 0 {
		{
			p.SetState(248)
			p.expr(0)
		}

	}
	{
		p.SetState(251)
		p.Match(VGrammarPUNTO_Y_COMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(252)
			p.Stmt_asignar()
		}

	}
	{
		p.SetState(255)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(256)
			p.Stmt()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *ForStmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ForStmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ForStmtContext) RW_IN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IN, 0)
}

func (s *ForStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ForStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) COMA() antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, 0)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VGrammarRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(266)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(270)
		p.Match(VGrammarRW_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(271)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(272)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(275)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(276)
			p.Stmt()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RangoNumContext struct {
	RangeContext
}

func NewRangoNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangoNumContext {
	var p = new(RangoNumContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *RangoNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoNumContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangoNumContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangoNumContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *RangoNumContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *RangoNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterRangoNum(s)
	}
}

func (s *RangoNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitRangoNum(s)
	}
}

func (s *RangoNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitRangoNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VGrammarRULE_range)
	localctx = NewRangoNumContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.expr(0)
	}
	{
		p.SetState(285)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarar_funcionContext is an interface to support dynamic dispatch.
type IDeclarar_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarar_funcionContext differentiates from other interfaces.
	IsDeclarar_funcionContext()
}

type Declarar_funcionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarar_funcionContext() *Declarar_funcionContext {
	var p = new(Declarar_funcionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_funcion
	return p
}

func InitEmptyDeclarar_funcionContext(p *Declarar_funcionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_funcion
}

func (*Declarar_funcionContext) IsDeclarar_funcionContext() {}

func NewDeclarar_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declarar_funcionContext {
	var p = new(Declarar_funcionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_declarar_funcion

	return p
}

func (s *Declarar_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Declarar_funcionContext) CopyAll(ctx *Declarar_funcionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declarar_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionDecleradaContext struct {
	Declarar_funcionContext
}

func NewFuncionDecleradaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionDecleradaContext {
	var p = new(FuncionDecleradaContext)

	InitEmptyDeclarar_funcionContext(&p.Declarar_funcionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declarar_funcionContext))

	return p
}

func (s *FuncionDecleradaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionDecleradaContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionDecleradaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionDecleradaContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionDecleradaContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionDecleradaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionDecleradaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionDecleradaContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *FuncionDecleradaContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionDecleradaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionDecleradaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionDecleradaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionDeclerada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Declarar_funcion() (localctx IDeclarar_funcionContext) {
	localctx = NewDeclarar_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VGrammarRULE_declarar_funcion)
	var _la int

	localctx = NewFuncionDecleradaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(293)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(296)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2322168557870592) != 0 {
		{
			p.SetState(297)
			p.Tipo()
		}

	}
	{
		p.SetState(300)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(301)
			p.Stmt()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamar_funcionContext is an interface to support dynamic dispatch.
type ILlamar_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLlamar_funcionContext differentiates from other interfaces.
	IsLlamar_funcionContext()
}

type Llamar_funcionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamar_funcionContext() *Llamar_funcionContext {
	var p = new(Llamar_funcionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_llamar_funcion
	return p
}

func InitEmptyLlamar_funcionContext(p *Llamar_funcionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_llamar_funcion
}

func (*Llamar_funcionContext) IsLlamar_funcionContext() {}

func NewLlamar_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamar_funcionContext {
	var p = new(Llamar_funcionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_llamar_funcion

	return p
}

func (s *Llamar_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamar_funcionContext) CopyAll(ctx *Llamar_funcionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Llamar_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamar_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LlamarFuncionContext struct {
	Llamar_funcionContext
}

func NewLlamarFuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamarFuncionContext {
	var p = new(LlamarFuncionContext)

	InitEmptyLlamar_funcionContext(&p.Llamar_funcionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Llamar_funcionContext))

	return p
}

func (s *LlamarFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamarFuncionContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *LlamarFuncionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *LlamarFuncionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *LlamarFuncionContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *LlamarFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLlamarFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Llamar_funcion() (localctx ILlamar_funcionContext) {
	localctx = NewLlamar_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VGrammarRULE_llamar_funcion)
	var _la int

	localctx = NewLlamarFuncionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.PatronId()
	}
	{
		p.SetState(310)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&774058334216192) != 0 {
		{
			p.SetState(311)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(314)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametrosContext is an interface to support dynamic dispatch.
type ILista_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_parametrosContext differentiates from other interfaces.
	IsLista_parametrosContext()
}

type Lista_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametrosContext() *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros
	return p
}

func InitEmptyLista_parametrosContext(p *Lista_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros
}

func (*Lista_parametrosContext) IsLista_parametrosContext() {}

func NewLista_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_parametros

	return p
}

func (s *Lista_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametrosContext) CopyAll(ctx *Lista_parametrosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaParametrosContext struct {
	Lista_parametrosContext
}

func NewListaParametrosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaParametrosContext {
	var p = new(ListaParametrosContext)

	InitEmptyLista_parametrosContext(&p.Lista_parametrosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_parametrosContext))

	return p
}

func (s *ListaParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParametrosContext) AllParametro_fun() []IParametro_funContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametro_funContext); ok {
			len++
		}
	}

	tst := make([]IParametro_funContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametro_funContext); ok {
			tst[i] = t.(IParametro_funContext)
			i++
		}
	}

	return tst
}

func (s *ListaParametrosContext) Parametro_fun(i int) IParametro_funContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametro_funContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametro_funContext)
}

func (s *ListaParametrosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaParametrosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaParametros(s)
	}
}

func (s *ListaParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaParametros(s)
	}
}

func (s *ListaParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_parametros() (localctx ILista_parametrosContext) {
	localctx = NewLista_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VGrammarRULE_lista_parametros)
	var _la int

	var _alt int

	localctx = NewListaParametrosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Parametro_fun()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(317)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(318)
				p.Parametro_fun()
			}

		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(324)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametro_funContext is an interface to support dynamic dispatch.
type IParametro_funContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametro_funContext differentiates from other interfaces.
	IsParametro_funContext()
}

type Parametro_funContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametro_funContext() *Parametro_funContext {
	var p = new(Parametro_funContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametro_fun
	return p
}

func InitEmptyParametro_funContext(p *Parametro_funContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametro_fun
}

func (*Parametro_funContext) IsParametro_funContext() {}

func NewParametro_funContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametro_funContext {
	var p = new(Parametro_funContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_parametro_fun

	return p
}

func (s *Parametro_funContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametro_funContext) CopyAll(ctx *Parametro_funContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Parametro_funContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametro_funContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametroFunContext struct {
	Parametro_funContext
}

func NewParametroFunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametroFunContext {
	var p = new(ParametroFunContext)

	InitEmptyParametro_funContext(&p.Parametro_funContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parametro_funContext))

	return p
}

func (s *ParametroFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroFunContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ParametroFunContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ParametroFunContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ParametroFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParametroFun(s)
	}
}

func (s *ParametroFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParametroFun(s)
	}
}

func (s *ParametroFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParametroFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Parametro_fun() (localctx IParametro_funContext) {
	localctx = NewParametro_funContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VGrammarRULE_parametro_fun)
	localctx = NewParametroFunContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(327)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(330)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Tipo()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_argumentosContext is an interface to support dynamic dispatch.
type ILista_argumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_argumentosContext differentiates from other interfaces.
	IsLista_argumentosContext()
}

type Lista_argumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_argumentosContext() *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_argumentos
	return p
}

func InitEmptyLista_argumentosContext(p *Lista_argumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_argumentos
}

func (*Lista_argumentosContext) IsLista_argumentosContext() {}

func NewLista_argumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_argumentos

	return p
}

func (s *Lista_argumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_argumentosContext) CopyAll(ctx *Lista_argumentosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_argumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_argumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaArgumentosContext struct {
	Lista_argumentosContext
}

func NewListaArgumentosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaArgumentosContext {
	var p = new(ListaArgumentosContext)

	InitEmptyLista_argumentosContext(&p.Lista_argumentosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_argumentosContext))

	return p
}

func (s *ListaArgumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaArgumentosContext) AllArgumento_fun() []IArgumento_funContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumento_funContext); ok {
			len++
		}
	}

	tst := make([]IArgumento_funContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumento_funContext); ok {
			tst[i] = t.(IArgumento_funContext)
			i++
		}
	}

	return tst
}

func (s *ListaArgumentosContext) Argumento_fun(i int) IArgumento_funContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumento_funContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumento_funContext)
}

func (s *ListaArgumentosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaArgumentosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaArgumentosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaArgumentos(s)
	}
}

func (s *ListaArgumentosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaArgumentos(s)
	}
}

func (s *ListaArgumentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaArgumentos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_argumentos() (localctx ILista_argumentosContext) {
	localctx = NewLista_argumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VGrammarRULE_lista_argumentos)
	var _la int

	var _alt int

	localctx = NewListaArgumentosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Argumento_fun()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(334)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(335)
				p.Argumento_fun()
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(341)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumento_funContext is an interface to support dynamic dispatch.
type IArgumento_funContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgumento_funContext differentiates from other interfaces.
	IsArgumento_funContext()
}

type Argumento_funContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumento_funContext() *Argumento_funContext {
	var p = new(Argumento_funContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_argumento_fun
	return p
}

func InitEmptyArgumento_funContext(p *Argumento_funContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_argumento_fun
}

func (*Argumento_funContext) IsArgumento_funContext() {}

func NewArgumento_funContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argumento_funContext {
	var p = new(Argumento_funContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_argumento_fun

	return p
}

func (s *Argumento_funContext) GetParser() antlr.Parser { return s.parser }

func (s *Argumento_funContext) CopyAll(ctx *Argumento_funContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Argumento_funContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argumento_funContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionArgContext struct {
	Argumento_funContext
}

func NewFuncionArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionArgContext {
	var p = new(FuncionArgContext)

	InitEmptyArgumento_funContext(&p.Argumento_funContext)
	p.parser = parser
	p.CopyAll(ctx.(*Argumento_funContext))

	return p
}

func (s *FuncionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionArgContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *FuncionArgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncionArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionArgContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *FuncionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionArg(s)
	}
}

func (s *FuncionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionArg(s)
	}
}

func (s *FuncionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Argumento_fun() (localctx IArgumento_funContext) {
	localctx = NewArgumento_funContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VGrammarRULE_argumento_fun)
	localctx = NewFuncionArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(344)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(VGrammarDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(348)
			p.PatronId()
		}

	case 2:
		{
			p.SetState(349)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarar_structContext is an interface to support dynamic dispatch.
type IDeclarar_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarar_structContext differentiates from other interfaces.
	IsDeclarar_structContext()
}

type Declarar_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarar_structContext() *Declarar_structContext {
	var p = new(Declarar_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_struct
	return p
}

func InitEmptyDeclarar_structContext(p *Declarar_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_struct
}

func (*Declarar_structContext) IsDeclarar_structContext() {}

func NewDeclarar_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declarar_structContext {
	var p = new(Declarar_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_declarar_struct

	return p
}

func (s *Declarar_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Declarar_structContext) CopyAll(ctx *Declarar_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declarar_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararStructContext struct {
	Declarar_structContext
}

func NewDeclararStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararStructContext {
	var p = new(DeclararStructContext)

	InitEmptyDeclarar_structContext(&p.Declarar_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declarar_structContext))

	return p
}

func (s *DeclararStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararStructContext) RW_STRUCT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRUCT, 0)
}

func (s *DeclararStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *DeclararStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *DeclararStructContext) AllPropiedad_struct() []IPropiedad_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropiedad_structContext); ok {
			len++
		}
	}

	tst := make([]IPropiedad_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropiedad_structContext); ok {
			tst[i] = t.(IPropiedad_structContext)
			i++
		}
	}

	return tst
}

func (s *DeclararStructContext) Propiedad_struct(i int) IPropiedad_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropiedad_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropiedad_structContext)
}

func (s *DeclararStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararStruct(s)
	}
}

func (s *DeclararStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararStruct(s)
	}
}

func (s *DeclararStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Declarar_struct() (localctx IDeclarar_structContext) {
	localctx = NewDeclarar_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VGrammarRULE_declarar_struct)
	var _la int

	localctx = NewDeclararStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(VGrammarRW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2322168557870592) != 0 {
		{
			p.SetState(355)
			p.Propiedad_struct()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(361)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropiedad_structContext is an interface to support dynamic dispatch.
type IPropiedad_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPropiedad_structContext differentiates from other interfaces.
	IsPropiedad_structContext()
}

type Propiedad_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropiedad_structContext() *Propiedad_structContext {
	var p = new(Propiedad_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_propiedad_struct
	return p
}

func InitEmptyPropiedad_structContext(p *Propiedad_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_propiedad_struct
}

func (*Propiedad_structContext) IsPropiedad_structContext() {}

func NewPropiedad_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Propiedad_structContext {
	var p = new(Propiedad_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_propiedad_struct

	return p
}

func (s *Propiedad_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Propiedad_structContext) CopyAll(ctx *Propiedad_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Propiedad_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Propiedad_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropiedadStructContext struct {
	Propiedad_structContext
}

func NewPropiedadStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropiedadStructContext {
	var p = new(PropiedadStructContext)

	InitEmptyPropiedad_structContext(&p.Propiedad_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Propiedad_structContext))

	return p
}

func (s *PropiedadStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropiedadStructContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *PropiedadStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *PropiedadStructContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *PropiedadStructContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PropiedadStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropiedadStruct(s)
	}
}

func (s *PropiedadStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropiedadStruct(s)
	}
}

func (s *PropiedadStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropiedadStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Propiedad_struct() (localctx IPropiedad_structContext) {
	localctx = NewPropiedad_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VGrammarRULE_propiedad_struct)
	var _la int

	localctx = NewPropiedadStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Tipo()
	}
	{
		p.SetState(364)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarOP_DECLARACION {
		{
			p.SetState(365)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.expr(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICrear_structContext is an interface to support dynamic dispatch.
type ICrear_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCrear_structContext differentiates from other interfaces.
	IsCrear_structContext()
}

type Crear_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrear_structContext() *Crear_structContext {
	var p = new(Crear_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_crear_struct
	return p
}

func InitEmptyCrear_structContext(p *Crear_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_crear_struct
}

func (*Crear_structContext) IsCrear_structContext() {}

func NewCrear_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Crear_structContext {
	var p = new(Crear_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_crear_struct

	return p
}

func (s *Crear_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Crear_structContext) CopyAll(ctx *Crear_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Crear_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Crear_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CrearStructContext struct {
	Crear_structContext
}

func NewCrearStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CrearStructContext {
	var p = new(CrearStructContext)

	InitEmptyCrear_structContext(&p.Crear_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Crear_structContext))

	return p
}

func (s *CrearStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrearStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *CrearStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *CrearStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *CrearStructContext) Lista_parametros_init() ILista_parametros_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametros_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametros_initContext)
}

func (s *CrearStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterCrearStruct(s)
	}
}

func (s *CrearStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitCrearStruct(s)
	}
}

func (s *CrearStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitCrearStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Crear_struct() (localctx ICrear_structContext) {
	localctx = NewCrear_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VGrammarRULE_crear_struct)
	var _la int

	localctx = NewCrearStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(371)
			p.Lista_parametros_init()
		}

	}
	{
		p.SetState(374)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametros_initContext is an interface to support dynamic dispatch.
type ILista_parametros_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_parametros_initContext differentiates from other interfaces.
	IsLista_parametros_initContext()
}

type Lista_parametros_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametros_initContext() *Lista_parametros_initContext {
	var p = new(Lista_parametros_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros_init
	return p
}

func InitEmptyLista_parametros_initContext(p *Lista_parametros_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros_init
}

func (*Lista_parametros_initContext) IsLista_parametros_initContext() {}

func NewLista_parametros_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametros_initContext {
	var p = new(Lista_parametros_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_parametros_init

	return p
}

func (s *Lista_parametros_initContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametros_initContext) CopyAll(ctx *Lista_parametros_initContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_parametros_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametros_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaParametrosInitContext struct {
	Lista_parametros_initContext
}

func NewListaParametrosInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaParametrosInitContext {
	var p = new(ListaParametrosInitContext)

	InitEmptyLista_parametros_initContext(&p.Lista_parametros_initContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_parametros_initContext))

	return p
}

func (s *ListaParametrosInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParametrosInitContext) AllParametros_init_struct() []IParametros_init_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametros_init_structContext); ok {
			len++
		}
	}

	tst := make([]IParametros_init_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametros_init_structContext); ok {
			tst[i] = t.(IParametros_init_structContext)
			i++
		}
	}

	return tst
}

func (s *ListaParametrosInitContext) Parametros_init_struct(i int) IParametros_init_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametros_init_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametros_init_structContext)
}

func (s *ListaParametrosInitContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaParametrosInitContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaParametrosInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaParametrosInit(s)
	}
}

func (s *ListaParametrosInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaParametrosInit(s)
	}
}

func (s *ListaParametrosInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaParametrosInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_parametros_init() (localctx ILista_parametros_initContext) {
	localctx = NewLista_parametros_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VGrammarRULE_lista_parametros_init)
	var _la int

	var _alt int

	localctx = NewListaParametrosInitContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Parametros_init_struct()
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(377)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(378)
				p.Parametros_init_struct()
			}

		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(384)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametros_init_structContext is an interface to support dynamic dispatch.
type IParametros_init_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametros_init_structContext differentiates from other interfaces.
	IsParametros_init_structContext()
}

type Parametros_init_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametros_init_structContext() *Parametros_init_structContext {
	var p = new(Parametros_init_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametros_init_struct
	return p
}

func InitEmptyParametros_init_structContext(p *Parametros_init_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametros_init_struct
}

func (*Parametros_init_structContext) IsParametros_init_structContext() {}

func NewParametros_init_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametros_init_structContext {
	var p = new(Parametros_init_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_parametros_init_struct

	return p
}

func (s *Parametros_init_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametros_init_structContext) CopyAll(ctx *Parametros_init_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Parametros_init_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametros_init_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametrosInitStructContext struct {
	Parametros_init_structContext
}

func NewParametrosInitStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametrosInitStructContext {
	var p = new(ParametrosInitStructContext)

	InitEmptyParametros_init_structContext(&p.Parametros_init_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parametros_init_structContext))

	return p
}

func (s *ParametrosInitStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosInitStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ParametrosInitStructContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *ParametrosInitStructContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParametrosInitStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParametrosInitStruct(s)
	}
}

func (s *ParametrosInitStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParametrosInitStruct(s)
	}
}

func (s *ParametrosInitStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParametrosInitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Parametros_init_struct() (localctx IParametros_init_structContext) {
	localctx = NewParametros_init_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VGrammarRULE_parametros_init_struct)
	localctx = NewParametrosInitStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_sliceContext is an interface to support dynamic dispatch.
type ILista_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_sliceContext differentiates from other interfaces.
	IsLista_sliceContext()
}

type Lista_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_sliceContext() *Lista_sliceContext {
	var p = new(Lista_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_slice
	return p
}

func InitEmptyLista_sliceContext(p *Lista_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_slice
}

func (*Lista_sliceContext) IsLista_sliceContext() {}

func NewLista_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_sliceContext {
	var p = new(Lista_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_slice

	return p
}

func (s *Lista_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_sliceContext) CopyAll(ctx *Lista_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaSliceContext struct {
	Lista_sliceContext
}

func NewListaSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaSliceContext {
	var p = new(ListaSliceContext)

	InitEmptyLista_sliceContext(&p.Lista_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_sliceContext))

	return p
}

func (s *ListaSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaSliceContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ListaSliceContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ListaSliceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListaSliceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListaSliceContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaSliceContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaSlice(s)
	}
}

func (s *ListaSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaSlice(s)
	}
}

func (s *ListaSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_slice() (localctx ILista_sliceContext) {
	localctx = NewLista_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VGrammarRULE_lista_slice)
	var _la int

	var _alt int

	localctx = NewListaSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&774058334216192) != 0 {
		{
			p.SetState(392)
			p.expr(0)
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(393)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(394)
					p.expr(0)
				}

			}
			p.SetState(399)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(402)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(405)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItem_sliceContext is an interface to support dynamic dispatch.
type IItem_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsItem_sliceContext differentiates from other interfaces.
	IsItem_sliceContext()
}

type Item_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_sliceContext() *Item_sliceContext {
	var p = new(Item_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_item_slice
	return p
}

func InitEmptyItem_sliceContext(p *Item_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_item_slice
}

func (*Item_sliceContext) IsItem_sliceContext() {}

func NewItem_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_sliceContext {
	var p = new(Item_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_item_slice

	return p
}

func (s *Item_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_sliceContext) CopyAll(ctx *Item_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Item_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ItemSliceContext struct {
	Item_sliceContext
}

func NewItemSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ItemSliceContext {
	var p = new(ItemSliceContext)

	InitEmptyItem_sliceContext(&p.Item_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Item_sliceContext))

	return p
}

func (s *ItemSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemSliceContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *ItemSliceContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *ItemSliceContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *ItemSliceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ItemSliceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ItemSliceContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *ItemSliceContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *ItemSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterItemSlice(s)
	}
}

func (s *ItemSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitItemSlice(s)
	}
}

func (s *ItemSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitItemSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Item_slice() (localctx IItem_sliceContext) {
	localctx = NewItem_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VGrammarRULE_item_slice)
	var _alt int

	localctx = NewItemSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.PatronId()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(408)
				p.Match(VGrammarCORCHETE_IZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(409)
				p.expr(0)
			}
			{
				p.SetState(410)
				p.Match(VGrammarCORCHETE_DER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProp_sliceContext is an interface to support dynamic dispatch.
type IProp_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsProp_sliceContext differentiates from other interfaces.
	IsProp_sliceContext()
}

type Prop_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProp_sliceContext() *Prop_sliceContext {
	var p = new(Prop_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_prop_slice
	return p
}

func InitEmptyProp_sliceContext(p *Prop_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_prop_slice
}

func (*Prop_sliceContext) IsProp_sliceContext() {}

func NewProp_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prop_sliceContext {
	var p = new(Prop_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_prop_slice

	return p
}

func (s *Prop_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Prop_sliceContext) CopyAll(ctx *Prop_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Prop_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prop_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropSliceContext struct {
	Prop_sliceContext
}

func NewPropSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropSliceContext {
	var p = new(PropSliceContext)

	InitEmptyProp_sliceContext(&p.Prop_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Prop_sliceContext))

	return p
}

func (s *PropSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *PropSliceContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *PropSliceContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *PropSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropSlice(s)
	}
}

func (s *PropSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropSlice(s)
	}
}

func (s *PropSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Prop_slice() (localctx IProp_sliceContext) {
	localctx = NewProp_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VGrammarRULE_prop_slice)
	localctx = NewPropSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Item_slice()
	}
	{
		p.SetState(417)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.PatronId()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFun_sliceContext is an interface to support dynamic dispatch.
type IFun_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFun_sliceContext differentiates from other interfaces.
	IsFun_sliceContext()
}

type Fun_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_sliceContext() *Fun_sliceContext {
	var p = new(Fun_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_fun_slice
	return p
}

func InitEmptyFun_sliceContext(p *Fun_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_fun_slice
}

func (*Fun_sliceContext) IsFun_sliceContext() {}

func NewFun_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_sliceContext {
	var p = new(Fun_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_fun_slice

	return p
}

func (s *Fun_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_sliceContext) CopyAll(ctx *Fun_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Fun_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionSliceContext struct {
	Fun_sliceContext
}

func NewFuncionSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionSliceContext {
	var p = new(FuncionSliceContext)

	InitEmptyFun_sliceContext(&p.Fun_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Fun_sliceContext))

	return p
}

func (s *FuncionSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *FuncionSliceContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *FuncionSliceContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *FuncionSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionSlice(s)
	}
}

func (s *FuncionSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionSlice(s)
	}
}

func (s *FuncionSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Fun_slice() (localctx IFun_sliceContext) {
	localctx = NewFun_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VGrammarRULE_fun_slice)
	localctx = NewFuncionSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Item_slice()
	}
	{
		p.SetState(421)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.Llamar_funcion()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	RW_INT() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	Tipos_slices() ITipos_slicesContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TipoContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *TipoContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *TipoContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *TipoContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *TipoContext) Tipos_slices() ITipos_slicesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipos_slicesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipos_slicesContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VGrammarRULE_tipo)
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(425)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_FLOAT64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(426)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(427)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(428)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarCORCHETE_IZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(429)
			p.Tipos_slices()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipos_slicesContext is an interface to support dynamic dispatch.
type ITipos_slicesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipos_slicesContext differentiates from other interfaces.
	IsTipos_slicesContext()
}

type Tipos_slicesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipos_slicesContext() *Tipos_slicesContext {
	var p = new(Tipos_slicesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipos_slices
	return p
}

func InitEmptyTipos_slicesContext(p *Tipos_slicesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipos_slices
}

func (*Tipos_slicesContext) IsTipos_slicesContext() {}

func NewTipos_slicesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipos_slicesContext {
	var p = new(Tipos_slicesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_tipos_slices

	return p
}

func (s *Tipos_slicesContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipos_slicesContext) CopyAll(ctx *Tipos_slicesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Tipos_slicesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipos_slicesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorSimpleContext struct {
	Tipos_slicesContext
}

func NewVectorSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorSimpleContext {
	var p = new(VectorSimpleContext)

	InitEmptyTipos_slicesContext(&p.Tipos_slicesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipos_slicesContext))

	return p
}

func (s *VectorSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorSimpleContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, 0)
}

func (s *VectorSimpleContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, 0)
}

func (s *VectorSimpleContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VectorSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorSimple(s)
	}
}

func (s *VectorSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorSimple(s)
	}
}

func (s *VectorSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrizDobleContext struct {
	Tipos_slicesContext
}

func NewMatrizDobleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrizDobleContext {
	var p = new(MatrizDobleContext)

	InitEmptyTipos_slicesContext(&p.Tipos_slicesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipos_slicesContext))

	return p
}

func (s *MatrizDobleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrizDobleContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *MatrizDobleContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *MatrizDobleContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *MatrizDobleContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *MatrizDobleContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *MatrizDobleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMatrizDoble(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Tipos_slices() (localctx ITipos_slicesContext) {
	localctx = NewTipos_slicesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, VGrammarRULE_tipos_slices)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Tipo()
		}

	case 2:
		localctx = NewMatrizDobleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Tipo()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralExpContext struct {
	ExprContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLiteralExp(s)
	}
}

func (s *LiteralExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLiteralExp(s)
	}
}

func (s *LiteralExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLiteralExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExpContext struct {
	ExprContext
}

func NewIdExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpContext {
	var p = new(IdExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *IdExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIdExp(s)
	}
}

func (s *IdExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIdExp(s)
	}
}

func (s *IdExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIdExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CrearStructExpContext struct {
	ExprContext
}

func NewCrearStructExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CrearStructExpContext {
	var p = new(CrearStructExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CrearStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrearStructExpContext) Crear_struct() ICrear_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICrear_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICrear_structContext)
}

func (s *CrearStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterCrearStructExp(s)
	}
}

func (s *CrearStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitCrearStructExp(s)
	}
}

func (s *CrearStructExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitCrearStructExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioExpContext struct {
	ExprContext
	op antlr.Token
}

func NewUnarioExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioExpContext {
	var p = new(UnarioExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnarioExpContext) GetOp() antlr.Token { return s.op }

func (s *UnarioExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnarioExpContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_NOT, 0)
}

func (s *UnarioExpContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *UnarioExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterUnarioExp(s)
	}
}

func (s *UnarioExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitUnarioExp(s)
	}
}

func (s *UnarioExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitUnarioExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ItemSliceExpContext struct {
	ExprContext
}

func NewItemSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ItemSliceExpContext {
	var p = new(ItemSliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ItemSliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemSliceExpContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *ItemSliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterItemSliceExp(s)
	}
}

func (s *ItemSliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitItemSliceExp(s)
	}
}

func (s *ItemSliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitItemSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorFuncExpContext struct {
	ExprContext
}

func NewVectorFuncExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorFuncExpContext {
	var p = new(VectorFuncExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorFuncExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorFuncExpContext) Fun_slice() IFun_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFun_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFun_sliceContext)
}

func (s *VectorFuncExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorFuncExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceExpContext struct {
	ExprContext
}

func NewSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceExpContext {
	var p = new(SliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExpContext) Lista_slice() ILista_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_sliceContext)
}

func (s *SliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSliceExp(s)
	}
}

func (s *SliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSliceExp(s)
	}
}

func (s *SliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropSliceExpContext struct {
	ExprContext
}

func NewPropSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropSliceExpContext {
	var p = new(PropSliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PropSliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropSliceExpContext) Prop_slice() IProp_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProp_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProp_sliceContext)
}

func (s *PropSliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropSliceExp(s)
	}
}

func (s *PropSliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropSliceExp(s)
	}
}

func (s *PropSliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LlamarFuncionExpContext struct {
	ExprContext
}

func NewLlamarFuncionExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamarFuncionExpContext {
	var p = new(LlamarFuncionExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LlamarFuncionExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamarFuncionExpContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *LlamarFuncionExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLlamarFuncionExp(s)
	}
}

func (s *LlamarFuncionExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLlamarFuncionExp(s)
	}
}

func (s *LlamarFuncionExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLlamarFuncionExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinarioExpContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewBinarioExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinarioExpContext {
	var p = new(BinarioExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BinarioExpContext) GetOp() antlr.Token { return s.op }

func (s *BinarioExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinarioExpContext) GetLeft() IExprContext { return s.left }

func (s *BinarioExpContext) GetRight() IExprContext { return s.right }

func (s *BinarioExpContext) SetLeft(v IExprContext) { s.left = v }

func (s *BinarioExpContext) SetRight(v IExprContext) { s.right = v }

func (s *BinarioExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinarioExpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinarioExpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinarioExpContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *BinarioExpContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIV, 0)
}

func (s *BinarioExpContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MOD, 0)
}

func (s *BinarioExpContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUMA, 0)
}

func (s *BinarioExpContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *BinarioExpContext) OP_MENORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORQ, 0)
}

func (s *BinarioExpContext) OP_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORIGUAL, 0)
}

func (s *BinarioExpContext) OP_MAYORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORQ, 0)
}

func (s *BinarioExpContext) OP_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORIGUAL, 0)
}

func (s *BinarioExpContext) OP_IGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_IGUAL, 0)
}

func (s *BinarioExpContext) OP_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIFERENTE, 0)
}

func (s *BinarioExpContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_AND, 0)
}

func (s *BinarioExpContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_OR, 0)
}

func (s *BinarioExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBinarioExp(s)
	}
}

func (s *BinarioExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBinarioExp(s)
	}
}

func (s *BinarioExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBinarioExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentecisExpContext struct {
	ExprContext
}

func NewParentecisExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentecisExpContext {
	var p = new(ParentecisExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParentecisExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentecisExpContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *ParentecisExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParentecisExpContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *ParentecisExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParentecisExp(s)
	}
}

func (s *ParentecisExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParentecisExp(s)
	}
}

func (s *ParentecisExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParentecisExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *VGrammar) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, VGrammarRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentecisExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(443)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.expr(0)
		}
		{
			p.SetState(445)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLlamarFuncionExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(447)
			p.Llamar_funcion()
		}

	case 3:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(448)
			p.PatronId()
		}

	case 4:
		localctx = NewItemSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(449)
			p.Item_slice()
		}

	case 5:
		localctx = NewPropSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(450)
			p.Prop_slice()
		}

	case 6:
		localctx = NewVectorFuncExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(451)
			p.Fun_slice()
		}

	case 7:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(452)
			p.Literal()
		}

	case 8:
		localctx = NewSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(453)
			p.Lista_slice()
		}

	case 9:
		localctx = NewCrearStructExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(454)
			p.Crear_struct()
		}

	case 10:
		localctx = NewUnarioExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(455)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_RESTA || _la == VGrammarOP_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(456)
			p.expr(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(459)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(460)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(461)

					var _x = p.expr(7)

					localctx.(*BinarioExpContext).right = _x
				}

			case 2:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(463)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_SUMA || _la == VGrammarOP_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(464)

					var _x = p.expr(6)

					localctx.(*BinarioExpContext).right = _x
				}

			case 3:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(465)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(466)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(467)

					var _x = p.expr(5)

					localctx.(*BinarioExpContext).right = _x
				}

			case 4:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(468)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(469)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_IGUAL || _la == VGrammarOP_DIFERENTE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(470)

					var _x = p.expr(4)

					localctx.(*BinarioExpContext).right = _x
				}

			case 5:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(471)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(472)

					var _m = p.Match(VGrammarOP_AND)

					localctx.(*BinarioExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(473)

					var _x = p.expr(3)

					localctx.(*BinarioExpContext).right = _x
				}

			case 6:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(474)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(475)

					var _m = p.Match(VGrammarOP_OR)

					localctx.(*BinarioExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(476)

					var _x = p.expr(2)

					localctx.(*BinarioExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLiteralContext struct {
	LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarBOOL_LITERAL, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilLiteralContext struct {
	LiteralContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarNIL_LITERAL, 0)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VGrammarRULE_literal)
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(VGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(483)
			p.Match(VGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(484)
			p.Match(VGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarBOOL_LITERAL:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(485)
			p.Match(VGrammarBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarNIL_LITERAL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(486)
			p.Match(VGrammarNIL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatronIdContext is an interface to support dynamic dispatch.
type IPatronIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPatronIdContext differentiates from other interfaces.
	IsPatronIdContext()
}

type PatronIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatronIdContext() *PatronIdContext {
	var p = new(PatronIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_patronId
	return p
}

func InitEmptyPatronIdContext(p *PatronIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_patronId
}

func (*PatronIdContext) IsPatronIdContext() {}

func NewPatronIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatronIdContext {
	var p = new(PatronIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_patronId

	return p
}

func (s *PatronIdContext) GetParser() antlr.Parser { return s.parser }

func (s *PatronIdContext) CopyAll(ctx *PatronIdContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PatronIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatronIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ID_PatronContext struct {
	PatronIdContext
}

func NewID_PatronContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ID_PatronContext {
	var p = new(ID_PatronContext)

	InitEmptyPatronIdContext(&p.PatronIdContext)
	p.parser = parser
	p.CopyAll(ctx.(*PatronIdContext))

	return p
}

func (s *ID_PatronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ID_PatronContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ID_PatronContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ID_PatronContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *ID_PatronContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *ID_PatronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterID_Patron(s)
	}
}

func (s *ID_PatronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitID_Patron(s)
	}
}

func (s *ID_PatronContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitID_Patron(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) PatronId() (localctx IPatronIdContext) {
	localctx = NewPatronIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, VGrammarRULE_patronId)
	var _alt int

	localctx = NewID_PatronContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(490)
				p.Match(VGrammarPUNTO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(491)
				p.Match(VGrammarID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 33:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VGrammar) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
