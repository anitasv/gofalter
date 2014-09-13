package main

type CompileError string

func (c CompileError) Error() string {
	return string(c)
}
