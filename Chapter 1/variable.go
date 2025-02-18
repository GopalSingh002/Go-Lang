/*
Varibales  in Go:......................................................
in go lang, we delcare the variable in two ways:
syntax:

1. var variableName Datatype

		or
2. varibaleName:=value

In 1st way, we can assign value or left for future to be assign, but the variable is assigned the data type
In 2nd way, you have to provide the value to the variable, so the compiler will assign the datatype to the variable according to the value.
if the initializer(value ) is present the type can be omitted
ex:
*/

package main

import "fmt"

func main() {
	var a = 2  //1st way
	var b int  //1st way, declaring
	b = 3      //1st way, defining
	c := a + b //3nd way
	var x float32
	x = 3

	fmt.Println("Result: ", x, c)
}
