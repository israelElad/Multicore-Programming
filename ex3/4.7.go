// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place.
package main

import (
	"fmt"
	"unicode/utf8"
)

// reverses a slice of bytes in place.
func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// reverses a bytes slice containing UTF8 runes(each rune occupies 1-4 bytes).
// first we reverse each rune. Finally when we're done we reverse the whole slice.
func UTF8Reverse(bRunes []byte) []byte {
	i := 0
	//iterate over each rune
	for i <= len(bRunes)-1 {
		_, currentRuneSize := utf8.DecodeRune(bRunes[i:])
		currentRune := bRunes[i : i+currentRuneSize]
		reverse(currentRune)
		i += currentRuneSize
	}
	reverse(bRunes)
	return bRunes
}

func main() {
	s := "Hello, 世界"
	b := []byte(s)
	fmt.Println(string(b))
	b = UTF8Reverse(b)
	fmt.Println(string(b))
}
