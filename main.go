package main

import (
	"fmt"

	"github.com/Amror/ReGompiler/postfix"
)

func main() {
	output := postfix.FormatInfix("abc|bcd")
	fmt.Println(output)
	output2 := postfix.ToPostfix(output)
	pt := &output2
	fmt.Println(pt.String())
}
