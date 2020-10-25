package main

import (
	"fmt"

	"github.com/Amror/ReGompiler/postfix"
)

func main() {
	output := postfix.FormatInfix("a(a|b)*b")
	fmt.Println(string(output))
	output2 := postfix.ToPostfix(output)
	pt := &output2
	fmt.Println(pt.String())
	fmt.Println(postfix.QueueToString(*pt))
}
