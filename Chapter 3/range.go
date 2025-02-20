/*
Range;.........................................................
range in go lang, work like a iterator over data structures like: array, slice, string, map
- range return two things:
index, value

syntax:
range var // (var must be a iteratable)

we can skip the index or value by:
for _,val:=range iterVar

for index, _:=range itervar   OR for index := range iterVar
*/
package main

import "fmt"

func main() {
	elements := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	for i, _ := range elements {
		fmt.Println(i, "index")
	}

	for i, v := range elements {

		fmt.Println(i, "index, value", v)
	}

}
