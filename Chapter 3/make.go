/*
make() function..............................................................
We can create a slice with the use of make() function
- this is how we can create a dynamic size array
- When you create the slice with make, initially it store the 0's as elements whatever the size you pass in the function
- ex:
a:=make([]int, 5)
[]int: type
5: capacity


make(type, length, capacity)


*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	b := a[1:3]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	/*
		the reason behind the capacity of b is 4 is:
		-slice b is sliced from a from index 1 to 3,
		- slicing concept in go says, how much array the copied one can still store,
		like we provided it index 1 to 3(excluded), and in the slice a , we can still go till 4,
		because the slice b underlying the sama array as slice a.

		currently b store a[1,2](they are indexes not values)
		but b can still store a[1,2,3,4](they are indexes not values)
		thats why b's capacity is 4

	*/
}
