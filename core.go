package main

import (
	"fmt"
)

func (env Env) Lookup(ident LispSymbol) (LispExpr, error) {
	for _, item := range env {
		if item.ident == ident {
			return item.expr, nil
		}
	}
	return nil, CompileError(fmt.Sprintf("Identifier: %s not defined", ident))
}

func (l LispList) Eval(env Env) (LispExpr, error) {
	if len(l) == 0 { // nil
		return l, nil
	}

	return l[0].Call(l[1:], env)
}

func (l LispSymbol) Eval(env Env) (LispExpr, error) {
	return env.Lookup(l)
}

func (fnName LispSymbol) Call(args LispList, env Env) (LispExpr, error) {
	switch fnName {
	case "car":
		result, err := args[0].Eval(env)
		if err != nil {
			return nil, err
		}
		return result.Car()
	case "cdr":
		result, err := args[0].Eval(env)
		if err != nil {
			return nil, err
		}
		return result.Cdr()
	case "atom":
		result, err := args[0].Eval(env)
		if err != nil {
			return nil, err
		}
		return result.Atom(), nil
	case "cons":
		result1, err := args[1].Eval(env)
		if err != nil {
			return nil, err
		}
		result2, err := args[0].Eval(env)
		if err != nil {
			return nil, err
		}
		return result1.Cons(result2)
	case "quote":
		return args[0], nil // The magical non eval call
	case "cond":
		for _, predarg := range args {
			pL, ok := predarg.(LispList)
			if !ok || len(pL) != 2 {
				return nil, CompileError(fmt.Sprintf("cond() expects pair args %s", predarg))
			}
			pred, err := pL[0].Eval(env)
			if err != nil {
				return nil, err
			}
			if !pred.IsNil() {
				return pL[1].Eval(env)
			}
		}
	}
	// label evaluation
	lambda, err := fnName.Eval(env)
	if err != nil {
		return nil, err
	}
	return lambda.Call(args, env)
}

func (cmd LispList) Call(args LispList, env Env) (LispExpr, error) {
	instr, ok := cmd[0].(LispSymbol)
	if !ok {
		return nil, CompileError(fmt.Sprintf("Invalid syntax ((( is not permitted in a call sequence %s %s", cmd, args))
	}

	switch instr {
	case "label":
		// ( (label F (lambda (v1 v2 v3) e)) e1 e2 e3)
		// cmd = (label F (lambda (v1 v2 v3) e))
		// args = (e1 e2 e3)
		fnName, ok := cmd[1].(LispSymbol) // F
		if !ok {
			return nil, CompileError(fmt.Sprintf("invalid label syntax %s %s", cmd, args))
		}
		lambda, ok := cmd[2].(LispList) // (lambda (v1 v2 v3) e)
		if !ok {
			return nil, CompileError(fmt.Sprintf("invalid label syntax %s %s", cmd, args))
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
			return nil, CompileError(fmt.Sprintf("invalid lamda syntax %s %s", cmd, args))
		}
		lambdaExpr := cmd[2] // e

		augEnv := make(Env, 0, len(lambdaVars))
		for i, lv := range lambdaVars {
			lvId, ok := lv.(LispSymbol)
			if !ok {
				return nil, CompileError(fmt.Sprintf("Not a valid variable for lambda: %s", lv))
			}
			e, err := args[i].Eval(env)
			if err != nil {
				return nil, err
			}
			augEnv = append(augEnv, NewAssoc(lvId, e))
		}

		env = env.Augment(augEnv)
		return lambdaExpr.Eval(env)
	}

	return nil, CompileError(fmt.Sprintf("lambda or label expected", cmd))
}
