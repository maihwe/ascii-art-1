package main

import (
	"strings"
	
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	input = strings.ReplaceAll(input, `\n`, "\n")

	onlyNewLine := true

	for _, ch := range input {
		if ch != '\n' {
			onlyNewLine = false
			break
		}
	}
	if onlyNewLine {
		return input
	}

	parts := strings.Split(input, "\n")
	var result strings.Builder

	for _, part := range parts {
		if part == "" {
			result.WriteString("\n")
			continue
		}
		maxrows := 0

		for _, ch := range part {
			if rows := len(banner[ch]); rows > maxrows {
				maxrows = rows
			}
		}
		for row := 0; row < maxrows; row++ {
			for _, char := range part {
				result.WriteString(banner[char][row])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}