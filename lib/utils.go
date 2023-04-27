package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func GetArgs() (string, string, string, string, string, string, bool, bool) {
	input := ""
	asciiFont := "standard"
	outputFile := ""
	isReverse := false
	errors := false
	align := "left"
	colorize := ""
	colorCode := ""
	args := os.Args[1:]
	for i, arg := range args {
		if len(arg) > 9 && arg[:9] == "--output=" {
			outputFile = strings.Split(arg, `=`)[1]
		} else if len(arg) > 8 && arg[:8] == "--align=" {
			align = strings.Split(arg, `=`)[1]
		} else if len(arg) > 10 && arg[:10] == "--reverse=" {
			fileName := strings.Split(arg, `=`)[1]
			_text, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("❌ ERROR: File not found")
				errors = true
			}
			input = string(_text)
			if len(input) == 0 {
				fmt.Println("❌ ERROR: Bad file format")
				errors = true
			}
			isReverse = true
		} else if i+1 <= len(args)-1 && len(arg) > 8 && arg[:8] == "--color=" {
			arr := strings.Split(arg, `=`)
			_colorCode := arr[1]
			if _colorCode[0] == '#' {
				colorCode = RGBToANSI(HexToRGB(_colorCode))
			} else if len(_colorCode) >= 12 && _colorCode[:4] == "rgb(" {
				colorCode = RGBToANSI(_colorCode)
			} else {
				_color, ok := ansiColors[_colorCode]
				if ok {
					colorCode = _color
				} else {
					fmt.Println("❌ ERROR: The program don't handle that color try using rgb or hex notation")
					errors = true
				}
			}
			if i+1 == len(args)-1 {
				input = args[i+1]
				break
			} else if i+2 == len(args)-1 {
				colorize = args[i+1]
				input = args[i+2]
				break
			} else if i+3 == len(args)-1 {
				colorize = args[i+1]
				input = args[i+2]
				asciiFont = args[i+3]
				break
			} else {
				fmt.Println("❌ ERROR: The color flag must be placed at last")
				errors = true
			}
		} else if i == len(args)-1 {
			if input != "" || isReverse {
				asciiFont = arg
			} else {
				input = arg
			}
		} else {
			input = arg
		}
	}
	if asciiFont != "standard" && asciiFont != "thinkertoy" && asciiFont != "shadow" && asciiFont != "zigzag" && asciiFont != "htag" {
		fmt.Println("❌ ERROR: The program don't handle that color try using rgb or hex notation")
		errors = true
	}
	if input == "" && len(os.Args) > 1 {
		errors = true
	}
	return input, asciiFont, outputFile, align, colorCode, colorize, isReverse, errors
}

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
	return line+7 <= len(text)-1 &&
		len(text[line]) > 0 && text[line][col] == ' ' &&
		len(text[line+1]) > 0 && text[line+1][col] == ' ' &&
		len(text[line+2]) > 0 && text[line+2][col] == ' ' &&
		len(text[line+3]) > 0 && text[line+3][col] == ' ' &&
		len(text[line+4]) > 0 && text[line+4][col] == ' ' &&
		len(text[line+5]) > 0 && text[line+5][col] == ' ' &&
		len(text[line+6]) > 0 && text[line+6][col] == ' ' &&
		len(text[line+7]) > 0 && text[line+7][col] == ' '
}

func IsAsciiSpace(text [][]rune, line, col int) bool {
	return col+6 <= len(text[line])-1 && IsCharacterDelimiter(text, line, col+1) &&
		IsCharacterDelimiter(text, line, col+2) &&
		IsCharacterDelimiter(text, line, col+3) &&
		IsCharacterDelimiter(text, line, col+4) &&
		IsCharacterDelimiter(text, line, col+5) &&
		IsCharacterDelimiter(text, line, col+6)
}

func PrintAscii(output string) {
	fmt.Print(output)
}

func PrintUsageError() {
	fmt.Println("Usage: go run . [OPTIONS] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . something")
	fmt.Println("Example: go run . something <font>")
	fmt.Println("Example: go run . --reverse=<fileName>")
	fmt.Println("Example: go run . --align=<position> something")
	fmt.Println("Example: go run . --align=<position> something <font>")
	fmt.Println("Example: go run . --output=<fileName> something <font>")
	fmt.Println("Example: go run . --color=<color> <letters to be colored> something")
	fmt.Println()
	fmt.Println("🚨 When you want to combine flag take in mind that the color tag must be the last of them")
	fmt.Println("🚨 Reverse is a text generate with ZIGZAG font is not handle yet")
}

func PrintLogo() {
	_content, err := os.ReadFile("assets/01-logo")
	content := string(_content)
	if err != nil {
		fmt.Println("❌ ERROR: exit when reading logo's file")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(content)
}

func ParseFile(name string, isJustifying bool) map[int][][]rune {
	_content, err := os.ReadFile(name)
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
		if actualChar != 32 && isJustifying {
			line = strings.ReplaceAll(line, " ", "R")
		}
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
