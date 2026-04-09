package internal

import (
	"strings"
)

const punctuationMarks = ".,!?:;"

func isPunct(token string) bool {
	if token == "" {
		return false
	}
	for _, r := range token {
		if !strings.ContainsRune(punctuationMarks, r) {
			return false
		}
	}
	return true
}

func fixPunctuation(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for _, tok := range tokens {
		if tok == "" {
			continue
		}
		if isPunct(tok) && len(result) > 0 {
			result[len(result)-1] = result[len(result)-1] + tok
		} else {
			result = append(result, tok)
		}
	}

	return result
}