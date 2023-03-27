package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func IsValid(text []string) bool {
	for _, word := range text {
		for _, char := range word {
			if char < 32 || char > 127 {
				return false
			}
		}
	}
	return true
}

func ParseFile(name string) map[int][][]rune {
	_content, err := os.ReadFile("templates/"+name)
	content := string(_content)
	if err != nil {
		fmt.Println("ERROR: exit when reading file")
		fmt.Println(err)
		os.Exit(1)
	}
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	asciiCharacters := map[int][][]rune{}
	character := [][]rune{}
	actualChar := 32
	for i := 1; i < len(lines); i++ {
		if i % 9 == 0 {
			asciiCharacters[actualChar] = character
			actualChar++
			character = [][]rune{}
			continue
		}
		line := lines[i]
		character = append(character, []rune(line))
	}
	return asciiCharacters
}

func PrintAsciiArt(text []string, asciiCharacters map[int][][]rune) {
	for i, word := range text {
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					line := string(asciiCharacters[int(char)][j])
					fmt.Print(line)
				}
				fmt.Println()
			}
		}
		if i < len(text)-1 {
			fmt.Println()
		}
	}
}
