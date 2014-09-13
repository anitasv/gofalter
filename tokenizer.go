package main

import (
	"bufio"
	"strings"
	"unicode/utf8"
)

// Taken from pkg/bufio/scan.go

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

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
		if !isSpace(r) {
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
		if isSpace(r) {
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
