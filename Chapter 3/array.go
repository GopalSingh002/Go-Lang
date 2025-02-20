/*
Array in go lang are the collection or squence of similar type of elements
Syntax:
var arrayname [length]datatype

	or

var arrayname [length]datatype{values}

-accessing element
arrayname[index]


if we define lenght: array
if dont: slice
*/

package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var arr1 [10]int
	for i := 0; i < 10; i++ {
		arr1[i] = i + 2
	}
	fmt.Println(arr, arr1)
}
