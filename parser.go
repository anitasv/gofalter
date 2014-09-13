package main

import (
	"bufio"
)

func Parse(scanner *bufio.Scanner) (LispExpr, bool) {

	for scanner.Scan() {
		tok := scanner.Text()

		if tok == "(" {
			return ConsumeList(scanner)
		} else {
			return LispSymbol(tok), true
		}
	}

	return nil, false
}

func ConsumeList(scanner *bufio.Scanner) (LispList, bool) {
	lispList := LispList(make([]LispExpr, 0))

	for {
		expr, stat := Parse(scanner)
		if !stat {
			break
		}
		t, ok := expr.(LispSymbol)
		if ok {
			if t == ")" {
				return lispList, true
			}
		}
		lispList = append(lispList, expr)
	}
	return lispList, false
}
