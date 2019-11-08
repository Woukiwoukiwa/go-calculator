package parser

import (
	"regexp"
	"unicode"
)

type isFunc func(rune) bool

func Tokenize(s string) []string {
	runes := []rune(s)
	tokens := []string{}

	for len(runes) > 0 {
		// Skip white space
		_, runes = readWhile(runes, isWhitespace)

		if len(runes) == 0 {
			break
		}

		r, shifted := shift(runes)
		runes = shifted

		if isNumber(r) {
			s, readed := readWhile(runes, isNumber)
			runes = readed
			tokens = append(tokens, string(r)+s)
		} else if isOperator(string(r)) {
			tokens = append(tokens, string(r))
		}
	}
	return tokens
}

func readWhile(runes []rune, predicate isFunc) (string, []rune) {
	var s = ""
	for len(runes) > 0 && predicate(runes[0]) {
		c, shifted := shift(runes)
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

func shift(runes []rune) (rune, []rune) {
	var r rune
	if len(runes) > 0 {
		r = runes[0]
		runes = runes[1:]
	}
	return r, runes
}
