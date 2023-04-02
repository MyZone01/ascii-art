package main

import (
	ascii "ascii_art/lib"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		output := ""
		input, typeAscii, outputFile, align, color, colorize, isReverse := ascii.GetArgs()
		asciiCharacters := ascii.ParseFile("fonts/" + typeAscii + ".txt", align == ascii.ALIGN_JUSTIFY && !isReverse)
		if isReverse {
			output = ascii.ConvertArtToText(input, align, color, colorize, asciiCharacters)
		} else {
			output = ascii.ConvertTextToArt(input, align, color, colorize, asciiCharacters)
		}
		if outputFile != "" {
			ascii.SaveFile(outputFile, output)
		} else {
			ascii.PrintAscii(output)
		}
	}
}
