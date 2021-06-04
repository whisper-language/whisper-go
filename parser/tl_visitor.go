// Code generated from D:/private/tiny-language-antlr4\TL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // TL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TLParser.
type TLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TLParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by TLParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by TLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TLParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by TLParser#buildInIdentifierFunctionCall.
	VisitBuildInIdentifierFunctionCall(ctx *BuildInIdentifierFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#identifierFunctionCall.
	VisitIdentifierFunctionCall(ctx *IdentifierFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#printlnFunctionCall.
	VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#printFunctionCall.
	VisitPrintFunctionCall(ctx *PrintFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#assertFunctionCall.
	VisitAssertFunctionCall(ctx *AssertFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#sizeFunctionCall.
	VisitSizeFunctionCall(ctx *SizeFunctionCallContext) interface{}

	// Visit a parse tree produced by TLParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by TLParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by TLParser#elseIfStat.
	VisitElseIfStat(ctx *ElseIfStatContext) interface{}

	// Visit a parse tree produced by TLParser#elseStat.
	VisitElseStat(ctx *ElseStatContext) interface{}

	// Visit a parse tree produced by TLParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by TLParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by TLParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by TLParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by TLParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by TLParser#boolExpression.
	VisitBoolExpression(ctx *BoolExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#numberExpression.
	VisitNumberExpression(ctx *NumberExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#unaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#powerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#eqExpression.
	VisitEqExpression(ctx *EqExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#inExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#stringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#expressionExpression.
	VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#compExpression.
	VisitCompExpression(ctx *CompExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#nullExpression.
	VisitNullExpression(ctx *NullExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#multExpression.
	VisitMultExpression(ctx *MultExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#listExpression.
	VisitListExpression(ctx *ListExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#ternaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#inputExpression.
	VisitInputExpression(ctx *InputExpressionContext) interface{}

	// Visit a parse tree produced by TLParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by TLParser#indexes.
	VisitIndexes(ctx *IndexesContext) interface{}
}
