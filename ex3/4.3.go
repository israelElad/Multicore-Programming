//Rewrite reverse to use an array pointer instead of a slice.
package main

import (
	"fmt"
)

// reverse reverses a slice of ints in place.
func reverse(ptr *[6]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a, "\n\n")
	reverse(&a)
	fmt.Println(a, "\n\n") // "[5 4 3 2 1 0]"
}
