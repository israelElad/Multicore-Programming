//Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

import "fmt"

func remove_adj_duplicates(slice []string) []string {
	index := 0                  //keep track of where to write the strings without duplicates in the slice
	for _, str := range slice { //no need for the index while traversing the slice
		if str == slice[index] { //an adjacent duplicate
			continue
		}
		//not duplicated(at least not adjacently)
		index++
		slice[index] = str
	}
	return slice[:index+1] //return a slice with the relevant elements only

}

func main() {
	slice := []string{"a", "c", "c", "c", "ab", "c", "c", "a", "c"} //slice declaration(no size specified)
	slice = remove_adj_duplicates(slice)
	fmt.Println(slice)
}
