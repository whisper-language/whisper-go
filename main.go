package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"whisper/go/parser"
)

type TLVisitor struct {
	*parser.BaseTLVisitor
}

func (this *TLVisitor) VisitTerminal(node antlr.TerminalNode) {
	fmt.Println(node.GetText())
}

func (this *TLVisitor) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Println(node.GetText())
}

func (this *TLVisitor) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func NewTLVisitor() *TLVisitor {
	return new(TLVisitor)
}

func (this *TLVisitor) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream("./case/main.whl")
	lexer := parser.NewTLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewTLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree :=p.Parse();
	antlr.ParseTreeWalkerDefault.Walk(NewTLVisitor(), tree)
}