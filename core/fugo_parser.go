// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package core // Fugo
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FugoParser struct {
	*antlr.BaseParser
}

var fugoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fugoParserInit() {
	staticData := &fugoParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'['", "']'", "'{'", "'}'", "'#{'", "'''", "'`'",
		"'~'", "'~@'", "'^'", "'@'", "'#'", "'#('", "'#^'", "'#''", "'#+'",
		"'#_'", "':'", "", "", "", "", "", "", "", "", "", "'nil'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "STRING", "FLOAT", "HEX", "BIN", "LONG", "BIGN", "CHAR_U",
		"CHAR_NAMED", "CHAR_ANY", "NIL", "BOOLEAN", "SYMBOL", "NS_SYMBOL", "PARAM_NAME",
		"TRASH",
	}
	staticData.ruleNames = []string{
		"file_", "form", "forms", "list_", "vector", "map_", "set_", "reader_macro",
		"quote", "backtick", "unquote", "unquote_splicing", "tag", "deref",
		"gensym", "lambda_", "meta_data", "var_quote", "host_expr", "discard",
		"dispatch", "regex", "literal", "string_", "hex_", "bin_", "bign", "number",
		"character", "named_char", "any_char", "u_hex_quad", "nil_", "keyword",
		"simple_keyword", "macro_keyword", "symbol", "simple_sym", "ns_symbol",
		"param_name",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 256, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 5, 0, 82, 8, 0, 10, 0,
		12, 0, 85, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 94, 8,
		1, 1, 2, 5, 2, 97, 8, 2, 10, 2, 12, 2, 100, 9, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 114, 8, 5, 10,
		5, 12, 5, 117, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 140, 8, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 166, 8, 15, 10, 15, 12, 15, 169,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 178, 8,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 205, 8, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 220, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 225, 8, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 237, 8, 33,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 248,
		8, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 0, 0, 40, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 0, 0, 253, 0, 83, 1, 0, 0, 0, 2, 93, 1, 0, 0, 0, 4, 98, 1, 0, 0, 0,
		6, 101, 1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12, 120, 1,
		0, 0, 0, 14, 139, 1, 0, 0, 0, 16, 141, 1, 0, 0, 0, 18, 144, 1, 0, 0, 0,
		20, 147, 1, 0, 0, 0, 22, 150, 1, 0, 0, 0, 24, 153, 1, 0, 0, 0, 26, 157,
		1, 0, 0, 0, 28, 160, 1, 0, 0, 0, 30, 163, 1, 0, 0, 0, 32, 172, 1, 0, 0,
		0, 34, 179, 1, 0, 0, 0, 36, 182, 1, 0, 0, 0, 38, 186, 1, 0, 0, 0, 40, 189,
		1, 0, 0, 0, 42, 193, 1, 0, 0, 0, 44, 204, 1, 0, 0, 0, 46, 206, 1, 0, 0,
		0, 48, 208, 1, 0, 0, 0, 50, 210, 1, 0, 0, 0, 52, 212, 1, 0, 0, 0, 54, 219,
		1, 0, 0, 0, 56, 224, 1, 0, 0, 0, 58, 226, 1, 0, 0, 0, 60, 228, 1, 0, 0,
		0, 62, 230, 1, 0, 0, 0, 64, 232, 1, 0, 0, 0, 66, 236, 1, 0, 0, 0, 68, 238,
		1, 0, 0, 0, 70, 241, 1, 0, 0, 0, 72, 247, 1, 0, 0, 0, 74, 249, 1, 0, 0,
		0, 76, 251, 1, 0, 0, 0, 78, 253, 1, 0, 0, 0, 80, 82, 3, 2, 1, 0, 81, 80,
		1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 0, 0, 1, 87, 1, 1, 0,
		0, 0, 88, 94, 3, 44, 22, 0, 89, 94, 3, 6, 3, 0, 90, 94, 3, 8, 4, 0, 91,
		94, 3, 10, 5, 0, 92, 94, 3, 14, 7, 0, 93, 88, 1, 0, 0, 0, 93, 89, 1, 0,
		0, 0, 93, 90, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 3,
		1, 0, 0, 0, 95, 97, 3, 2, 1, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0,
		98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 5, 1, 0, 0, 0, 100, 98, 1,
		0, 0, 0, 101, 102, 5, 1, 0, 0, 102, 103, 3, 4, 2, 0, 103, 104, 5, 2, 0,
		0, 104, 7, 1, 0, 0, 0, 105, 106, 5, 3, 0, 0, 106, 107, 3, 4, 2, 0, 107,
		108, 5, 4, 0, 0, 108, 9, 1, 0, 0, 0, 109, 115, 5, 5, 0, 0, 110, 111, 3,
		2, 1, 0, 111, 112, 3, 2, 1, 0, 112, 114, 1, 0, 0, 0, 113, 110, 1, 0, 0,
		0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 5, 6, 0, 0, 119, 11, 1,
		0, 0, 0, 120, 121, 5, 7, 0, 0, 121, 122, 3, 4, 2, 0, 122, 123, 5, 6, 0,
		0, 123, 13, 1, 0, 0, 0, 124, 140, 3, 30, 15, 0, 125, 140, 3, 32, 16, 0,
		126, 140, 3, 42, 21, 0, 127, 140, 3, 34, 17, 0, 128, 140, 3, 36, 18, 0,
		129, 140, 3, 12, 6, 0, 130, 140, 3, 24, 12, 0, 131, 140, 3, 38, 19, 0,
		132, 140, 3, 40, 20, 0, 133, 140, 3, 26, 13, 0, 134, 140, 3, 16, 8, 0,
		135, 140, 3, 18, 9, 0, 136, 140, 3, 20, 10, 0, 137, 140, 3, 22, 11, 0,
		138, 140, 3, 28, 14, 0, 139, 124, 1, 0, 0, 0, 139, 125, 1, 0, 0, 0, 139,
		126, 1, 0, 0, 0, 139, 127, 1, 0, 0, 0, 139, 128, 1, 0, 0, 0, 139, 129,
		1, 0, 0, 0, 139, 130, 1, 0, 0, 0, 139, 131, 1, 0, 0, 0, 139, 132, 1, 0,
		0, 0, 139, 133, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0,
		139, 136, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140,
		15, 1, 0, 0, 0, 141, 142, 5, 8, 0, 0, 142, 143, 3, 2, 1, 0, 143, 17, 1,
		0, 0, 0, 144, 145, 5, 9, 0, 0, 145, 146, 3, 2, 1, 0, 146, 19, 1, 0, 0,
		0, 147, 148, 5, 10, 0, 0, 148, 149, 3, 2, 1, 0, 149, 21, 1, 0, 0, 0, 150,
		151, 5, 11, 0, 0, 151, 152, 3, 2, 1, 0, 152, 23, 1, 0, 0, 0, 153, 154,
		5, 12, 0, 0, 154, 155, 3, 2, 1, 0, 155, 156, 3, 2, 1, 0, 156, 25, 1, 0,
		0, 0, 157, 158, 5, 13, 0, 0, 158, 159, 3, 2, 1, 0, 159, 27, 1, 0, 0, 0,
		160, 161, 5, 32, 0, 0, 161, 162, 5, 14, 0, 0, 162, 29, 1, 0, 0, 0, 163,
		167, 5, 15, 0, 0, 164, 166, 3, 2, 1, 0, 165, 164, 1, 0, 0, 0, 166, 169,
		1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0,
		0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 5, 2, 0, 0, 171, 31, 1, 0, 0, 0,
		172, 177, 5, 16, 0, 0, 173, 174, 3, 10, 5, 0, 174, 175, 3, 2, 1, 0, 175,
		178, 1, 0, 0, 0, 176, 178, 3, 2, 1, 0, 177, 173, 1, 0, 0, 0, 177, 176,
		1, 0, 0, 0, 178, 33, 1, 0, 0, 0, 179, 180, 5, 17, 0, 0, 180, 181, 3, 72,
		36, 0, 181, 35, 1, 0, 0, 0, 182, 183, 5, 18, 0, 0, 183, 184, 3, 2, 1, 0,
		184, 185, 3, 2, 1, 0, 185, 37, 1, 0, 0, 0, 186, 187, 5, 19, 0, 0, 187,
		188, 3, 2, 1, 0, 188, 39, 1, 0, 0, 0, 189, 190, 5, 14, 0, 0, 190, 191,
		3, 72, 36, 0, 191, 192, 3, 2, 1, 0, 192, 41, 1, 0, 0, 0, 193, 194, 5, 14,
		0, 0, 194, 195, 3, 46, 23, 0, 195, 43, 1, 0, 0, 0, 196, 205, 3, 46, 23,
		0, 197, 205, 3, 54, 27, 0, 198, 205, 3, 56, 28, 0, 199, 205, 3, 64, 32,
		0, 200, 205, 5, 31, 0, 0, 201, 205, 3, 66, 33, 0, 202, 205, 3, 72, 36,
		0, 203, 205, 3, 78, 39, 0, 204, 196, 1, 0, 0, 0, 204, 197, 1, 0, 0, 0,
		204, 198, 1, 0, 0, 0, 204, 199, 1, 0, 0, 0, 204, 200, 1, 0, 0, 0, 204,
		201, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 45, 1,
		0, 0, 0, 206, 207, 5, 21, 0, 0, 207, 47, 1, 0, 0, 0, 208, 209, 5, 23, 0,
		0, 209, 49, 1, 0, 0, 0, 210, 211, 5, 24, 0, 0, 211, 51, 1, 0, 0, 0, 212,
		213, 5, 26, 0, 0, 213, 53, 1, 0, 0, 0, 214, 220, 5, 22, 0, 0, 215, 220,
		3, 48, 24, 0, 216, 220, 3, 50, 25, 0, 217, 220, 3, 52, 26, 0, 218, 220,
		5, 25, 0, 0, 219, 214, 1, 0, 0, 0, 219, 215, 1, 0, 0, 0, 219, 216, 1, 0,
		0, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 55, 1, 0, 0, 0,
		221, 225, 3, 58, 29, 0, 222, 225, 3, 62, 31, 0, 223, 225, 3, 60, 30, 0,
		224, 221, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 223, 1, 0, 0, 0, 225,
		57, 1, 0, 0, 0, 226, 227, 5, 28, 0, 0, 227, 59, 1, 0, 0, 0, 228, 229, 5,
		29, 0, 0, 229, 61, 1, 0, 0, 0, 230, 231, 5, 27, 0, 0, 231, 63, 1, 0, 0,
		0, 232, 233, 5, 30, 0, 0, 233, 65, 1, 0, 0, 0, 234, 237, 3, 70, 35, 0,
		235, 237, 3, 68, 34, 0, 236, 234, 1, 0, 0, 0, 236, 235, 1, 0, 0, 0, 237,
		67, 1, 0, 0, 0, 238, 239, 5, 20, 0, 0, 239, 240, 3, 72, 36, 0, 240, 69,
		1, 0, 0, 0, 241, 242, 5, 20, 0, 0, 242, 243, 5, 20, 0, 0, 243, 244, 3,
		72, 36, 0, 244, 71, 1, 0, 0, 0, 245, 248, 3, 76, 38, 0, 246, 248, 3, 74,
		37, 0, 247, 245, 1, 0, 0, 0, 247, 246, 1, 0, 0, 0, 248, 73, 1, 0, 0, 0,
		249, 250, 5, 32, 0, 0, 250, 75, 1, 0, 0, 0, 251, 252, 5, 33, 0, 0, 252,
		77, 1, 0, 0, 0, 253, 254, 5, 34, 0, 0, 254, 79, 1, 0, 0, 0, 12, 83, 93,
		98, 115, 139, 167, 177, 204, 219, 224, 236, 247,
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

// FugoParserInit initializes any static state used to implement FugoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFugoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FugoParserInit() {
	staticData := &fugoParserStaticData
	staticData.once.Do(fugoParserInit)
}

// NewFugoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFugoParser(input antlr.TokenStream) *FugoParser {
	FugoParserInit()
	this := new(FugoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fugoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// FugoParser tokens.
const (
	FugoParserEOF        = antlr.TokenEOF
	FugoParserT__0       = 1
	FugoParserT__1       = 2
	FugoParserT__2       = 3
	FugoParserT__3       = 4
	FugoParserT__4       = 5
	FugoParserT__5       = 6
	FugoParserT__6       = 7
	FugoParserT__7       = 8
	FugoParserT__8       = 9
	FugoParserT__9       = 10
	FugoParserT__10      = 11
	FugoParserT__11      = 12
	FugoParserT__12      = 13
	FugoParserT__13      = 14
	FugoParserT__14      = 15
	FugoParserT__15      = 16
	FugoParserT__16      = 17
	FugoParserT__17      = 18
	FugoParserT__18      = 19
	FugoParserT__19      = 20
	FugoParserSTRING     = 21
	FugoParserFLOAT      = 22
	FugoParserHEX        = 23
	FugoParserBIN        = 24
	FugoParserLONG       = 25
	FugoParserBIGN       = 26
	FugoParserCHAR_U     = 27
	FugoParserCHAR_NAMED = 28
	FugoParserCHAR_ANY   = 29
	FugoParserNIL        = 30
	FugoParserBOOLEAN    = 31
	FugoParserSYMBOL     = 32
	FugoParserNS_SYMBOL  = 33
	FugoParserPARAM_NAME = 34
	FugoParserTRASH      = 35
)

// FugoParser rules.
const (
	FugoParserRULE_file_            = 0
	FugoParserRULE_form             = 1
	FugoParserRULE_forms            = 2
	FugoParserRULE_list_            = 3
	FugoParserRULE_vector           = 4
	FugoParserRULE_map_             = 5
	FugoParserRULE_set_             = 6
	FugoParserRULE_reader_macro     = 7
	FugoParserRULE_quote            = 8
	FugoParserRULE_backtick         = 9
	FugoParserRULE_unquote          = 10
	FugoParserRULE_unquote_splicing = 11
	FugoParserRULE_tag              = 12
	FugoParserRULE_deref            = 13
	FugoParserRULE_gensym           = 14
	FugoParserRULE_lambda_          = 15
	FugoParserRULE_meta_data        = 16
	FugoParserRULE_var_quote        = 17
	FugoParserRULE_host_expr        = 18
	FugoParserRULE_discard          = 19
	FugoParserRULE_dispatch         = 20
	FugoParserRULE_regex            = 21
	FugoParserRULE_literal          = 22
	FugoParserRULE_string_          = 23
	FugoParserRULE_hex_             = 24
	FugoParserRULE_bin_             = 25
	FugoParserRULE_bign             = 26
	FugoParserRULE_number           = 27
	FugoParserRULE_character        = 28
	FugoParserRULE_named_char       = 29
	FugoParserRULE_any_char         = 30
	FugoParserRULE_u_hex_quad       = 31
	FugoParserRULE_nil_             = 32
	FugoParserRULE_keyword          = 33
	FugoParserRULE_simple_keyword   = 34
	FugoParserRULE_macro_keyword    = 35
	FugoParserRULE_symbol           = 36
	FugoParserRULE_simple_sym       = 37
	FugoParserRULE_ns_symbol        = 38
	FugoParserRULE_param_name       = 39
)

// IFile_Context is an interface to support dynamic dispatch.
type IFile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_Context differentiates from other interfaces.
	IsFile_Context()
}

type File_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_Context() *File_Context {
	var p = new(File_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_file_
	return p
}

func (*File_Context) IsFile_Context() {}

func NewFile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_Context {
	var p = new(File_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_file_

	return p
}

func (s *File_Context) GetParser() antlr.Parser { return s.parser }

func (s *File_Context) EOF() antlr.TerminalNode {
	return s.GetToken(FugoParserEOF, 0)
}

func (s *File_Context) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *File_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterFile_(s)
	}
}

func (s *File_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitFile_(s)
	}
}

func (p *FugoParser) File_() (localctx IFile_Context) {
	this := p
	_ = this

	localctx = NewFile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FugoParserRULE_file_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359738282) != 0 {
		{
			p.SetState(80)
			p.Form()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(FugoParserEOF)
	}

	return localctx
}

// IFormContext is an interface to support dynamic dispatch.
type IFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormContext differentiates from other interfaces.
	IsFormContext()
}

type FormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormContext() *FormContext {
	var p = new(FormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_form
	return p
}

func (*FormContext) IsFormContext() {}

func NewFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormContext {
	var p = new(FormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_form

	return p
}

func (s *FormContext) GetParser() antlr.Parser { return s.parser }

func (s *FormContext) Literal() ILiteralContext {
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

func (s *FormContext) List_() IList_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *FormContext) Vector() IVectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *FormContext) Map_() IMap_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *FormContext) Reader_macro() IReader_macroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReader_macroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReader_macroContext)
}

func (s *FormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterForm(s)
	}
}

func (s *FormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitForm(s)
	}
}

func (p *FugoParser) Form() (localctx IFormContext) {
	this := p
	_ = this

	localctx = NewFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FugoParserRULE_form)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.List_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Vector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Map_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.Reader_macro()
		}

	}

	return localctx
}

// IFormsContext is an interface to support dynamic dispatch.
type IFormsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormsContext differentiates from other interfaces.
	IsFormsContext()
}

type FormsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormsContext() *FormsContext {
	var p = new(FormsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_forms
	return p
}

func (*FormsContext) IsFormsContext() {}

func NewFormsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormsContext {
	var p = new(FormsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_forms

	return p
}

func (s *FormsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormsContext) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *FormsContext) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *FormsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterForms(s)
	}
}

func (s *FormsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitForms(s)
	}
}

func (p *FugoParser) Forms() (localctx IFormsContext) {
	this := p
	_ = this

	localctx = NewFormsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FugoParserRULE_forms)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359738282) != 0 {
		{
			p.SetState(95)
			p.Form()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) Forms() IFormsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *FugoParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FugoParserRULE_list_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(FugoParserT__0)
	}
	{
		p.SetState(102)
		p.Forms()
	}
	{
		p.SetState(103)
		p.Match(FugoParserT__1)
	}

	return localctx
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_vector
	return p
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) Forms() IFormsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *FugoParser) Vector() (localctx IVectorContext) {
	this := p
	_ = this

	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FugoParserRULE_vector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(FugoParserT__2)
	}
	{
		p.SetState(106)
		p.Forms()
	}
	{
		p.SetState(107)
		p.Match(FugoParserT__3)
	}

	return localctx
}

// IMap_Context is an interface to support dynamic dispatch.
type IMap_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_Context differentiates from other interfaces.
	IsMap_Context()
}

type Map_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_Context() *Map_Context {
	var p = new(Map_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_map_
	return p
}

func (*Map_Context) IsMap_Context() {}

func NewMap_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_Context {
	var p = new(Map_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_map_

	return p
}

func (s *Map_Context) GetParser() antlr.Parser { return s.parser }

func (s *Map_Context) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *Map_Context) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *Map_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterMap_(s)
	}
}

func (s *Map_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitMap_(s)
	}
}

func (p *FugoParser) Map_() (localctx IMap_Context) {
	this := p
	_ = this

	localctx = NewMap_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FugoParserRULE_map_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(FugoParserT__4)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359738282) != 0 {
		{
			p.SetState(110)
			p.Form()
		}
		{
			p.SetState(111)
			p.Form()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(FugoParserT__5)
	}

	return localctx
}

// ISet_Context is an interface to support dynamic dispatch.
type ISet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_Context differentiates from other interfaces.
	IsSet_Context()
}

type Set_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_Context() *Set_Context {
	var p = new(Set_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_set_
	return p
}

func (*Set_Context) IsSet_Context() {}

func NewSet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_Context {
	var p = new(Set_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_set_

	return p
}

func (s *Set_Context) GetParser() antlr.Parser { return s.parser }

func (s *Set_Context) Forms() IFormsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *Set_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterSet_(s)
	}
}

func (s *Set_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitSet_(s)
	}
}

func (p *FugoParser) Set_() (localctx ISet_Context) {
	this := p
	_ = this

	localctx = NewSet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FugoParserRULE_set_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(FugoParserT__6)
	}
	{
		p.SetState(121)
		p.Forms()
	}
	{
		p.SetState(122)
		p.Match(FugoParserT__5)
	}

	return localctx
}

// IReader_macroContext is an interface to support dynamic dispatch.
type IReader_macroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReader_macroContext differentiates from other interfaces.
	IsReader_macroContext()
}

type Reader_macroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReader_macroContext() *Reader_macroContext {
	var p = new(Reader_macroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_reader_macro
	return p
}

func (*Reader_macroContext) IsReader_macroContext() {}

func NewReader_macroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reader_macroContext {
	var p = new(Reader_macroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_reader_macro

	return p
}

func (s *Reader_macroContext) GetParser() antlr.Parser { return s.parser }

func (s *Reader_macroContext) Lambda_() ILambda_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambda_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambda_Context)
}

func (s *Reader_macroContext) Meta_data() IMeta_dataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_dataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_dataContext)
}

func (s *Reader_macroContext) Regex() IRegexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexContext)
}

func (s *Reader_macroContext) Var_quote() IVar_quoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_quoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_quoteContext)
}

func (s *Reader_macroContext) Host_expr() IHost_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHost_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHost_exprContext)
}

func (s *Reader_macroContext) Set_() ISet_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_Context)
}

func (s *Reader_macroContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Reader_macroContext) Discard() IDiscardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiscardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiscardContext)
}

func (s *Reader_macroContext) Dispatch() IDispatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDispatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDispatchContext)
}

func (s *Reader_macroContext) Deref() IDerefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDerefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDerefContext)
}

func (s *Reader_macroContext) Quote() IQuoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuoteContext)
}

func (s *Reader_macroContext) Backtick() IBacktickContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBacktickContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBacktickContext)
}

func (s *Reader_macroContext) Unquote() IUnquoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnquoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnquoteContext)
}

func (s *Reader_macroContext) Unquote_splicing() IUnquote_splicingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnquote_splicingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnquote_splicingContext)
}

func (s *Reader_macroContext) Gensym() IGensymContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGensymContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGensymContext)
}

func (s *Reader_macroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reader_macroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reader_macroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterReader_macro(s)
	}
}

func (s *Reader_macroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitReader_macro(s)
	}
}

func (p *FugoParser) Reader_macro() (localctx IReader_macroContext) {
	this := p
	_ = this

	localctx = NewReader_macroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FugoParserRULE_reader_macro)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Lambda_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Meta_data()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Regex()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Var_quote()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Host_expr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Set_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.Tag()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(131)
			p.Discard()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(132)
			p.Dispatch()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(133)
			p.Deref()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(134)
			p.Quote()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(135)
			p.Backtick()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(136)
			p.Unquote()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(137)
			p.Unquote_splicing()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(138)
			p.Gensym()
		}

	}

	return localctx
}

// IQuoteContext is an interface to support dynamic dispatch.
type IQuoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoteContext differentiates from other interfaces.
	IsQuoteContext()
}

type QuoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoteContext() *QuoteContext {
	var p = new(QuoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_quote
	return p
}

func (*QuoteContext) IsQuoteContext() {}

func NewQuoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuoteContext {
	var p = new(QuoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_quote

	return p
}

func (s *QuoteContext) GetParser() antlr.Parser { return s.parser }

func (s *QuoteContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *QuoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterQuote(s)
	}
}

func (s *QuoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitQuote(s)
	}
}

func (p *FugoParser) Quote() (localctx IQuoteContext) {
	this := p
	_ = this

	localctx = NewQuoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FugoParserRULE_quote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FugoParserT__7)
	}
	{
		p.SetState(142)
		p.Form()
	}

	return localctx
}

// IBacktickContext is an interface to support dynamic dispatch.
type IBacktickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBacktickContext differentiates from other interfaces.
	IsBacktickContext()
}

type BacktickContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBacktickContext() *BacktickContext {
	var p = new(BacktickContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_backtick
	return p
}

func (*BacktickContext) IsBacktickContext() {}

func NewBacktickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BacktickContext {
	var p = new(BacktickContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_backtick

	return p
}

func (s *BacktickContext) GetParser() antlr.Parser { return s.parser }

func (s *BacktickContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *BacktickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BacktickContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BacktickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterBacktick(s)
	}
}

func (s *BacktickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitBacktick(s)
	}
}

func (p *FugoParser) Backtick() (localctx IBacktickContext) {
	this := p
	_ = this

	localctx = NewBacktickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FugoParserRULE_backtick)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(FugoParserT__8)
	}
	{
		p.SetState(145)
		p.Form()
	}

	return localctx
}

// IUnquoteContext is an interface to support dynamic dispatch.
type IUnquoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquoteContext differentiates from other interfaces.
	IsUnquoteContext()
}

type UnquoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquoteContext() *UnquoteContext {
	var p = new(UnquoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_unquote
	return p
}

func (*UnquoteContext) IsUnquoteContext() {}

func NewUnquoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnquoteContext {
	var p = new(UnquoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_unquote

	return p
}

func (s *UnquoteContext) GetParser() antlr.Parser { return s.parser }

func (s *UnquoteContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *UnquoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnquoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnquoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterUnquote(s)
	}
}

func (s *UnquoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitUnquote(s)
	}
}

func (p *FugoParser) Unquote() (localctx IUnquoteContext) {
	this := p
	_ = this

	localctx = NewUnquoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FugoParserRULE_unquote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(FugoParserT__9)
	}
	{
		p.SetState(148)
		p.Form()
	}

	return localctx
}

// IUnquote_splicingContext is an interface to support dynamic dispatch.
type IUnquote_splicingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquote_splicingContext differentiates from other interfaces.
	IsUnquote_splicingContext()
}

type Unquote_splicingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquote_splicingContext() *Unquote_splicingContext {
	var p = new(Unquote_splicingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_unquote_splicing
	return p
}

func (*Unquote_splicingContext) IsUnquote_splicingContext() {}

func NewUnquote_splicingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unquote_splicingContext {
	var p = new(Unquote_splicingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_unquote_splicing

	return p
}

func (s *Unquote_splicingContext) GetParser() antlr.Parser { return s.parser }

func (s *Unquote_splicingContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *Unquote_splicingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unquote_splicingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unquote_splicingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterUnquote_splicing(s)
	}
}

func (s *Unquote_splicingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitUnquote_splicing(s)
	}
}

func (p *FugoParser) Unquote_splicing() (localctx IUnquote_splicingContext) {
	this := p
	_ = this

	localctx = NewUnquote_splicingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FugoParserRULE_unquote_splicing)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(FugoParserT__10)
	}
	{
		p.SetState(151)
		p.Form()
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *TagContext) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *FugoParser) Tag() (localctx ITagContext) {
	this := p
	_ = this

	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FugoParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(FugoParserT__11)
	}
	{
		p.SetState(154)
		p.Form()
	}
	{
		p.SetState(155)
		p.Form()
	}

	return localctx
}

// IDerefContext is an interface to support dynamic dispatch.
type IDerefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerefContext differentiates from other interfaces.
	IsDerefContext()
}

type DerefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerefContext() *DerefContext {
	var p = new(DerefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_deref
	return p
}

func (*DerefContext) IsDerefContext() {}

func NewDerefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerefContext {
	var p = new(DerefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_deref

	return p
}

func (s *DerefContext) GetParser() antlr.Parser { return s.parser }

func (s *DerefContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DerefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterDeref(s)
	}
}

func (s *DerefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitDeref(s)
	}
}

func (p *FugoParser) Deref() (localctx IDerefContext) {
	this := p
	_ = this

	localctx = NewDerefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FugoParserRULE_deref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(FugoParserT__12)
	}
	{
		p.SetState(158)
		p.Form()
	}

	return localctx
}

// IGensymContext is an interface to support dynamic dispatch.
type IGensymContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGensymContext differentiates from other interfaces.
	IsGensymContext()
}

type GensymContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGensymContext() *GensymContext {
	var p = new(GensymContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_gensym
	return p
}

func (*GensymContext) IsGensymContext() {}

func NewGensymContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GensymContext {
	var p = new(GensymContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_gensym

	return p
}

func (s *GensymContext) GetParser() antlr.Parser { return s.parser }

func (s *GensymContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(FugoParserSYMBOL, 0)
}

func (s *GensymContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GensymContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GensymContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterGensym(s)
	}
}

func (s *GensymContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitGensym(s)
	}
}

func (p *FugoParser) Gensym() (localctx IGensymContext) {
	this := p
	_ = this

	localctx = NewGensymContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FugoParserRULE_gensym)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(FugoParserSYMBOL)
	}
	{
		p.SetState(161)
		p.Match(FugoParserT__13)
	}

	return localctx
}

// ILambda_Context is an interface to support dynamic dispatch.
type ILambda_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambda_Context differentiates from other interfaces.
	IsLambda_Context()
}

type Lambda_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambda_Context() *Lambda_Context {
	var p = new(Lambda_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_lambda_
	return p
}

func (*Lambda_Context) IsLambda_Context() {}

func NewLambda_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_Context {
	var p = new(Lambda_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_lambda_

	return p
}

func (s *Lambda_Context) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_Context) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *Lambda_Context) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *Lambda_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterLambda_(s)
	}
}

func (s *Lambda_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitLambda_(s)
	}
}

func (p *FugoParser) Lambda_() (localctx ILambda_Context) {
	this := p
	_ = this

	localctx = NewLambda_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FugoParserRULE_lambda_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(FugoParserT__14)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359738282) != 0 {
		{
			p.SetState(164)
			p.Form()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(FugoParserT__1)
	}

	return localctx
}

// IMeta_dataContext is an interface to support dynamic dispatch.
type IMeta_dataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_dataContext differentiates from other interfaces.
	IsMeta_dataContext()
}

type Meta_dataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_dataContext() *Meta_dataContext {
	var p = new(Meta_dataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_meta_data
	return p
}

func (*Meta_dataContext) IsMeta_dataContext() {}

func NewMeta_dataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_dataContext {
	var p = new(Meta_dataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_meta_data

	return p
}

func (s *Meta_dataContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_dataContext) Map_() IMap_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *Meta_dataContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *Meta_dataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_dataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_dataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterMeta_data(s)
	}
}

func (s *Meta_dataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitMeta_data(s)
	}
}

func (p *FugoParser) Meta_data() (localctx IMeta_dataContext) {
	this := p
	_ = this

	localctx = NewMeta_dataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FugoParserRULE_meta_data)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(FugoParserT__15)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(173)
			p.Map_()
		}
		{
			p.SetState(174)
			p.Form()
		}

	case 2:
		{
			p.SetState(176)
			p.Form()
		}

	}

	return localctx
}

// IVar_quoteContext is an interface to support dynamic dispatch.
type IVar_quoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_quoteContext differentiates from other interfaces.
	IsVar_quoteContext()
}

type Var_quoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_quoteContext() *Var_quoteContext {
	var p = new(Var_quoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_var_quote
	return p
}

func (*Var_quoteContext) IsVar_quoteContext() {}

func NewVar_quoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_quoteContext {
	var p = new(Var_quoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_var_quote

	return p
}

func (s *Var_quoteContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_quoteContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Var_quoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_quoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_quoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterVar_quote(s)
	}
}

func (s *Var_quoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitVar_quote(s)
	}
}

func (p *FugoParser) Var_quote() (localctx IVar_quoteContext) {
	this := p
	_ = this

	localctx = NewVar_quoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FugoParserRULE_var_quote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(FugoParserT__16)
	}
	{
		p.SetState(180)
		p.Symbol()
	}

	return localctx
}

// IHost_exprContext is an interface to support dynamic dispatch.
type IHost_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHost_exprContext differentiates from other interfaces.
	IsHost_exprContext()
}

type Host_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHost_exprContext() *Host_exprContext {
	var p = new(Host_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_host_expr
	return p
}

func (*Host_exprContext) IsHost_exprContext() {}

func NewHost_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Host_exprContext {
	var p = new(Host_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_host_expr

	return p
}

func (s *Host_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Host_exprContext) AllForm() []IFormContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormContext); ok {
			len++
		}
	}

	tst := make([]IFormContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormContext); ok {
			tst[i] = t.(IFormContext)
			i++
		}
	}

	return tst
}

func (s *Host_exprContext) Form(i int) IFormContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
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

	return t.(IFormContext)
}

func (s *Host_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Host_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Host_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterHost_expr(s)
	}
}

func (s *Host_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitHost_expr(s)
	}
}

func (p *FugoParser) Host_expr() (localctx IHost_exprContext) {
	this := p
	_ = this

	localctx = NewHost_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FugoParserRULE_host_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(FugoParserT__17)
	}
	{
		p.SetState(183)
		p.Form()
	}
	{
		p.SetState(184)
		p.Form()
	}

	return localctx
}

// IDiscardContext is an interface to support dynamic dispatch.
type IDiscardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiscardContext differentiates from other interfaces.
	IsDiscardContext()
}

type DiscardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiscardContext() *DiscardContext {
	var p = new(DiscardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_discard
	return p
}

func (*DiscardContext) IsDiscardContext() {}

func NewDiscardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiscardContext {
	var p = new(DiscardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_discard

	return p
}

func (s *DiscardContext) GetParser() antlr.Parser { return s.parser }

func (s *DiscardContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DiscardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiscardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiscardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterDiscard(s)
	}
}

func (s *DiscardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitDiscard(s)
	}
}

func (p *FugoParser) Discard() (localctx IDiscardContext) {
	this := p
	_ = this

	localctx = NewDiscardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FugoParserRULE_discard)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(FugoParserT__18)
	}
	{
		p.SetState(187)
		p.Form()
	}

	return localctx
}

// IDispatchContext is an interface to support dynamic dispatch.
type IDispatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDispatchContext differentiates from other interfaces.
	IsDispatchContext()
}

type DispatchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDispatchContext() *DispatchContext {
	var p = new(DispatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_dispatch
	return p
}

func (*DispatchContext) IsDispatchContext() {}

func NewDispatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DispatchContext {
	var p = new(DispatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_dispatch

	return p
}

func (s *DispatchContext) GetParser() antlr.Parser { return s.parser }

func (s *DispatchContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *DispatchContext) Form() IFormContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DispatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DispatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DispatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterDispatch(s)
	}
}

func (s *DispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitDispatch(s)
	}
}

func (p *FugoParser) Dispatch() (localctx IDispatchContext) {
	this := p
	_ = this

	localctx = NewDispatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FugoParserRULE_dispatch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(FugoParserT__13)
	}
	{
		p.SetState(190)
		p.Symbol()
	}
	{
		p.SetState(191)
		p.Form()
	}

	return localctx
}

// IRegexContext is an interface to support dynamic dispatch.
type IRegexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexContext differentiates from other interfaces.
	IsRegexContext()
}

type RegexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexContext() *RegexContext {
	var p = new(RegexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_regex
	return p
}

func (*RegexContext) IsRegexContext() {}

func NewRegexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexContext {
	var p = new(RegexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_regex

	return p
}

func (s *RegexContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterRegex(s)
	}
}

func (s *RegexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitRegex(s)
	}
}

func (p *FugoParser) Regex() (localctx IRegexContext) {
	this := p
	_ = this

	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FugoParserRULE_regex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(FugoParserT__13)
	}
	{
		p.SetState(194)
		p.String_()
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *LiteralContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LiteralContext) Character() ICharacterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterContext)
}

func (s *LiteralContext) Nil_() INil_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INil_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INil_Context)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FugoParserBOOLEAN, 0)
}

func (s *LiteralContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *LiteralContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *LiteralContext) Param_name() IParam_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_nameContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *FugoParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FugoParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FugoParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.String_()
		}

	case FugoParserFLOAT, FugoParserHEX, FugoParserBIN, FugoParserLONG, FugoParserBIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Number()
		}

	case FugoParserCHAR_U, FugoParserCHAR_NAMED, FugoParserCHAR_ANY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(198)
			p.Character()
		}

	case FugoParserNIL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.Nil_()
		}

	case FugoParserBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(200)
			p.Match(FugoParserBOOLEAN)
		}

	case FugoParserT__19:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(201)
			p.Keyword()
		}

	case FugoParserSYMBOL, FugoParserNS_SYMBOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(202)
			p.Symbol()
		}

	case FugoParserPARAM_NAME:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(203)
			p.Param_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(FugoParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *FugoParser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FugoParserRULE_string_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(FugoParserSTRING)
	}

	return localctx
}

// IHex_Context is an interface to support dynamic dispatch.
type IHex_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHex_Context differentiates from other interfaces.
	IsHex_Context()
}

type Hex_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHex_Context() *Hex_Context {
	var p = new(Hex_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_hex_
	return p
}

func (*Hex_Context) IsHex_Context() {}

func NewHex_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hex_Context {
	var p = new(Hex_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_hex_

	return p
}

func (s *Hex_Context) GetParser() antlr.Parser { return s.parser }

func (s *Hex_Context) HEX() antlr.TerminalNode {
	return s.GetToken(FugoParserHEX, 0)
}

func (s *Hex_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hex_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hex_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterHex_(s)
	}
}

func (s *Hex_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitHex_(s)
	}
}

func (p *FugoParser) Hex_() (localctx IHex_Context) {
	this := p
	_ = this

	localctx = NewHex_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FugoParserRULE_hex_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(FugoParserHEX)
	}

	return localctx
}

// IBin_Context is an interface to support dynamic dispatch.
type IBin_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_Context differentiates from other interfaces.
	IsBin_Context()
}

type Bin_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_Context() *Bin_Context {
	var p = new(Bin_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_bin_
	return p
}

func (*Bin_Context) IsBin_Context() {}

func NewBin_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_Context {
	var p = new(Bin_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_bin_

	return p
}

func (s *Bin_Context) GetParser() antlr.Parser { return s.parser }

func (s *Bin_Context) BIN() antlr.TerminalNode {
	return s.GetToken(FugoParserBIN, 0)
}

func (s *Bin_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterBin_(s)
	}
}

func (s *Bin_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitBin_(s)
	}
}

func (p *FugoParser) Bin_() (localctx IBin_Context) {
	this := p
	_ = this

	localctx = NewBin_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FugoParserRULE_bin_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(FugoParserBIN)
	}

	return localctx
}

// IBignContext is an interface to support dynamic dispatch.
type IBignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBignContext differentiates from other interfaces.
	IsBignContext()
}

type BignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBignContext() *BignContext {
	var p = new(BignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_bign
	return p
}

func (*BignContext) IsBignContext() {}

func NewBignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BignContext {
	var p = new(BignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_bign

	return p
}

func (s *BignContext) GetParser() antlr.Parser { return s.parser }

func (s *BignContext) BIGN() antlr.TerminalNode {
	return s.GetToken(FugoParserBIGN, 0)
}

func (s *BignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterBign(s)
	}
}

func (s *BignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitBign(s)
	}
}

func (p *FugoParser) Bign() (localctx IBignContext) {
	this := p
	_ = this

	localctx = NewBignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FugoParserRULE_bign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(FugoParserBIGN)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(FugoParserFLOAT, 0)
}

func (s *NumberContext) Hex_() IHex_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHex_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHex_Context)
}

func (s *NumberContext) Bin_() IBin_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_Context)
}

func (s *NumberContext) Bign() IBignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBignContext)
}

func (s *NumberContext) LONG() antlr.TerminalNode {
	return s.GetToken(FugoParserLONG, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *FugoParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FugoParserRULE_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FugoParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(FugoParserFLOAT)
		}

	case FugoParserHEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Hex_()
		}

	case FugoParserBIN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Bin_()
		}

	case FugoParserBIGN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Bign()
		}

	case FugoParserLONG:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.Match(FugoParserLONG)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICharacterContext is an interface to support dynamic dispatch.
type ICharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterContext differentiates from other interfaces.
	IsCharacterContext()
}

type CharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterContext() *CharacterContext {
	var p = new(CharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_character
	return p
}

func (*CharacterContext) IsCharacterContext() {}

func NewCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterContext {
	var p = new(CharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_character

	return p
}

func (s *CharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterContext) Named_char() INamed_charContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamed_charContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamed_charContext)
}

func (s *CharacterContext) U_hex_quad() IU_hex_quadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IU_hex_quadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IU_hex_quadContext)
}

func (s *CharacterContext) Any_char() IAny_charContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAny_charContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAny_charContext)
}

func (s *CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterCharacter(s)
	}
}

func (s *CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitCharacter(s)
	}
}

func (p *FugoParser) Character() (localctx ICharacterContext) {
	this := p
	_ = this

	localctx = NewCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FugoParserRULE_character)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FugoParserCHAR_NAMED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Named_char()
		}

	case FugoParserCHAR_U:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.U_hex_quad()
		}

	case FugoParserCHAR_ANY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Any_char()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamed_charContext is an interface to support dynamic dispatch.
type INamed_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamed_charContext differentiates from other interfaces.
	IsNamed_charContext()
}

type Named_charContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamed_charContext() *Named_charContext {
	var p = new(Named_charContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_named_char
	return p
}

func (*Named_charContext) IsNamed_charContext() {}

func NewNamed_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Named_charContext {
	var p = new(Named_charContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_named_char

	return p
}

func (s *Named_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Named_charContext) CHAR_NAMED() antlr.TerminalNode {
	return s.GetToken(FugoParserCHAR_NAMED, 0)
}

func (s *Named_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Named_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Named_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterNamed_char(s)
	}
}

func (s *Named_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitNamed_char(s)
	}
}

func (p *FugoParser) Named_char() (localctx INamed_charContext) {
	this := p
	_ = this

	localctx = NewNamed_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FugoParserRULE_named_char)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(FugoParserCHAR_NAMED)
	}

	return localctx
}

// IAny_charContext is an interface to support dynamic dispatch.
type IAny_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_charContext differentiates from other interfaces.
	IsAny_charContext()
}

type Any_charContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_charContext() *Any_charContext {
	var p = new(Any_charContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_any_char
	return p
}

func (*Any_charContext) IsAny_charContext() {}

func NewAny_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_charContext {
	var p = new(Any_charContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_any_char

	return p
}

func (s *Any_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Any_charContext) CHAR_ANY() antlr.TerminalNode {
	return s.GetToken(FugoParserCHAR_ANY, 0)
}

func (s *Any_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterAny_char(s)
	}
}

func (s *Any_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitAny_char(s)
	}
}

func (p *FugoParser) Any_char() (localctx IAny_charContext) {
	this := p
	_ = this

	localctx = NewAny_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FugoParserRULE_any_char)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(FugoParserCHAR_ANY)
	}

	return localctx
}

// IU_hex_quadContext is an interface to support dynamic dispatch.
type IU_hex_quadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsU_hex_quadContext differentiates from other interfaces.
	IsU_hex_quadContext()
}

type U_hex_quadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyU_hex_quadContext() *U_hex_quadContext {
	var p = new(U_hex_quadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_u_hex_quad
	return p
}

func (*U_hex_quadContext) IsU_hex_quadContext() {}

func NewU_hex_quadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *U_hex_quadContext {
	var p = new(U_hex_quadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_u_hex_quad

	return p
}

func (s *U_hex_quadContext) GetParser() antlr.Parser { return s.parser }

func (s *U_hex_quadContext) CHAR_U() antlr.TerminalNode {
	return s.GetToken(FugoParserCHAR_U, 0)
}

func (s *U_hex_quadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *U_hex_quadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *U_hex_quadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterU_hex_quad(s)
	}
}

func (s *U_hex_quadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitU_hex_quad(s)
	}
}

func (p *FugoParser) U_hex_quad() (localctx IU_hex_quadContext) {
	this := p
	_ = this

	localctx = NewU_hex_quadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FugoParserRULE_u_hex_quad)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(FugoParserCHAR_U)
	}

	return localctx
}

// INil_Context is an interface to support dynamic dispatch.
type INil_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNil_Context differentiates from other interfaces.
	IsNil_Context()
}

type Nil_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNil_Context() *Nil_Context {
	var p = new(Nil_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_nil_
	return p
}

func (*Nil_Context) IsNil_Context() {}

func NewNil_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nil_Context {
	var p = new(Nil_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_nil_

	return p
}

func (s *Nil_Context) GetParser() antlr.Parser { return s.parser }

func (s *Nil_Context) NIL() antlr.TerminalNode {
	return s.GetToken(FugoParserNIL, 0)
}

func (s *Nil_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nil_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nil_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterNil_(s)
	}
}

func (s *Nil_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitNil_(s)
	}
}

func (p *FugoParser) Nil_() (localctx INil_Context) {
	this := p
	_ = this

	localctx = NewNil_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FugoParserRULE_nil_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(FugoParserNIL)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Macro_keyword() IMacro_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacro_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacro_keywordContext)
}

func (s *KeywordContext) Simple_keyword() ISimple_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_keywordContext)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *FugoParser) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FugoParserRULE_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Macro_keyword()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Simple_keyword()
		}

	}

	return localctx
}

// ISimple_keywordContext is an interface to support dynamic dispatch.
type ISimple_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_keywordContext differentiates from other interfaces.
	IsSimple_keywordContext()
}

type Simple_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_keywordContext() *Simple_keywordContext {
	var p = new(Simple_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_simple_keyword
	return p
}

func (*Simple_keywordContext) IsSimple_keywordContext() {}

func NewSimple_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_keywordContext {
	var p = new(Simple_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_simple_keyword

	return p
}

func (s *Simple_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_keywordContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Simple_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterSimple_keyword(s)
	}
}

func (s *Simple_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitSimple_keyword(s)
	}
}

func (p *FugoParser) Simple_keyword() (localctx ISimple_keywordContext) {
	this := p
	_ = this

	localctx = NewSimple_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FugoParserRULE_simple_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(FugoParserT__19)
	}
	{
		p.SetState(239)
		p.Symbol()
	}

	return localctx
}

// IMacro_keywordContext is an interface to support dynamic dispatch.
type IMacro_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacro_keywordContext differentiates from other interfaces.
	IsMacro_keywordContext()
}

type Macro_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacro_keywordContext() *Macro_keywordContext {
	var p = new(Macro_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_macro_keyword
	return p
}

func (*Macro_keywordContext) IsMacro_keywordContext() {}

func NewMacro_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_keywordContext {
	var p = new(Macro_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_macro_keyword

	return p
}

func (s *Macro_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_keywordContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Macro_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Macro_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterMacro_keyword(s)
	}
}

func (s *Macro_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitMacro_keyword(s)
	}
}

func (p *FugoParser) Macro_keyword() (localctx IMacro_keywordContext) {
	this := p
	_ = this

	localctx = NewMacro_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FugoParserRULE_macro_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(FugoParserT__19)
	}
	{
		p.SetState(242)
		p.Match(FugoParserT__19)
	}
	{
		p.SetState(243)
		p.Symbol()
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Ns_symbol() INs_symbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INs_symbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INs_symbolContext)
}

func (s *SymbolContext) Simple_sym() ISimple_symContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_symContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_symContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *FugoParser) Symbol() (localctx ISymbolContext) {
	this := p
	_ = this

	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FugoParserRULE_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FugoParserNS_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Ns_symbol()
		}

	case FugoParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Simple_sym()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_symContext is an interface to support dynamic dispatch.
type ISimple_symContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_symContext differentiates from other interfaces.
	IsSimple_symContext()
}

type Simple_symContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_symContext() *Simple_symContext {
	var p = new(Simple_symContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_simple_sym
	return p
}

func (*Simple_symContext) IsSimple_symContext() {}

func NewSimple_symContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_symContext {
	var p = new(Simple_symContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_simple_sym

	return p
}

func (s *Simple_symContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_symContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(FugoParserSYMBOL, 0)
}

func (s *Simple_symContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_symContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_symContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterSimple_sym(s)
	}
}

func (s *Simple_symContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitSimple_sym(s)
	}
}

func (p *FugoParser) Simple_sym() (localctx ISimple_symContext) {
	this := p
	_ = this

	localctx = NewSimple_symContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FugoParserRULE_simple_sym)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(FugoParserSYMBOL)
	}

	return localctx
}

// INs_symbolContext is an interface to support dynamic dispatch.
type INs_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNs_symbolContext differentiates from other interfaces.
	IsNs_symbolContext()
}

type Ns_symbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNs_symbolContext() *Ns_symbolContext {
	var p = new(Ns_symbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_ns_symbol
	return p
}

func (*Ns_symbolContext) IsNs_symbolContext() {}

func NewNs_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ns_symbolContext {
	var p = new(Ns_symbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_ns_symbol

	return p
}

func (s *Ns_symbolContext) GetParser() antlr.Parser { return s.parser }

func (s *Ns_symbolContext) NS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(FugoParserNS_SYMBOL, 0)
}

func (s *Ns_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ns_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ns_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterNs_symbol(s)
	}
}

func (s *Ns_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitNs_symbol(s)
	}
}

func (p *FugoParser) Ns_symbol() (localctx INs_symbolContext) {
	this := p
	_ = this

	localctx = NewNs_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FugoParserRULE_ns_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(FugoParserNS_SYMBOL)
	}

	return localctx
}

// IParam_nameContext is an interface to support dynamic dispatch.
type IParam_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParam_nameContext differentiates from other interfaces.
	IsParam_nameContext()
}

type Param_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_nameContext() *Param_nameContext {
	var p = new(Param_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FugoParserRULE_param_name
	return p
}

func (*Param_nameContext) IsParam_nameContext() {}

func NewParam_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_nameContext {
	var p = new(Param_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FugoParserRULE_param_name

	return p
}

func (s *Param_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_nameContext) PARAM_NAME() antlr.TerminalNode {
	return s.GetToken(FugoParserPARAM_NAME, 0)
}

func (s *Param_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.EnterParam_name(s)
	}
}

func (s *Param_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FugoListener); ok {
		listenerT.ExitParam_name(s)
	}
}

func (p *FugoParser) Param_name() (localctx IParam_nameContext) {
	this := p
	_ = this

	localctx = NewParam_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FugoParserRULE_param_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(FugoParserPARAM_NAME)
	}

	return localctx
}
