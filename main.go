package main

import (
	ascii "ascii_art/lib"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		output := ""
		input, asciiFont, outputFile, align, color, colorize, isReverse, isError := ascii.GetArgs()
		if !isError {
			asciiCharacters := ascii.ParseFile("fonts/"+asciiFont+".txt", align == ascii.ALIGN_JUSTIFY && !isReverse)
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
		} else {
			ascii.PrintUsageError()
		}
	} else {
		ascii.PrintLogo()
		ascii.PrintUsageError()
	}
}
