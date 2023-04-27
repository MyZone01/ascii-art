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
	font     string
	input    string
	expected int
}

func TestAsciiPrint(t *testing.T) {
	tests := []Test{
		{name: "Basic 1", font: "standard", input: `Hello\n`, expected: 1},
		{name: "Basic 2", font: "standard", input: "hello", expected: 2},
		{name: "Basic 3", font: "standard", input: "HeLlO", expected: 3},
		{name: "Basic 4", font: "standard", input: "Hello There", expected: 4},
		{name: "Basic 5", font: "standard", input: "{Hello There}", expected: 5},
		{name: "Basic 6", font: "standard", input: `Hello\nThere`, expected: 6},
		{name: "Basic 7", font: "standard", input: `Hello\n\nThere`, expected: 7},
	}
	file, err := os.ReadFile("expected.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	expected := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "£\n")
	
	for _, test := range tests {
		asciiCharacters := ascii.ParseFile("../fonts/"+test.font+".txt", false)
		output := ascii.ConvertTextToArt(test.input, "left", "", "", asciiCharacters)
		if output != expected[test.expected] {
			t.Errorf("🚩Test failed: %s\n Inputted: %s\n Expected:\n%s\n Output:\n%s", test.name, test.input, expected[test.expected], output)
		} else {
			fmt.Printf("✅ Test succeeded: %s\n", test.name)
		}

	}
}
