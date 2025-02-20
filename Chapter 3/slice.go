/*
Slices in Go
Where the arrays are fixed size, slice are more flexible iwth with size,
we can append the slices.
Every slice have two components:
Lenght:- Current number of elements inside the slice
Capacity:- capacity refers to how many elements we can store in the slice till now,

- A slice have a max capacity of 2Xcapacity
-In start the slice have 0 len and 0 cap
suppose you have a slice of len 4, capacity 4.
the max capacity became 8.


var s []int
len(s) //0
cap(0) //0

this slice call nil slice





WHY 2x:
reason behind it, whenever the new element is added to the slice, that make the overflow from the capacity
a new slice is created with a new capacity, with the added element.
Now the concern can came of space, like the max capacity till now is 4, you added one more element, instead of 5 it take 8 space.
But this is the trade-off between the time and space,
everytime slice overflow, all the element are copied in the another slice with new len and cap, suppose doing this work
everytime take a lot of time, so it double the current size, so that, the copy time can be reduce.

we can find the length and capacity of the slice by:
len(slice)
cap(slice)

suppose you have a slice of len 3, we can extend it to 10 by:
slice[:10]
or if we need to drop it, we can :

slice[:8],
from first:
slice[1:8]
, slice will store the element 1 to 8-1
slicing property






Syntax:
[]T , T=type

var := [size]datatype{values}

var varname []datatype


Slicing

arr[start:end]
end is excluded
*/

package main

import "fmt"

func main() {
	slice1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//slicing
	var slice3 []int = slice1[2:8]
	fmt.Println(slice1)
	fmt.Println(slice3)
	fmt.Println("Slices before")

	/*
	   Slices are like references to arrays........................................
	   A slice does not store any data, they are just storing an array,
	   all the slices sharing the same array, if the one will get changed, all the slices will reflect the changes:

	*/

	fmt.Println("changing the value of slice1 index 2 to 2000")
	slice1[2] = 2000
	fmt.Println("Slices After")
	fmt.Println(slice1)
	fmt.Println(slice3)
}
