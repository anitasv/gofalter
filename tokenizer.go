package main

import (
	"bufio"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isParan(r rune) bool {
	return r == ')' || r == '('
}

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0

	var r rune
	var width int
	for width = 0; start < len(data); start += width {
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if isParan(r) {
		return start + width, data[start:(start + 1)], nil
	}

	// Scan until space or paren and return the word.

	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
		if isParan(r) {
			// Not adding i + width because we want to keep Parenthesis back in
			// the buffer for the next token.
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return 0, nil, nil
}

func NewLispScanner(str string) *bufio.Scanner {
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanTokens)
	return scanner
}
