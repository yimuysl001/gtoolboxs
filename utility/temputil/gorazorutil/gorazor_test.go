package gorazorutil

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestRazo(t *testing.T) {
	str := "select [ DB=123] * FROM [TABLE]"

	// Define a regular expression pattern to match content within square brackets with an equal sign and optional spaces.
	regex := regexp.MustCompile(`\[\s*(\w+)\s*=\s*([^]]+)\s*]`)

	// Find all matches in the string.
	matches := regex.FindAllStringSubmatch(str, -1)

	// Iterate through the matches and extract the content.
	for _, match := range matches {
		// match[1] contains the key inside the square brackets (e.g., "DB")
		// match[2] contains the value inside the square brackets (e.g., "123")
		key := strings.TrimSpace(match[1])
		value := strings.TrimSpace(match[2])

		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	allString := regex.ReplaceAllString(str, "")
	fmt.Println(allString)
}
