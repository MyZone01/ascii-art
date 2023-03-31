package ascii_art

import (
	"strings"
	"syscall"
	"unsafe"
)

// Function to get the terminal width size
func GetTerminalWidth() int {
	winsize := struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}{}
	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&winsize)))
	return int(winsize.Col)
}

// Function to align text to the left
func AlignLeft(text string, width int) string {
	return text + strings.Repeat(" ", width-len(text))
}

// Function to align text to the center
func AlignCenter(text string, width int) string {
	padding := width - len(text)
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return strings.Repeat(" ", leftPadding) + text + strings.Repeat(" ", rightPadding)
}

// Function to align text to the right
func AlignRight(text string, width int) string {
	return strings.Repeat(" ", width-len(text)) + text
}

// Function to justify text
func AlignJustify(text string, width int) string {
	words := strings.Split(text, "      ")
	wordsCount := len(words)
	textSize := 0
	for _, word := range words {
		textSize += len(word)
	}

	// If there's only one word or the width is smaller than the length of the
	// text, return the text aligned to the left
	if wordsCount == 1 || textSize >= width {
		return AlignLeft(text, width)
	}

	numGaps := wordsCount - 1
	spaceCount := width - textSize
	gapSize := spaceCount / (numGaps)
	extraSpaces := spaceCount % (numGaps)

	// Build the justified text
	var justifiedText string
	for i, word := range words {
		justifiedText += word
		if i < numGaps {
			justifiedText += strings.Repeat(" ", gapSize)
			if i < extraSpaces {
				justifiedText += " "
			}
		}
	}

	return strings.ReplaceAll(justifiedText, "R", " ")
}
