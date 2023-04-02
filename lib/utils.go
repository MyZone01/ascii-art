package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

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
				colorCode = RGBToANSI(HexToRGB(_colorCode))
			} else if len(_colorCode) >= 12 && _colorCode[:4] == "rgb(" {
				colorCode = RGBToANSI(_colorCode)
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
	return line+7 <= len(text)-1 && text[line][col] == ' ' && text[line+1][col] == ' ' && text[line+2][col] == ' ' && text[line+3][col] == ' ' && text[line+4][col] == ' ' && text[line+5][col] == ' ' && text[line+6][col] == ' ' && text[line+7][col] == ' '
}

func PrintAscii(output string) {
	fmt.Print(output)
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
