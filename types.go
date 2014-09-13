package main

import (
	"bytes"
)

type LispExpr interface {
	Atom() LispExpr

	Nil() LispExpr

	Car() LispExpr

	Cdr() LispList

	Cons(expr LispExpr) LispList

	Eval(env Env) LispExpr

	Call(args LispList, env Env) LispExpr

	Eq(expr LispExpr) LispExpr

	// Implementing Equaler interface
	Equal(expr LispExpr) bool

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

func (s LispSymbol) Car() LispExpr {
	panic("car is not permitted for symbol")
}

func (l LispList) Car() LispExpr {
	if len(l) > 0 {
		return l[0]
	} else {
		panic("car on nil not permitted")
	}
}

func (s LispSymbol) Cons(expr LispExpr) LispList {
	panic("cons is not permitted for symbol")
}

func (l LispList) Cons(expr LispExpr) LispList {
	newL := make(LispList, len(l)+1, len(l)+1)
	copy(newL[1:], l)
	newL[0] = expr
	return newL
}

func (s LispSymbol) Cdr() LispList {
	panic("cdr is not permitted for symbol")
}

func (l LispList) Cdr() LispList {
	if len(l) > 0 {
		return l[1:]
	} else {
		panic("cdr not permitted on nil")
	}
}

func (l LispSymbol) Equal(expr LispExpr) bool {
	other, ok := expr.(LispSymbol)
	if ok {
		if l == other {
			return true
		}
	}
	return false
}

func (l LispList) Equal(expr LispExpr) bool {
	other, ok := expr.(LispList)
	if ok {
		if len(other) != len(l) {
			return false
		}
		for i, a := range l {
			b := other[i]
			if !a.Equal(b) {
				return false
			}
		}
		return true
	}
	return false
}

func (l LispSymbol) Eq(expr LispExpr) LispExpr {
	if l.Equal(expr) {
		return LISP_TRUE
	} else {
		return LISP_FALSE
	}
}

func (l LispList) Eq(expr LispExpr) LispExpr {
	if l.Equal(expr) {
		return LISP_TRUE
	} else {
		return LISP_FALSE
	}
}
