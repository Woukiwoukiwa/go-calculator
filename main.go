package main

import (
	"fmt"
	"os"

	"github.com/Woukiwoukiwa/go-calculator/calc"
	"github.com/Woukiwoukiwa/go-calculator/parser"
)

func main() {
	argsWithoutProg := os.Args[1]

	infix := parser.Tokenize(argsWithoutProg)
	tokens := parser.PostFixExpression(infix)
	result := calc.Evaluate(tokens)

	fmt.Println(result)
}
