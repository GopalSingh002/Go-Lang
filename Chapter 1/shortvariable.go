/*
Short Variable Declaration...........................................
:=, can work only in a function, not outside the function
var := value : known as short variable declaration

ex:
*/

package main

import "fmt"

var b = 10 //Declared outside the function
// c :=11 //This line will throw an error

func main() {
	a := 5 //Declared inside the function
	sum := a + b
	fmt.Println("result: ", sum)

}
