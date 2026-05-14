package main

import (
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for i := range 8 {
		var want strings.Builder

		for _, ch := range input {
			art, ok := banner[ch]
			if !ok {
				art = banner[' ']
			}
			want.WriteString(art[i])
		}
		result[i] = want.String()
	}
	return result
}