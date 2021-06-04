// Code generated from D:/private/tiny-language-antlr4\TL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // TL

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTLVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitBuildInIdentifierFunctionCall(ctx *BuildInIdentifierFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIdentifierFunctionCall(ctx *IdentifierFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitPrintlnFunctionCall(ctx *PrintlnFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitPrintFunctionCall(ctx *PrintFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitAssertFunctionCall(ctx *AssertFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitSizeFunctionCall(ctx *SizeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitElseIfStat(ctx *ElseIfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitElseStat(ctx *ElseStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIdList(ctx *IdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitBoolExpression(ctx *BoolExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitNumberExpression(ctx *NumberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitEqExpression(ctx *EqExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitStringExpression(ctx *StringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitCompExpression(ctx *CompExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitNullExpression(ctx *NullExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitMultExpression(ctx *MultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitListExpression(ctx *ListExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitInputExpression(ctx *InputExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTLVisitor) VisitIndexes(ctx *IndexesContext) interface{} {
	return v.VisitChildren(ctx)
}
