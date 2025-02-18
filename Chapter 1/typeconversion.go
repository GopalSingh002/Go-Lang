/*
Type conversion...................................................
In go language, we need to do the conversion explicitly, unlike in other languages,
we need to do to explicitly.

Syntax:
T(v)
T: Type
v: variable or value
ex:
*/

package main

import "fmt"

func main() {
	var i int = 22
	var f float32 = 3.14
	var sum int
	// sum = i + f //this line give error as type mismatch
	// sum = float32(i) + f //Same coz, sum var is declared as in

	sum = i + int(f) //This line work fine, as float is converted to int, and the resulting var also a int type
	fmt.Println("Result: ", sum)

}

/*
Type inference...................................................
as we already know, we can ommit the type in the declaration either with the use of := or var varname = value
, compiler assign the data type to the variable automatically according to the value in the right side.
*/
