package calc

import (
	"log"
	"strconv"

	"github.com/Woukiwoukiwa/go-calculator/parser"
)

func Evaluate(tokens []parser.Token) int {
	result := 0

	stack := []int{}

	for len(tokens) > 0 {
		t, shifted := parser.ShiftToken(tokens)
		tokens = shifted

		if t.Type == parser.NUM {
			intValue, err := strconv.Atoi(t.Value)
			if err != nil {
				log.Fatalf("cast error %s", t.Value)
			}
			stack = append(stack, intValue)
		} else {
			// Case of OPE, get operands
			l, lPoped := popInt(stack)
			stack = lPoped

			r, rPoped := popInt(stack)
			stack = rPoped

			switch t.Value {
			case "+":
				calc := r + l
				stack = append(stack, calc)
			case "-":
				calc := r - l
				stack = append(stack, calc)
			case "*":
				calc := r * l
				stack = append(stack, calc)
			case "/":
				calc := r / l
				stack = append(stack, calc)
			case "%":
				calc := r % l
				stack = append(stack, calc)
			}
		}
	}

	result, _ = popInt(stack)

	return result
}

func popInt(values []int) (int, []int) {
	var v int
	if len(values) > 0 {
		v = values[len(values)-1]
		values = values[:len(values)-1]
	}
	return v, values
}
