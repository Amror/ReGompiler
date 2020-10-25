package postfix

import (
	"github.com/Amror/ReGompiler/structures"
)

var operatorPrecedence = map[string]int{"|": 1, ".": 2, "*": 3, "+": 3, "?": 3}
var operators = []string{"*", "+", "|", "?", "."}

func FormatInfix(exp string) string {
	newexp := ""

	for pos, char := range exp {
		char := string(char)
		if pos+1 < len(exp) {
			char2 := string(exp[pos+1])
			newexp += char
			if char != "(" && char2 != ")" && char != "|" && !InArray(char2, operators) {
				newexp += "."
			}
		}
	}
	newexp += string(exp[len(exp)-1])
	return newexp
}

func InArray(val string, list []string) bool {
	for _, word := range list {
		if val == word {
			return true
		}
	}
	return false
}

func higherPrecendence(op1, op2 string) bool {
	return operatorPrecedence[op1] >= operatorPrecedence[op2]
}

func ToPostfix(infix string) structures.Queue {
	outputQ := &structures.Queue{}
	opStack := &structures.Stack{}

	for _, char := range infix {
		char := string(char)
		isOperator := InArray(char, operators)
		if isOperator {
			peek := opStack.Peek()
			if peek == "(" {
				opStack.Push(char)
			} else {
				for opStack.Count() != 0 && higherPrecendence(peek.(string), char) {
					outputQ.Insert(opStack.Pop())
					peek = opStack.Peek()
				}
				opStack.Push(char)
			}
		} else {
			// char is a symbol
			switch char {
			case "(":
				opStack.Push(char)
			case ")":
				for opStack.Peek() != "(" {
					outputQ.Insert(opStack.Pop())
				}
				opStack.Pop()
			default:
				outputQ.Insert(char)
			}

		}
	}

	for opStack.Count() != 0 {
		outputQ.Insert(opStack.Pop())
	}

	return *outputQ
}

func QueueToString(q structures.Queue) string {
	output := ""
	for i := 0; i < q.Count(); i++ {
		output += q.Remove().(string)
	}
	return output
}
