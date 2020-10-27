package postfix

import (
	"errors"

	"github.com/Amror/ReGompiler/structures"
)

var operatorPrecedence = map[rune]int{'|': 1, '.': 2, '*': 3, '+': 3, '?': 3}
var operators = []rune{'*', '+', '|', '?', '.'}

func FormatInfix(exp string) ([]rune, error) {
	rexp := []rune(exp)
	newrexp := make([]rune, 0, len(rexp))

	for pos, char := range exp {
		if char == '.' {
			return []rune{}, errors.New("Concatenation operator inserted explicitly")
		}
		if pos+1 < len(rexp) {
			char2 := rexp[pos+1]
			newrexp = append(newrexp, char)
			if char != '(' && char2 != ')' && char != '|' && !InArray(char2, operators) {
				newrexp = append(newrexp, '.')
			}
		}

	}
	newrexp = append(newrexp, rexp[len(rexp)-1])
	return newrexp, nil
}

func InArray(val rune, list []rune) bool {
	for _, char := range list {
		if val == char {
			return true
		}
	}
	return false
}

func higherPrecendence(op1, op2 rune) bool {
	return operatorPrecedence[op1] >= operatorPrecedence[op2]
}

func ToPostfix(infix []rune) (structures.Queue, error) {
	outputQ := &structures.Queue{}
	opStack := &structures.Stack{}

	for _, char := range infix {
		isOperator := InArray(char, operators)
		if isOperator {
			peek := opStack.Peek()
			if peek == '(' {
				opStack.Push(char)
			} else {
				for opStack.Count() != 0 && higherPrecendence(peek.(rune), char) {
					outputQ.Insert(opStack.Pop())
					peek = opStack.Peek()
				}
				opStack.Push(char)
			}
		} else {
			// char is a symbol
			switch char {
			case '(':
				opStack.Push(char)
			case ')':
				for opStack.Count() != 0 && opStack.Peek() != '(' {
					outputQ.Insert(opStack.Pop())
				}
				if opStack.Count() == 0 {
					return structures.Queue{}, errors.New("Unmatched parentheses")
				}
				opStack.Pop()
			default:
				outputQ.Insert(char)
			}

		}
	}

	var toInsert rune
	for opStack.Count() != 0 {
		toInsert = opStack.Pop().(rune)
		if toInsert == '(' {
			return structures.Queue{}, errors.New("Unmatched parentheses")
		}
		outputQ.Insert(toInsert)
	}

	return *outputQ, nil
}

func QueueToString(q structures.Queue) string {
	length := q.Count()
	output := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		output = append(output, q.Remove().(rune))
	}
	return string(output)
}
