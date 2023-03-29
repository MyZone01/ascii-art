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
	words := strings.Fields(text)
	wordsCount := len(words)

	// If there's only one word or the width is smaller than the length of the
	// text, return the text aligned to the left
	if wordsCount == 1 || len(text) >= width {
		return AlignLeft(text, width)
	}

	spaceCount := width - len(text)
	spacesPerWord := spaceCount / (wordsCount - 1)
	extraSpaces := spaceCount % (wordsCount - 1)

	// Add the first word
	justifiedText := words[0]

	// Add the spaces between the words
	for i := 1; i < wordsCount; i++ {
		spaces := strings.Repeat(" ", spacesPerWord)
		if i <= extraSpaces {
			spaces += " "
		}
		justifiedText += spaces + words[i]
	}

	return justifiedText
}
