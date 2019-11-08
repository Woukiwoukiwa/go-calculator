package parser

var PrecedenceOperators = map[string]int{
	"(":   10,
	"+":   20,
	"-":   20,
	"/":   30,
	"*":   30,
	"\\/": 30,
	"^":   40,
}

// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func InfixExpression(tokens []Token) []Token {
	stack := []Token{}
	pending := []Token{}

	for len(tokens) > 0 {
		t, shifted := shiftToken(tokens)
		tokens = shifted

		tokenPrecedence := PrecedenceOperators[t.Value]

		stackPrecedence := 0
		if len(stack) > 0 {
			stackPrecedence = PrecedenceOperators[stack[len(stack)-1].Value]
		}

		if t.Type == OPE {
			var ope Token

			ope, poped := popToken(stack)
			stack = poped
			for ope.Value != "(" {
				pending = append(pending, ope)
				ope, poped = popToken(stack)
				pending = poped
			}
		} else if t.Type == NUM {
			pending = append(pending, t)
		} else if t.Type == OPE && (!(len(stack) > 0) || t.Value == "(" || tokenPrecedence > stackPrecedence) {
			stack = append(stack, t)
		} else {
			for tokenPrecedence <= stackPrecedence {
				token, poped := popToken(stack)
				stack = poped

				pending = append(pending, token)

				stackPrecedence = 0
				if len(stack) > 0 {
					stackPrecedence = PrecedenceOperators[stack[len(stack)-1].Value]
				}
			}

			stack = append(stack, t)
		}
	}

	for len(stack) > 0 {
		token, poped := popToken(stack)
		stack = poped

		pending = append(pending, token)
	}

	return pending
}

func shiftToken(tokens []Token) (Token, []Token) {
	var t Token
	if len(tokens) > 0 {
		t = tokens[0]
		tokens = tokens[1:]
	}
	return t, tokens
}

func popToken(tokens []Token) (Token, []Token) {
	var t Token
	if len(tokens) > 0 {
		t = tokens[len(tokens)-1]
		tokens = tokens[:len(tokens)-1]
	}
	return t, tokens
}
