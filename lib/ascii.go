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
	_content, err := os.ReadFile("templates/" + name)
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
		if i%9 == 0 {
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

func ConvertTextToAsciiArt(text []string, asciiCharacters map[int][][]rune) string {
	result := ""
	for _, word := range text {
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					result += string(asciiCharacters[int(char)][j])
				}
				result += "\n"
			}
		} else if word != "\n" {
			result += "\n"
		}
	}
	return result
}

func ConvertAsciiArtToText(_text string, asciiCharacters map[int][][]rune) string {
	result := ""
	text := [][]rune{}
	_lines := strings.Split(strings.ReplaceAll(_text, "\r\n", "\n"), "\n")
	for _, l := range _lines {
		text = append(text, []rune(l))
	}
	previousIndex := 0
	nbSuccessiveSpace := 0
	isSuccessiveSpace := false
	for i := range text[0] {
		if IsCharacterDelimiter(text, 0, i) {
			nbSuccessiveSpace++
			if !isSuccessiveSpace {
				result += GetMatchingCharacter(text, asciiCharacters, previousIndex, i+1, 0)
				nbSuccessiveSpace = 0
			} else if nbSuccessiveSpace == 6 {
				result += " "
				nbSuccessiveSpace = 0
			}
			previousIndex = i + 1
			isSuccessiveSpace = true
		} else {
			isSuccessiveSpace = false
		}
	}
	return result
}

func GetMatchingCharacter(text [][]rune, asciiCharacters map[int][][]rune, firstCol, lastCol, line int) string {
	for key, char := range asciiCharacters {
		if len(char[0]) == len(text[line][firstCol:lastCol]) {
			if string(char[0]) == string(text[line][firstCol:lastCol]) &&
				string(char[1]) == string(text[line+1][firstCol:lastCol]) &&
				string(char[2]) == string(text[line+2][firstCol:lastCol]) &&
				string(char[3]) == string(text[line+3][firstCol:lastCol]) &&
				string(char[4]) == string(text[line+4][firstCol:lastCol]) &&
				string(char[5]) == string(text[line+5][firstCol:lastCol]) &&
				string(char[6]) == string(text[line+6][firstCol:lastCol]) &&
				string(char[7]) == string(text[line+7][firstCol:lastCol]) {
				return string(rune(key))
			}
		}
	}
	return "€"
}

func IsCharacterDelimiter(text [][]rune, line, col int) bool {
	return text[line][col] == ' ' && text[line+1][col] == ' ' && text[line+2][col] == ' ' && text[line+3][col] == ' ' && text[line+4][col] == ' ' && text[line+5][col] == ' ' && text[line+6][col] == ' ' && text[line+7][col] == ' '
}

func PrintText(text [][]rune) {
	for _, l := range text {
		for _, char := range l {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
