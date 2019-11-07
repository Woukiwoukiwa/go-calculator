package parser

import (
	"regexp"
	"unicode"
)

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNum(r rune) bool {
	return unicode.IsDigit(r)
}
func isOp(r string) bool {
	var validOp = regexp.MustCompile(`\(|\)|\-|\+|\/|\*|\^|\%`)
	return validOp.MatchString(r)
}
