package main

import (
	"fmt"
)

func (env Env) Lookup(ident LispSymbol) LispExpr {
	for _, item := range env {
		if item.ident == ident {
			return item.expr
		}
	}
	panic(fmt.Sprintf("Identifier: %s not defined", ident))
}

func (l LispList) Eval(env Env) LispExpr {
	if len(l) == 0 { // nil
		return l
	}

	return l[0].Call(l[1:], env)
}

func (l LispSymbol) Eval(env Env) LispExpr {
	return env.Lookup(l)
}

func (fnName LispSymbol) Call(args LispList, env Env) LispExpr {
	switch fnName {
	case "car":
		return args[0].Eval(env).Car()
	case "cdr":
		return args[0].Eval(env).Cdr()
	case "atom":
		return args[0].Eval(env).Atom()
	case "cons":
		return args[1].Eval(env).Cons(args[0].Eval(env))
  case "equal":
    return args[0].Eval(env).Eq(args[1].Eval(env))
	case "quote":
		return args[0] // The magical non eval call
	case "cond":
		for _, predarg := range args {
			pL, ok := predarg.(LispList)
			if !ok || len(pL) != 2 {
				panic(fmt.Sprintf("cond() expects pair args %s", predarg))
			}
			pred := pL[0].Eval(env)
			if !pred.IsNil() {
				return pL[1].Eval(env)
			}
		}
	}
	// label evaluation
	lambda := fnName.Eval(env)
	return lambda.Call(args, env)
}

func (cmd LispList) Call(args LispList, env Env) LispExpr {
	instr, ok := cmd[0].(LispSymbol)
	if !ok {
		panic(fmt.Sprintf("Invalid syntax ((( is not permitted in a call sequence %s %s", cmd, args))
	}

	switch instr {
	case "label":
		// ( (label F (lambda (v1 v2 v3) e)) e1 e2 e3)
		// cmd = (label F (lambda (v1 v2 v3) e))
		// args = (e1 e2 e3)
		fnName, ok := cmd[1].(LispSymbol) // F
		if !ok {
			panic(fmt.Sprintf("invalid label syntax %s %s", cmd, args))
		}
		lambda, ok := cmd[2].(LispList) // (lambda (v1 v2 v3) e)
		if !ok {
			panic(fmt.Sprintf("invalid label syntax %s %s", cmd, args))
		}

		augEnv := Env{NewAssoc(fnName, lambda)}

		env = env.Augment(augEnv)

		return lambda.Call(args, env)

	case "lambda":
		// ((lambda (v1 v2 v3) e) e1 e2 e3)
		// cmd = (lambda (v1 v2 v3) e)
		// args = (e1 e2 e3)
		lambdaVars, ok := cmd[1].(LispList) // (v1 v2 v3)
		if !ok {
			panic(fmt.Sprintf("invalid lamda syntax %s %s", cmd, args))
		}
		lambdaExpr := cmd[2] // e

		augEnv := make(Env, 0, len(lambdaVars))
		for i, lv := range lambdaVars {
			lvId, ok := lv.(LispSymbol)
			if !ok {
				panic(fmt.Sprintf("Not a valid variable for lambda: %s", lv))
			}
			e := args[i].Eval(env)
			augEnv = append(augEnv, NewAssoc(lvId, e))
		}

		env = env.Augment(augEnv)
		return lambdaExpr.Eval(env)
	}

	panic(fmt.Sprintf("lambda or label expected", cmd))
}
