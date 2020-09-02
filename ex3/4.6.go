//Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func utf8SpacesToASCIISpaces(str []byte) []byte {
	for i := 0; i <= len(str)-1; {
		strSliceFromI := str[i:]
		currentRune, size := utf8.DecodeRune(strSliceFromI)
		if unicode.IsSpace(currentRune) {

			//put ASCII space in i index, and copy all the characters after the UTF space after it.
			str[i] = ' '
			afterUTFSpace := str[i+size:]
			afterASCIISpace := str[i+1:]
			copy(afterASCIISpace, afterUTFSpace)

			str = str[:len(str)+1-size]
			i++
		} else {
			i += size
		}
	}
	return str
}

func main() {
	str := "Hello　世界"
	fmt.Println(str)
	str = string(utf8SpacesToASCIISpaces([]byte(str)))
	fmt.Println(str)
}
