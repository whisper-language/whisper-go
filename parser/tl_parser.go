// Code generated from D:/private/tiny-language-antlr4\TL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // TL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 267,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 45, 10,
	3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3,
	5, 3, 5, 5, 5, 69, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 77,
	10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 83, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 89, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 107, 10, 6, 3, 7, 3, 7,
	7, 7, 111, 10, 7, 12, 7, 14, 7, 114, 11, 7, 3, 7, 5, 7, 117, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	141, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 167, 10, 14, 12, 14, 14, 14,
	170, 11, 14, 3, 15, 3, 15, 3, 15, 7, 15, 175, 10, 15, 12, 15, 14, 15, 178,
	11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 5, 16, 190, 10, 16, 3, 16, 3, 16, 5, 16, 194, 10, 16, 3, 16, 3,
	16, 5, 16, 198, 10, 16, 3, 16, 3, 16, 5, 16, 202, 10, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 208, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 213, 10,
	16, 3, 16, 5, 16, 216, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 248, 10, 16, 12, 16, 14, 16, 251, 11, 16, 3,
	17, 3, 17, 5, 17, 255, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	6, 18, 263, 10, 18, 13, 18, 14, 18, 264, 3, 18, 2, 3, 30, 19, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 6, 3, 2, 31,
	33, 3, 2, 29, 30, 4, 2, 23, 24, 27, 28, 3, 2, 21, 22, 2, 298, 2, 38, 3,
	2, 2, 2, 4, 46, 3, 2, 2, 2, 6, 64, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 106,
	3, 2, 2, 2, 12, 108, 3, 2, 2, 2, 14, 118, 3, 2, 2, 2, 16, 124, 3, 2, 2,
	2, 18, 131, 3, 2, 2, 2, 20, 136, 3, 2, 2, 2, 22, 147, 3, 2, 2, 2, 24, 157,
	3, 2, 2, 2, 26, 163, 3, 2, 2, 2, 28, 171, 3, 2, 2, 2, 30, 215, 3, 2, 2,
	2, 32, 252, 3, 2, 2, 2, 34, 262, 3, 2, 2, 2, 36, 39, 5, 4, 3, 2, 37, 39,
	5, 6, 4, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 41, 7, 2, 2, 3, 41, 3, 3, 2, 2, 2, 42, 45, 5, 6, 4, 2, 43, 45, 5, 20,
	11, 2, 44, 42, 3, 2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46,
	44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 53, 3, 2, 2, 2, 48, 46, 3, 2, 2,
	2, 49, 50, 7, 11, 2, 2, 50, 51, 5, 30, 16, 2, 51, 52, 7, 40, 2, 2, 52,
	54, 3, 2, 2, 2, 53, 49, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 5, 3, 2, 2,
	2, 55, 56, 5, 8, 5, 2, 56, 57, 7, 40, 2, 2, 57, 65, 3, 2, 2, 2, 58, 59,
	5, 10, 6, 2, 59, 60, 7, 40, 2, 2, 60, 65, 3, 2, 2, 2, 61, 65, 5, 12, 7,
	2, 62, 65, 5, 22, 12, 2, 63, 65, 5, 24, 13, 2, 64, 55, 3, 2, 2, 2, 64,
	58, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 63, 3, 2, 2,
	2, 65, 7, 3, 2, 2, 2, 66, 68, 7, 48, 2, 2, 67, 69, 5, 34, 18, 2, 68, 67,
	3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 7, 41, 2, 2,
	71, 72, 5, 30, 16, 2, 72, 9, 3, 2, 2, 2, 73, 74, 7, 47, 2, 2, 74, 76, 7,
	38, 2, 2, 75, 77, 5, 28, 15, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 78, 3, 2, 2, 2, 78, 107, 7, 39, 2, 2, 79, 80, 7, 48, 2, 2, 80, 82,
	7, 38, 2, 2, 81, 83, 5, 28, 15, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2,
	2, 83, 84, 3, 2, 2, 2, 84, 107, 7, 39, 2, 2, 85, 86, 7, 3, 2, 2, 86, 88,
	7, 38, 2, 2, 87, 89, 5, 30, 16, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 107, 7, 39, 2, 2, 91, 92, 7, 4, 2, 2, 92, 93,
	7, 38, 2, 2, 93, 94, 5, 30, 16, 2, 94, 95, 7, 39, 2, 2, 95, 107, 3, 2,
	2, 2, 96, 97, 7, 6, 2, 2, 97, 98, 7, 38, 2, 2, 98, 99, 5, 30, 16, 2, 99,
	100, 7, 39, 2, 2, 100, 107, 3, 2, 2, 2, 101, 102, 7, 7, 2, 2, 102, 103,
	7, 38, 2, 2, 103, 104, 5, 30, 16, 2, 104, 105, 7, 39, 2, 2, 105, 107, 3,
	2, 2, 2, 106, 73, 3, 2, 2, 2, 106, 79, 3, 2, 2, 2, 106, 85, 3, 2, 2, 2,
	106, 91, 3, 2, 2, 2, 106, 96, 3, 2, 2, 2, 106, 101, 3, 2, 2, 2, 107, 11,
	3, 2, 2, 2, 108, 112, 5, 14, 8, 2, 109, 111, 5, 16, 9, 2, 110, 109, 3,
	2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2,
	2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 117, 5, 18, 10, 2,
	116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 13, 3, 2, 2, 2, 118, 119,
	7, 9, 2, 2, 119, 120, 5, 30, 16, 2, 120, 121, 7, 34, 2, 2, 121, 122, 5,
	4, 3, 2, 122, 123, 7, 35, 2, 2, 123, 15, 3, 2, 2, 2, 124, 125, 7, 10, 2,
	2, 125, 126, 7, 9, 2, 2, 126, 127, 5, 30, 16, 2, 127, 128, 7, 34, 2, 2,
	128, 129, 5, 4, 3, 2, 129, 130, 7, 35, 2, 2, 130, 17, 3, 2, 2, 2, 131,
	132, 7, 10, 2, 2, 132, 133, 7, 34, 2, 2, 133, 134, 5, 4, 3, 2, 134, 135,
	7, 35, 2, 2, 135, 19, 3, 2, 2, 2, 136, 137, 7, 8, 2, 2, 137, 138, 7, 48,
	2, 2, 138, 140, 7, 38, 2, 2, 139, 141, 5, 26, 14, 2, 140, 139, 3, 2, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 39, 2, 2, 143,
	144, 7, 34, 2, 2, 144, 145, 5, 4, 3, 2, 145, 146, 7, 35, 2, 2, 146, 21,
	3, 2, 2, 2, 147, 148, 7, 12, 2, 2, 148, 149, 7, 48, 2, 2, 149, 150, 7,
	41, 2, 2, 150, 151, 5, 30, 16, 2, 151, 152, 7, 14, 2, 2, 152, 153, 5, 30,
	16, 2, 153, 154, 7, 34, 2, 2, 154, 155, 5, 4, 3, 2, 155, 156, 7, 35, 2,
	2, 156, 23, 3, 2, 2, 2, 157, 158, 7, 13, 2, 2, 158, 159, 5, 30, 16, 2,
	159, 160, 7, 34, 2, 2, 160, 161, 5, 4, 3, 2, 161, 162, 7, 35, 2, 2, 162,
	25, 3, 2, 2, 2, 163, 168, 7, 48, 2, 2, 164, 165, 7, 42, 2, 2, 165, 167,
	7, 48, 2, 2, 166, 164, 3, 2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2,
	2, 2, 168, 169, 3, 2, 2, 2, 169, 27, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2,
	171, 176, 5, 30, 16, 2, 172, 173, 7, 42, 2, 2, 173, 175, 5, 30, 16, 2,
	174, 172, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176,
	177, 3, 2, 2, 2, 177, 29, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 179, 180, 8,
	16, 1, 2, 180, 181, 7, 30, 2, 2, 181, 216, 5, 30, 16, 22, 182, 183, 7,
	26, 2, 2, 183, 216, 5, 30, 16, 21, 184, 216, 7, 46, 2, 2, 185, 216, 7,
	45, 2, 2, 186, 216, 7, 18, 2, 2, 187, 189, 5, 10, 6, 2, 188, 190, 5, 34,
	18, 2, 189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 216, 3, 2, 2, 2,
	191, 193, 5, 32, 17, 2, 192, 194, 5, 34, 18, 2, 193, 192, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 216, 3, 2, 2, 2, 195, 197, 7, 48, 2, 2, 196, 198,
	5, 34, 18, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 216, 3,
	2, 2, 2, 199, 201, 7, 49, 2, 2, 200, 202, 5, 34, 18, 2, 201, 200, 3, 2,
	2, 2, 201, 202, 3, 2, 2, 2, 202, 216, 3, 2, 2, 2, 203, 204, 7, 38, 2, 2,
	204, 205, 5, 30, 16, 2, 205, 207, 7, 39, 2, 2, 206, 208, 5, 34, 18, 2,
	207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 216, 3, 2, 2, 2, 209,
	210, 7, 5, 2, 2, 210, 212, 7, 38, 2, 2, 211, 213, 7, 49, 2, 2, 212, 211,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 216, 7, 39,
	2, 2, 215, 179, 3, 2, 2, 2, 215, 182, 3, 2, 2, 2, 215, 184, 3, 2, 2, 2,
	215, 185, 3, 2, 2, 2, 215, 186, 3, 2, 2, 2, 215, 187, 3, 2, 2, 2, 215,
	191, 3, 2, 2, 2, 215, 195, 3, 2, 2, 2, 215, 199, 3, 2, 2, 2, 215, 203,
	3, 2, 2, 2, 215, 209, 3, 2, 2, 2, 216, 249, 3, 2, 2, 2, 217, 218, 12, 20,
	2, 2, 218, 219, 7, 25, 2, 2, 219, 248, 5, 30, 16, 20, 220, 221, 12, 19,
	2, 2, 221, 222, 9, 2, 2, 2, 222, 248, 5, 30, 16, 20, 223, 224, 12, 18,
	2, 2, 224, 225, 9, 3, 2, 2, 225, 248, 5, 30, 16, 19, 226, 227, 12, 17,
	2, 2, 227, 228, 9, 4, 2, 2, 228, 248, 5, 30, 16, 18, 229, 230, 12, 16,
	2, 2, 230, 231, 9, 5, 2, 2, 231, 248, 5, 30, 16, 17, 232, 233, 12, 15,
	2, 2, 233, 234, 7, 20, 2, 2, 234, 248, 5, 30, 16, 16, 235, 236, 12, 14,
	2, 2, 236, 237, 7, 19, 2, 2, 237, 248, 5, 30, 16, 15, 238, 239, 12, 13,
	2, 2, 239, 240, 7, 43, 2, 2, 240, 241, 5, 30, 16, 2, 241, 242, 7, 44, 2,
	2, 242, 243, 5, 30, 16, 14, 243, 248, 3, 2, 2, 2, 244, 245, 12, 12, 2,
	2, 245, 246, 7, 17, 2, 2, 246, 248, 5, 30, 16, 13, 247, 217, 3, 2, 2, 2,
	247, 220, 3, 2, 2, 2, 247, 223, 3, 2, 2, 2, 247, 226, 3, 2, 2, 2, 247,
	229, 3, 2, 2, 2, 247, 232, 3, 2, 2, 2, 247, 235, 3, 2, 2, 2, 247, 238,
	3, 2, 2, 2, 247, 244, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2, 249, 247, 3, 2,
	2, 2, 249, 250, 3, 2, 2, 2, 250, 31, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2,
	252, 254, 7, 36, 2, 2, 253, 255, 5, 28, 15, 2, 254, 253, 3, 2, 2, 2, 254,
	255, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 7, 37, 2, 2, 257, 33,
	3, 2, 2, 2, 258, 259, 7, 36, 2, 2, 259, 260, 5, 30, 16, 2, 260, 261, 7,
	37, 2, 2, 261, 263, 3, 2, 2, 2, 262, 258, 3, 2, 2, 2, 263, 264, 3, 2, 2,
	2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 35, 3, 2, 2, 2, 28,
	38, 44, 46, 53, 64, 68, 76, 82, 88, 106, 112, 116, 140, 168, 176, 189,
	193, 197, 201, 207, 212, 215, 247, 249, 254, 264,
}
var literalNames = []string{
	"", "'println'", "'print'", "'input'", "'assert'", "'size'", "'func'",
	"'if'", "'else'", "'return'", "'for'", "'while'", "'to'", "'do'", "'end'",
	"'in'", "'null'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'^'",
	"'!'", "'>'", "'<'", "'+'", "'-'", "'*'", "'/'", "'%'", "'{'", "'}'", "'['",
	"']'", "'('", "')'", "';'", "'='", "','", "'?'", "':'",
}
var symbolicNames = []string{
	"", "Println", "Print", "Input", "Assert", "Size", "Def", "If", "Else",
	"Return", "For", "While", "To", "Do", "End", "In", "Null", "Or", "And",
	"Equals", "NEquals", "GTEquals", "LTEquals", "Pow", "Excl", "GT", "LT",
	"Add", "Subtract", "Multiply", "Divide", "Modulus", "OBrace", "CBrace",
	"OBracket", "CBracket", "OParen", "CParen", "SColon", "Assign", "Comma",
	"QMark", "Colon", "Bool", "Number", "BuildIdentifier", "Identifier", "P_String",
	"Comment", "Space",
}

var ruleNames = []string{
	"parse", "block", "statement", "assignment", "functionCall", "ifStatement",
	"ifStat", "elseIfStat", "elseStat", "functionDecl", "forStatement", "whileStatement",
	"idList", "exprList", "expression", "list", "indexes",
}

type TLParser struct {
	*antlr.BaseParser
}

// NewTLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *TLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTLParser(input antlr.TokenStream) *TLParser {
	this := new(TLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TL.g4"

	return this
}

// TLParser tokens.
const (
	TLParserEOF             = antlr.TokenEOF
	TLParserPrintln         = 1
	TLParserPrint           = 2
	TLParserInput           = 3
	TLParserAssert          = 4
	TLParserSize            = 5
	TLParserDef             = 6
	TLParserIf              = 7
	TLParserElse            = 8
	TLParserReturn          = 9
	TLParserFor             = 10
	TLParserWhile           = 11
	TLParserTo              = 12
	TLParserDo              = 13
	TLParserEnd             = 14
	TLParserIn              = 15
	TLParserNull            = 16
	TLParserOr              = 17
	TLParserAnd             = 18
	TLParserEquals          = 19
	TLParserNEquals         = 20
	TLParserGTEquals        = 21
	TLParserLTEquals        = 22
	TLParserPow             = 23
	TLParserExcl            = 24
	TLParserGT              = 25
	TLParserLT              = 26
	TLParserAdd             = 27
	TLParserSubtract        = 28
	TLParserMultiply        = 29
	TLParserDivide          = 30
	TLParserModulus         = 31
	TLParserOBrace          = 32
	TLParserCBrace          = 33
	TLParserOBracket        = 34
	TLParserCBracket        = 35
	TLParserOParen          = 36
	TLParserCParen          = 37
	TLParserSColon          = 38
	TLParserAssign          = 39
	TLParserComma           = 40
	TLParserQMark           = 41
	TLParserColon           = 42
	TLParserBool            = 43
	TLParserNumber          = 44
	TLParserBuildIdentifier = 45
	TLParserIdentifier      = 46
	TLParserP_String        = 47
	TLParserComment         = 48
	TLParserSpace           = 49
)

// TLParser rules.
const (
	TLParserRULE_parse          = 0
	TLParserRULE_block          = 1
	TLParserRULE_statement      = 2
	TLParserRULE_assignment     = 3
	TLParserRULE_functionCall   = 4
	TLParserRULE_ifStatement    = 5
	TLParserRULE_ifStat         = 6
	TLParserRULE_elseIfStat     = 7
	TLParserRULE_elseStat       = 8
	TLParserRULE_functionDecl   = 9
	TLParserRULE_forStatement   = 10
	TLParserRULE_whileStatement = 11
	TLParserRULE_idList         = 12
	TLParserRULE_exprList       = 13
	TLParserRULE_expression     = 14
	TLParserRULE_list           = 15
	TLParserRULE_indexes        = 16
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(TLParserEOF, 0)
}

func (s *ParseContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParseContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TLParserRULE_parse)

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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(34)
			p.Block()
		}

	case 2:
		{
			p.SetState(35)
			p.Statement()
		}

	}
	{
		p.SetState(38)
		p.Match(TLParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllFunctionDecl() []IFunctionDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclContext)(nil)).Elem())
	var tst = make([]IFunctionDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclContext)
		}
	}

	return tst
}

func (s *BlockContext) FunctionDecl(i int) IFunctionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *BlockContext) Return() antlr.TerminalNode {
	return s.GetToken(TLParserReturn, 0)
}

func (s *BlockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockContext) SColon() antlr.TerminalNode {
	return s.GetToken(TLParserSColon, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TLParserRULE_block)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserPrintln)|(1<<TLParserPrint)|(1<<TLParserAssert)|(1<<TLParserSize)|(1<<TLParserDef)|(1<<TLParserIf)|(1<<TLParserFor)|(1<<TLParserWhile))) != 0) || _la == TLParserBuildIdentifier || _la == TLParserIdentifier {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TLParserPrintln, TLParserPrint, TLParserAssert, TLParserSize, TLParserIf, TLParserFor, TLParserWhile, TLParserBuildIdentifier, TLParserIdentifier:
			{
				p.SetState(40)
				p.Statement()
			}

		case TLParserDef:
			{
				p.SetState(41)
				p.FunctionDecl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TLParserReturn {
		{
			p.SetState(47)
			p.Match(TLParserReturn)
		}
		{
			p.SetState(48)
			p.expression(0)
		}
		{
			p.SetState(49)
			p.Match(TLParserSColon)
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) SColon() antlr.TerminalNode {
	return s.GetToken(TLParserSColon, 0)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TLParserRULE_statement)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Assignment()
		}
		{
			p.SetState(54)
			p.Match(TLParserSColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.FunctionCall()
		}
		{
			p.SetState(57)
			p.Match(TLParserSColon)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
			p.ForStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)
			p.WhileStatement()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, 0)
}

func (s *AssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(TLParserAssign, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TLParserRULE_assignment)
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
		p.SetState(64)
		p.Match(TLParserIdentifier)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TLParserOBracket {
		{
			p.SetState(65)
			p.Indexes()
		}

	}
	{
		p.SetState(68)
		p.Match(TLParserAssign)
	}
	{
		p.SetState(69)
		p.expression(0)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) CopyFrom(ctx *FunctionCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssertFunctionCallContext struct {
	*FunctionCallContext
}

func NewAssertFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssertFunctionCallContext {
	var p = new(AssertFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *AssertFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertFunctionCallContext) Assert() antlr.TerminalNode {
	return s.GetToken(TLParserAssert, 0)
}

func (s *AssertFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *AssertFunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssertFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *AssertFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitAssertFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type SizeFunctionCallContext struct {
	*FunctionCallContext
}

func NewSizeFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SizeFunctionCallContext {
	var p = new(SizeFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *SizeFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeFunctionCallContext) Size() antlr.TerminalNode {
	return s.GetToken(TLParserSize, 0)
}

func (s *SizeFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *SizeFunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SizeFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *SizeFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitSizeFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnFunctionCallContext struct {
	*FunctionCallContext
}

func NewPrintlnFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnFunctionCallContext {
	var p = new(PrintlnFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *PrintlnFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnFunctionCallContext) Println() antlr.TerminalNode {
	return s.GetToken(TLParserPrintln, 0)
}

func (s *PrintlnFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *PrintlnFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *PrintlnFunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintlnFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitPrintlnFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type BuildInIdentifierFunctionCallContext struct {
	*FunctionCallContext
}

func NewBuildInIdentifierFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuildInIdentifierFunctionCallContext {
	var p = new(BuildInIdentifierFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *BuildInIdentifierFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInIdentifierFunctionCallContext) BuildIdentifier() antlr.TerminalNode {
	return s.GetToken(TLParserBuildIdentifier, 0)
}

func (s *BuildInIdentifierFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *BuildInIdentifierFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *BuildInIdentifierFunctionCallContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *BuildInIdentifierFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitBuildInIdentifierFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierFunctionCallContext struct {
	*FunctionCallContext
}

func NewIdentifierFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierFunctionCallContext {
	var p = new(IdentifierFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *IdentifierFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, 0)
}

func (s *IdentifierFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *IdentifierFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *IdentifierFunctionCallContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *IdentifierFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIdentifierFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintFunctionCallContext struct {
	*FunctionCallContext
}

func NewPrintFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintFunctionCallContext {
	var p = new(PrintFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *PrintFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintFunctionCallContext) Print() antlr.TerminalNode {
	return s.GetToken(TLParserPrint, 0)
}

func (s *PrintFunctionCallContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *PrintFunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintFunctionCallContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *PrintFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitPrintFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TLParserRULE_functionCall)
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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TLParserBuildIdentifier:
		localctx = NewBuildInIdentifierFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(TLParserBuildIdentifier)
		}
		{
			p.SetState(72)
			p.Match(TLParserOParen)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserPrintln)|(1<<TLParserPrint)|(1<<TLParserInput)|(1<<TLParserAssert)|(1<<TLParserSize)|(1<<TLParserNull)|(1<<TLParserExcl)|(1<<TLParserSubtract))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(TLParserOBracket-34))|(1<<(TLParserOParen-34))|(1<<(TLParserBool-34))|(1<<(TLParserNumber-34))|(1<<(TLParserBuildIdentifier-34))|(1<<(TLParserIdentifier-34))|(1<<(TLParserP_String-34)))) != 0) {
			{
				p.SetState(73)
				p.ExprList()
			}

		}
		{
			p.SetState(76)
			p.Match(TLParserCParen)
		}

	case TLParserIdentifier:
		localctx = NewIdentifierFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(TLParserIdentifier)
		}
		{
			p.SetState(78)
			p.Match(TLParserOParen)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserPrintln)|(1<<TLParserPrint)|(1<<TLParserInput)|(1<<TLParserAssert)|(1<<TLParserSize)|(1<<TLParserNull)|(1<<TLParserExcl)|(1<<TLParserSubtract))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(TLParserOBracket-34))|(1<<(TLParserOParen-34))|(1<<(TLParserBool-34))|(1<<(TLParserNumber-34))|(1<<(TLParserBuildIdentifier-34))|(1<<(TLParserIdentifier-34))|(1<<(TLParserP_String-34)))) != 0) {
			{
				p.SetState(79)
				p.ExprList()
			}

		}
		{
			p.SetState(82)
			p.Match(TLParserCParen)
		}

	case TLParserPrintln:
		localctx = NewPrintlnFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.Match(TLParserPrintln)
		}
		{
			p.SetState(84)
			p.Match(TLParserOParen)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserPrintln)|(1<<TLParserPrint)|(1<<TLParserInput)|(1<<TLParserAssert)|(1<<TLParserSize)|(1<<TLParserNull)|(1<<TLParserExcl)|(1<<TLParserSubtract))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(TLParserOBracket-34))|(1<<(TLParserOParen-34))|(1<<(TLParserBool-34))|(1<<(TLParserNumber-34))|(1<<(TLParserBuildIdentifier-34))|(1<<(TLParserIdentifier-34))|(1<<(TLParserP_String-34)))) != 0) {
			{
				p.SetState(85)
				p.expression(0)
			}

		}
		{
			p.SetState(88)
			p.Match(TLParserCParen)
		}

	case TLParserPrint:
		localctx = NewPrintFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(TLParserPrint)
		}
		{
			p.SetState(90)
			p.Match(TLParserOParen)
		}
		{
			p.SetState(91)
			p.expression(0)
		}
		{
			p.SetState(92)
			p.Match(TLParserCParen)
		}

	case TLParserAssert:
		localctx = NewAssertFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(TLParserAssert)
		}
		{
			p.SetState(95)
			p.Match(TLParserOParen)
		}
		{
			p.SetState(96)
			p.expression(0)
		}
		{
			p.SetState(97)
			p.Match(TLParserCParen)
		}

	case TLParserSize:
		localctx = NewSizeFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.Match(TLParserSize)
		}
		{
			p.SetState(100)
			p.Match(TLParserOParen)
		}
		{
			p.SetState(101)
			p.expression(0)
		}
		{
			p.SetState(102)
			p.Match(TLParserCParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfStat() IIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *IfStatementContext) AllElseIfStat() []IElseIfStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStatContext)(nil)).Elem())
	var tst = make([]IElseIfStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStatContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStat(i int) IElseIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStatContext)
}

func (s *IfStatementContext) ElseStat() IElseStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TLParserRULE_ifStatement)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.IfStat()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(107)
				p.ElseIfStat()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TLParserElse {
		{
			p.SetState(113)
			p.ElseStat()
		}

	}

	return localctx
}

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_ifStat
	return p
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) If() antlr.TerminalNode {
	return s.GetToken(TLParserIf, 0)
}

func (s *IfStatContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *IfStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TLParserRULE_ifStat)

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
		p.SetState(116)
		p.Match(TLParserIf)
	}
	{
		p.SetState(117)
		p.expression(0)
	}
	{
		p.SetState(118)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(119)
		p.Block()
	}
	{
		p.SetState(120)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IElseIfStatContext is an interface to support dynamic dispatch.
type IElseIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatContext differentiates from other interfaces.
	IsElseIfStatContext()
}

type ElseIfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatContext() *ElseIfStatContext {
	var p = new(ElseIfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_elseIfStat
	return p
}

func (*ElseIfStatContext) IsElseIfStatContext() {}

func NewElseIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatContext {
	var p = new(ElseIfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_elseIfStat

	return p
}

func (s *ElseIfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatContext) Else() antlr.TerminalNode {
	return s.GetToken(TLParserElse, 0)
}

func (s *ElseIfStatContext) If() antlr.TerminalNode {
	return s.GetToken(TLParserIf, 0)
}

func (s *ElseIfStatContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStatContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *ElseIfStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfStatContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *ElseIfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitElseIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) ElseIfStat() (localctx IElseIfStatContext) {
	localctx = NewElseIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TLParserRULE_elseIfStat)

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
		p.SetState(122)
		p.Match(TLParserElse)
	}
	{
		p.SetState(123)
		p.Match(TLParserIf)
	}
	{
		p.SetState(124)
		p.expression(0)
	}
	{
		p.SetState(125)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(126)
		p.Block()
	}
	{
		p.SetState(127)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IElseStatContext is an interface to support dynamic dispatch.
type IElseStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatContext differentiates from other interfaces.
	IsElseStatContext()
}

type ElseStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatContext() *ElseStatContext {
	var p = new(ElseStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_elseStat
	return p
}

func (*ElseStatContext) IsElseStatContext() {}

func NewElseStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatContext {
	var p = new(ElseStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_elseStat

	return p
}

func (s *ElseStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatContext) Else() antlr.TerminalNode {
	return s.GetToken(TLParserElse, 0)
}

func (s *ElseStatContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *ElseStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseStatContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *ElseStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitElseStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) ElseStat() (localctx IElseStatContext) {
	localctx = NewElseStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TLParserRULE_elseStat)

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
		p.SetState(129)
		p.Match(TLParserElse)
	}
	{
		p.SetState(130)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(131)
		p.Block()
	}
	{
		p.SetState(132)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_functionDecl
	return p
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) Def() antlr.TerminalNode {
	return s.GetToken(TLParserDef, 0)
}

func (s *FunctionDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, 0)
}

func (s *FunctionDeclContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *FunctionDeclContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *FunctionDeclContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *FunctionDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *FunctionDeclContext) IdList() IIdListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TLParserRULE_functionDecl)
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
		p.SetState(134)
		p.Match(TLParserDef)
	}
	{
		p.SetState(135)
		p.Match(TLParserIdentifier)
	}
	{
		p.SetState(136)
		p.Match(TLParserOParen)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TLParserIdentifier {
		{
			p.SetState(137)
			p.IdList()
		}

	}
	{
		p.SetState(140)
		p.Match(TLParserCParen)
	}
	{
		p.SetState(141)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(142)
		p.Block()
	}
	{
		p.SetState(143)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(TLParserFor, 0)
}

func (s *ForStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, 0)
}

func (s *ForStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(TLParserAssign, 0)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) To() antlr.TerminalNode {
	return s.GetToken(TLParserTo, 0)
}

func (s *ForStatementContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TLParserRULE_forStatement)

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
		p.SetState(145)
		p.Match(TLParserFor)
	}
	{
		p.SetState(146)
		p.Match(TLParserIdentifier)
	}
	{
		p.SetState(147)
		p.Match(TLParserAssign)
	}
	{
		p.SetState(148)
		p.expression(0)
	}
	{
		p.SetState(149)
		p.Match(TLParserTo)
	}
	{
		p.SetState(150)
		p.expression(0)
	}
	{
		p.SetState(151)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(152)
		p.Block()
	}
	{
		p.SetState(153)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(TLParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) OBrace() antlr.TerminalNode {
	return s.GetToken(TLParserOBrace, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) CBrace() antlr.TerminalNode {
	return s.GetToken(TLParserCBrace, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TLParserRULE_whileStatement)

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
		p.SetState(155)
		p.Match(TLParserWhile)
	}
	{
		p.SetState(156)
		p.expression(0)
	}
	{
		p.SetState(157)
		p.Match(TLParserOBrace)
	}
	{
		p.SetState(158)
		p.Block()
	}
	{
		p.SetState(159)
		p.Match(TLParserCBrace)
	}

	return localctx
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_idList
	return p
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(TLParserIdentifier)
}

func (s *IdListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, i)
}

func (s *IdListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(TLParserComma)
}

func (s *IdListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(TLParserComma, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TLParserRULE_idList)
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
		p.SetState(161)
		p.Match(TLParserIdentifier)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TLParserComma {
		{
			p.SetState(162)
			p.Match(TLParserComma)
		}
		{
			p.SetState(163)
			p.Match(TLParserIdentifier)
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(TLParserComma)
}

func (s *ExprListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(TLParserComma, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TLParserRULE_exprList)
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
		p.SetState(169)
		p.expression(0)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TLParserComma {
		{
			p.SetState(170)
			p.Match(TLParserComma)
		}
		{
			p.SetState(171)
			p.expression(0)
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExpressionContext struct {
	*ExpressionContext
}

func NewBoolExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) Bool() antlr.TerminalNode {
	return s.GetToken(TLParserBool, 0)
}

func (s *BoolExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitBoolExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExpressionContext struct {
	*ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(TLParserNumber, 0)
}

func (s *NumberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitNumberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TLParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Excl() antlr.TerminalNode {
	return s.GetToken(TLParserExcl, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpressionContext struct {
	*ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(TLParserOr, 0)
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExpressionContext struct {
	*ExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Subtract() antlr.TerminalNode {
	return s.GetToken(TLParserSubtract, 0)
}

func (s *UnaryMinusExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitUnaryMinusExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerExpressionContext struct {
	*ExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowerExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExpressionContext) Pow() antlr.TerminalNode {
	return s.GetToken(TLParserPow, 0)
}

func (s *PowerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitPowerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewEqExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExpressionContext {
	var p = new(EqExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(TLParserEquals, 0)
}

func (s *EqExpressionContext) NEquals() antlr.TerminalNode {
	return s.GetToken(TLParserNEquals, 0)
}

func (s *EqExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitEqExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpressionContext struct {
	*ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(TLParserAnd, 0)
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(TLParserIn, 0)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExpressionContext struct {
	*ExpressionContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) P_String() antlr.TerminalNode {
	return s.GetToken(TLParserP_String, 0)
}

func (s *StringExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *StringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionExpressionContext struct {
	*ExpressionContext
}

func NewExpressionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionExpressionContext {
	var p = new(ExpressionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionExpressionContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *ExpressionExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionExpressionContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *ExpressionExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *ExpressionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitExpressionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) Add() antlr.TerminalNode {
	return s.GetToken(TLParserAdd, 0)
}

func (s *AddExpressionContext) Subtract() antlr.TerminalNode {
	return s.GetToken(TLParserSubtract, 0)
}

func (s *AddExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitAddExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewCompExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExpressionContext {
	var p = new(CompExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompExpressionContext) GetOp() antlr.Token { return s.op }

func (s *CompExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) GTEquals() antlr.TerminalNode {
	return s.GetToken(TLParserGTEquals, 0)
}

func (s *CompExpressionContext) LTEquals() antlr.TerminalNode {
	return s.GetToken(TLParserLTEquals, 0)
}

func (s *CompExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(TLParserGT, 0)
}

func (s *CompExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(TLParserLT, 0)
}

func (s *CompExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitCompExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExpressionContext struct {
	*ExpressionContext
}

func NewNullExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExpressionContext {
	var p = new(NullExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NullExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(TLParserNull, 0)
}

func (s *NullExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitNullExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMultExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpressionContext {
	var p = new(MultExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) Multiply() antlr.TerminalNode {
	return s.GetToken(TLParserMultiply, 0)
}

func (s *MultExpressionContext) Divide() antlr.TerminalNode {
	return s.GetToken(TLParserDivide, 0)
}

func (s *MultExpressionContext) Modulus() antlr.TerminalNode {
	return s.GetToken(TLParserModulus, 0)
}

func (s *MultExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitMultExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListExpressionContext struct {
	*ExpressionContext
}

func NewListExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExpressionContext {
	var p = new(ListExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExpressionContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ListExpressionContext) Indexes() IIndexesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *ListExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitListExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExpressionContext struct {
	*ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) QMark() antlr.TerminalNode {
	return s.GetToken(TLParserQMark, 0)
}

func (s *TernaryExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(TLParserColon, 0)
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitTernaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InputExpressionContext struct {
	*ExpressionContext
}

func NewInputExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InputExpressionContext {
	var p = new(InputExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InputExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputExpressionContext) Input() antlr.TerminalNode {
	return s.GetToken(TLParserInput, 0)
}

func (s *InputExpressionContext) OParen() antlr.TerminalNode {
	return s.GetToken(TLParserOParen, 0)
}

func (s *InputExpressionContext) CParen() antlr.TerminalNode {
	return s.GetToken(TLParserCParen, 0)
}

func (s *InputExpressionContext) P_String() antlr.TerminalNode {
	return s.GetToken(TLParserP_String, 0)
}

func (s *InputExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitInputExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TLParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, TLParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(178)
			p.Match(TLParserSubtract)
		}
		{
			p.SetState(179)
			p.expression(20)
		}

	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(180)
			p.Match(TLParserExcl)
		}
		{
			p.SetState(181)
			p.expression(19)
		}

	case 3:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Match(TLParserNumber)
		}

	case 4:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(183)
			p.Match(TLParserBool)
		}

	case 5:
		localctx = NewNullExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(184)
			p.Match(TLParserNull)
		}

	case 6:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(185)
			p.FunctionCall()
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(186)
				p.Indexes()
			}

		}

	case 7:
		localctx = NewListExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(189)
			p.List()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(190)
				p.Indexes()
			}

		}

	case 8:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(193)
			p.Match(TLParserIdentifier)
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(194)
				p.Indexes()
			}

		}

	case 9:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Match(TLParserP_String)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(198)
				p.Indexes()
			}

		}

	case 10:
		localctx = NewExpressionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(201)
			p.Match(TLParserOParen)
		}
		{
			p.SetState(202)
			p.expression(0)
		}
		{
			p.SetState(203)
			p.Match(TLParserCParen)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(204)
				p.Indexes()
			}

		}

	case 11:
		localctx = NewInputExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(207)
			p.Match(TLParserInput)
		}
		{
			p.SetState(208)
			p.Match(TLParserOParen)
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TLParserP_String {
			{
				p.SetState(209)
				p.Match(TLParserP_String)
			}

		}
		{
			p.SetState(212)
			p.Match(TLParserCParen)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(216)
					p.Match(TLParserPow)
				}
				{
					p.SetState(217)
					p.expression(18)
				}

			case 2:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(219)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserMultiply)|(1<<TLParserDivide)|(1<<TLParserModulus))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(220)
					p.expression(18)
				}

			case 3:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(222)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TLParserAdd || _la == TLParserSubtract) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(223)
					p.expression(17)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(225)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserGTEquals)|(1<<TLParserLTEquals)|(1<<TLParserGT)|(1<<TLParserLT))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(226)
					p.expression(16)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(228)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TLParserEquals || _la == TLParserNEquals) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(229)
					p.expression(15)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(231)
					p.Match(TLParserAnd)
				}
				{
					p.SetState(232)
					p.expression(14)
				}

			case 7:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(234)
					p.Match(TLParserOr)
				}
				{
					p.SetState(235)
					p.expression(13)
				}

			case 8:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(237)
					p.Match(TLParserQMark)
				}
				{
					p.SetState(238)
					p.expression(0)
				}
				{
					p.SetState(239)
					p.Match(TLParserColon)
				}
				{
					p.SetState(240)
					p.expression(12)
				}

			case 9:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TLParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(243)
					p.Match(TLParserIn)
				}
				{
					p.SetState(244)
					p.expression(11)
				}

			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) OBracket() antlr.TerminalNode {
	return s.GetToken(TLParserOBracket, 0)
}

func (s *ListContext) CBracket() antlr.TerminalNode {
	return s.GetToken(TLParserCBracket, 0)
}

func (s *ListContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TLParserRULE_list)
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
		p.SetState(250)
		p.Match(TLParserOBracket)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TLParserPrintln)|(1<<TLParserPrint)|(1<<TLParserInput)|(1<<TLParserAssert)|(1<<TLParserSize)|(1<<TLParserNull)|(1<<TLParserExcl)|(1<<TLParserSubtract))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(TLParserOBracket-34))|(1<<(TLParserOParen-34))|(1<<(TLParserBool-34))|(1<<(TLParserNumber-34))|(1<<(TLParserBuildIdentifier-34))|(1<<(TLParserIdentifier-34))|(1<<(TLParserP_String-34)))) != 0) {
		{
			p.SetState(251)
			p.ExprList()
		}

	}
	{
		p.SetState(254)
		p.Match(TLParserCBracket)
	}

	return localctx
}

// IIndexesContext is an interface to support dynamic dispatch.
type IIndexesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexesContext differentiates from other interfaces.
	IsIndexesContext()
}

type IndexesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexesContext() *IndexesContext {
	var p = new(IndexesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TLParserRULE_indexes
	return p
}

func (*IndexesContext) IsIndexesContext() {}

func NewIndexesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexesContext {
	var p = new(IndexesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TLParserRULE_indexes

	return p
}

func (s *IndexesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexesContext) AllOBracket() []antlr.TerminalNode {
	return s.GetTokens(TLParserOBracket)
}

func (s *IndexesContext) OBracket(i int) antlr.TerminalNode {
	return s.GetToken(TLParserOBracket, i)
}

func (s *IndexesContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IndexesContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexesContext) AllCBracket() []antlr.TerminalNode {
	return s.GetTokens(TLParserCBracket)
}

func (s *IndexesContext) CBracket(i int) antlr.TerminalNode {
	return s.GetToken(TLParserCBracket, i)
}

func (s *IndexesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TLVisitor:
		return t.VisitIndexes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TLParser) Indexes() (localctx IIndexesContext) {
	localctx = NewIndexesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TLParserRULE_indexes)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(256)
				p.Match(TLParserOBracket)
			}
			{
				p.SetState(257)
				p.expression(0)
			}
			{
				p.SetState(258)
				p.Match(TLParserCBracket)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

func (p *TLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TLParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
