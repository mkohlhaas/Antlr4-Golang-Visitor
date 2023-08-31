.PHONY: all examples gen-parser clean

all: examples

gen-parser:
	antlr4 -Dlanguage=Go -o parser -no-listener -visitor Calc.g4

examples: gen-parser
	go build

clean:
	@rm -rf parser
	@rm -f ./calc
