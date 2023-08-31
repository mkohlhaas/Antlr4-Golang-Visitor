// Code generated from Calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by CalcParser#Addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by CalcParser#Number.
	VisitNumber(ctx *NumberContext) interface{}
}
