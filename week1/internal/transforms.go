package internal

import (
	"strconv"
	"strings"
	"unicode"
)

func toUpper(word string) string {
	return strings.ToUpper(word)
}

func toLower(word string) string {
	return strings.ToLower(word)
}

func toCap(word string) string {
	if word == "" {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func fromHex(word string) string {
	n, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

func fromBin(word string) string {
	n, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}