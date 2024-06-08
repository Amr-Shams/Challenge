package challenge

import (
	"testing"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func TestFromFile(t *testing.T) {
	// Set the path of the input file
	viper.Set("input", "/Users/amraly/Projects/aoc2024/challenge/day2/input.txt")

	// Call the function
	input := FromFile()

	// Read all lines from the input
	var lines []string
	for line := range input.Lines() {
		lines = append(lines, line)
	}

	// Read the test file directly
	content, err := os.ReadFile("/Users/amraly/Projects/aoc2024/challenge/day2/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Split the content by lines
	expectedLines := strings.Split(string(content), "\n")

	// Compare the lines read by the function and the lines in the file
	for i, line := range lines {
		if i >= len(expectedLines) {
			t.Errorf("Extra line in FromFile(): %v", line)
			continue
		}
		if line != expectedLines[i] {
			t.Errorf("FromFile() line %d = %v, want %v", i, line, expectedLines[i])
		}
	}
}