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

func TestIsNum(t *testing.T) {
	// Asserts
	assert.Assert(t, isNum('1'))
	assert.Assert(t, !isNum('a'))
	assert.Assert(t, !isNum('.'))
}

func TestIsOp(t *testing.T) {
	// Asserts
	assert.Assert(t, isOp("("))
	assert.Assert(t, isOp(")"))
	assert.Assert(t, isOp("-"))
	assert.Assert(t, isOp("+"))
	assert.Assert(t, isOp("/"))
	assert.Assert(t, isOp("*"))
	assert.Assert(t, isOp("^"))
	assert.Assert(t, isOp("%"))
	assert.Assert(t, !isOp("a"))
	assert.Assert(t, !isOp("1"))
	assert.Assert(t, !isOp("."))
}
