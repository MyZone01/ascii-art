package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func ConvertTextToArt(_text, align, color, colorize string, asciiCharacters map[int][][]rune) string {
	result := ""
	colorGap := 0
	text := strings.Split(_text, `\n`)
	if !IsValid(text) {
		fmt.Println("❌ ERROR: Argument containing unknown characters")
		os.Exit(1)
	}
	width := GetTerminalWidth()

	for _, line := range text {
		if line != "" {
			buffer := ""
			for j := 0; j < 8; j++ {
				for _, _char := range line {
					char := string(asciiCharacters[int(_char)][j])
					char, isColored := Colorize(colorize, string(_char), char, color)
					if isColored {
						colorGap += len(color) + 4
					}
					buffer += char
				}
				if len(buffer) > 0 {
					switch align {
					case ALIGN_LEFT:
						result += buffer
						colorGap = 0
					case ALIGN_CENTER:
						result += AlignCenter(buffer, width, colorGap)
						colorGap = 0
					case ALIGN_RIGHT:
						result += AlignRight(buffer, width, colorGap)
						colorGap = 0
					case ALIGN_JUSTIFY:
						result += AlignJustify(strings.ReplaceAll(buffer, "      ", " "), width, colorGap)
						colorGap = 0
					default:
						fmt.Fprintln(os.Stderr, "Invalid alignment type")
						os.Exit(1)
					}
				} else {
					result += buffer
				}
				buffer = ""
				result += "\n"
			}
		} else if line != "\n" {
			result += "\n"
		}
	}
	return result
}

func ConvertArtToText(_text, align, color, colorize string, asciiCharacters map[int][][]rune) string {
	result := ""
	colorGap := 0
	text := [][]rune{}
	width := GetTerminalWidth()
	_lines := strings.Split(strings.ReplaceAll(_text, "\r\n", "\n"), "\n")
	for _, l := range _lines {
		text = append(text, []rune(l))
	}
	for i := 0; i < len(text); i = i + 8 {
		previousIndex := 0
		nbSuccessiveSpace := 0
		buffer:=""
		for j := range text[i] {
			if IsCharacterDelimiter(text, i, j) {
				nbSuccessiveSpace++
				if nbSuccessiveSpace == 1 {
					char := GetMatchingCharacter(text, asciiCharacters, previousIndex, j+1, i)
					char, isColored := Colorize(colorize, char, char, color)
					if isColored {
						colorGap += len(color)
					}
					buffer += char
				} else if nbSuccessiveSpace == 7 {
					buffer += " "
					nbSuccessiveSpace = 0
				}
				previousIndex = j + 1
			} else {
				nbSuccessiveSpace = 0
			}
		}
		switch align {
		case ALIGN_LEFT:
			result += buffer
			colorGap = 0
		case ALIGN_CENTER:
			result += AlignCenter(buffer, width, colorGap)
			colorGap = 0
		case ALIGN_RIGHT:
			result += AlignRight(buffer, width, colorGap)
			colorGap = 0
		case ALIGN_JUSTIFY:
			result += AlignJustify(strings.ReplaceAll(buffer, "      ", " "), width, colorGap)
			colorGap = 0
		default:
			fmt.Fprintln(os.Stderr, "Invalid alignment type")
			os.Exit(1)
		}
		if i+8 < len(text)-1 {
			result += `\n`
		}
	}
	return result + "\n"
}

func GetMatchingCharacter(text [][]rune, asciiCharacters map[int][][]rune, firstCol, lastCol, line int) string {
	for key, asciiChar := range asciiCharacters {
		if len(asciiChar[0]) == len(text[line][firstCol:lastCol]) && line+7 <= len(text) {
			if string(asciiChar[0]) == string(text[line][firstCol:lastCol]) &&
				string(asciiChar[1]) == string(text[line+1][firstCol:lastCol]) &&
				string(asciiChar[2]) == string(text[line+2][firstCol:lastCol]) &&
				string(asciiChar[3]) == string(text[line+3][firstCol:lastCol]) &&
				string(asciiChar[4]) == string(text[line+4][firstCol:lastCol]) &&
				string(asciiChar[5]) == string(text[line+5][firstCol:lastCol]) &&
				string(asciiChar[6]) == string(text[line+6][firstCol:lastCol]) &&
				string(asciiChar[7]) == string(text[line+7][firstCol:lastCol]) {
				return string(rune(key))
			}
		}
	}
	fmt.Println("❌ Error: BAD ASCII FORMAT")
	os.Exit(1)
	return ""
}
