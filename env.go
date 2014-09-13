package main

import (
	"fmt"
)

type Assoc struct {
	ident LispSymbol

	expr LispExpr
}

func NewAssoc(ident LispSymbol, expr LispExpr) Assoc {
	var a Assoc
	a.ident = ident
	a.expr = expr
	return a
}

func NewAssocPrimitive(ident string, expr string) Assoc {
	parsedExpr, ok := Parse(NewLispScanner(expr))
	if ok {
		return NewAssoc(LispSymbol(ident), parsedExpr)
	} else {
		panic(fmt.Sprintf("Internal Error, Can't parse %s %s", expr, parsedExpr))
	}
}

type Env []Assoc

func (env Env) Augment(aug Env) Env {
	newEnvLen := len(env) + len(aug)
	newEnv := make(Env, newEnvLen, newEnvLen)
	copy(newEnv[len(aug):], env)
	copy(newEnv, aug)
	return newEnv
}
