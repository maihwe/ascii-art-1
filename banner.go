package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(banner string) (map[rune][]string, error) {
	file, err := os.ReadFile(banner)
	if err != nil {
		return nil, fmt.Errorf("Error: reading file")
	}

	LoadMap := make(map[rune][]string)
	content := string(file)

	if len(content) == 0 {
		return nil, fmt.Errorf("Empty file")
	}

	lines := strings.Split(content, "\n")

	if len(lines) < 855 {
		return nil, fmt.Errorf("bad file format")
	}

	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9

		if start+1 > len(lines) {
			break
		}
		LoadMap[rune(i)] = lines[start+1 : start+9]
	}
	return LoadMap, nil
}
