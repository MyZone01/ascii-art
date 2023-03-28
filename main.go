package main

import (
	ascii "ascii_art/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		arg := os.Args[1]
		text := ""
		typeAscii := "standard"
		if len(os.Args) > 2 && (os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "thinkertoy" || os.Args[len(os.Args)-1] == "shadow") {
			typeAscii = os.Args[len(os.Args)-1]
		}
		asciiCharacters := ascii.ParseFile(typeAscii + ".txt")
		if len(arg) > 10 && arg[:10] == "--reverse=" {
			filePath := strings.Split(arg, `=`)[1]
			_text, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("❌ ERROR: File not found")
				os.Exit(1)
			}
			text = ascii.ConvertAsciiArtToText(string(_text), asciiCharacters)
			fmt.Println(text)
		} else {
			_text := []string{}
			if len(arg) > 9 && arg[:9] == "--output=" && len(os.Args) > 2 {
				_text = strings.Split(os.Args[2], `\n`)
			} else {
				_text = strings.Split(os.Args[1], `\n`)
			}
			if ascii.IsValid(_text) {
				text = ascii.ConvertTextToAsciiArt(_text, asciiCharacters)
				if len(arg) > 9 && arg[:9] == "--output=" && len(os.Args) > 2 {
					outputFileName := strings.Split(arg, `=`)[1]
					file, err := os.Create(outputFileName)
					if err != nil {
						fmt.Println("❌ ERROR: Output file creation error")
					}
					file.WriteString(text)
					file.Close()
				} else {
					fmt.Print(text)
				}
			} else {
				fmt.Println("❌ ERROR: Argument containing unknown characters")
				os.Exit(1)
			}
		}
	}
}
