// Modify the echo program to also print os.Args[0]
package main
import (
	"fmt"
 	"os" 
 )

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	} 
	fmt.Println(s) 
}