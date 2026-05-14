package main

import (
	"fmt"
)

func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r < 32 || r > 126 {
			return r, fmt.Errorf("Unsupported character: %c", r)
		}
	}
	return 0, nil
}