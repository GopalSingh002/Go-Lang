/*
In go lang, we can use slicing property of an array

[start:end]
end is excluded
[0:n]: from 0 to n-1
[:] : from 0 to n-1
[0:] : from 0 to n-1
[:] : from 0 to n-1
ex:
*/
package main

import "fmt"

func main() {
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := arr1[1:6] // this is a slice not a array
	fmt.Println("Array with sub slice:")
	fmt.Println("arrays before change in arr1")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("Array after change in arr1 only")
	arr1[2] = 1000
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("Array with sub array")
	//if we want to make a subarray, not a slice from the array one
	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //this is an array
	var arr4 [8]int
	copy(arr4[:], arr3[1:6]) // it will copy the arr3 in array 4
	fmt.Println("arrays Before changing in array3")
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println("Arrays after changing in arr3")
	arr3[4] = 12000
	fmt.Println(arr3)
	fmt.Println(arr4)

	//As we can see in the arr4, that the values are 0, this is because the array is created of fixed size, and filled with the default values of int

}
