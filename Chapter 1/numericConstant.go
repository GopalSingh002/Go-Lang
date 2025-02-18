/*
Numeric Constants..................................................
they are the high precision values that do no have a fixed type. they take
any type depending on the context in which they are used.


untyped constants are not assigned a specific type until they are used in an expression.
Go will choose the type based on the operation


ex: 1<<100, shifting the bit by 100
an int in go store 64 bits,, large constants like 1<<100, when you try to fit into max value an int can hold
, if the limit exceed, a compiler error will occur because it is too large to fit in int
*/

package main

import "fmt"

// functions to convert into diff dt
func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	const big = 1 << 100
	const small = big >> 99

	//specifiying them into the datatype like int, float64 to see which is enough to fit the number
	fmt.Println("Small to int: ", needInt(small))
	fmt.Println("Small to float: ", needFloat(small))
	fmt.Println("big to float: ", needFloat(big))

	//if we try to convert big to int, it will throw error, because the value cannot fit

}
