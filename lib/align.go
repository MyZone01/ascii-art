package ascii_art

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
	"unicode/utf8"
)

const (
	ALIGN_LEFT    = "left"
	ALIGN_CENTER  = "center"
	ALIGN_RIGHT   = "right"
	ALIGN_JUSTIFY = "justify"
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
func AlignLeft(text string, width, colorGap int) string {
	gap := width-utf8.RuneCountInString(text)+colorGap
	if gap < 0 {
		fmt.Println("Can't align text to right the gap value is negative", gap)
		os.Exit(1)
	}
	return text + strings.Repeat(" ", gap)
}

// Function to align text to the center
func AlignCenter(text string, width, colorGap int) string {
	padding := width - utf8.RuneCountInString(text) + colorGap
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return strings.Repeat(" ", leftPadding) + text + strings.Repeat(" ", rightPadding)
}

// Function to align text to the right
func AlignRight(text string, width, colorGap int) string {
	gap := width-utf8.RuneCountInString(text)+colorGap
	if gap < 0 {
		fmt.Println("Can't align text to right the gap value is negative", gap)
		os.Exit(1)
	}
	return strings.Repeat(" ", gap) + text
}

// Function to justify text
func AlignJustify(text string, width, colorGap int) string {
	words := strings.Split(text, " ")
	wordsCount := len(words)
	textSize := 0
	for _, word := range words {
		textSize += utf8.RuneCountInString(word)
	}

	// If there's only one word or the width is smaller than the length of the
	// text, return the text aligned to the left
	if wordsCount == 1 || textSize >= width {
		text := strings.ReplaceAll(text, "R", " ")
		return AlignLeft(text, width, colorGap)
	}

	numGaps := wordsCount - 1
	spaceCount := width - textSize + colorGap
	gapSize := spaceCount / (numGaps)
	if gapSize < 0 {
		fmt.Println("Can't align text to right the gap value is negative", gapSize)
		os.Exit(1)
	}
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
