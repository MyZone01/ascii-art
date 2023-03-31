package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

const (
	ALIGN_LEFT    = "left"
	ALIGN_CENTER  = "center"
	ALIGN_RIGHT   = "right"
	ALIGN_JUSTIFY = "justify"
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

func IsCharacterDelimiter(text [][]rune, line, col int) bool {
	return text[line][col] == ' ' && text[line+1][col] == ' ' && text[line+2][col] == ' ' && text[line+3][col] == ' ' && text[line+4][col] == ' ' && text[line+5][col] == ' ' && text[line+6][col] == ' ' && text[line+7][col] == ' '
}

func GetArgs() (string, string, string, string, string, string, bool) {
	input := ""
	typeAscii := "standard"
	outputFile := ""
	isReverse := false
	align := "left"
	colorize := ""
	colorCode := ""
	for _, arg := range os.Args {
		if len(arg) > 9 && arg[:9] == "--output=" {
			outputFile = strings.Split(arg, `=`)[1]
		} else if len(arg) > 8 && arg[:8] == "--align=" {
			align = strings.Split(arg, `=`)[1]
		} else if len(arg) > 10 && arg[:10] == "--reverse=" {
			fileName := strings.Split(arg, `=`)[1]
			_text, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("❌ ERROR: File not found")
				os.Exit(1)
			}
			input = string(_text)
			isReverse = true
		} else if len(arg) > 8 && arg[:8] == "--color=" {
			arr := strings.Split(arg, `=`)
			_colorCode := arr[1]
			if _colorCode[0] == '#' {
				colorCode = ParseRGB(HexToRGB(_colorCode))
			} else if len(_colorCode) >= 12 && _colorCode[:4] == "rgb(" {
				colorCode = ParseRGB(_colorCode)
			} else {
				colorCode = ansiColors[_colorCode]
			}
			if len(arr) == 3 {
				colorize = arr[2]
			}
		} else if arg == "standard" || arg == "thinkertoy" || arg == "shadow" {
			typeAscii = arg
		} else {
			input = arg
		}
	}
	return input, typeAscii, outputFile, align, colorCode, colorize, isReverse
}

func PrintAscii(output string) {
	fmt.Print(output)
}

func ParseFile(name string) map[int][][]rune {
	_content, err := os.ReadFile("templates/" + name + ".txt")
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

func SaveFile(fileName string, text string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("❌ ERROR: Output file creation error")
	}
	file.WriteString(text)
	file.Close()
}

func ConvertTextToArt(_text, align, color, colorize string, isColorizing bool, asciiCharacters map[int][][]rune) string {
	result := ""
	text := strings.Split(_text, `\n`)
	if !IsValid(text) {
		fmt.Println("❌ ERROR: Argument containing unknown characters")
		os.Exit(1)
	}
	// Get the terminal width size
	width := GetTerminalWidth()

	for _, line := range text {
		if line != "" {
			buffer := ""
			for j := 0; j < 8; j++ {
				for _, _char := range line {
					char := string(asciiCharacters[int(_char)][j])
					char = Colorize(colorize, string(_char), isColorizing, char, color)
					buffer += char
				}
				if len(buffer) > 0 {
					switch align {
					case ALIGN_LEFT:
						result += buffer
					case ALIGN_CENTER:
						result += AlignCenter(buffer, width)
					case ALIGN_RIGHT:
						result += AlignRight(buffer, width)
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

func ConvertArtToText(_text, algin, color, colorize string, isColorizing bool, asciiCharacters map[int][][]rune) string {
	result := ""
	text := [][]rune{}
	_lines := strings.Split(strings.ReplaceAll(_text, "\r\n", "\n"), "\n")
	for _, l := range _lines {
		text = append(text, []rune(l))
	}
	for i := 0; i < len(text); i = i + 8 {
		previousIndex := 0
		nbSuccessiveSpace := 0
		for j := range text[i] {
			if IsCharacterDelimiter(text, i, j) {
				nbSuccessiveSpace++
				if nbSuccessiveSpace == 1 {
					char := GetMatchingCharacter(text, asciiCharacters, previousIndex, j+1, i)
					char = Colorize(colorize, char, isColorizing, char, color)
					result += char
				} else if nbSuccessiveSpace == 7 {
					result += " "
					nbSuccessiveSpace = 0
				}
				previousIndex = j + 1
			} else {
				nbSuccessiveSpace = 0
			}
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
	return "£€"
}
