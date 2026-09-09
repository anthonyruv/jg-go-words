// Package words holds small text utilities used to qualify Juggernaut against a real Go repository.
package words

import (
	"strings"

	"github.com/google/uuid"
)

// Count returns the number of whitespace-separated words in text.
func Count(text string) int {
	return len(strings.Fields(text))
}

// Reverse returns the whitespace-separated words in text in reverse order.
func Reverse(text string) string {
	words := strings.Fields(text)
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}
	return strings.Join(words, " ")
}

// Tag returns a stable-looking identifier for text (a UUIDv5 in the URL namespace).
func Tag(text string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(text)).String()
}
