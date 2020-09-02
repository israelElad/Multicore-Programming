package main
import (
	"fmt"
	"os"
	"time"
	"strings"
)
func main() {
	start := time.Now()
	var s, sep string
	for i:=0; i<1000000; i++{
		s = ""
		sep= ""
		for _, arg := range os.Args[1:] { 
			s += sep + arg
			sep = " "
		}
	}
	fmt.Printf("%s\n",s)
	fmt.Printf("%.2fs elapsed after the 100000 run\n", time.Since(start).Seconds())

	start2 := time.Now()
	for i:=0; i<1000000; i++{
		s = " "
		s = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("%s\n",s)
	fmt.Printf("%.2fs elapsed after the 100000 run\n", time.Since(start2).Seconds())
}