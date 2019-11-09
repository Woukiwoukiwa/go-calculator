package parser

// PrecedenceOperators define Precedence of operators
var PrecedenceOperators = map[string]int{
	"(": 10,
	"+": 20,
	"-": 20,
	"/": 30,
	"*": 30,
	"^": 40,
}

// PostFixExpression return postfix expression
func PostFixExpression(infix []Token) []Token {
	opeStack := []Token{}
	tokens := []Token{}

	for len(infix) > 0 {
		t, shifted := ShiftToken(infix)
		infix = shifted

		tokenPrecedence := PrecedenceOperators[t.Value]

		// Get last OPE from stack
		stackPrecedence := 0
		if len(opeStack) > 0 {
			stackPrecedence = PrecedenceOperators[opeStack[len(opeStack)-1].Value]
		}

		if t.Type == NUM {
			// Case of NUM
			tokens = append(tokens, t)
		} else if t.Type == OPE && len(opeStack) == 0 {
			// Case First OPE in stack
			opeStack = append(opeStack, t)
		} else if t.Type == OPE && tokenPrecedence > stackPrecedence {
			// Case of token predecende is bigger of stack precedence
			opeStack = append(opeStack, t)
		} else if t.Type == OPE && t.Value == "(" {
			// Case of open bracket
			opeStack = append(opeStack, t)
		} else if t.Type == OPE && t.Value == ")" {
			// Case of closed bracket
			var ope Token

			ope, poped := popToken(opeStack)
			opeStack = poped
			for ope.Value != "(" && len(opeStack) != 0 {
				tokens = append(tokens, ope)

				ope, poped = popToken(opeStack)
				opeStack = poped
			}
		} else {
			// Case of stack predence is bigger of token precedence
			for tokenPrecedence <= stackPrecedence {
				token, poped := popToken(opeStack)
				opeStack = poped

				tokens = append(tokens, token)

				stackPrecedence = 0
				if len(opeStack) > 0 {
					stackPrecedence = PrecedenceOperators[opeStack[len(opeStack)-1].Value]
				}
			}

			opeStack = append(opeStack, t)
		}
	}

	// Add ope stack
	for len(opeStack) > 0 {
		token, poped := popToken(opeStack)
		opeStack = poped

		tokens = append(tokens, token)
	}

	return tokens
}

// ShiftToken shift token array
func ShiftToken(tokens []Token) (Token, []Token) {
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
