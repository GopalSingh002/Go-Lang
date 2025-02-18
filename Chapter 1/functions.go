/*
Function:
A function can take n arguments
Syntax:
func functionName(arguments datatype) returnType{
	code
	return
}
func main(){
	functionName(parameters)	//function call
}
ex:
*/

package main

import (
	"fmt"
	"math"
)

/*
	when more than one arguments with the same parameters have to be pass, you dont need to mention for all,

instead you can write it in the last argument
ex:
*/
func calc(x, y, z int) int {
	sum := x + y + z
	return sum
}

func area(r float64) float64 { //Function
	return math.Pi * r * r // Return statement
}

func main() {
	fmt.Printf("Area of circle with radius: 4: ")
	fmt.Println(area(4)) //Callling statement
	fmt.Println("sum: ", calc(2, 3, 4))

}
