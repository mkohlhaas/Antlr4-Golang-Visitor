package main

import (
	"calc/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseCalcVisitor
}

func (v *Visitor) Visit(t antlr.ParseTree) any {
	return t.Accept(v)
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) any {
	fmt.Println("StartContext")
	return nil
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
