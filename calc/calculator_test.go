package calc

import (
	"testing"

	"gotest.tools/assert"

	"github.com/Woukiwoukiwa/go-calculator/parser"
)

func TestEvaluate1(t *testing.T) {
	infix := parser.Tokenize("1 + 2")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 3)
}

func TestEvaluate2(t *testing.T) {
	infix := parser.Tokenize("5 - 4")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 1)
}

func TestEvaluate3(t *testing.T) {
	infix := parser.Tokenize("3 * 2")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 6)
}

func TestEvaluate4(t *testing.T) {
	infix := parser.Tokenize("4 / 2")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 2)
}

func TestEvaluate5(t *testing.T) {
	infix := parser.Tokenize("5 % 2")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 1)
}

func TestEvaluate6(t *testing.T) {
	infix := parser.Tokenize("((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1))")
	tokens := parser.PostFixExpression(infix)
	result := Evaluate(tokens)

	// Asserts
	assert.Assert(t, result == 5)
}
