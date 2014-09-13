package main

import (
	"fmt"
)

func main() {
	testCode := "((label ff (lambda (x) (cond ((atom x) x) ((quote T) (ff (car x)))))) " +
		"(cons (cons (car (quote (cd ()))) (cdr (quote (ef gh)))) nil))"

	scanner := NewLispScanner(testCode)

	expr, ok := Parse(scanner)
	if ok {
		fmt.Println(expr)

		env := NewEnv()

		result := expr.Eval(env)
		fmt.Println(result)
	} else {
		fmt.Println("Parsing Failed")
	}
}
