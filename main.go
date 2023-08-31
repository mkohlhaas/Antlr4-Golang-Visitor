package main

import (
	"calc/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseCalcVisitor
}

var _ parser.CalcVisitor = &Visitor{}

// /////////////////////////////////
// THIS FUNCTION IS NEVER CALLED  //
// /////////////////////////////////
func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	fmt.Println("StartContext")
	return v.VisitChildren(ctx)
}

func main() {
	input, _ := antlr.NewFileStream("./test_calc.txt")
	lexer := parser.NewCalcLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)
	s := p.Start_()
	v := &Visitor{}
	v.Visit(s)
}
