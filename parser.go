package main

const (
  Nil
  Int
  Float
  String
  Variable
  List
)


type LispType rune

type LispExpr interface {
  String() string

  Atom() bool

  Eval(ram LispList) LispExpr
}

type LispNil struct {}
type LispInt int64
type LispFloat float64
type LispString string
type LispVariable string
type LispList val []LispExpr

// String() definitions for all of them
func (n* LispNil) String() {
  return 'nil'
}

func (l *LispList) String() string {
  out := "("
  for item := range l {
    out = append(out, " ")
    out = append(out, item.String())
  }
  out = append(out, ")")
  return out
}

// Atom() definition

func (LispList* l) Eval() LispExpr {
}

func (n* LispNil) Atom() {
  return true
}

func (n* LispNil) Eval(ram LispList) LispExpr {
  return n
}


func (i* LispInt) Atom() bool {
  return true
}

func (i* LispInt) Eval(ram LispList) LispExpr {
  return i
}


func (i* LispInt) Atom() bool {
  return true
}

func (i* LispInt) Eval(ram LispList) LispExpr {
  return i
}
