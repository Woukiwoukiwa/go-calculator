package parser

import (
	"testing"

	"gotest.tools/assert"
)

func TestInfixExpression1(t *testing.T) {
	infix := Tokenize("15 + 12")
	tokens := PostFixExpression(infix)

	// Asserts
	assert.Assert(t, tokens[0].Value == "15")
	assert.Assert(t, tokens[1].Value == "12")
	assert.Assert(t, tokens[2].Value == "+")
}

func TestInfixExpression2(t *testing.T) {
	infix := Tokenize("15 + 12 * 2")
	tokens := PostFixExpression(infix)

	// Asserts
	assert.Assert(t, tokens[0].Value == "15")
	assert.Assert(t, tokens[1].Value == "12")
	assert.Assert(t, tokens[2].Value == "2")
	assert.Assert(t, tokens[3].Value == "*")
	assert.Assert(t, tokens[4].Value == "+")
}

func TestInfixExpression3(t *testing.T) {
	infix := Tokenize("(15 + 12) * 2")
	tokens := PostFixExpression(infix)

	// Asserts
	assert.Assert(t, tokens[0].Value == "15")
	assert.Assert(t, tokens[1].Value == "12")
	assert.Assert(t, tokens[2].Value == "+")
	assert.Assert(t, tokens[3].Value == "2")
	assert.Assert(t, tokens[4].Value == "*")
}

func TestInfixExpression4(t *testing.T) {
	infix := Tokenize("2+15*2+2")
	tokens := PostFixExpression(infix)

	// Asserts
	assert.Assert(t, tokens[0].Value == "2")
	assert.Assert(t, tokens[1].Value == "15")
	assert.Assert(t, tokens[2].Value == "2")
	assert.Assert(t, tokens[3].Value == "*")
	assert.Assert(t, tokens[4].Value == "+")
	assert.Assert(t, tokens[5].Value == "2")
	assert.Assert(t, tokens[6].Value == "+")
}

func TestInfixExpression5(t *testing.T) {
	infix := Tokenize("((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1))")
	tokens := PostFixExpression(infix)

	// Asserts
	assert.Assert(t, tokens[0].Value == "15")
	assert.Assert(t, tokens[1].Value == "7")
	assert.Assert(t, tokens[2].Value == "1")
	assert.Assert(t, tokens[3].Value == "1")
	assert.Assert(t, tokens[4].Value == "+")
	assert.Assert(t, tokens[5].Value == "-")
	assert.Assert(t, tokens[6].Value == "/")
	assert.Assert(t, tokens[7].Value == "3")
	assert.Assert(t, tokens[8].Value == "*")
	assert.Assert(t, tokens[9].Value == "2")
	assert.Assert(t, tokens[10].Value == "1")
	assert.Assert(t, tokens[11].Value == "1")
	assert.Assert(t, tokens[12].Value == "+")
	assert.Assert(t, tokens[13].Value == "+")
	assert.Assert(t, tokens[14].Value == "-")
}
