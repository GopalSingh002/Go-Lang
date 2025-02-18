/*
These are the flow control statements
If statements are like its for loops, not working but the syntax

Syntax:
if condition{
	code
}

block of if run only if the condition is true
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		//v:=math.Pow(x,n) we declared v variable for if block that chech if the x^n is < lim or not
		return v
	}
	return lim
}

func main() {
	a := 10
	//if the number is even = true, else = false
	if a%2 == 0 {
		fmt.Println("The number is true")
	}

	/*
		like for loop. if can also a short statement like:
		- the if statement can start with a short statement to execute before condition
		- variable declared in this are only if block scoped
		ex: pow(x,n,lim) function
	*/
	fmt.Println(pow(3, 4, 20)) //limit is return as v = 81
	fmt.Println(pow(3, 2, 10)) // return 9

	/*
		else: this statement execute only when the if condition is fail
		- all the variable declared in the if block are available in any of the else block
	*/

	if a := 10; a%2 != 0 {
		fmt.Println("The number is odd")
	} else {
		fmt.Println(a, "is even")
	}

}
