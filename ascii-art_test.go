package main

import (
    "os"
    "reflect"
    "testing"
)

func TestLoadBanner(t *testing.T) {
    // Create a temporary file with enough lines
    content := ""
    for i := 0; i < 855; i++ {
        content += "line\n"
    }
    tmpFile := "test_banner.txt"
    if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpFile)

    banner, err := LoadBanner(tmpFile)
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
    if banner == nil {
        t.Errorf("expected banner map, got nil")
    }
}

func TestGenerateArt(t *testing.T) {
    banner := map[rune][]string{
        'A': {"A1", "A2"},
        'B': {"B1", "B2"},
    }

    result := GenerateArt("AB", banner)
    if result == "" {
        t.Errorf("expected non-empty result")
    }

    if GenerateArt("", banner) != "" {
        t.Errorf("expected empty string for empty input")
    }
}

func TestRenderLine(t *testing.T) {
    banner := map[rune][]string{
        'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
        ' ': {"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
    }

    result := RenderLine("A", banner)
    if len(result) != 8 {
        t.Errorf("expected 8 lines, got %d", len(result))
    }
    if result[0] != "A1" {
        t.Errorf("expected A1, got %s", result[0])
    }
}

func TestValidateInput(t *testing.T) {
    if _, err := ValidateInput("ABC"); err != nil {
        t.Errorf("expected no error, got %v", err)
    }

    if _, err := ValidateInput("你好"); err == nil {
        t.Errorf("expected error for unsupported character")
    }
}

func TestSplitInput(t *testing.T) {
    input := "Hello\\nWorld"
    result := SplitInput(input)

    if !reflect.DeepEqual(result, []string{"Hello", "World"}) {
        t.Errorf("unexpected split result: %v", result)
    }
}
