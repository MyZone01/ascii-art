package ascii_art_test

import (
	ascii "ascii_art/lib"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Test struct {
	name     string
	input    string
	expected int
}

func TestAsciiPrint(t *testing.T) {
	tests := []Test{
		{name: "Basic 0", input: "", expected: 1},
		{name: "Basic 1", input: `\n`, expected: 1},
		{name: "Basic 2", input: `Hello\n`, expected: 2},
		{name: "Basic 3", input: "hello", expected: 3},
		{name: "Basic 4", input: "HeLlO", expected: 4},
		{name: "Basic 5", input: "Hello There", expected: 5},
		{name: "Basic 6", input: "{Hello There}", expected: 6},
		{name: "Basic 7", input: `Hello\nThere`, expected: 7},
		{name: "Basic 8", input: `Hello\n\nThere`, expected: 8},
	}
	file, err := os.ReadFile("expected.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	expected := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "£\n")
	asciiCharacters := ascii.ParseFile("../fonts/standard.txt", false)

	for _, test := range tests {
		output := "\n" + ascii.ConvertTextToArt(test.input, "left", "", "", asciiCharacters)
		if strings.Compare(output, expected[test.expected]) == 0 {
			t.Errorf("🚩Test failed: %s\n Inputed: %s\n Expected: %s\n Output: %s", test.name, test.input, expected[test.expected], output)
		} else {
			fmt.Printf("✅ Test succeeded: %s\n", test.name)
		}

	}
}
