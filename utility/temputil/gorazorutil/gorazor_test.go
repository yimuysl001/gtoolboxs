package gorazorutil

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"
	"strings"
	"testing"
)

func TestRazo(t *testing.T) {
	str := "select [ DB=123=] * FROM [TABLE]"

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

func TestB(t *testing.T) {
	err2 := g.Redis().SetEX(context.Background(), "test123", map[string]interface{}{"qw": 123}, 30)
	fmt.Println(err2)
	ex, err := g.Redis().Get(context.Background(), "test123")
	fmt.Println(ex, err)
}
