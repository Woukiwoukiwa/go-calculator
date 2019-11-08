package parser

import (
	"testing"

	"gotest.tools/assert"
)

func TestInfixExpression(t *testing.T) {
	tokens := Tokenize("((15 ÷ (7 − (1 + 1))) × 3) − (2 + (1 + 1))")
	infix := InfixExpression(tokens)

	// Asserts
	assert.Assert(t, infix[0].Value == "15")
	assert.Assert(t, infix[1].Value == "7")
	assert.Assert(t, infix[2].Value == "1")
}
