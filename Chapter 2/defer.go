/*
Defer.................................................................
- Defer in go lang is used to schedule the execute of the function according to its surrounding functions
- defer function executed at their time, but return only after its surrounding functions are returned
- Execution order: They executed in LIFO manner
suppose you  have 3 defer function, but they executed in the LIFO order after the surrounding function returns
- defer functions are mostly used for the cleanup purpose.

ex:
*/

package main

import "fmt"

func de() {
	fmt.Println(("in function before defer"))
	defer fmt.Println("Function defer")
}

func main() {
	fmt.Println("Before defer")
	defer fmt.Println("First defer")
	de()
	defer fmt.Println("Second defer")
	fmt.Println("after defer")
}
