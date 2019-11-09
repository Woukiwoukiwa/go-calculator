package parser

import (
	"regexp"
	"unicode"
)

type isFunc func(rune) bool

type LexerType int

const (
	OPE    LexerType = iota //0
	NUM                     //1
	Erable                  //2
)

type Token struct {
	Type  LexerType
	Value string
}

func Tokenize(s string) []Token {
	runes := []rune(s)
	tokens := []Token{}

	for len(runes) > 0 {
		// Skip white space
		_, runes = readWhile(runes, isWhitespace)

		if len(runes) == 0 {
			break
		}

		r, shifted := shiftRune(runes)
		runes = shifted

		if isNumber(r) {
			s, readed := readWhile(runes, isNumber)
			runes = readed
			token := Token{
				Type:  NUM,
				Value: string(r) + s,
			}
			tokens = append(tokens, token)
		} else if isOperator(string(r)) {
			token := Token{
				Type:  OPE,
				Value: string(r),
			}
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func readWhile(runes []rune, predicate isFunc) (string, []rune) {
	var s = ""
	for len(runes) > 0 && predicate(runes[0]) {
		c, shifted := shiftRune(runes)
		runes = shifted
		s += string(c)
	}
	return s, runes
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}
func isOperator(r string) bool {
	var validOp = regexp.MustCompile(`\(|\)|\-|\+|\/|\*|\^|\%`)
	return validOp.MatchString(r)
}

func shiftRune(runes []rune) (rune, []rune) {
	var r rune
	if len(runes) > 0 {
		r = runes[0]
		runes = runes[1:]
	}
	return r, runes
}
