package internal

import "strings"

const vowelsAndH = "aeiouAEIOUhH"

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			next := words[i+1]
			if next == "" {
				continue
			}
			firstChar := rune(next[0])
			if strings.ContainsRune(vowelsAndH, firstChar) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return words
}