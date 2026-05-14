package main 

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Args := os.Args[1:]

	if len(Args) < 1 || len(Args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		os.Exit(1)
	}

	rawInput := Args[0]
	bannerName := "standard"

	if len(Args) == 2 {
		bannerName = Args[1]
	}

	bannerName = strings.TrimSuffix(bannerName, ",txt")
	filename := bannerName + ".txt"

	banner, err := LoadBanner(filename)
	if err != nil {
		fmt.Printf("Error could not find or reading file: '%s'\n", filename)
		os.Exit(1)
	}

	var filtered strings.Builder
	for i := 0; i < len(rawInput); i++ {
		r := rune(rawInput[i])

		if r == '\\' && i+1 < len(rawInput) && rawInput[i+1] == 'n' {
			filtered.WriteString("\\n")
			i++
			continue
		}

		if _, ok := banner[r]; ok {
			filtered.WriteRune(r)
		}
	}

	finalInput := filtered.String()
	if finalInput == "" {
		return
	}
	fmt.Print(GenerateArt(finalInput, banner))
}