package main

import (
	ascii "ascii_art/lib"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		output := ""
		input, typeAscii, outputFile, align, isReverse := ascii.GetArgs()
		asciiCharacters := ascii.ParseFile(typeAscii + ".txt")
		if isReverse {
			output = ascii.ConvertArtToText(input, align, asciiCharacters)
		} else {
			output = ascii.ConvertTextToArt(input, align, asciiCharacters)
		}
		if outputFile != "" {
			ascii.SaveFile(outputFile, output)
		} else {
			ascii.PrintAscii(output)
		}
	}
}
