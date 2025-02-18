/*
Switch in go
- switch are the easiest way to replacement of multiple if-else block
- switch execute the condition and execute the case whose value is equal to the condition result
- in Go lang switch, we don't need to write break statement explicitly for each case,
go lang manage it by its own
- order of execution of switch from top to bottom, suppose, if two case are given,
case1
cass2
case3
1 and 3 both match with the condition result,
but the case1 is already found it didn't go further, stop the further case evaluation.



Syntax:
switch condition{
case "":
	code
case " ":
	code
default:
	code
}

ex:
*/

package main

import "fmt"

func main() {
	switch 10 % 2 {
	case 0:
		fmt.Println("Number is even")
	case 1:
		fmt.Println("Number is odd")
	default:
		fmt.Println("Invalid number")
	}
}
