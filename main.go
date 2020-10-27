package main

import (
	"fmt"

	"github.com/Amror/ReGompiler/nfa"
)

func main() {
	var re, str string
	for {
		fmt.Println("Enter regex expression")
		fmt.Scanln(&re)
		fmt.Println("Enter string to match")
		fmt.Scanln(&str)
		matched, err := nfa.Match(re, str)
		if matched && err == nil {
			fmt.Println("Matched")
		} else if !matched && err == nil {
			fmt.Println("Not Matched")
		} else {
			fmt.Println(err)
		}
	}
}
