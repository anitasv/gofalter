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

		env := Env{
			*NewAssocPrimitive("nil", "()"),
			*NewAssocPrimitive("T", "T"),
			*NewAssocPrimitive("cadr", "(lambda (x) (car (cdr x)))"),
			*NewAssocPrimitive("caddr", "(lambda (x) (car (cdr (cdr x))))"),
			*NewAssocPrimitive("caar", "(lambda (x) (car (car x)))"),
			*NewAssocPrimitive("cadar", "(lambda (x) (car (cdr (car x))))"),
			*NewAssocPrimitive("caddar", "(lambda (x) (car (cdr (cdr (car x)))))"),
			*NewAssocPrimitive("null", "(lambda (x) (equal x nil))"),
		}
		result, err := expr.Eval(env)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result)
	} else {
		fmt.Println("Parsing Failed")
	}
}
