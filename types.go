package main

import (
	"bytes"
)

type LispExpr interface {
	Atom() LispExpr

	Nil() LispExpr

	Car() (LispExpr, error)

	Cdr() (LispList, error)

	Cons(expr LispExpr) (LispList, error)

	Eval(env Env) (LispExpr, error)

	Call(args LispList, env Env) (LispExpr, error)

	// Debugging and final printing purpose.
	String() string

	// For speedy Cond() evaluation
	IsNil() bool
}

type LispSymbol string
type LispList []LispExpr

var LISP_TRUE = LispSymbol("T")
var LISP_FALSE = new(LispList)

func (s LispSymbol) String() string {
	return string(s)
}

func (l LispList) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	first := true
	for _, item := range l {
		if !first {
			out.WriteString(" ")
		} else {
			first = false
		}
		out.WriteString(item.String())
	}
	out.WriteString(")")
	return out.String()
}

func (s LispSymbol) Atom() LispExpr {
	return LISP_TRUE
}

func (l LispList) Atom() LispExpr {
	if l.IsNil() {
		return LISP_TRUE
	} else {
		return LISP_FALSE
	}
}

func (s LispSymbol) IsNil() bool {
	return false
}

func (l LispList) IsNil() bool {
	return len(l) == 0
}

func (s LispSymbol) Nil() LispExpr {
	return LISP_FALSE
}

func (l LispList) Nil() LispExpr {
	return l.Atom()
}

func (s LispSymbol) Car() (LispExpr, error) {
	return nil, CompileError("car is not permitted for symbol")
}

func (l LispList) Car() (LispExpr, error) {
	if len(l) > 0 {
		return l[0], nil
	} else {
		return nil, CompileError("car on nil not permitted")
	}
}

func (s LispSymbol) Cons(expr LispExpr) (LispList, error) {
	return nil, CompileError("cons is not permitted for symbol")
}

func (l LispList) Cons(expr LispExpr) (LispList, error) {
	newL := make(LispList, len(l)+1, len(l)+1)
	copy(newL[1:], l)
	newL[0] = expr
	return newL, nil
}

func (s LispSymbol) Cdr() (LispList, error) {
	return nil, CompileError("cdr is not permitted for symbol")
}

func (l LispList) Cdr() (LispList, error) {
	if len(l) > 0 {
		return l[1:], nil
	} else {
		return nil, CompileError("cdr not permitted on nil")
	}
}
