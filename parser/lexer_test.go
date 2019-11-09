package parser

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsWhitespace(t *testing.T) {
	// Asserts
	assert.Assert(t, isWhitespace(' '))
	assert.Assert(t, isWhitespace('\n'))
	assert.Assert(t, !isWhitespace('a'))
	assert.Assert(t, !isWhitespace('.'))
}

func TestIsNumber(t *testing.T) {
	// Asserts
	assert.Assert(t, isNumber('1'))
	assert.Assert(t, !isNumber('a'))
	assert.Assert(t, !isNumber('.'))
}

func TestIsOperator(t *testing.T) {
	// Asserts
	assert.Assert(t, isOperator("("))
	assert.Assert(t, isOperator(")"))
	assert.Assert(t, isOperator("-"))
	assert.Assert(t, isOperator("+"))
	assert.Assert(t, isOperator("/"))
	assert.Assert(t, isOperator("*"))
	assert.Assert(t, isOperator("^"))
	assert.Assert(t, isOperator("%"))
	assert.Assert(t, !isOperator("a"))
	assert.Assert(t, !isOperator("1"))
	assert.Assert(t, !isOperator("."))
}

func TestReadWhile(t *testing.T) {
	numRunes := []rune("123a456")
	num, _ := readWhile(numRunes, isNumber)

	spaceRunes := []rune("   123")
	_, withoutSpace := readWhile(spaceRunes, isWhitespace)

	// Asserts
	assert.Assert(t, num == "123")
	assert.Assert(t, string(withoutSpace) == "123")
}

func TestTokenize(t *testing.T) {
	tokens := Tokenize("    12+14")

	noToken := Tokenize("    ")

	// Asserts
	assert.Assert(t, tokens[0].Type == NUM)
	assert.Assert(t, tokens[0].Value == "12")
	assert.Assert(t, tokens[1].Type == OPE)
	assert.Assert(t, tokens[1].Value == "+")
	assert.Assert(t, tokens[2].Type == NUM)
	assert.Assert(t, tokens[2].Value == "14")

	assert.Assert(t, len(noToken) == 0)
}
