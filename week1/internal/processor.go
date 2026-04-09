package internal

import (
	"strconv"
	"strings"
)

func Process(input string) string {
	tokens := tokenize(input)
	tokens = applyCommands(tokens)
	tokens = fixPunctuation(tokens)
	tokens = fixQuotes(tokens)
	tokens = fixArticles(tokens)
	return strings.Join(tokens, " ")
}

func tokenize(input string) []string {
	input = normalizeCommands(input)

	raw := strings.Fields(input)
	var tokens []string

	for _, field := range raw {
		if isCommand(field) {
			tokens = append(tokens, field)
			continue
		}
		tokens = append(tokens, splitPunct(field)...)
	}

	return tokens
}

func normalizeCommands(input string) string {
	var sb strings.Builder
	i := 0
	for i < len(input) {
		if input[i] == '(' {
			j := i + 1
			for j < len(input) && input[j] != ')' {
				j++
			}
			if j < len(input) {
				inner := input[i+1 : j]
				inner = strings.ReplaceAll(inner, " ", "")
				sb.WriteByte('(')
				sb.WriteString(inner)
				sb.WriteByte(')')
				i = j + 1
			} else {
				sb.WriteByte(input[i])
				i++
			}
		} else {
			sb.WriteByte(input[i])
			i++
		}
	}
	return sb.String()
}

func splitPunct(s string) []string {
	var result []string

	start := 0
	for start < len(s) && isPunctByte(s[start]) {
		start++
	}
	if start > 0 {
		result = append(result, s[:start])
	}

	end := len(s)
	for end > start && isPunctByte(s[end-1]) {
		end--
	}

	if start < end {
		result = append(result, s[start:end])
	}

	if end < len(s) {
		result = append(result, s[end:])
	}

	return result
}

func isPunctByte(b byte) bool {
	return strings.ContainsRune(".,!?:;", rune(b))
}

func isCommand(s string) bool {
	if len(s) < 4 || s[0] != '(' || s[len(s)-1] != ')' {
		return false
	}
	inner := s[1 : len(s)-1]
	parts := strings.SplitN(inner, ",", 2)
	cmd := strings.TrimSpace(parts[0])
	switch cmd {
	case "up", "low", "cap", "hex", "bin":
		return true
	}
	return false
}

func parseCommand(s string) (cmd string, n int) {
	inner := s[1 : len(s)-1]
	parts := strings.SplitN(inner, ",", 2)
	cmd = strings.TrimSpace(parts[0])
	n = 1
	if len(parts) == 2 {
		count, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil && count > 0 {
			n = count
		}
	}
	return
}

func applyCommands(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for _, tok := range tokens {
		if !isCommand(tok) {
			result = append(result, tok)
			continue
		}

		cmd, n := parseCommand(tok)

		wordIndices := collectWordIndices(result, n)

		switch cmd {
		case "up":
			for _, idx := range wordIndices {
				result[idx] = toUpper(result[idx])
			}
		case "low":
			for _, idx := range wordIndices {
				result[idx] = toLower(result[idx])
			}
		case "cap":
			for _, idx := range wordIndices {
				result[idx] = toCap(result[idx])
			}
		case "hex":
			if len(wordIndices) > 0 {
				idx := wordIndices[len(wordIndices)-1]
				result[idx] = fromHex(result[idx])
			}
		case "bin":
			if len(wordIndices) > 0 {
				idx := wordIndices[len(wordIndices)-1]
				result[idx] = fromBin(result[idx])
			}
		}
	}

	return result
}

func collectWordIndices(tokens []string, n int) []int {
	var indices []int
	for i := len(tokens) - 1; i >= 0 && len(indices) < n; i-- {
		if !isPunct(tokens[i]) {
			indices = append([]int{i}, indices...)
		}
	}
	return indices
}

func fixQuotes(tokens []string) []string {
	result := make([]string, 0, len(tokens))
	i := 0
	for i < len(tokens) {
		if tokens[i] == "'" {
			j := i + 1
			for j < len(tokens) && tokens[j] != "'" {
				j++
			}
			if j < len(tokens) {
				inner := strings.Join(tokens[i+1:j], " ")
				result = append(result, "'"+inner+"'")
				i = j + 1
			} else {
				result = append(result, tokens[i])
				i++
			}
		} else {
			result = append(result, tokens[i])
			i++
		}
	}
	return result
}