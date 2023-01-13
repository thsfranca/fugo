// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FugoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var fugolexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fugolexerLexerInit() {
	staticData := &fugolexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "STRING", "FLOAT", "FLOAT_TAIL", "FLOAT_DECIMAL",
		"FLOAT_EXP", "HEXD", "HEX", "BIN", "LONG", "BIGN", "CHAR_U", "CHAR_NAMED",
		"CHAR_ANY", "NIL", "BOOLEAN", "SYMBOL", "NS_SYMBOL", "PARAM_NAME", "NAME",
		"SYMBOL_HEAD", "SYMBOL_REST", "WS", "COMMENT", "TRASH",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 358, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 141, 8, 20, 10, 20, 12, 20, 144, 9, 20, 1, 20, 1, 20, 1,
		21, 3, 21, 149, 8, 21, 1, 21, 4, 21, 152, 8, 21, 11, 21, 12, 21, 153, 1,
		21, 1, 21, 3, 21, 158, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 169, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 174,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 181, 8, 22, 1, 23, 1,
		23, 4, 23, 185, 8, 23, 11, 23, 12, 23, 186, 1, 24, 1, 24, 3, 24, 191, 8,
		24, 1, 24, 4, 24, 194, 8, 24, 11, 24, 12, 24, 195, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 4, 26, 203, 8, 26, 11, 26, 12, 26, 204, 1, 27, 1, 27, 1,
		27, 4, 27, 210, 8, 27, 11, 27, 12, 27, 211, 1, 28, 3, 28, 215, 8, 28, 1,
		28, 4, 28, 218, 8, 28, 11, 28, 12, 28, 219, 1, 28, 3, 28, 223, 8, 28, 1,
		29, 3, 29, 226, 8, 29, 1, 29, 4, 29, 229, 8, 29, 11, 29, 12, 29, 230, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 281, 8, 31, 1, 32, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 299, 8, 34, 1, 35, 1, 35, 3, 35, 303, 8,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 5, 37, 312, 8, 37,
		10, 37, 12, 37, 315, 9, 37, 1, 37, 3, 37, 318, 8, 37, 1, 38, 1, 38, 5,
		38, 322, 8, 38, 10, 38, 12, 38, 325, 9, 38, 1, 38, 1, 38, 4, 38, 329, 8,
		38, 11, 38, 12, 38, 330, 5, 38, 333, 8, 38, 10, 38, 12, 38, 336, 9, 38,
		1, 39, 1, 39, 1, 40, 1, 40, 3, 40, 342, 8, 40, 1, 41, 1, 41, 1, 42, 1,
		42, 5, 42, 348, 8, 42, 10, 42, 12, 42, 351, 9, 42, 1, 43, 1, 43, 3, 43,
		355, 8, 43, 1, 43, 1, 43, 0, 0, 44, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 0, 47, 0, 49, 0,
		51, 0, 53, 23, 55, 24, 57, 25, 59, 26, 61, 27, 63, 28, 65, 29, 67, 30,
		69, 31, 71, 32, 73, 33, 75, 34, 77, 0, 79, 0, 81, 0, 83, 0, 85, 0, 87,
		35, 1, 0, 14, 1, 0, 34, 34, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101, 3, 0,
		48, 57, 65, 70, 97, 102, 2, 0, 88, 88, 120, 120, 2, 0, 66, 66, 98, 98,
		1, 0, 48, 49, 2, 0, 76, 76, 108, 108, 2, 0, 78, 78, 110, 110, 3, 0, 48,
		57, 68, 70, 100, 102, 14, 0, 9, 10, 13, 13, 32, 32, 34, 35, 37, 37, 39,
		41, 44, 44, 47, 58, 64, 64, 91, 91, 93, 94, 96, 96, 123, 123, 125, 126,
		2, 0, 46, 46, 48, 57, 4, 0, 9, 10, 13, 13, 32, 32, 44, 44, 2, 0, 10, 10,
		13, 13, 384, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0, 3, 91, 1, 0, 0, 0, 5, 93, 1, 0, 0,
		0, 7, 95, 1, 0, 0, 0, 9, 97, 1, 0, 0, 0, 11, 99, 1, 0, 0, 0, 13, 101, 1,
		0, 0, 0, 15, 104, 1, 0, 0, 0, 17, 106, 1, 0, 0, 0, 19, 108, 1, 0, 0, 0,
		21, 110, 1, 0, 0, 0, 23, 113, 1, 0, 0, 0, 25, 115, 1, 0, 0, 0, 27, 117,
		1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 122, 1, 0, 0, 0, 33, 125, 1, 0, 0,
		0, 35, 128, 1, 0, 0, 0, 37, 131, 1, 0, 0, 0, 39, 134, 1, 0, 0, 0, 41, 136,
		1, 0, 0, 0, 43, 173, 1, 0, 0, 0, 45, 180, 1, 0, 0, 0, 47, 182, 1, 0, 0,
		0, 49, 188, 1, 0, 0, 0, 51, 197, 1, 0, 0, 0, 53, 199, 1, 0, 0, 0, 55, 206,
		1, 0, 0, 0, 57, 214, 1, 0, 0, 0, 59, 225, 1, 0, 0, 0, 61, 234, 1, 0, 0,
		0, 63, 241, 1, 0, 0, 0, 65, 282, 1, 0, 0, 0, 67, 285, 1, 0, 0, 0, 69, 298,
		1, 0, 0, 0, 71, 302, 1, 0, 0, 0, 73, 304, 1, 0, 0, 0, 75, 308, 1, 0, 0,
		0, 77, 319, 1, 0, 0, 0, 79, 337, 1, 0, 0, 0, 81, 341, 1, 0, 0, 0, 83, 343,
		1, 0, 0, 0, 85, 345, 1, 0, 0, 0, 87, 354, 1, 0, 0, 0, 89, 90, 5, 40, 0,
		0, 90, 2, 1, 0, 0, 0, 91, 92, 5, 41, 0, 0, 92, 4, 1, 0, 0, 0, 93, 94, 5,
		91, 0, 0, 94, 6, 1, 0, 0, 0, 95, 96, 5, 93, 0, 0, 96, 8, 1, 0, 0, 0, 97,
		98, 5, 123, 0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 125, 0, 0, 100, 12, 1,
		0, 0, 0, 101, 102, 5, 35, 0, 0, 102, 103, 5, 123, 0, 0, 103, 14, 1, 0,
		0, 0, 104, 105, 5, 39, 0, 0, 105, 16, 1, 0, 0, 0, 106, 107, 5, 96, 0, 0,
		107, 18, 1, 0, 0, 0, 108, 109, 5, 126, 0, 0, 109, 20, 1, 0, 0, 0, 110,
		111, 5, 126, 0, 0, 111, 112, 5, 64, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114,
		5, 94, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 5, 64, 0, 0, 116, 26, 1, 0,
		0, 0, 117, 118, 5, 35, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 35, 0, 0,
		120, 121, 5, 40, 0, 0, 121, 30, 1, 0, 0, 0, 122, 123, 5, 35, 0, 0, 123,
		124, 5, 94, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 35, 0, 0, 126, 127,
		5, 39, 0, 0, 127, 34, 1, 0, 0, 0, 128, 129, 5, 35, 0, 0, 129, 130, 5, 43,
		0, 0, 130, 36, 1, 0, 0, 0, 131, 132, 5, 35, 0, 0, 132, 133, 5, 95, 0, 0,
		133, 38, 1, 0, 0, 0, 134, 135, 5, 58, 0, 0, 135, 40, 1, 0, 0, 0, 136, 142,
		5, 34, 0, 0, 137, 141, 8, 0, 0, 0, 138, 139, 5, 92, 0, 0, 139, 141, 5,
		34, 0, 0, 140, 137, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 144, 1, 0, 0,
		0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 145, 146, 5, 34, 0, 0, 146, 42, 1, 0, 0, 0, 147, 149,
		5, 45, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 1, 0,
		0, 0, 150, 152, 7, 1, 0, 0, 151, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		174, 3, 45, 22, 0, 156, 158, 5, 45, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 5, 73, 0, 0, 160, 161, 5, 110,
		0, 0, 161, 162, 5, 102, 0, 0, 162, 163, 5, 105, 0, 0, 163, 164, 5, 110,
		0, 0, 164, 165, 5, 105, 0, 0, 165, 166, 5, 116, 0, 0, 166, 174, 5, 121,
		0, 0, 167, 169, 5, 45, 0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 171, 5, 78, 0, 0, 171, 172, 5, 97, 0, 0, 172,
		174, 5, 78, 0, 0, 173, 148, 1, 0, 0, 0, 173, 157, 1, 0, 0, 0, 173, 168,
		1, 0, 0, 0, 174, 44, 1, 0, 0, 0, 175, 176, 3, 47, 23, 0, 176, 177, 3, 49,
		24, 0, 177, 181, 1, 0, 0, 0, 178, 181, 3, 47, 23, 0, 179, 181, 3, 49, 24,
		0, 180, 175, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181,
		46, 1, 0, 0, 0, 182, 184, 5, 46, 0, 0, 183, 185, 7, 1, 0, 0, 184, 183,
		1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 48, 1, 0, 0, 0, 188, 190, 7, 2, 0, 0, 189, 191, 5, 45, 0, 0,
		190, 189, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192,
		194, 7, 1, 0, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 50, 1, 0, 0, 0, 197, 198, 7, 3,
		0, 0, 198, 52, 1, 0, 0, 0, 199, 200, 5, 48, 0, 0, 200, 202, 7, 4, 0, 0,
		201, 203, 3, 51, 25, 0, 202, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 54, 1, 0, 0, 0, 206, 207, 5,
		48, 0, 0, 207, 209, 7, 5, 0, 0, 208, 210, 7, 6, 0, 0, 209, 208, 1, 0, 0,
		0, 210, 211, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212,
		56, 1, 0, 0, 0, 213, 215, 5, 45, 0, 0, 214, 213, 1, 0, 0, 0, 214, 215,
		1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 218, 7, 1, 0, 0, 217, 216, 1, 0,
		0, 0, 218, 219, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 222, 1, 0, 0, 0, 221, 223, 7, 7, 0, 0, 222, 221, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 58, 1, 0, 0, 0, 224, 226, 5, 45, 0, 0, 225, 224,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 228, 1, 0, 0, 0, 227, 229, 7, 1,
		0, 0, 228, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 7, 8, 0, 0, 233,
		60, 1, 0, 0, 0, 234, 235, 5, 92, 0, 0, 235, 236, 5, 117, 0, 0, 236, 237,
		7, 9, 0, 0, 237, 238, 3, 51, 25, 0, 238, 239, 3, 51, 25, 0, 239, 240, 3,
		51, 25, 0, 240, 62, 1, 0, 0, 0, 241, 280, 5, 92, 0, 0, 242, 243, 5, 110,
		0, 0, 243, 244, 5, 101, 0, 0, 244, 245, 5, 119, 0, 0, 245, 246, 5, 108,
		0, 0, 246, 247, 5, 105, 0, 0, 247, 248, 5, 110, 0, 0, 248, 281, 5, 101,
		0, 0, 249, 250, 5, 114, 0, 0, 250, 251, 5, 101, 0, 0, 251, 252, 5, 116,
		0, 0, 252, 253, 5, 117, 0, 0, 253, 254, 5, 114, 0, 0, 254, 281, 5, 110,
		0, 0, 255, 256, 5, 115, 0, 0, 256, 257, 5, 112, 0, 0, 257, 258, 5, 97,
		0, 0, 258, 259, 5, 99, 0, 0, 259, 281, 5, 101, 0, 0, 260, 261, 5, 116,
		0, 0, 261, 262, 5, 97, 0, 0, 262, 281, 5, 98, 0, 0, 263, 264, 5, 102, 0,
		0, 264, 265, 5, 111, 0, 0, 265, 266, 5, 114, 0, 0, 266, 267, 5, 109, 0,
		0, 267, 268, 5, 102, 0, 0, 268, 269, 5, 101, 0, 0, 269, 270, 5, 101, 0,
		0, 270, 281, 5, 100, 0, 0, 271, 272, 5, 98, 0, 0, 272, 273, 5, 97, 0, 0,
		273, 274, 5, 99, 0, 0, 274, 275, 5, 107, 0, 0, 275, 276, 5, 115, 0, 0,
		276, 277, 5, 112, 0, 0, 277, 278, 5, 97, 0, 0, 278, 279, 5, 99, 0, 0, 279,
		281, 5, 101, 0, 0, 280, 242, 1, 0, 0, 0, 280, 249, 1, 0, 0, 0, 280, 255,
		1, 0, 0, 0, 280, 260, 1, 0, 0, 0, 280, 263, 1, 0, 0, 0, 280, 271, 1, 0,
		0, 0, 281, 64, 1, 0, 0, 0, 282, 283, 5, 92, 0, 0, 283, 284, 9, 0, 0, 0,
		284, 66, 1, 0, 0, 0, 285, 286, 5, 110, 0, 0, 286, 287, 5, 105, 0, 0, 287,
		288, 5, 108, 0, 0, 288, 68, 1, 0, 0, 0, 289, 290, 5, 116, 0, 0, 290, 291,
		5, 114, 0, 0, 291, 292, 5, 117, 0, 0, 292, 299, 5, 101, 0, 0, 293, 294,
		5, 102, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5, 108, 0, 0, 296, 297,
		5, 115, 0, 0, 297, 299, 5, 101, 0, 0, 298, 289, 1, 0, 0, 0, 298, 293, 1,
		0, 0, 0, 299, 70, 1, 0, 0, 0, 300, 303, 2, 46, 47, 0, 301, 303, 3, 77,
		38, 0, 302, 300, 1, 0, 0, 0, 302, 301, 1, 0, 0, 0, 303, 72, 1, 0, 0, 0,
		304, 305, 3, 77, 38, 0, 305, 306, 5, 47, 0, 0, 306, 307, 3, 71, 35, 0,
		307, 74, 1, 0, 0, 0, 308, 317, 5, 37, 0, 0, 309, 313, 2, 49, 57, 0, 310,
		312, 2, 48, 57, 0, 311, 310, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311,
		1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 318, 1, 0, 0, 0, 315, 313, 1, 0,
		0, 0, 316, 318, 5, 38, 0, 0, 317, 309, 1, 0, 0, 0, 317, 316, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 76, 1, 0, 0, 0, 319, 323, 3, 79, 39, 0, 320,
		322, 3, 81, 40, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 334, 1, 0, 0, 0, 325, 323, 1, 0,
		0, 0, 326, 328, 5, 58, 0, 0, 327, 329, 3, 81, 40, 0, 328, 327, 1, 0, 0,
		0, 329, 330, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331,
		333, 1, 0, 0, 0, 332, 326, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334, 332,
		1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 78, 1, 0, 0, 0, 336, 334, 1, 0,
		0, 0, 337, 338, 8, 10, 0, 0, 338, 80, 1, 0, 0, 0, 339, 342, 3, 79, 39,
		0, 340, 342, 7, 11, 0, 0, 341, 339, 1, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342,
		82, 1, 0, 0, 0, 343, 344, 7, 12, 0, 0, 344, 84, 1, 0, 0, 0, 345, 349, 5,
		59, 0, 0, 346, 348, 8, 13, 0, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1, 0,
		0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 86, 1, 0, 0, 0,
		351, 349, 1, 0, 0, 0, 352, 355, 3, 83, 41, 0, 353, 355, 3, 85, 42, 0, 354,
		352, 1, 0, 0, 0, 354, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357,
		6, 43, 0, 0, 357, 88, 1, 0, 0, 0, 30, 0, 140, 142, 148, 153, 157, 168,
		173, 180, 186, 190, 195, 204, 211, 214, 219, 222, 225, 230, 280, 298, 302,
		313, 317, 323, 330, 334, 341, 349, 354, 1, 0, 1, 0,
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

// FugoLexerInit initializes any static state used to implement FugoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFugoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FugoLexerInit() {
	staticData := &fugolexerLexerStaticData
	staticData.once.Do(fugolexerLexerInit)
}

// NewFugoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFugoLexer(input antlr.CharStream) *FugoLexer {
	FugoLexerInit()
	l := new(FugoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &fugolexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Fugo.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FugoLexer tokens.
const (
	FugoLexerT__0       = 1
	FugoLexerT__1       = 2
	FugoLexerT__2       = 3
	FugoLexerT__3       = 4
	FugoLexerT__4       = 5
	FugoLexerT__5       = 6
	FugoLexerT__6       = 7
	FugoLexerT__7       = 8
	FugoLexerT__8       = 9
	FugoLexerT__9       = 10
	FugoLexerT__10      = 11
	FugoLexerT__11      = 12
	FugoLexerT__12      = 13
	FugoLexerT__13      = 14
	FugoLexerT__14      = 15
	FugoLexerT__15      = 16
	FugoLexerT__16      = 17
	FugoLexerT__17      = 18
	FugoLexerT__18      = 19
	FugoLexerT__19      = 20
	FugoLexerSTRING     = 21
	FugoLexerFLOAT      = 22
	FugoLexerHEX        = 23
	FugoLexerBIN        = 24
	FugoLexerLONG       = 25
	FugoLexerBIGN       = 26
	FugoLexerCHAR_U     = 27
	FugoLexerCHAR_NAMED = 28
	FugoLexerCHAR_ANY   = 29
	FugoLexerNIL        = 30
	FugoLexerBOOLEAN    = 31
	FugoLexerSYMBOL     = 32
	FugoLexerNS_SYMBOL  = 33
	FugoLexerPARAM_NAME = 34
	FugoLexerTRASH      = 35
)
