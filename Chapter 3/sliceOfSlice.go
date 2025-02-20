/*
Slice of Slice
2d slices
its more like storing slices inside a slice
Syntax:

var:=[][]type{
	[]type{values}
	[]type{values}
	.
	.
	.
	[]type{values}
}

Accessing 2d slices:
slice[r][c]
if we need to print each row then we can: slice[r]

*/

package main

import "fmt"

func main() {
	sudo := [][]string{
		[]string{"1", "_", "_", "_"},
		[]string{"_", "_", "_", "_"},
		[]string{"_", "_", "2", "_"},
		[]string{"_", "4", "_", "_"},
	}
	for i := 0; i < len(sudo); i++ {
		fmt.Println(sudo[i])
	}
	//now we can add more values in the sudo
	sudo[0][1] = "2"
	sudo[0][2] = "4"
	sudo[0][3] = "3"
	sudo[1][1] = "2"
	sudo[1][2] = "3"
	sudo[1][3] = "4"
	fmt.Println("2d slice after insertions")

	for i := 0; i < len(sudo); i++ {
		fmt.Println(sudo[i])
	}

}
