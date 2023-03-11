package main

import (
	ascii "ascii_art/lib"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		text := strings.Split(os.Args[1], `\n`)
		asciiCharacters := ascii.ParseFile("standard.txt")
		ascii.PrintAsciiArt(text, asciiCharacters)
	}
}
