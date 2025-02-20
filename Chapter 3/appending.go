/*
Appending to a slice..................................................

append(slice, elements.........n)

*/

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before append: \nlen: ", len(s), "cap: ", cap(s), "slice: ", s)
	s = append(s, 10, 20, 30, 40, 50)
	fmt.Println("Slice after append: \nlen: ", len(s), "cap: ", cap(s), "slice: ", s)
}
