package nfa

import (
	"errors"

	"github.com/Amror/ReGompiler/postfix"

	"github.com/Amror/ReGompiler/structures"
)

type State struct {
	symbol     rune
	out1, out2 *State
}

type NFA struct {
	start *State
	match *State
}

func PostfixToNFA(q structures.Queue) (*NFA, error) {
	nfaStack := &structures.Stack{}

	for q.Count() != 0 {
		char := q.Remove().(rune)

		switch char {

		case '|':
			item2 := nfaStack.Pop().(*NFA)
			item1 := nfaStack.Pop().(*NFA)
			match := State{}
			start := State{out1: item1.start, out2: item2.start}
			item1.match.out1 = &match
			item2.match.out1 = &match
			nfaStack.Push(&NFA{start: &start, match: &match})

		case '.':
			item2 := nfaStack.Pop().(*NFA)
			item1 := nfaStack.Pop().(*NFA)
			item1.match.out1 = item2.start
			nfaStack.Push(&NFA{start: item1.start, match: item2.match})

		case '?':
			item := nfaStack.Pop().(*NFA)
			start := State{out1: item.start, out2: item.match}
			nfaStack.Push(&NFA{start: &start, match: item.match})

		case '*':
			item := nfaStack.Pop().(*NFA)
			match := State{}
			start := State{out1: item.start, out2: &match}
			item.match.out1 = item.start
			item.match.out2 = &match
			nfaStack.Push(&NFA{start: &start, match: &match})

		case '+':
			item := nfaStack.Pop().(*NFA)
			match := State{}
			start := State{out1: item.start, out2: &match}
			item.match.out1 = &start
			nfaStack.Push(&NFA{start: item.start, match: &match})

		default:
			match := State{}
			start := State{symbol: char, out1: &match}
			nfaStack.Push(&NFA{start: &start, match: &match})
		}
	}

	if nfaStack.Count() != 1 {
		return nil, errors.New("An error occurred while creating NFA")
	}

	return nfaStack.Pop().(*NFA), nil
}

func Match(re, str string) (bool, error) {
	if len(str) == 0 {
		return false, errors.New("regular expression cannot be empty")
	}
	infixf, error1 := postfix.FormatInfix(re)
	if error1 != nil {
		return false, error1
	}
	postfix, error2 := postfix.ToPostfix(infixf)
	if error2 != nil {
		return false, error2
	}
	nfa, error3 := PostfixToNFA(postfix)
	if error3 != nil {
		return false, error3
	}

	current := []*State{}
	next := []*State{}

	current = AddState(current[:], nfa.start, nfa.match)
	for _, r := range str {
		for _, c := range current {
			if c.symbol == r {
				next = AddState(next[:], c.out1, nfa.match)
			}
		}
		current, next = next, []*State{}
	}

	for _, c := range current {
		if c == nfa.match {
			return true, nil
		}
	}

	return false, nil
}

func AddState(l []*State, s *State, e *State) []*State {
	l = append(l, s)
	if s != e && s.symbol == 0 {
		l = AddState(l, s.out1, e)
		if s.out2 != nil {
			l = AddState(l, s.out2, e)
		}
	}
	return l
}
