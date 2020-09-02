//Exercise 4.4: Write a version of rotate that operates in a single pass.

package main

import "fmt"

func rotate_n_left(slice []int, n int) {
	if n < 0 {
		fmt.Println("Slice unchanged. Invalid number of rotations!")
		return
	}
	sliceLen := len(slice)
	shiftsCount := n % sliceLen
	// append unknown number of elements to the end of the original slice (that's why we're using "...") and put the result in tempSlice
	tempSlice := append(slice, slice[:shiftsCount]...)
	// take the relevant elements from the tempSlice
	rotatedSlice := tempSlice[shiftsCount : shiftsCount+sliceLen]
	//copy the rotated slice to the original. no need to return it as the underlying array was changed by reference and the capacity/length remains the same.
	copy(slice, rotatedSlice)
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	rotate_n_left(slice, 2)
	fmt.Println(slice)
}
