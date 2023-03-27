package main

import (
	ascii "ascii_art/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		text := strings.Split(os.Args[1], `\n`)
		if ascii.IsValid(text) {
			asciiCharacters := ascii.ParseFile("standard.txt")
			ascii.PrintAsciiArt(text, asciiCharacters)
		} else {
			fmt.Println("❌ ERROR: Argument containing unknown characters")
		}
	}
}
